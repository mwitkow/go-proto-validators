// Copyright 2016 Michal Witkowski. All Rights Reserved.
// See LICENSE for licensing terms.

package validator

type Validator interface {
	Validate() error
}

func CallValidatorIfExists(candidate interface{}) error {
	if validator, ok := candidate.(Validator); ok {
		return validator.Validate()
	}
	return nil
}
