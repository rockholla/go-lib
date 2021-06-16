package marshal

import (
	"testing"
)

type testStruct struct {
	FieldOne              string            `yaml:"fieldOne" comment:"field one comment"`
	FieldTwo              string            `yaml:"fieldTwo" comment:"field two comment"`
	EmbeddedObject        *embeddedStruct   `yaml:"embeddedObject" comment:"embedded object comment"`
	EmbeddedList          []*embeddedStruct `yaml:"embeddedList" comment:"embedded list comment"`
	FieldOmittedZeroValue string            `yaml:"fieldOmittedZeroValue,omitempty"`
	FieldOmittedNoMarshal string            `yaml:"-"`
}

type embeddedStruct struct {
	EmbeddedFieldOne string `yaml:"embeddedFieldOne" comment:"embedded field one"`
}

func TestYAMLWithComments(t *testing.T) {
	expectedResult := `fieldOne: "" # field one comment
fieldTwo: "" # field two comment
embeddedObject: # embedded object comment
  embeddedFieldOne: "" # embedded field one
embeddedList: # embedded list comment
- embeddedFieldOne: "" # embedded field one
- embeddedFieldOne: "" # embedded field one
- # embedded field one
   embeddedFieldOne: |
    Multi-line
    text
`
	testStructInstance := &testStruct{
		EmbeddedObject: &embeddedStruct{},
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
	}
	result, err := YAMLWithComments(testStructInstance, 0)
	if err != nil {
		t.Errorf("Got unexpected error in TestYAMLWithComments: %s", err)
	}
	if result != expectedResult {
		t.Errorf("Got unexpected result from TestYAMLWithComments:\n%s", result)
	}
}
