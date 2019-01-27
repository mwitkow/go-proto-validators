package errors

import (
	"fmt"
	"strings"
)

// ValidatorFieldError is a generic struct that can be used for better error usage in tests and in code.
type ValidatorFieldError struct {
	nestedErr error
	fieldName string
	errType   Types
}

// Error returns the error as a string
func (f *ValidatorFieldError) Error() string {
	return fmt.Sprintf("%s: %s: %s", f.fieldName, f.errType.String(), f.nestedErr.Error())
}

// GetFieldName extracts the field name from the error message.
func GetFieldName(err string) string {
	s := strings.Split(err, ": ")
	if len(s) != 3 {
		return ""
	}
	return s[0]
}

// GetType extracts the errors.Types name from the error message.
func GetType(err string) string {
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

// GetErrorWithTopField ...
func GetErrorWithTopField(name string, err error) error {
	return fmt.Errorf(fmt.Sprintf("%s.%s", name, err.Error()))
}

// FieldError wraps a given Validator error providing a message call stack.
func FieldError(fieldName string, Type Types, err error) error {
	return &ValidatorFieldError{
		nestedErr: err,
		fieldName: fieldName,
		errType:   Type,
	}
}
