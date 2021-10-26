package marshal

import (
	"testing"

	"github.com/andreyvit/diff"
)

type testStruct struct {
	FieldOne              string                     `yaml:"fieldOne" comment:"field one comment"`
	FieldTwo              string                     `yaml:"fieldTwo" comment:"field two comment"`
	NumberField           int                        `yaml:"numberField" comment:"number field comment"`
	EmbeddedObject        *embeddedStruct            `yaml:"embeddedObject" comment:"embedded object comment"`
	EmbeddedObjectTwo     *embeddedStruct            `yaml:"embeddedObjectTwo" comment:"embedded object two comment"`
	EmbeddedList          []*embeddedStruct          `yaml:"embeddedList" comment:"embedded list comment"`
	FieldOmittedZeroValue string                     `yaml:"fieldOmittedZeroValue,omitempty"`
	FieldOmittedNoMarshal string                     `yaml:"-"`
	MapField              map[string]string          `yaml:"mapField" comment:"map field comment"`
	MapComplexField       map[string]*embeddedStruct `yaml:"mapComplexField" comment:"map complex field"`
	MapFieldZeroValue     map[string]string          `yaml:"mapFieldZeroValue,omitempty"`
}

type embeddedStruct struct {
	EmbeddedFieldOne string `yaml:"embeddedFieldOne" comment:"embedded field one"`
}

func TestYAMLWithComments(t *testing.T) {
	expectedResult := `fieldOne: "field one value" # field one comment
fieldTwo: "field two value" # field two comment
numberField: 14 # number field comment
embeddedObject: # embedded object comment
  embeddedFieldOne: "" # embedded field one
embeddedObjectTwo: # embedded object two comment
  embeddedFieldOne: | # embedded field one
    Multi-line
    indented some amount to be trimmed
embeddedList: # embedded list comment
- embeddedFieldOne: "" # embedded field one
- embeddedFieldOne: "" # embedded field one
- embeddedFieldOne: | # embedded field one
    Multi-line
    text
mapField: # map field comment
  one: "1"
  two: "2"
mapComplexField: # map complex field
  one:
    embeddedFieldOne: "one" # embedded field one
`
	expectedResultAlt := `fieldOne: "field one value" # field one comment
fieldTwo: "field two value" # field two comment
numberField: 14 # number field comment
embeddedObject: # embedded object comment
  embeddedFieldOne: "" # embedded field one
embeddedObjectTwo: # embedded object two comment
  embeddedFieldOne: | # embedded field one
    Multi-line
    indented some amount to be trimmed
embeddedList: # embedded list comment
- embeddedFieldOne: "" # embedded field one
- embeddedFieldOne: "" # embedded field one
- embeddedFieldOne: | # embedded field one
    Multi-line
    text
mapField: # map field comment
  two: "2"
  one: "1"
mapComplexField: # map complex field
  one:
    embeddedFieldOne: "one" # embedded field one
`
	testStructInstance := &testStruct{
		FieldOne:       "field one value",
		FieldTwo:       "field two value",
		NumberField:    14,
		EmbeddedObject: &embeddedStruct{},
		EmbeddedObjectTwo: &embeddedStruct{
			EmbeddedFieldOne: `
    Multi-line
    indented some amount to be trimmed
`,
		},
		EmbeddedList: []*embeddedStruct{
			&embeddedStruct{},
			&embeddedStruct{},
			&embeddedStruct{
				EmbeddedFieldOne: `
Multi-line
text
`,
			},
		},
		MapField: map[string]string{
			"one": "1",
			"two": "2",
		},
		MapComplexField: map[string]*embeddedStruct{
			"one": &embeddedStruct{
				EmbeddedFieldOne: "one",
			},
		},
		MapFieldZeroValue: map[string]string{},
	}
	result, err := YAMLWithComments(testStructInstance, 0)
	if err != nil {
		t.Errorf("Got unexpected error in TestYAMLWithComments: %s", err)
	}
	if r, e, ea := result, expectedResult, expectedResultAlt; r != e && r != ea {
		t.Errorf("Got unexpected result from TestYAMLWithComments:\n%v%v", diff.LineDiff(e, r), diff.LineDiff(ea, r))
	}
}
