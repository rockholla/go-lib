package marshal

// TODO: this is probably a bit hacky, not even enough test cases to fully validate,
//       but practically the best option I've found so far

import (
	"fmt"
	"reflect"
	"regexp"
	"sort"
	"strings"
)

const (
	indentation = 2
)

// YAMLWithComments will marshal an interface, respecting a "comment" metadata/tag on a property/field as a comment for that property/field
func YAMLWithComments(data interface{}, atIndent int) (string, error) {
	var result string

	// based on our depth of the tree, we'll set our indent
	indent := strings.Repeat(" ", atIndent)

	// our reusable anon function here for processing values of different types
	processValue := func(value reflect.Value, comment string) error {
		switch value.Kind() {
		case reflect.Struct, reflect.Ptr, reflect.Map:
			if comment != "" {
				result = fmt.Sprintf("%s %s\n", result, comment)
			}
			nested, err := YAMLWithComments(value.Interface(), atIndent+indentation)
			if err != nil {
				return err
			}
			result = fmt.Sprintf("%s\n%s", result, nested)
		case reflect.Slice:
			if value.Len() == 0 {
				result = fmt.Sprintf("%s [] %s\n", result, comment)
			} else {
				result = fmt.Sprintf("%s %s\n", result, comment)
				for i := 0; i < value.Len(); i++ {
					result = fmt.Sprintf("%s%s-", result, indent)
					nested, err := YAMLWithComments(value.Index(i).Interface(), atIndent+indentation)
					if err != nil {
						return err
					}
					nested = strings.TrimLeft(nested, " ")
					result = fmt.Sprintf("%s %s", result, nested)
				}
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64, reflect.Bool:
			result = fmt.Sprintf("%s %v %s", result, value, comment)
		default:
			if strings.Contains(value.String(), "\n") {
				result = fmt.Sprintf("%s | %s\n", result, comment)
				multiline := ""
				for _, line := range strings.Split(value.String(), "\n") {
					line = strings.TrimSpace(line)
					if line == "" {
						multiline = fmt.Sprintf("%s\n", multiline)
					} else {
						multiline = fmt.Sprintf("%s%s  %s\n", multiline, indent, line)
					}
				}
				result = fmt.Sprintf("%s%s", result, multiline)
			} else {
				if comment != "" {
					comment = fmt.Sprintf(" %s", comment)
				}
				result = fmt.Sprintf("%s \"%v\"%s", result, value, comment)
			}
		}
		result = fmt.Sprintf("%s\n", result)
		return nil
	}

	// use reflection to construct our YAML string
	dataValue := reflect.ValueOf(data)
	if dataValue.Kind() == reflect.Ptr {
		if dataValue.IsNil() {
			return result, nil
		}
		dataValue = dataValue.Elem()
	}
	switch dataValue.Kind() {
	case reflect.Struct:
		for i := 0; i < dataValue.NumField(); i++ {
			fieldValue := dataValue.Field(i)
			fieldType := dataValue.Type().Field(i)
			comment, _ := fieldType.Tag.Lookup("comment")
			yamlKeyValue, _ := fieldType.Tag.Lookup("yaml")
			yamlKeyValueParts := strings.Split(yamlKeyValue, ",")
			if containsOmitEmpty(yamlKeyValueParts) && (fieldValue.IsZero() || isEmptyMap(fieldValue)) || comment == "exclude" || yamlKeyValueParts[0] == "-" {
				continue
			}
			result = fmt.Sprintf("%s%s%s:", result, indent, yamlKeyValueParts[0])
			if comment != "" {
				comment = fmt.Sprintf("# %s", comment)
			}
			if err := processValue(fieldValue, comment); err != nil {
				return result, err
			}
		}
	case reflect.Map:
		sortedStringKeys := []string{}
		stringKeysToMapValues := map[string]reflect.Value{}
		for _, key := range dataValue.MapKeys() {
			sortedStringKeys = append(sortedStringKeys, key.String())
			stringKeysToMapValues[key.String()] = dataValue.MapIndex(key)
		}
		sort.Strings(sortedStringKeys)
		for _, key := range sortedStringKeys {
			result = fmt.Sprintf("%s%s%s:", result, indent, key)
			if err := processValue(stringKeysToMapValues[key], ""); err != nil {
				return result, err
			}
		}
	default:
		if err := processValue(dataValue, ""); err != nil {
			return result, err
		}
	}

	reCompact, _ := regexp.Compile("(?m)\\n{2,}")
	result = reCompact.ReplaceAllString(result, "\n")
	return result, nil
}

func containsOmitEmpty(yamlTagValueSplit []string) bool {
	for _, valueItem := range yamlTagValueSplit {
		if valueItem == "omitempty" {
			return true
		}
	}
	return false
}

func isEmptyMap(v reflect.Value) bool {
	if v.Kind() == reflect.Map && len(v.MapKeys()) == 0 {
		return true
	}
	return false
}
