package util

import (
	"errors"
	"reflect"
	"strings"
)

const (
	fieldMaskDelimiter string = "."
	jsonTagDelimiter   string = ","
)

var (
	errInvalidMessage error = errors.New("Invalid Message")
	errEmptyMessage   error = errors.New("Message is empty")
)

// Validator is a general interface that allows a message to be validated.
type Validator interface {
	Validate([]string) error
}

// CallValidatorIfExists is used to call the validator for the embedded message if it exists.
// It generates the fieldmask for the sub-field before calling it.
// The conditional field shouldBeCalled is used to prevent this function from calling the sub validator based on the parent fieldmask.
func CallValidatorIfExists(candidate interface{}, topLevelPath string, fullPaths []string) error {
	if validator, ok := candidate.(Validator); ok {
		return validator.Validate(getFieldMaskForEmbeddedFields(topLevelPath, fullPaths))
	}
	return nil
}

type fieldError struct {
	fieldStack []string
	nestedErr  error
}

// Error returns the error as a string
func (f *fieldError) Error() string {
	return "invalid field " + strings.Join(f.fieldStack, ".") + ": " + f.nestedErr.Error()
}

// FieldError wraps a given Validator error providing a message call stack.
func FieldError(fieldName string, err error) error {
	if fErr, ok := err.(*fieldError); ok {
		fErr.fieldStack = append([]string{fieldName}, fErr.fieldStack...)
		return err
	}
	return &fieldError{
		fieldStack: []string{fieldName},
		nestedErr:  err,
	}
}

// GetFieldsToValidate extracts the names of fields for the corresponding fieldmasks.
// If the fieldmask is empty, all the fields are returned.
// This works only on protobuf generated structs or structs that have JSON tags in the format "fieldname,omitempty".
func GetFieldsToValidate(i interface{}, paths []string) ([]string, error) {
	val := reflect.ValueOf(i).Elem()
	if !val.IsValid() || val.Type().NumField() == 0 {
		return []string{}, errInvalidMessage
	}
	fields := []string{}
	for i := 0; i < val.Type().NumField(); i++ {
		jsonTag := val.Type().Field(i).Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue
		}
		// Split the tag since protobuf json tags are in the format "fieldname,omitempty".
		s := strings.Split(jsonTag, jsonTagDelimiter)
		if len(s) != 2 {
			return []string{}, errInvalidMessage
		}
		// Add all fields to the list to be validated if no fieldmask paths are specified.
		if len(paths) == 0 {
			fields = append(fields, val.Type().Field(i).Name)
			continue
		}
		// Add a field if it a part of the supplied list.
		for _, st := range paths {
			if s[0] == st {
				fields = append(fields, val.Type().Field(i).Name)
				break
			}
		}
	}
	return fields, nil
}

// ShouldBeValidated checks if the given field is a part of the list of fields to be validated.
// This list is created using "GetFieldsToValidate".
func ShouldBeValidated(name string, fieldNames []string) bool {
	// The name as passed by the generator would be in the format this.FieldName
	s := strings.Split(name, fieldMaskDelimiter)
	if len(s) != 2 {
		// When it's malformed, validate anyway since it's difficult validate errors here.
		return true
	}
	for _, fieldName := range fieldNames {
		if s[1] == fieldName {
			return true
		}
	}
	return false
}

// getFieldMaskForEmbeddedFields returns a new FieldMask path for fields inside an embedded message.
func getFieldMaskForEmbeddedFields(topLevelMask string, paths []string) []string {
	var subFields strings.Builder
	embeddedFields := []string{}
	for _, path := range paths {
		subFields.Reset()
		if path == "" {
			continue //Sanity check for empty paths
		}
		s := strings.Split(path, fieldMaskDelimiter)
		if len(s) < 2 || s[0] != topLevelMask {
			continue
		}
		// Join the rest of the sub-fields back into a single string.
		for i := 1; i < len(s); i++ {
			if s[i] == "" {
				// If empty strings are encountered, exit loop and invalidate the entire string
				subFields.Reset()
				break
			}
			if subFields.String() != "" {
				// Add the dot for all subsequent characters
				subFields.WriteString(fieldMaskDelimiter)
			}
			subFields.WriteString(s[i])
		}
		if subFields.String() != "" {
			embeddedFields = append(embeddedFields, subFields.String())
		}
	}
	return embeddedFields
}

func GetFieldMaskForRepeatedFields(i interface{}, paths []string) ([]string, error) {
	return nil, nil
}

func GetFieldMaskForOneOfFields(i interface{}, paths []string) ([]string, error) {
	return nil, nil
}
