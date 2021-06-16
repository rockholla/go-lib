package marshal

import (
	"testing"
)

type testStruct struct {
	FieldOne       string            `yaml:"fieldOne" help:"field one help"`
	FieldTwo       string            `yaml:"fieldTwo" help:"field two help"`
	EmbeddedObject *embeddedStruct   `yaml:"embeddedObject" help:"embedded object help"`
	EmbeddedList   []*embeddedStruct `yaml:"embeddedList" help:"embedded list help"`
}

type embeddedStruct struct {
	EmbeddedFieldOne string `yaml:"embeddedFieldOne" help:"embedded field one"`
}

func TestYAMLWithComments(t *testing.T) {
	expectedResult := `fieldOne: "" # field one help
fieldTwo: "" # field two help
embeddedObject: # embedded object help
  embeddedFieldOne: "" # embedded field one
embeddedList: # embedded list help
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
