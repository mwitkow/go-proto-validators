package util

import (
	"fmt"
	"reflect"
	"testing"
)

var allFields = map[string]string{"some_string": "SomeString", "some_int": "SomeInt", "some_double": "SomeDouble", "some_repeated": "SomeRepeated", "some_embedded": "SomeEmbedded"}
var completeFieldMask = []string{"some_string", "some_int", "some_double", "some_repeated", "some_embedded"}
var completeEmbeddedFieldMask = []string{"some_string", "some_int", "some_double", "some_repeated", "some_embedded.ids", "some_embedded.ids.version", "some_embedded.ids.version.timestamp", "some_embedded.value", "some_embedded.name"}
var embeddedFields = []string{"ids", "ids.version", "ids.version.timestamp", "value", "name"}

// Embedded is an embedded message test structure.
type Embedded struct {
	Identifier           string   `protobuf:"bytes,1,opt,name=Identifier,proto3" json:"Identifier,omitempty"`
	SomeValue            int64    `protobuf:"varint,2,opt,name=SomeValue,proto3" json:"SomeValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type TestMessage struct {
	SomeString           string    `protobuf:"bytes,1,opt,name=SomeString,proto3" json:"some_string,omitempty"`
	SomeInt              uint32    `protobuf:"varint,6,opt,name=SomeInt,proto3" json:"some_int,omitempty"`
	SomeDouble           float64   `protobuf:"fixed64,24,opt,name=SomeDouble,proto3" json:"some_double,omitempty"`
	SomeRepeated         []int32   `protobuf:"varint,33,rep,packed,name=SomeRepeated" json:"some_repeated,omitempty"`
	SomeEmbedded         *Embedded `protobuf:"bytes,10,opt,name=SomeEmbedded" json:"some_embedded,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

type BadStruct struct {
	ID string `json:"id,omitempty,repeatable"`
}

type CustomTagStruct struct {
	SomeString string `json:"some_string"`
}

func (m *TestMessage) Validate(fieldMask []string) error {
	return nil
}
func (m *Embedded) Validate(fieldMask []string) error {
	return nil
}

func TestGetFieldsToValidate(t *testing.T) {

	for _, tc := range []struct {
		Name                string
		InputMessage        interface{}
		InputFieldMaskPaths []string
		ExpectedFields      map[string]string
		ExpectedError       interface{}
	}{
		{
			Name:                "NilFieldMask",
			InputMessage:        &TestMessage{},
			InputFieldMaskPaths: nil,
			ExpectedFields:      map[string]string{},
			ExpectedError:       nil,
		},
		{
			Name:                "OneField",
			InputMessage:        &TestMessage{},
			InputFieldMaskPaths: []string{"some_int"},
			ExpectedFields:      map[string]string{"some_int": "SomeInt"},
			ExpectedError:       nil,
		},
		{
			Name:                "InvalidFieldMask",
			InputMessage:        &TestMessage{},
			InputFieldMaskPaths: []string{"somesome_int"},
			ExpectedFields:      map[string]string{},
			ExpectedError:       nil,
		},
		{
			Name:                "FullFieldMask",
			InputMessage:        &TestMessage{},
			InputFieldMaskPaths: completeFieldMask,
			ExpectedFields:      allFields,
			ExpectedError:       nil,
		},
		{
			Name:                "InvalidStructTag",
			InputMessage:        &BadStruct{},
			InputFieldMaskPaths: completeFieldMask,
			ExpectedFields:      map[string]string{},
			ExpectedError:       errInvalidMessage,
		},
		{
			Name:                "CustomStructTag",
			InputMessage:        &CustomTagStruct{},
			InputFieldMaskPaths: completeFieldMask,
			ExpectedFields:      map[string]string{"some_string": "SomeString"},
			ExpectedError:       nil,
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			res, err := GetFieldsToValidate(tc.InputMessage, tc.InputFieldMaskPaths)
			if err != tc.ExpectedError {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(res, tc.ExpectedFields) {
				t.Fatalf("Invalid field array received: %s", res)
			}
		})
	}

}

func TestShouldBeValidated(t *testing.T) {
	for _, tc := range []struct {
		Name           string
		InputField     string
		ValidFields    map[string]string
		ExpectedResult bool
	}{
		{
			Name:           "ShouldBeValidated",
			InputField:     "SomeInt",
			ValidFields:    allFields,
			ExpectedResult: true,
		},
		{
			Name:           "ShouldNotBeValidated",
			InputField:     "SomeInt",
			ValidFields:    map[string]string{"some_string": "SomeString", "some_double": "SomeDouble", "some_repeated": "SomeRepeated", "some_embedded": "SomeEmbedded"},
			ExpectedResult: false,
		},
		{
			Name:           "MalFormedInput",
			InputField:     "this.SomeInt",
			ValidFields:    map[string]string{"some_string": "SomeString", "some_double": "SomeDouble", "some_repeated": "SomeRepeated", "some_embedded": "SomeEmbedded"},
			ExpectedResult: false,
		},
		{
			Name:           "NoFieldsToValidate",
			InputField:     "SomeInt",
			ValidFields:    map[string]string{},
			ExpectedResult: false,
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			res := ShouldBeValidated(tc.InputField, tc.ValidFields)
			if tc.ExpectedResult != res {
				t.Fatal(fmt.Sprintf("Expected %v, received %v", tc.ExpectedResult, res))
			}
		})
	}
}

