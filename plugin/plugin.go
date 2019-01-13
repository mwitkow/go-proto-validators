// Copyright 2016 Michal Witkowski. All Rights Reserved.
// See LICENSE for licensing terms.

/*

The validator plugin generates a Validate method for each message.
By default, if none of the message's fields are annotated with the gogo validator annotation, it returns a nil.
In case some of the fields are annotated, the Validate function returns nil upon sucessful validation, or an error
describing why the validation failed.
The Validate method is called recursively for all submessage of the message.

TODO(michal): ADD COMMENTS.

Equal is enabled using the following extensions:

  - equal
  - equal_all

While VerboseEqual is enable dusing the following extensions:

  - verbose_equal
  - verbose_equal_all

The equal plugin also generates a test given it is enabled using one of the following extensions:

  - testgen
  - testgen_all

Let us look at:

  github.com/gogo/protobuf/test/example/example.proto

Btw all the output can be seen at:

  github.com/gogo/protobuf/test/example/*

The following message:



given to the equal plugin, will generate the following code:



and the following test code:


*/
package plugin

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/TheThingsIndustries/go-proto-validators"
	"github.com/TheThingsIndustries/go-proto-validators/errors"
	"github.com/gogo/protobuf/gogoproto"
	"github.com/gogo/protobuf/proto"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"github.com/gogo/protobuf/vanity"
)

type plugin struct {
	*generator.Generator
	generator.PluginImports
	regexPkg      generator.Single
	fmtPkg        generator.Single
	protoPkg      generator.Single
	validatorPkg  generator.Single
	errorPkg      generator.Single
	useGogoImport bool
}

// NewPlugin ...
func NewPlugin(useGogoImport bool) generator.Plugin {
	return &plugin{useGogoImport: useGogoImport}
}

func (p *plugin) Name() string {
	return "validator"
}

func (p *plugin) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *plugin) Generate(file *generator.FileDescriptor) {
	if !p.useGogoImport {
		vanity.TurnOffGogoImport(file.FileDescriptorProto)
	}
	p.PluginImports = generator.NewPluginImports(p.Generator)
	p.regexPkg = p.NewImport("regexp")
	p.fmtPkg = p.NewImport("fmt")
	p.validatorPkg = p.NewImport("github.com/TheThingsIndustries/go-proto-validators/util")
	p.errorPkg = p.NewImport("github.com/TheThingsIndustries/go-proto-validators/errors")

	for _, msg := range file.Messages() {
		if msg.DescriptorProto.GetOptions().GetMapEntry() {
			continue
		}
		p.generateRegexVars(file, msg)
		if gogoproto.IsProto3(file.FileDescriptorProto) {
			p.generateProto3Message(file, msg)
		}
	}
}

func getFieldValidatorIfAny(field *descriptor.FieldDescriptorProto) *validator.FieldValidator {
	if field.Options != nil {
		v, err := proto.GetExtension(field.Options, validator.E_Field)
		if err == nil && v.(*validator.FieldValidator) != nil {
			return (v.(*validator.FieldValidator))
		}
	}
	return nil
}

func (p *plugin) isSupportedInt(field *descriptor.FieldDescriptorProto) bool {
	switch *(field.Type) {
	case descriptor.FieldDescriptorProto_TYPE_INT32, descriptor.FieldDescriptorProto_TYPE_INT64:
		return true
	case descriptor.FieldDescriptorProto_TYPE_UINT32, descriptor.FieldDescriptorProto_TYPE_UINT64:
		return true
	case descriptor.FieldDescriptorProto_TYPE_SINT32, descriptor.FieldDescriptorProto_TYPE_SINT64:
		return true
	}
	return false
}

func (p *plugin) isSupportedFloat(field *descriptor.FieldDescriptorProto) bool {
	switch *(field.Type) {
	case descriptor.FieldDescriptorProto_TYPE_FLOAT, descriptor.FieldDescriptorProto_TYPE_DOUBLE:
		return true
	case descriptor.FieldDescriptorProto_TYPE_FIXED32, descriptor.FieldDescriptorProto_TYPE_FIXED64:
		return true
	case descriptor.FieldDescriptorProto_TYPE_SFIXED32, descriptor.FieldDescriptorProto_TYPE_SFIXED64:
		return true
	}
	return false
}

func (p *plugin) generateRegexVars(file *generator.FileDescriptor, message *generator.Descriptor) {
	ccTypeName := generator.CamelCaseSlice(message.TypeName())
	for _, field := range message.Field {
		validator := getFieldValidatorIfAny(field)
		if validator != nil && validator.Regex != nil {
			fieldName := p.GetFieldName(message, field)
			p.P(`var `, p.regexName(ccTypeName, fieldName), ` = `, p.regexPkg.Use(), `.MustCompile(`, "`", *validator.Regex, "`", `)`)
		}
	}
}

