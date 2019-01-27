package errors

import (
	"errors"
	"reflect"
	"testing"
)

func TestErrorMethods(t *testing.T) {
	errString := "SomeInt: INT_GT: field must be greater than '500'"

	errType := GetType(errString)
	if errType != "INT_GT" {
		t.Fatal(t)
	}

	errFieldName := GetFieldName(errString)
	if errFieldName != "SomeInt" {
		t.Fatal(t)
	}

	if errFieldName := GetFieldName("errString"); errFieldName != "" {
		t.Fatal(errFieldName)
	}
	if errDesc := GetErrorDescripton("errString:Type"); errDesc != "" {
		t.Fatal(errDesc)
	}

	errDesc := GetErrorDescripton(errString)
	if errDesc != "field must be greater than '500'" {
		t.Fatal(t)
	}

	nilErrType := GetType("")
	if nilErrType != "" {
		t.Fatal(t)
	}

	if err := GetErrorWithTopField("top_field", errors.New("inner_field: INT_GT: field must be greater than '500'")); err.Error() != "top_field.inner_field: INT_GT: field must be greater than '500'" {
		t.Fatal(err)
	}

	if err := FieldError("top_field", Types_INT_GT, errors.New("field must be greater than '500'")); reflect.DeepEqual(err, ValidatorFieldError{nestedErr: errors.New("field must be greater than '500'"), fieldName: "top_field", errType: Types_INT_GT}) {
		t.Fatal(err)
	}

}
