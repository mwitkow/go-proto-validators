# Golang ProtoBuf Validator Compiler

[![Travis Build](https://travis-ci.org/mwitkow/go-proto-validators.svg)](https://travis-ci.org/mwitkow/go-proto-validators)
[![Apache 2.0 License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)

A `protoc` plugin that generates `Validate([]paths) error` functions on Go proto `struct`s based on field options inside `.proto` 
files.

It incorporates [fieldmasks](https://developers.google.com/protocol-buffers/docs/reference/csharp/class/google/protobuf/well-known-types/field-mask) into the validators thereby providing optional validation on fields. By default, no fields are validated. This allows for each field that needs to be validated  to be specified by the fieldmask.

## Paint me a code picture

Let's take the following `proto3` snippet:

```proto
syntax = "proto3";
package validator.examples;
import "github.com/mwitkow/go-proto-validators/validator.proto";

message InnerMessage {
  // some_integer can only be in range (1, 100).
  int32 some_integer = 1 [(validator.field) = {int_gt: 0, int_lt: 100}];
  // some_float can only be in range (0;1).
  double some_float = 2 [(validator.field) = {float_gte: 0, float_lte: 1}];
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
func (this *InnerMessage) Validate(paths []string) error {
	toBeValidated, err := github_com_TheThingsIndustries_go_proto_validators_util.GetFieldsToValidate(this, paths)
	if err != nil {
		return err
	}
	_ = toBeValidated

	if !(this.SomeInteger > 0) && (github_com_TheThingsIndustries_go_proto_validators_util.ShouldBeValidated("this.SomeInteger", toBeValidated)) {
		return github_com_TheThingsIndustries_go_proto_validators_errors.FieldError(github_com_TheThingsIndustries_go_proto_validators_util.GetProtoNameForField("SomeInteger", toBeValidated), github_com_TheThingsIndustries_go_proto_validators_errors.Types_INT_GT, fmt.Errorf(`field must be greater than '0'`))
	}
	if !(this.SomeInteger < 100) && (github_com_TheThingsIndustries_go_proto_validators_util.ShouldBeValidated("this.SomeInteger", toBeValidated)) {
		return github_com_TheThingsIndustries_go_proto_validators_errors.FieldError(github_com_TheThingsIndustries_go_proto_validators_util.GetProtoNameForField("SomeInteger", toBeValidated), github_com_TheThingsIndustries_go_proto_validators_errors.Types_INT_LT, fmt.Errorf(`field must be lesser than '100'`))
	}
	return nil
}

var _regex_OuterMessage_ImportantString = regexp.MustCompile(`^[a-z]{2,5}$`)

func (this *OuterMessage) Validate(paths []string) error {
	toBeValidated, err := github_com_TheThingsIndustries_go_proto_validators_util.GetFieldsToValidate(this, paths)
	if err != nil {
		return err
	}
	_ = toBeValidated

	if !_regex_OuterMessage_ImportantString.MatchString(this.ImportantString) && (github_com_TheThingsIndustries_go_proto_validators_util.ShouldBeValidated("this.ImportantString", toBeValidated)) {
		return github_com_TheThingsIndustries_go_proto_validators_errors.FieldError(github_com_TheThingsIndustries_go_proto_validators_util.GetProtoNameForField("ImportantString", toBeValidated), github_com_TheThingsIndustries_go_proto_validators_errors.Types_STRING_REGEX, fmt.Errorf(`field must be a string conforming to the regex "^[a-z]{2,5}$"`))
	}
	if nil == this.Inner {
		return github_com_TheThingsIndustries_go_proto_validators_errors.FieldError(github_com_TheThingsIndustries_go_proto_validators_util.GetProtoNameForField("Inner", toBeValidated), github_com_TheThingsIndustries_go_proto_validators_errors.Types_MSG_EXISTS, fmt.Errorf("message must exist"))
	}
	if (this.Inner != nil) && (github_com_TheThingsIndustries_go_proto_validators_util.ShouldBeValidated("this.Inner", toBeValidated)) {
		if err := github_com_TheThingsIndustries_go_proto_validators_util.CallValidatorIfExists(this.Inner, github_com_TheThingsIndustries_go_proto_validators_util.GetProtoNameForField("this.Inner", toBeValidated), paths); err != nil {
			return github_com_TheThingsIndustries_go_proto_validators_errors.GetErrorWithTopField(github_com_TheThingsIndustries_go_proto_validators_util.GetProtoNameForField("this.Inner", toBeValidated), err)
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
go get github.com/mwitkow/go-proto-validators/protoc-gen-govalidators
```

Check the [Makefile](/Makefile) for installing this plugin.

The following is an example of using this plugin as part of your proto generation.

```sh
protoc  \
	--proto_path=${GOPATH}/src \
	--proto_path=${GOPATH}/src/github.com/gogo/protobuf/protobuf \
	--proto_path=. \
	--gogo_out=. \
	--govalidators_out=gogoimport=true:. \
	*.proto
```

Basically the magical incantation (apart from includes) is the `--govalidators_out`. That triggers the 
`protoc-gen-govalidators` plugin to generate `mymessage.validator.pb.go`. That's it :)

###License

`go-proto-validators` is released under the Apache 2.0 license. See the [LICENSE](LICENSE) file for details.



