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
	errInvalidMessage = errors.New("Invalid Message")
	errEmptyMessage   = errors.New("Message is empty")
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

// GetFieldsToValidate extracts the names of fields for the corresponding fieldmasks.
// If the fieldmask is empty, an empty map is returned which means that nothing will be validated.
func GetFieldsToValidate(i interface{}, paths []string) (map[string]string, error) {
	if len(paths) == 0 {
		return map[string]string{}, nil
	}
	val := reflect.ValueOf(i).Elem()
	if !val.IsValid() || val.Type().NumField() == 0 {
		return map[string]string{}, errInvalidMessage
	}
	topPaths := []string{}
	for _, path := range paths {
		s := strings.Split(path, fieldMaskDelimiter)
		if len(s) != 0 {
			topPaths = append(topPaths, s[0])
		}
	}
	fields := make(map[string]string)
	for i := 0; i < val.Type().NumField(); i++ {
		jsonTag := val.Type().Field(i).Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue
		}
		s := strings.Split(jsonTag, jsonTagDelimiter)
		if len(s) > 2 {
			return map[string]string{}, errInvalidMessage
		}
		// Add a field if it a part of the supplied list.
		for _, st := range topPaths {
			if s[0] == st {
				fields[s[0]] = val.Type().Field(i).Name
				break
			}
		}
	}
	return fields, nil
}

// ShouldBeValidated checks if the given field is a part of the list of fields to be validated.
// This list is created using "GetFieldsToValidate".
func ShouldBeValidated(name string, fields map[string]string) bool {
	for _, fieldName := range fields {
		if name == fieldName {
			return true
		}
	}
	return false
}

// GetTopNameForField retrieves the top field name for the field.
func GetTopNameForField(name string, i interface{}) string {
	if name == "" || i == nil {
		return ""
	}
	names := strings.Split(name, ".")
	if len(names) != 2 {
		return ""
	}
	val := reflect.ValueOf(i).Elem()
	if !val.IsValid() || val.Type().NumField() == 0 {
		return ""
	}
	for i := 0; i < val.Type().NumField(); i++ {
		if name == val.Type().Field(i).Name {
			jsonTag := val.Type().Field(i).Tag.Get("json")
			if jsonTag == "" || jsonTag == "-" {
				return ""
			}
			s := strings.Split(jsonTag, jsonTagDelimiter)
			if len(s) > 2 {
				return ""
			}
			return s[0]
		}
	}
	return ""
}

// getFieldMaskForEmbeddedFields returns a new FieldMask path for fields inside an embedded message.
func getFieldMaskForEmbeddedFields(topLevelMask string, paths []string) []string {
	var subFields strings.Builder
	embeddedFields := []string{}
	for _, path := range paths {
		subFields.Reset()
		if path == "" {
			continue
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

// GetFieldMaskForOneOfFields ...
func GetFieldMaskForOneOfFields(i interface{}, paths []string) ([]string, error) {
	return nil, nil
}
