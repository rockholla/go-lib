package marshal

// TODO: this is probably a bit hacky, not even enough test cases to fully validate,
//       but practically the best option I've found so far

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

const (
	indentation = 2
)

// YAMLWithComments will marshal an interface, respecting a "help" metadata/tag on a property/field as a comment for that property/field
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
				result = fmt.Sprintf("%s []", result)
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
			result = fmt.Sprintf("%s %v", result, value)
		default:
			if strings.Contains(value.String(), "\n") {
				result = fmt.Sprintf("%s\n %s |\n", comment, result)
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
				result = fmt.Sprintf("%s \"%v\" %s", result, value, comment)
			}
		}
		result = fmt.Sprintf("%s\n", result)
		return nil
	}

	// use reflection to construct our YAML string
	dataTypeOf := reflect.TypeOf(data)
	if dataTypeOf.Elem().Kind() == reflect.Struct {
		dataValueOf := reflect.ValueOf(data)
		if dataValueOf.IsNil() {
			return result, nil
		}
		dataValueOf = dataValueOf.Elem()
		for i := 0; i < dataValueOf.NumField(); i++ {
			fieldValue := dataValueOf.Field(i)
			fieldType := dataValueOf.Type().Field(i)
			help, _ := fieldType.Tag.Lookup("help")
			yamlKeyValue, _ := fieldType.Tag.Lookup("yaml")
			if help == "exclude" || yamlKeyValue == "-" {
				continue
			}
			yamlKeyValue = strings.Split(yamlKeyValue, ",")[0]
			result = fmt.Sprintf("%s%s%s:", result, indent, yamlKeyValue)
			if help != "" {
				help = fmt.Sprintf("# %s", help)
			}
			if err := processValue(fieldValue, help); err != nil {
				return result, err
			}
		}
	} else {
		if err := processValue(reflect.ValueOf(data), ""); err != nil {
			return result, err
		}
	}

	reCompact, _ := regexp.Compile("(?m)\\n{2,}")
	result = reCompact.ReplaceAllString(result, "\n")
	return result, nil
}