func (p *plugin) generateProto3Message(file *generator.FileDescriptor, message *generator.Descriptor) {
	ccTypeName := generator.CamelCaseSlice(message.TypeName())
	p.P(`func (this *`, ccTypeName, `) Validate(paths []string) error {`)
	p.In()
	p.P(`toBeValidated, err := `, p.validatorPkg.Use(), `.GetFieldsToValidate(this, paths)`)
	p.P(`if err != nil {`)
	p.In()
	p.P(`return err`)
	p.Out()
	p.P(`}`)
	p.P("\n")

	for _, field := range message.Field {
		fieldValidator := getFieldValidatorIfAny(field)
		if fieldValidator == nil && !field.IsMessage() {
			continue
		}
		isOneOf := field.OneofIndex != nil
		fieldName := p.GetOneOfFieldName(message, field)
		variableName := "this." + fieldName
		repeated := field.IsRepeated()
		// Golang's proto3 has no concept of unset primitive fields
		nullable := (gogoproto.IsNullable(field) || !gogoproto.ImportsGoGoProto(file.FileDescriptorProto)) && field.IsMessage()
		if p.fieldIsProto3Map(file, message, field) {
			p.P(`// Validation of proto3 map<> fields is unsupported.`)
			continue
		}
		if isOneOf {
			p.In()
			oneOfName := p.GetFieldName(message, field)
			oneOfType := p.OneOfTypeName(message, field)
			//if x, ok := m.GetType().(*OneOfMessage3_OneInt); ok {
			p.P(`if oneOfNester, ok := this.Get` + oneOfName + `().(* ` + oneOfType + `); ok {`)
			variableName = "oneOfNester." + p.GetOneOfFieldName(message, field)
		}
		if repeated {
			p.generateRepeatedCountValidator(variableName, ccTypeName, fieldName, fieldValidator)
			if field.IsMessage() || p.validatorWithNonRepeatedConstraint(fieldValidator) {
				p.P(`for _, item := range `, variableName, `{`)
				p.In()
			}
		} else if fieldValidator != nil {
			if fieldValidator.RepeatedCountMin != nil {
				fmt.Fprintf(os.Stderr, "WARNING: field %v.%v is not repeated, validator.min_elts has no effects\n", ccTypeName, fieldName)
			}
			if fieldValidator.RepeatedCountMax != nil {
				fmt.Fprintf(os.Stderr, "WARNING: field %v.%v is not repeated, validator.max_elts has no effects\n", ccTypeName, fieldName)
			}
		}
		if field.IsString() {
			p.generateStringValidator(variableName, ccTypeName, fieldName, fieldValidator)
		} else if p.isSupportedInt(field) {
			p.generateIntValidator(variableName, ccTypeName, fieldName, fieldValidator)
		} else if p.isSupportedFloat(field) {
			p.generateFloatValidator(variableName, ccTypeName, fieldName, fieldValidator)
		} else if field.IsBytes() {
			p.generateLengthValidator(variableName, ccTypeName, fieldName, fieldValidator)
		} else if field.IsMessage() {
			if p.validatorWithMessageExists(fieldValidator) {
				if nullable && !repeated {
					p.P(`if nil == `, variableName, `{`)
					p.In()
					p.P(`return `, p.errorPkg.Use(), `.FieldError("`, fieldName, `",`, p.errorPkg.Use(), `.Types_`, errors.Types_MSG_EXISTS.String(), `, `, p.fmtPkg.Use(), `.Errorf("message must exist"))`)
					p.Out()
					p.P(`}`)
				} else if repeated {
					fmt.Fprintf(os.Stderr, "WARNING: field %v.%v is repeated, validator.msg_exists has no effect\n", ccTypeName, fieldName)
				} else if !nullable {
					fmt.Fprintf(os.Stderr, "WARNING: field %v.%v is a nullable=false, validator.msg_exists has no effect\n", ccTypeName, fieldName)
				}
			}
			if nullable {
				p.P(`if (`, variableName, ` != nil) && (`, p.validatorPkg.Use(), `.ShouldBeValidated("`, variableName, `", toBeValidated)){`)
				p.In()
			} else {
				// Non-nullable fields are created using [(gogoproto.nullable) = false].
				// These fields in proto3 store actual structs, we need pointers to operate on interfaces.
				// For these messages, the fieldmask will be ignored since they must exist.
				// So the validator of the message will be called regardless of the fieldmask.
				variableName = "&(" + variableName + ")"
			}

			if repeated {
				p.P(`if err := `, p.validatorPkg.Use(), `.CallValidatorIfExists(item,`, p.validatorPkg.Use(), `.GetTopNameForField("`, variableName, `", this), paths ); err != nil {`)
			} else {
				p.P(`if err := `, p.validatorPkg.Use(), `.CallValidatorIfExists(`, variableName, `,`, p.validatorPkg.Use(), `.GetTopNameForField("`, variableName, `", this), paths ); err != nil {`)
			}

			p.In()
			p.P(`return err`)
			p.Out()
			p.P(`}`)
			if nullable {
				p.Out()
				p.P(`}`)
			}
		}
		if repeated && (field.IsMessage() || p.validatorWithNonRepeatedConstraint(fieldValidator)) {
			// end the repeated loop
			p.Out()
			p.P(`}`)
		}
		if isOneOf {
			// end the oneof if statement
			p.Out()
			p.P(`}`)
		}
	}
	p.P(`return nil`)
	p.Out()
	p.P(`}`)
}

