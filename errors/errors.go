package errors

import (
	"fmt"
	"strings"
)

// ValidatorFieldError is a generic struct that can be used for better error usage in tests and in code.
type ValidatorFieldError struct {
	fieldStack []string
	nestedErr  error
	fieldName  string
	errType    Types
}

// Error returns the error as a string
func (f *ValidatorFieldError) Error() string {
	return fmt.Sprintf("FIELD_ERROR_TYPE_%s: %s: %s", f.errType.String(), strings.Join(f.fieldStack, "."), f.nestedErr.Error())
}

// GetFieldName extracts the field name from the error message.
func GetFieldName(err string) string {
	s := strings.Split(err, ": ")
	if len(s) != 3 {
		return ""
	}
	return s[1]
}

// GetErrorDescripton extracts the error stack from the error message.
func GetErrorDescripton(err string) string {
	s := strings.Split(err, ": ")
	if len(s) != 3 {
		return ""
	}
	return s[2]
}

// GetType extracts the errors.Types name from the error message.
func GetType(err string) string {
	s := strings.Split(err, ": ")
	if len(s) != 3 {
		return ""
	}
	return strings.Replace(s[0], "FIELD_ERROR_TYPE", "Types", 1)
}

// FieldError wraps a given Validator error providing a message call stack.
func FieldError(fieldName string, Type Types, err error) error {
	if fErr, ok := err.(*ValidatorFieldError); ok {
		fErr.fieldStack = append([]string{fieldName}, fErr.fieldStack...)
		fErr.fieldName = fieldName
		fErr.errType = Type
		return err
	}
	return &ValidatorFieldError{
		fieldStack: []string{fieldName},
		nestedErr:  err,
		fieldName:  fieldName,
		errType:    Type,
	}
}
