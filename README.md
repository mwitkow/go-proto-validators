# Golang ProtoBuf Validator Compiler

[![Travis Build](https://travis-ci.org/mwitkow/go-proto-validators.svg)](https://travis-ci.org/mwitkow/go-proto-validators)
[![Apache 2.0 License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)

A `protoc` plugin that generates `Validate() error` functions on Go proto `struct`s based on annotations inside `.proto` 
files. The validation functions are code-generated and thus don't suffer on performance from tag-based reflection on
deeply-nested messages.

## Paint me a code picture

Let's take the following `proto3` snippet:

```proto
syntax = "proto3";
package validator.examples;
import "github.com/mwitkow/go-proto-validators/validator.proto";

message InnerMessage {
  // some_integer can only be in range (1, 100).
  int32 some_integer = 1 [(validator.field) = {int_gt: 0, int_lt: 100}];
}

message OuterMessage {
  // important_string must be a lowercase alpha-numeric of 5 to 30 characters (RE2 syntax).
  string important_string = 1 [(validator.field) = {regex: "^[a-z]{2,5}$"}];
  // proto3 doesn't have `required`, the `msg_exist` enforces presence of InnerMessage.
  InnerMessage inner = 2 [(validator.field) = {msg_exists : true}];
}
```

First, the **`required` keyword is back** for `proto3`, under the guise of `msg_exists`. The painful `if-nil` checks are taken care of!

Second, the expected values in fields are now part of the contract `.proto` file. No more hunting down conditions in code!

Third, the generated code is understandable and has clear understandable error messages. Take a look:

```go
func (this *InnerMessage) Validate() error {
	if !(this.SomeInteger > 0) {
		return fmt.Errorf("validation error: InnerMessage.SomeInteger must be greater than '0'")
	}
	if !(this.SomeInteger < 100) {
		return fmt.Errorf("validation error: InnerMessage.SomeInteger must be less than '100'")
	}
	return nil
}

var _regex_OuterMessage_ImportantString = regexp.MustCompile("^[a-z]{2,5}$")

func (this *OuterMessage) Validate() error {
	if !_regex_OuterMessage_ImportantString.MatchString(this.ImportantString) {
		return fmt.Errorf("validation error: OuterMessage.ImportantString must conform to regex '^[a-z]{2,5}$'")
	}
	if nil == this.Inner {
		return fmt.Errorf("validation error: OuterMessage.Inner message must exist")
	}
	if this.Inner != nil {
		if err := validators.CallValidatorIfExists(this.Inner); err != nil {
			return err
		}
	}
	return nil
}
```

## Installing and using

The `protoc` compiler expects to find plugins named `proto-gen-XYZ` on the execution `$PATH`. So first:

```sh
export PATH=${PATH}:${GOPATH}/bin
```

Then, do the usual

```sh
go install github.com/mwitkow/go-proto-validators/protoc-gen-govalidators
```

TODO(mwitkow): Finish this section

###License

`go-proto-validators` is released under the Apache 2.0 license. See the [LICENSE](LICENSE) file for details.