func (p *plugin) generateIntValidator(variableName string, ccTypeName string, fieldName string, fv *validator.FieldValidator) {
	if fv.IntGt != nil {
		p.P(`if !(`, variableName, ` > `, fv.IntGt, `) && (`, p.validatorPkg.Use(), `.ShouldBeValidated("`, variableName, `", toBeValidated)) {`)
		p.In()
		errorStr := fmt.Sprintf(`be greater than '%d'`, fv.GetIntGt())
		p.generateErrorString(variableName, errors.Types_INT_GT, fieldName, errorStr, fv)
		p.Out()
		p.P(`}`)
	}
	if fv.IntLt != nil {
		p.P(`if !(`, variableName, ` < `, fv.IntLt, `) && (`, p.validatorPkg.Use(), `.ShouldBeValidated("`, variableName, `", toBeValidated)) {`)
		p.In()
		errorStr := fmt.Sprintf(`be lesser than '%d'`, fv.GetIntLt())
		p.generateErrorString(variableName, errors.Types_INT_LT, fieldName, errorStr, fv)
		p.Out()
		p.P(`}`)
	}
}

func (p *plugin) generateLengthValidator(variableName string, ccTypeName string, fieldName string, fv *validator.FieldValidator) {
	if fv.LengthGt != nil {
		p.P(`if !( len(`, variableName, `) > `, fv.LengthGt, `) && (`, p.validatorPkg.Use(), `.ShouldBeValidated("`, variableName, `", toBeValidated)) {`)
		p.In()
		errorStr := fmt.Sprintf(` be longer than '%d' elements`, fv.GetLengthGt())
		p.generateErrorString(variableName, errors.Types_LENGTH_GT, fieldName, errorStr, fv)
		p.Out()
		p.P(`}`)
	}

	if fv.LengthLt != nil {
		p.P(`if !( len(`, variableName, `) < `, fv.LengthLt, `) && (`, p.validatorPkg.Use(), `.ShouldBeValidated("`, variableName, `", toBeValidated)) {`)
		p.In()
		errorStr := fmt.Sprintf(` be shorter than '%d' elements`, fv.GetLengthLt())
		p.generateErrorString(variableName, errors.Types_LENGTH_LT, fieldName, errorStr, fv)
		p.Out()
		p.P(`}`)
	}

	if fv.LengthEq != nil {
		p.P(`if !( len(`, variableName, `) == `, fv.LengthEq, `) && (`, p.validatorPkg.Use(), `.ShouldBeValidated("`, variableName, `", toBeValidated)) {`)
		p.In()
		errorStr := fmt.Sprintf(` have length equal to '%d' elements`, fv.GetLengthEq())
		p.generateErrorString(variableName, errors.Types_LENGTH_EQ, fieldName, errorStr, fv)
		p.Out()
		p.P(`}`)
	}

}