func TestGetTopNameForField(t *testing.T) {
	for _, tc := range []struct {
		Name           string
		InputTopField  string
		InputMessage   interface{}
		ExpectedResult string
	}{
		{
			Name:           "ValidInputs",
			InputTopField:  "SomeInt",
			InputMessage:   &TestMessage{},
			ExpectedResult: "some_int",
		},
		{
			Name:           "NilMessage",
			InputTopField:  "SomeInt",
			InputMessage:   nil,
			ExpectedResult: "",
		},
		{
			Name:           "MalformedTopField",
			InputTopField:  "this.SomeInt",
			InputMessage:   &TestMessage{},
			ExpectedResult: "",
		},
		{
			Name:           "CustomStruct",
			InputTopField:  "SomeString",
			InputMessage:   &CustomTagStruct{},
			ExpectedResult: "some_string",
		},
		{
			Name:           "BadStruct",
			InputTopField:  "SomeInt",
			InputMessage:   &BadStruct{},
			ExpectedResult: "",
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			res := GetTopNameForField(tc.InputTopField, tc.InputMessage)
			if tc.ExpectedResult != res {
				t.Fatal(fmt.Sprintf("Expected %v, received %v", tc.ExpectedResult, res))
			}
		})
	}

}

func TestGetFieldMaskForEmbeddedFields(t *testing.T) {
	for _, tc := range []struct {
		Name              string
		TopLevelField     string
		InputFieldMask    []string
		ExpectedFieldMask []string
	}{
		{
			Name:              "EmptyTopField",
			TopLevelField:     "",
			InputFieldMask:    completeEmbeddedFieldMask,
			ExpectedFieldMask: []string{},
		},
		{
			Name:              "EmptyFieldMask",
			TopLevelField:     "ids",
			InputFieldMask:    nil,
			ExpectedFieldMask: []string{},
		},
		{
			Name:              "WithSomeEmptyPaths",
			TopLevelField:     "ids",
			InputFieldMask:    []string{"ids.name", "ids.version", "", "up_counter"},
			ExpectedFieldMask: []string{"name", "version"},
		},
		{
			Name:              "WithSomeMalformedFields",
			TopLevelField:     "ids",
			InputFieldMask:    []string{"ids.name", "ids.version", "", "up_counter", "ids.", ".ids", "ids.version.name.", "ids..version.name", "firmware,version"},
			ExpectedFieldMask: []string{"name", "version"},
		},
		{
			Name:              "NoMatch",
			TopLevelField:     "counter",
			InputFieldMask:    completeEmbeddedFieldMask,
			ExpectedFieldMask: []string{},
		},
		{
			Name:              "FullMatch",
			TopLevelField:     "some_embedded",
			InputFieldMask:    completeEmbeddedFieldMask,
			ExpectedFieldMask: embeddedFields,
		},
		{
			Name:              "NoMatch",
			TopLevelField:     "some_embedded",
			InputFieldMask:    []string{"some_embedded"},
			ExpectedFieldMask: []string{},
		},
		{
			Name:              "PartialMatch",
			TopLevelField:     "some_embedded",
			InputFieldMask:    []string{"some_embedded.ids"},
			ExpectedFieldMask: []string{"ids"},
		},
		{
			Name:              "EmptyFM",
			TopLevelField:     "some_embedded",
			InputFieldMask:    []string{},
			ExpectedFieldMask: []string{},
		},
		{
			Name:              "OnlyTopField",
			TopLevelField:     "some_embedded",
			InputFieldMask:    completeEmbeddedFieldMask,
			ExpectedFieldMask: embeddedFields,
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			res := getFieldMaskForEmbeddedFields(tc.TopLevelField, tc.InputFieldMask)
			if !reflect.DeepEqual(res, tc.ExpectedFieldMask) {
				t.Fatalf("Invalid FieldMasks received: %s", res)
			}

		})
	}
}

func TestCallIfValidatorExists(t *testing.T) {
	for _, tc := range []struct {
		Name         string
		Message      interface{}
		TopLevelPath string
		FullPaths    []string
	}{
		{
			Name:         "ValidFields",
			Message:      &Embedded{},
			TopLevelPath: "some_embedded",
			FullPaths:    completeEmbeddedFieldMask,
		},
		{
			Name:         "NoTopLevelPath",
			Message:      &Embedded{},
			TopLevelPath: "",
			FullPaths:    []string{},
		},
		{
			Name:         "NoFieldMask",
			Message:      &Embedded{},
			TopLevelPath: "some_embedded",
			FullPaths:    []string{},
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			err := CallValidatorIfExists(tc.Message, tc.TopLevelPath, tc.FullPaths)
			if err != nil {
				t.Fatal("Validator not called")
			}
		})
	}
}
