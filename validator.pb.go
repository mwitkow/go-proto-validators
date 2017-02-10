// Code generated by protoc-gen-gogo.
// source: validator.proto
// DO NOT EDIT!

/*
Package validator is a generated protocol buffer package.

It is generated from these files:
	validator.proto

It has these top-level messages:
	FieldValidator
*/
package validator

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type FieldValidator struct {
	// Uses a Golang RE2-syntax regex to match the field contents.
	Regex *string `protobuf:"bytes,1,opt,name=regex" json:"regex,omitempty"`
	// Field value of integer strictly greater than this value.
	IntGt *int64 `protobuf:"varint,2,opt,name=int_gt,json=intGt" json:"int_gt,omitempty"`
	// Field value of integer strictly smaller than this value.
	IntLt *int64 `protobuf:"varint,3,opt,name=int_lt,json=intLt" json:"int_lt,omitempty"`
	// Used for nested message types, requires that the message type exists.
	MsgExists *bool `protobuf:"varint,4,opt,name=msg_exists,json=msgExists" json:"msg_exists,omitempty"`
	// Human error specifies a user-customizable error that is visible to the user.
	HumanError *string `protobuf:"bytes,5,opt,name=human_error,json=humanError" json:"human_error,omitempty"`
	// Field value of double strictly greater than this value.
	// Note that this value can only take on a valid floating point
	// value. Use together with float_epsilon if you need something more specific.
	FloatGt *float64 `protobuf:"fixed64,6,opt,name=float_gt,json=floatGt" json:"float_gt,omitempty"`
	// Field value of double strictly smaller than this value.
	// Note that this value can only take on a valid floating point
	// value. Use together with float_epsilon if you need something more specific.
	FloatLt *float64 `protobuf:"fixed64,7,opt,name=float_lt,json=floatLt" json:"float_lt,omitempty"`
	// Field value of double describing the epsilon within which
	// any comparison should be considered to be true. For example,
	// when using float_gt = 0.35, using a float_epsilon of 0.05
	// would mean that any value above 0.30 is acceptable. It can be
	// thought of as a {float_value_condition} +- {float_epsilon}.
	// If unset, no correction for floating point inaccuracies in
	// comparisons will be attempted.
	FloatEpsilon *float64 `protobuf:"fixed64,8,opt,name=float_epsilon,json=floatEpsilon" json:"float_epsilon,omitempty"`
	// Floating-point value compared to which the field content should be greater or equal.
	FloatGte *float64 `protobuf:"fixed64,9,opt,name=float_gte,json=floatGte" json:"float_gte,omitempty"`
	// Floating-point value compared to which the field content should be smaller or equal.
	FloatLte *float64 `protobuf:"fixed64,10,opt,name=float_lte,json=floatLte" json:"float_lte,omitempty"`
	// Used for string fields, requires the string to be not empty (i.e different from "").
	StringNotEmpty *bool `protobuf:"varint,11,opt,name=string_not_empty,json=stringNotEmpty" json:"string_not_empty,omitempty"`
	// Repeated field with at least this number of elements.
	RepeatedCountMin *int64 `protobuf:"varint,12,opt,name=repeated_count_min,json=repeatedCountMin" json:"repeated_count_min,omitempty"`
	// Repeated field with at most this number of elements.
	RepeatedCountMax *int64 `protobuf:"varint,13,opt,name=repeated_count_max,json=repeatedCountMax" json:"repeated_count_max,omitempty"`
	// Field value of integer  should be greater or equal.
	IntGte *int64 `protobuf:"varint,14,opt,name=int_gte,json=intGte" json:"int_gte,omitempty"`
	// Field value of integer  should be smaller or equal.
	IntLte           *int64 `protobuf:"varint,15,opt,name=int_lte,json=intLte" json:"int_lte,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *FieldValidator) Reset()                    { *m = FieldValidator{} }
func (m *FieldValidator) String() string            { return proto.CompactTextString(m) }
func (*FieldValidator) ProtoMessage()               {}
func (*FieldValidator) Descriptor() ([]byte, []int) { return fileDescriptorValidator, []int{0} }

func (m *FieldValidator) GetRegex() string {
	if m != nil && m.Regex != nil {
		return *m.Regex
	}
	return ""
}

func (m *FieldValidator) GetIntGt() int64 {
	if m != nil && m.IntGt != nil {
		return *m.IntGt
	}
	return 0
}

func (m *FieldValidator) GetIntLt() int64 {
	if m != nil && m.IntLt != nil {
		return *m.IntLt
	}
	return 0
}

func (m *FieldValidator) GetMsgExists() bool {
	if m != nil && m.MsgExists != nil {
		return *m.MsgExists
	}
	return false
}

func (m *FieldValidator) GetHumanError() string {
	if m != nil && m.HumanError != nil {
		return *m.HumanError
	}
	return ""
}

func (m *FieldValidator) GetFloatGt() float64 {
	if m != nil && m.FloatGt != nil {
		return *m.FloatGt
	}
	return 0
}

func (m *FieldValidator) GetFloatLt() float64 {
	if m != nil && m.FloatLt != nil {
		return *m.FloatLt
	}
	return 0
}

func (m *FieldValidator) GetFloatEpsilon() float64 {
	if m != nil && m.FloatEpsilon != nil {
		return *m.FloatEpsilon
	}
	return 0
}

func (m *FieldValidator) GetFloatGte() float64 {
	if m != nil && m.FloatGte != nil {
		return *m.FloatGte
	}
	return 0
}

func (m *FieldValidator) GetFloatLte() float64 {
	if m != nil && m.FloatLte != nil {
		return *m.FloatLte
	}
	return 0
}

func (m *FieldValidator) GetStringNotEmpty() bool {
	if m != nil && m.StringNotEmpty != nil {
		return *m.StringNotEmpty
	}
	return false
}

func (m *FieldValidator) GetRepeatedCountMin() int64 {
	if m != nil && m.RepeatedCountMin != nil {
		return *m.RepeatedCountMin
	}
	return 0
}

func (m *FieldValidator) GetRepeatedCountMax() int64 {
	if m != nil && m.RepeatedCountMax != nil {
		return *m.RepeatedCountMax
	}
	return 0
}

func (m *FieldValidator) GetIntGte() int64 {
	if m != nil && m.IntGte != nil {
		return *m.IntGte
	}
	return 0
}

func (m *FieldValidator) GetIntLte() int64 {
	if m != nil && m.IntLte != nil {
		return *m.IntLte
	}
	return 0
}

var E_Field = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*FieldValidator)(nil),
	Field:         65020,
	Name:          "validator.field",
	Tag:           "bytes,65020,opt,name=field",
	Filename:      "validator.proto",
}

func init() {
	proto.RegisterType((*FieldValidator)(nil), "validator.FieldValidator")
	proto.RegisterExtension(E_Field)
}

func init() { proto.RegisterFile("validator.proto", fileDescriptorValidator) }

var fileDescriptorValidator = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x92, 0x4d, 0xaf, 0x93, 0x40,
	0x14, 0x86, 0x83, 0x95, 0x16, 0x0e, 0xf7, 0xf6, 0xde, 0x4c, 0x34, 0xce, 0xd5, 0xdc, 0x48, 0x74,
	0xc3, 0xc2, 0xd0, 0xc4, 0xa5, 0x4b, 0x0d, 0x76, 0x83, 0x1f, 0x61, 0xe1, 0xc2, 0x0d, 0xc1, 0x72,
	0x8a, 0x93, 0x0c, 0x33, 0x64, 0xe6, 0xd4, 0xe0, 0xcf, 0xf5, 0x77, 0xe8, 0xc2, 0x30, 0x08, 0xd4,
	0xa4, 0xcb, 0x79, 0x9e, 0xc3, 0x19, 0xe0, 0x7d, 0xe1, 0xe6, 0x47, 0x25, 0x45, 0x5d, 0x91, 0x36,
	0x69, 0x67, 0x34, 0x69, 0x16, 0xce, 0xe0, 0x69, 0xdc, 0x68, 0xdd, 0x48, 0xdc, 0x39, 0xf1, 0xed,
	0x74, 0xdc, 0xd5, 0x68, 0x0f, 0x46, 0x74, 0xf3, 0xf0, 0x8b, 0x5f, 0x2b, 0xd8, 0xbe, 0x17, 0x28,
	0xeb, 0x2f, 0xd3, 0x43, 0xec, 0x11, 0xf8, 0x06, 0x1b, 0xec, 0xb9, 0x17, 0x7b, 0x49, 0x58, 0x8c,
	0x07, 0xf6, 0x18, 0xd6, 0x42, 0x51, 0xd9, 0x10, 0x7f, 0x10, 0x7b, 0xc9, 0xaa, 0xf0, 0x85, 0xa2,
	0x3d, 0x4d, 0x58, 0x12, 0x5f, 0xcd, 0x38, 0x27, 0x76, 0x0f, 0xd0, 0xda, 0xa6, 0xc4, 0x5e, 0x58,
	0xb2, 0xfc, 0x61, 0xec, 0x25, 0x41, 0x11, 0xb6, 0xb6, 0xc9, 0x1c, 0x60, 0xcf, 0x21, 0xfa, 0x7e,
	0x6a, 0x2b, 0x55, 0xa2, 0x31, 0xda, 0x70, 0xdf, 0x5d, 0x04, 0x0e, 0x65, 0x03, 0x61, 0x77, 0x10,
	0x1c, 0xa5, 0xae, 0xdc, 0x7d, 0xeb, 0xd8, 0x4b, 0xbc, 0x62, 0xe3, 0xce, 0x7b, 0x5a, 0x94, 0x24,
	0xbe, 0x39, 0x53, 0x39, 0xb1, 0x97, 0x70, 0x3d, 0x2a, 0xec, 0xac, 0x90, 0x5a, 0xf1, 0xc0, 0xf9,
	0x2b, 0x07, 0xb3, 0x91, 0xb1, 0x67, 0x10, 0x4e, 0xab, 0x91, 0x87, 0x6e, 0x20, 0xf8, 0xb7, 0x1b,
	0x17, 0x29, 0x09, 0x39, 0x9c, 0xc9, 0x9c, 0x90, 0x25, 0x70, 0x6b, 0xc9, 0x08, 0xd5, 0x94, 0x4a,
	0x53, 0x89, 0x6d, 0x47, 0x3f, 0x79, 0xe4, 0x3e, 0x6d, 0x3b, 0xf2, 0x8f, 0x9a, 0xb2, 0x81, 0xb2,
	0x57, 0xc0, 0x0c, 0x76, 0x58, 0x11, 0xd6, 0xe5, 0x41, 0x9f, 0x14, 0x95, 0xad, 0x50, 0xfc, 0xca,
	0xfd, 0xa1, 0xdb, 0xc9, 0xbc, 0x1b, 0xc4, 0x07, 0xa1, 0x2e, 0x4d, 0x57, 0x3d, 0xbf, 0xbe, 0x34,
	0x5d, 0xf5, 0xec, 0x09, 0x6c, 0xc6, 0x20, 0x90, 0x6f, 0xdd, 0xc8, 0xda, 0x25, 0x81, 0x93, 0x18,
	0xde, 0xfc, 0x66, 0x16, 0x39, 0xe1, 0x9b, 0xcf, 0xe0, 0x1f, 0x87, 0x88, 0xd9, 0x7d, 0x3a, 0xf6,
	0x21, 0x9d, 0xfa, 0x90, 0xba, 0xe8, 0x3f, 0x75, 0x24, 0xb4, 0xb2, 0xfc, 0xcf, 0xef, 0x21, 0xc3,
	0xe8, 0xf5, 0x5d, 0xba, 0x54, 0xea, 0xff, 0x6e, 0x14, 0xe3, 0xa2, 0xb7, 0xd1, 0xd7, 0xa5, 0x64,
	0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x5f, 0xd4, 0xc1, 0xab, 0x81, 0x02, 0x00, 0x00,
}