func (p *plugin) generateFloatValidator(variableName string, ccTypeName string, fieldName string, fv *validator.FieldValidator) {
	upperIsStrict := true
	lowerIsStrict := true

	// First check for incompatible constraints (i.e flt_lt & flt_lte both defined, etc) and determine the real limits.
	if fv.FloatEpsilon != nil && fv.FloatLt == nil && fv.FloatGt == nil {
		fmt.Fprintf(os.Stderr, "WARNING: field %v.%v has no 'float_lt' or 'float_gt' field so setting 'float_epsilon' has no effect.", ccTypeName, fieldName)
	}
	if fv.FloatLt != nil && fv.FloatLte != nil {
		fmt.Fprintf(os.Stderr, "WARNING: field %v.%v has both 'float_lt' and 'float_lte' constraints, only the strictest will be used.", ccTypeName, fieldName)
		strictLimit := fv.GetFloatLt()
		if fv.FloatEpsilon != nil {
			strictLimit += fv.GetFloatEpsilon()
		}

		if fv.GetFloatLte() < strictLimit {
			upperIsStrict = false
		}
	} else if fv.FloatLte != nil {
		upperIsStrict = false
	}

	if fv.FloatGt != nil && fv.FloatGte != nil {
		fmt.Fprintf(os.Stderr, "WARNING: field %v.%v has both 'float_gt' and 'float_gte' constraints, only the strictest will be used.", ccTypeName, fieldName)
		strictLimit := fv.GetFloatGt()
		if fv.FloatEpsilon != nil {
			strictLimit -= fv.GetFloatEpsilon()
		}

		if fv.GetFloatGte() > strictLimit {
			lowerIsStrict = false
		}
	} else if fv.FloatGte != nil {
		lowerIsStrict = false
	}

	// Generate the constraint checking code.
	errorStr := ""
	compareStr := ""
	var errType errors.Types
	if fv.FloatGt != nil || fv.FloatGte != nil {
		compareStr = fmt.Sprint(`if !(`, variableName)
		if lowerIsStrict {
			errType = errors.Types_FLOAT_GT
			errorStr = fmt.Sprintf(`be strictly greater than '%g'`, fv.GetFloatGt())
			if fv.FloatEpsilon != nil {
				errType = errors.Types_FLOAT_ELIPSON
				errorStr += fmt.Sprintf(` with a tolerance of '%g'`, fv.GetFloatEpsilon())
				compareStr += fmt.Sprint(` + `, fv.GetFloatEpsilon())
			}
			compareStr += fmt.Sprint(` > `, fv.GetFloatGt(), `) {`)
		} else {
			errType = errors.Types_FLOAT_GTE
			errorStr = fmt.Sprintf(`be greater than or equal to '%g'`, fv.GetFloatGte())
			compareStr += fmt.Sprint(` >= `, fv.GetFloatGte(), `) {`)
		}
		p.P(compareStr)
		p.In()
		p.generateErrorString(variableName, errType, fieldName, errorStr, fv)
		p.Out()
		p.P(`}`)
	}

	if fv.FloatLt != nil || fv.FloatLte != nil {
		compareStr = fmt.Sprint(`if !(`, variableName)
		if upperIsStrict {
			errType = errors.Types_FLOAT_LT
			errorStr = fmt.Sprintf(`be strictly lower than '%g'`, fv.GetFloatLt())
			if fv.FloatEpsilon != nil {
				errType = errors.Types_FLOAT_ELIPSON
				errorStr += fmt.Sprintf(` with a tolerance of '%g'`, fv.GetFloatEpsilon())
				compareStr += fmt.Sprint(` - `, fv.GetFloatEpsilon())
			}
			compareStr += fmt.Sprint(` < `, fv.GetFloatLt(), `) {`)
		} else {
			errType = errors.Types_FLOAT_LTE
			errorStr = fmt.Sprintf(`be lower than or equal to '%g'`, fv.GetFloatLte())
			compareStr += fmt.Sprint(` <= `, fv.GetFloatLte(), `) {`)
		}
		p.P(compareStr)
		p.In()
		p.generateErrorString(variableName, errType, fieldName, errorStr, fv)
		p.Out()
		p.P(`}`)
	}
}

func (p *plugin) generateStringValidator(variableName string, ccTypeName string, fieldName string, fv *validator.FieldValidator) {
	if fv.Regex != nil {
		p.P(`if !`, p.regexName(ccTypeName, fieldName), `.MatchString(`, variableName, `) && (`, p.validatorPkg.Use(), `.ShouldBeValidated("`, variableName, `", toBeValidated)) {`)
		p.In()
		errorStr := "be a string conforming to the regex " + strconv.Quote(fv.GetRegex())
		p.generateErrorString(variableName, errors.Types_STRING_REGEX, fieldName, errorStr, fv)
		p.Out()
		p.P(`}`)
	}
	if fv.StringNotEmpty != nil && fv.GetStringNotEmpty() {
		p.P(`if (`, variableName, ` == "" ) && (`, p.validatorPkg.Use(), `.ShouldBeValidated("`, variableName, `", toBeValidated)) {`)
		p.In()
		errorStr := "not be an empty string"
		p.generateErrorString(variableName, errors.Types_STRING_NOT_EMPTY, fieldName, errorStr, fv)
		p.Out()
		p.P(`}`)
	}
	p.generateLengthValidator(variableName, ccTypeName, fieldName, fv)

}

