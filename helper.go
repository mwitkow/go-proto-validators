// Copyright 2016 Michal Witkowski. All Rights Reserved.
// See LICENSE for licensing terms.

package validator

import (
	"strings"
)

// Validator is a general interface that allows a message to be validated.
type Validator interface {
	Validate() *MultiError
}

func CallValidatorIfExists(candidate interface{}) *MultiError {
	if validator, ok := candidate.(Validator); ok {
		return validator.Validate()
	}
	return nil
}

type fieldError struct {
	fieldStack []string
	nestedErr  error
}
type ErrorDetail struct {
	Reason   string
	Domain   string
	Metadata map[string]string
}

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
