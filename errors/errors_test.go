package errors

import "testing"

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

	errDesc := GetErrorDescripton(errString)
	if errDesc != "field must be greater than '500'" {
		t.Fatal(t)
	}

	nilErrType := GetType("")
	if nilErrType != "" {
		t.Fatal(t)
	}

}
