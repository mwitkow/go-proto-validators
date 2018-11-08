// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/connectim/go-proto-validators/validator.proto

package validator

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

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
	// Field value of length greater than this value.
	LengthGt *int64 `protobuf:"varint,14,opt,name=length_gt,json=lengthGt" json:"length_gt,omitempty"`
	// Field value of length smaller than this value.
	LengthLt *int64 `protobuf:"varint,15,opt,name=length_lt,json=lengthLt" json:"length_lt,omitempty"`
	// Field value of integer strictly equal this value.
	LengthEq             *int64   `protobuf:"varint,16,opt,name=length_eq,json=lengthEq" json:"length_eq,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FieldValidator) Reset()         { *m = FieldValidator{} }
func (m *FieldValidator) String() string { return proto.CompactTextString(m) }
func (*FieldValidator) ProtoMessage()    {}
func (*FieldValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_eadb76d564aeaff6, []int{0}
}

func (m *FieldValidator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FieldValidator.Unmarshal(m, b)
}
func (m *FieldValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FieldValidator.Marshal(b, m, deterministic)
}
func (m *FieldValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FieldValidator.Merge(m, src)
}
func (m *FieldValidator) XXX_Size() int {
	return xxx_messageInfo_FieldValidator.Size(m)
}
func (m *FieldValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_FieldValidator.DiscardUnknown(m)
}

var xxx_messageInfo_FieldValidator proto.InternalMessageInfo

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

func (m *FieldValidator) GetLengthGt() int64 {
	if m != nil && m.LengthGt != nil {
		return *m.LengthGt
	}
	return 0
}

func (m *FieldValidator) GetLengthLt() int64 {
	if m != nil && m.LengthLt != nil {
		return *m.LengthLt
	}
	return 0
}

func (m *FieldValidator) GetLengthEq() int64 {
	if m != nil && m.LengthEq != nil {
		return *m.LengthEq
	}
	return 0
}

var E_Field = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*FieldValidator)(nil),
	Field:         65020,
	Name:          "validator.field",
	Tag:           "bytes,65020,opt,name=field",
	Filename:      "github.com/connectim/go-proto-validators/validator.proto",
}

func init() {
	proto.RegisterType((*FieldValidator)(nil), "validator.FieldValidator")
	proto.RegisterExtension(E_Field)
}

func init() {
	proto.RegisterFile("github.com/connectim/go-proto-validators/validator.proto", fileDescriptor_eadb76d564aeaff6)
}

var fileDescriptor_eadb76d564aeaff6 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4d, 0x8f, 0x94, 0x30,
	0x18, 0xc7, 0x83, 0xbb, 0xb3, 0x0b, 0x9d, 0xdd, 0x71, 0xd2, 0x68, 0xd2, 0xd5, 0x6c, 0x24, 0x7a,
	0xe1, 0xe0, 0x42, 0xe2, 0xc9, 0x78, 0xd4, 0xe0, 0x5c, 0xc6, 0x97, 0x70, 0xf0, 0xe0, 0x85, 0xb0,
	0xf0, 0x4c, 0xa7, 0x49, 0x69, 0xd9, 0xf6, 0xc1, 0xe0, 0x17, 0xf0, 0x4b, 0xeb, 0xc1, 0x50, 0xe4,
	0xc5, 0x64, 0x6e, 0xf4, 0xf7, 0xfb, 0xf3, 0x14, 0xda, 0x3f, 0x79, 0xcb, 0x05, 0x1e, 0xdb, 0xfb,
	0xb8, 0xd4, 0x75, 0x52, 0x6a, 0xa5, 0xa0, 0x44, 0x51, 0x27, 0x5c, 0xdf, 0x35, 0x46, 0xa3, 0xbe,
	0xfb, 0x51, 0x48, 0x51, 0x15, 0xa8, 0x8d, 0x4d, 0xa6, 0xc7, 0xd8, 0x29, 0x1a, 0x4c, 0xe0, 0x59,
	0xc8, 0xb5, 0xe6, 0x12, 0x12, 0x27, 0xee, 0xdb, 0x43, 0x52, 0x81, 0x2d, 0x8d, 0x68, 0xa6, 0xf0,
	0xcb, 0x5f, 0xe7, 0x64, 0xf3, 0x51, 0x80, 0xac, 0xbe, 0x8d, 0x2f, 0xd1, 0x27, 0x64, 0x65, 0x80,
	0x43, 0xc7, 0xbc, 0xd0, 0x8b, 0x82, 0x6c, 0x58, 0xd0, 0xa7, 0xe4, 0x42, 0x28, 0xcc, 0x39, 0xb2,
	0x47, 0xa1, 0x17, 0x9d, 0x65, 0x2b, 0xa1, 0x70, 0x87, 0x23, 0x96, 0xc8, 0xce, 0x26, 0xbc, 0x47,
	0x7a, 0x4b, 0x48, 0x6d, 0x79, 0x0e, 0x9d, 0xb0, 0x68, 0xd9, 0x79, 0xe8, 0x45, 0x7e, 0x16, 0xd4,
	0x96, 0xa7, 0x0e, 0xd0, 0x17, 0x64, 0x7d, 0x6c, 0xeb, 0x42, 0xe5, 0x60, 0x8c, 0x36, 0x6c, 0xe5,
	0x36, 0x22, 0x0e, 0xa5, 0x3d, 0xa1, 0x37, 0xc4, 0x3f, 0x48, 0x5d, 0xb8, 0xfd, 0x2e, 0x42, 0x2f,
	0xf2, 0xb2, 0x4b, 0xb7, 0xde, 0xe1, 0xac, 0x24, 0xb2, 0xcb, 0x85, 0xda, 0x23, 0x7d, 0x45, 0xae,
	0x07, 0x05, 0x8d, 0x15, 0x52, 0x2b, 0xe6, 0x3b, 0x7f, 0xe5, 0x60, 0x3a, 0x30, 0xfa, 0x9c, 0x04,
	0xe3, 0x68, 0x60, 0x81, 0x0b, 0xf8, 0xff, 0x66, 0xc3, 0x2c, 0x25, 0x02, 0x23, 0x0b, 0xb9, 0x47,
	0xa0, 0x11, 0xd9, 0x5a, 0x34, 0x42, 0xf1, 0x5c, 0x69, 0xcc, 0xa1, 0x6e, 0xf0, 0x27, 0x5b, 0xbb,
	0x5f, 0xdb, 0x0c, 0xfc, 0xb3, 0xc6, 0xb4, 0xa7, 0xf4, 0x35, 0xa1, 0x06, 0x1a, 0x28, 0x10, 0xaa,
	0xbc, 0xd4, 0xad, 0xc2, 0xbc, 0x16, 0x8a, 0x5d, 0xb9, 0x13, 0xda, 0x8e, 0xe6, 0x43, 0x2f, 0x3e,
	0x09, 0x75, 0x2a, 0x5d, 0x74, 0xec, 0xfa, 0x54, 0xba, 0xe8, 0xfa, 0x4f, 0x94, 0xa0, 0x38, 0x1e,
	0xfb, 0xb3, 0xd9, 0xb8, 0x90, 0x3f, 0x80, 0x1d, 0x2e, 0xa4, 0x44, 0xf6, 0x78, 0x29, 0xf7, 0x4b,
	0x09, 0x0f, 0x6c, 0xbb, 0x94, 0xe9, 0xc3, 0xbb, 0xaf, 0x64, 0x75, 0xe8, 0x7b, 0x40, 0x6f, 0xe3,
	0xa1, 0x34, 0xf1, 0x58, 0x9a, 0xd8, 0xf5, 0xe3, 0x4b, 0x83, 0x42, 0x2b, 0xcb, 0xfe, 0xfc, 0xee,
	0x2f, 0x7a, 0xfd, 0xe6, 0x26, 0x9e, 0x7b, 0xf7, 0x7f, 0x81, 0xb2, 0x61, 0xd0, 0xfb, 0xf5, 0xf7,
	0xb9, 0x89, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x51, 0x48, 0xb8, 0xc6, 0xcf, 0x02, 0x00, 0x00,
}