func (p *plugin) generateRepeatedCountValidator(variableName string, ccTypeName string, fieldName string, fv *validator.FieldValidator) {
	if fv == nil {
		return
	}
	if fv.RepeatedCountMin != nil {
		compareStr := fmt.Sprint(`if len(`, variableName, `) < `, fv.GetRepeatedCountMin(), ` {`)
		p.P(compareStr)
		p.In()
		errorStr := fmt.Sprint(`contain at least `, fv.GetRepeatedCountMin(), ` elements`)
		p.generateErrorString(variableName, errors.Types_REPEATED_COUNT_MIN, fieldName, errorStr, fv)
		p.Out()
		p.P(`}`)
	}
	if fv.RepeatedCountMax != nil {
		compareStr := fmt.Sprint(`if len(`, variableName, `) > `, fv.GetRepeatedCountMax(), ` {`)
		p.P(compareStr)
		p.In()
		errorStr := fmt.Sprint(`contain at most `, fv.GetRepeatedCountMax(), ` elements`)
		p.generateErrorString(variableName, errors.Types_REPEATED_COUNT_MAX, fieldName, errorStr, fv)
		p.Out()
		p.P(`}`)
	}
}

func (p *plugin) generateErrorString(variableName string, errType errors.Types, fieldName string, specificError string, fv *validator.FieldValidator) {
	if fv.GetHumanError() == "" {
		p.P(`return `, p.errorPkg.Use(), `.FieldError("`, fieldName, `",`, p.errorPkg.Use(), `.Types_`, errType.String(), `, `, p.fmtPkg.Use(), ".Errorf( `field must ", specificError, "`", `))`)
	} else {
		p.P(`return `, p.errorPkg.Use(), `.FieldError("`, fieldName, `",`, p.fmtPkg.Use(), ".Errorf(`", fv.GetHumanError(), "`))")
	}
}

func (p *plugin) fieldIsProto3Map(file *generator.FileDescriptor, message *generator.Descriptor, field *descriptor.FieldDescriptorProto) bool {
	// Context from descriptor.proto
	// Whether the message is an automatically generated map entry type for the
	// maps field.
	//
	// For maps fields:
	//     map<KeyType, ValueType> map_field = 1;
	// The parsed descriptor looks like:
	//     message MapFieldEntry {
	//         option map_entry = true;
	//         optional KeyType key = 1;
	//         optional ValueType value = 2;
	//     }
	//     repeated MapFieldEntry map_field = 1;
	//
	// Implementations may choose not to generate the map_entry=true message, but
	// use a native map in the target language to hold the keys and values.
	// The reflection APIs in such implementions still need to work as
	// if the field is a repeated message field.
	//
	// NOTE: Do not set the option in .proto files. Always use the maps syntax
	// instead. The option should only be implicitly set by the proto compiler
	// parser.
	if field.GetType() != descriptor.FieldDescriptorProto_TYPE_MESSAGE || !field.IsRepeated() {
		return false
	}
	typeName := field.GetTypeName()
	var msg *descriptor.DescriptorProto
	if strings.HasPrefix(typeName, ".") {
		// Fully qualified case, look up in global map, must work or fail badly.
		msg = p.ObjectNamed(field.GetTypeName()).(*generator.Descriptor).DescriptorProto
	} else {
		// Nested, relative case.
		msg = file.GetNestedMessage(message.DescriptorProto, field.GetTypeName())
	}
	return msg.GetOptions().GetMapEntry()
}

func (p *plugin) validatorWithAnyConstraint(fv *validator.FieldValidator) bool {
	if fv == nil {
		return false
	}

	// Need to use reflection in order to be future-proof for new types of constraints.
	v := reflect.ValueOf(fv)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() != nil {
			return true
		}
	}
	return false
}

func (p *plugin) validatorWithMessageExists(fv *validator.FieldValidator) bool {
	return fv != nil && fv.MsgExists != nil && *(fv.MsgExists)
}

func (p *plugin) validatorWithNonRepeatedConstraint(fv *validator.FieldValidator) bool {
	if fv == nil {
		return false
	}

	// Need to use reflection in order to be future-proof for new types of constraints.
	v := reflect.ValueOf(fv).Elem()
	for i := 0; i < v.NumField(); i++ {
		if v.Type().Field(i).Name != "RepeatedCountMin" && v.Type().Field(i).Name != "RepeatedCountMax" {
			return true
		}
	}
	return false
}

func (p *plugin) regexName(ccTypeName string, fieldName string) string {
	return "_regex_" + ccTypeName + "_" + fieldName
}
