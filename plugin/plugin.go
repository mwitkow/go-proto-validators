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

	"github.com/gogo/protobuf/gogoproto"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"github.com/mwitkow/go-proto-validators"
	"github.com/gogo/protobuf/proto"
)

func init() {
	generator.RegisterPlugin(NewPlugin())
}

type plugin struct {
	*generator.Generator
	generator.PluginImports
	regexPkg     generator.Single
	fmtPkg       generator.Single
	protoPkg     generator.Single
	validatorPkg generator.Single
}

func NewPlugin() generator.Plugin {
	return &plugin{}
}

func (p *plugin) Name() string {
	return "validator"
}

func (p *plugin) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *plugin) Generate(file *generator.FileDescriptor) {
	p.PluginImports = generator.NewPluginImports(p.Generator)
	p.regexPkg = p.NewImport("regexp")
	p.fmtPkg = p.NewImport("fmt")
	p.validatorPkg = p.NewImport("github.com/mwitkow/go-proto-validators")

	for _, msg := range file.Messages() {
		if msg.DescriptorProto.GetOptions().GetMapEntry() {
			continue
		}
		p.generateRegexVars(file, msg)
		if gogoproto.IsProto3(file.FileDescriptorProto) {
			p.generateProto3Message(file, msg)
		} else {
			p.generateProto2Message(file, msg)
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

func (p *plugin) generateRegexVars(file *generator.FileDescriptor, message *generator.Descriptor) {
	ccTypeName := generator.CamelCaseSlice(message.TypeName())
	for _, field := range message.Field {
		validator := getFieldValidatorIfAny(field)
		if validator != nil && validator.Regex != nil {
			fieldName := p.GetFieldName(message, field)
			p.P(`var `, p.regexName(ccTypeName, fieldName), ` = `, p.regexPkg.Use(), `.MustCompile("`, validator.Regex, `")`)
		}
	}
}

func (p *plugin) generateProto2Message(file *generator.FileDescriptor, message *generator.Descriptor) {
	ccTypeName := generator.CamelCaseSlice(message.TypeName())

	p.P(`func (this *`, ccTypeName, `) Validate() error {`)
	p.In()
	for _, field := range message.Field {
		fieldName := p.GetFieldName(message, field)
		fieldValudator := getFieldValidatorIfAny(field)
		if fieldValudator == nil && !field.IsMessage() {
			continue
		}
		if p.validatorWithMessageExists(fieldValudator) {
			fmt.Fprintf(os.Stderr, "WARNING: field %v.%v is a proto2 message, validator.msg_exists has no effect\n", ccTypeName, fieldName)
		}
		variableName := "this." + fieldName
		repeated := field.IsRepeated()
		nullable := gogoproto.IsNullable(field)
		if repeated {
			p.P(`for _, item := range `, variableName, `{`)
			p.In()
			variableName = "item"
		} else if nullable {
			p.P(`if `, variableName, ` != nil {`)
			p.In()
			variableName = "*(" + variableName + ")"
		}
		if field.IsString() {
			p.generateStringValidator(variableName, ccTypeName, fieldName, fieldValudator)
		} else if p.isSupportedInt(field) {
			p.generateIntValidator(variableName, ccTypeName, fieldName, fieldValudator)
		} else if field.IsMessage() {
			if repeated && nullable {
				variableName = "*(item)"
			}
			p.P(`if err := `, p.validatorPkg.Use(), `.CallValidatorIfExists(&(`, variableName, `)); err != nil {`)
			p.In()
			p.P(`return err`)
			p.Out()
			p.P(`}`)
		}
		if repeated {
			// end the repeated loop
			p.Out()
			p.P(`}`)
		} else if nullable {
			// end the if around nullable
			p.Out()
			p.P(`}`)
		}
	}
	p.P(`return nil`)
	p.Out()
	p.P(`}`)
}

func (p *plugin) generateProto3Message(file *generator.FileDescriptor, message *generator.Descriptor) {
	ccTypeName := generator.CamelCaseSlice(message.TypeName())
	p.P(`func (this *`, ccTypeName, `) Validate() error {`)
	p.In()
	for _, field := range message.Field {
		fieldValidator := getFieldValidatorIfAny(field)
		if fieldValidator == nil && !field.IsMessage() {
			continue
		}
		fieldName := p.GetFieldName(message, field)
		variableName := "this." + fieldName
		repeated := field.IsRepeated()
		nullable := gogoproto.IsNullable(field)
		if repeated {
			p.P(`for _, item := range `, variableName, `{`)
			p.In()
			variableName = "item"
		}
		if field.IsString() {
			p.generateStringValidator(variableName, ccTypeName, fieldName, fieldValidator)
		} else if p.isSupportedInt(field) {
			p.generateIntValidator(variableName, ccTypeName, fieldName, fieldValidator)
		} else if field.IsMessage() {
			if p.validatorWithMessageExists(fieldValidator) {
				if nullable && !repeated {
					p.P(`if nil == `, variableName, `{`)
					p.In()
					p.P(`return `, p.fmtPkg.Use(), `.Errorf("validation error: `, ccTypeName+"."+fieldName, ` message must exist")`)
					p.Out()
					p.P(`}`)
				} else if repeated {
					fmt.Fprintf(os.Stderr, "WARNING: field %v.%v is a repeated, validator.msg_exists has no effect\n", ccTypeName, fieldName)
				} else if !nullable {
					fmt.Fprintf(os.Stderr, "WARNING: field %v.%v is a nullable=false, validator.msg_exists has no effect\n", ccTypeName, fieldName)
				}
			}
			if nullable {
				p.P(`if `, variableName, ` != nil {`)
				p.In()
			} else {
				// non-nullable fields in proto3 store actual structs, we need pointers to operate on interaces
				variableName = "&(" + variableName + ")"
			}
			p.P(`if err := `, p.validatorPkg.Use(), `.CallValidatorIfExists(`, variableName, `); err != nil {`)
			p.In()
			p.P(`return err`)
			p.Out()
			p.P(`}`)
			if nullable {
				p.Out()
				p.P(`}`)
			}
		}
		if repeated {
			// end the repeated loop
			p.Out()
			p.P(`}`)
		}
	}
	p.P(`return nil`)
	p.Out()
	p.P(`}`)
}

func (p *plugin) generateIntValidator(variableName string, ccTypeName string, fieldName string, fv *validator.FieldValidator) {
	fieldIdentifier := ccTypeName + "." + fieldName
	if fv.IntGt != nil {
		p.P(`if !(`, variableName, ` > `, fv.IntGt, `){`)
		p.In()
		p.P(`return `, p.fmtPkg.Use(), `.Errorf("validation error: `, fieldIdentifier, ` must be greater than '`, fv.IntGt, `'")`)
		p.Out()
		p.P(`}`)
	}
	if fv.IntLt != nil {
		p.P(`if !(`, variableName, ` < `, fv.IntLt, `){`)
		p.In()
		p.P(`return `, p.fmtPkg.Use(), `.Errorf("validation error: `, fieldIdentifier, ` must be less than '`, fv.IntLt, `'")`)
		p.Out()
		p.P(`}`)
	}
}

func (p *plugin) generateStringValidator(variableName string, ccTypeName string, fieldName string, fv *validator.FieldValidator) {
	fieldIdentifier := ccTypeName + "." + fieldName
	if fv.Regex != nil {
		p.P(`if !`, p.regexName(ccTypeName, fieldName), `.MatchString(`, variableName, `) {`)
		p.In()
		p.P(`return `, p.fmtPkg.Use(), `.Errorf("validation error: `, fieldIdentifier, ` must conform to regex '`, fv.Regex, `'")`)
		p.Out()
		p.P(`}`)
	}
}

func (p *plugin) validatorWithMessageExists(fv *validator.FieldValidator) bool {
	return fv != nil && fv.MsgExists != nil && *(fv.MsgExists)
}

func (p *plugin) regexName(ccTypeName string, fieldName string) string {
	return "_regex_" + ccTypeName + "_" + fieldName
}
