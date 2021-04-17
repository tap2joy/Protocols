// Copyright 2016 Michal Witkowski. All Rights Reserved.
// See LICENSE for licensing terms.

// Protocol Buffers extensions for defining auto-generateable validators for messages.

// TODO(mwitkow): Add example.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.7.1
// source: grpc/validator/validator.proto

package validator

import (
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FieldValidator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Uses a Golang RE2-syntax regex to match the field contents.
	Regex string `protobuf:"bytes,1,opt,name=regex,proto3" json:"regex,omitempty"`
	// Field value of integer strictly greater than this value.
	IntGt int64 `protobuf:"varint,2,opt,name=int_gt,json=intGt,proto3" json:"int_gt,omitempty"`
	// Field value of integer strictly smaller than this value.
	IntLt int64 `protobuf:"varint,3,opt,name=int_lt,json=intLt,proto3" json:"int_lt,omitempty"`
	// Used for nested message types, requires that the message type exists.
	MsgExists bool `protobuf:"varint,4,opt,name=msg_exists,json=msgExists,proto3" json:"msg_exists,omitempty"`
	// Human error specifies a user-customizable error that is visible to the user.
	HumanError string `protobuf:"bytes,5,opt,name=human_error,json=humanError,proto3" json:"human_error,omitempty"`
	// Field value of double strictly greater than this value.
	// Note that this value can only take on a valid floating point
	// value. Use together with float_epsilon if you need something more specific.
	FloatGt float64 `protobuf:"fixed64,6,opt,name=float_gt,json=floatGt,proto3" json:"float_gt,omitempty"`
	// Field value of double strictly smaller than this value.
	// Note that this value can only take on a valid floating point
	// value. Use together with float_epsilon if you need something more specific.
	FloatLt float64 `protobuf:"fixed64,7,opt,name=float_lt,json=floatLt,proto3" json:"float_lt,omitempty"`
	// Field value of double describing the epsilon within which
	// any comparison should be considered to be true. For example,
	// when using float_gt = 0.35, using a float_epsilon of 0.05
	// would mean that any value above 0.30 is acceptable. It can be
	// thought of as a {float_value_condition} +- {float_epsilon}.
	// If unset, no correction for floating point inaccuracies in
	// comparisons will be attempted.
	FloatEpsilon float64 `protobuf:"fixed64,8,opt,name=float_epsilon,json=floatEpsilon,proto3" json:"float_epsilon,omitempty"`
	// Floating-point value compared to which the field content should be greater or equal.
	FloatGte float64 `protobuf:"fixed64,9,opt,name=float_gte,json=floatGte,proto3" json:"float_gte,omitempty"`
	// Floating-point value compared to which the field content should be smaller or equal.
	FloatLte float64 `protobuf:"fixed64,10,opt,name=float_lte,json=floatLte,proto3" json:"float_lte,omitempty"`
	// Used for string fields, requires the string to be not empty (i.e different from "").
	StringNotEmpty bool `protobuf:"varint,11,opt,name=string_not_empty,json=stringNotEmpty,proto3" json:"string_not_empty,omitempty"`
	// Repeated field with at least this number of elements.
	RepeatedCountMin int64 `protobuf:"varint,12,opt,name=repeated_count_min,json=repeatedCountMin,proto3" json:"repeated_count_min,omitempty"`
	// Repeated field with at most this number of elements.
	RepeatedCountMax int64 `protobuf:"varint,13,opt,name=repeated_count_max,json=repeatedCountMax,proto3" json:"repeated_count_max,omitempty"`
	// Field value of length greater than this value.
	LengthGt int64 `protobuf:"varint,14,opt,name=length_gt,json=lengthGt,proto3" json:"length_gt,omitempty"`
	// Field value of length smaller than this value.
	LengthLt int64 `protobuf:"varint,15,opt,name=length_lt,json=lengthLt,proto3" json:"length_lt,omitempty"`
	// Field value of length strictly equal to this value.
	LengthEq int64 `protobuf:"varint,16,opt,name=length_eq,json=lengthEq,proto3" json:"length_eq,omitempty"`
	// Requires that the value is in the enum.
	IsInEnum bool `protobuf:"varint,17,opt,name=is_in_enum,json=isInEnum,proto3" json:"is_in_enum,omitempty"`
	// Ensures that a string value is in UUID format.
	// uuid_ver specifies the valid UUID versions. Valid values are: 0-5.
	// If uuid_ver is 0 all UUID versions are accepted.
	UuidVer int32 `protobuf:"varint,18,opt,name=uuid_ver,json=uuidVer,proto3" json:"uuid_ver,omitempty"`
}

func (x *FieldValidator) Reset() {
	*x = FieldValidator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_validator_validator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldValidator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldValidator) ProtoMessage() {}

func (x *FieldValidator) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_validator_validator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldValidator.ProtoReflect.Descriptor instead.
func (*FieldValidator) Descriptor() ([]byte, []int) {
	return file_grpc_validator_validator_proto_rawDescGZIP(), []int{0}
}

func (x *FieldValidator) GetRegex() string {
	if x != nil {
		return x.Regex
	}
	return ""
}

func (x *FieldValidator) GetIntGt() int64 {
	if x != nil {
		return x.IntGt
	}
	return 0
}

func (x *FieldValidator) GetIntLt() int64 {
	if x != nil {
		return x.IntLt
	}
	return 0
}

func (x *FieldValidator) GetMsgExists() bool {
	if x != nil {
		return x.MsgExists
	}
	return false
}

func (x *FieldValidator) GetHumanError() string {
	if x != nil {
		return x.HumanError
	}
	return ""
}

func (x *FieldValidator) GetFloatGt() float64 {
	if x != nil {
		return x.FloatGt
	}
	return 0
}

func (x *FieldValidator) GetFloatLt() float64 {
	if x != nil {
		return x.FloatLt
	}
	return 0
}

func (x *FieldValidator) GetFloatEpsilon() float64 {
	if x != nil {
		return x.FloatEpsilon
	}
	return 0
}

func (x *FieldValidator) GetFloatGte() float64 {
	if x != nil {
		return x.FloatGte
	}
	return 0
}

func (x *FieldValidator) GetFloatLte() float64 {
	if x != nil {
		return x.FloatLte
	}
	return 0
}

func (x *FieldValidator) GetStringNotEmpty() bool {
	if x != nil {
		return x.StringNotEmpty
	}
	return false
}

func (x *FieldValidator) GetRepeatedCountMin() int64 {
	if x != nil {
		return x.RepeatedCountMin
	}
	return 0
}

func (x *FieldValidator) GetRepeatedCountMax() int64 {
	if x != nil {
		return x.RepeatedCountMax
	}
	return 0
}

func (x *FieldValidator) GetLengthGt() int64 {
	if x != nil {
		return x.LengthGt
	}
	return 0
}

func (x *FieldValidator) GetLengthLt() int64 {
	if x != nil {
		return x.LengthLt
	}
	return 0
}

func (x *FieldValidator) GetLengthEq() int64 {
	if x != nil {
		return x.LengthEq
	}
	return 0
}

func (x *FieldValidator) GetIsInEnum() bool {
	if x != nil {
		return x.IsInEnum
	}
	return false
}

func (x *FieldValidator) GetUuidVer() int32 {
	if x != nil {
		return x.UuidVer
	}
	return 0
}

type OneofValidator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Require that one of the oneof fields is set.
	Required bool `protobuf:"varint,1,opt,name=required,proto3" json:"required,omitempty"`
}

func (x *OneofValidator) Reset() {
	*x = OneofValidator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_validator_validator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OneofValidator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OneofValidator) ProtoMessage() {}

func (x *OneofValidator) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_validator_validator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OneofValidator.ProtoReflect.Descriptor instead.
func (*OneofValidator) Descriptor() ([]byte, []int) {
	return file_grpc_validator_validator_proto_rawDescGZIP(), []int{1}
}

func (x *OneofValidator) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

var file_grpc_validator_validator_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*FieldValidator)(nil),
		Field:         65020,
		Name:          "field",
		Tag:           "bytes,65020,opt,name=field",
		Filename:      "grpc/validator/validator.proto",
	},
	{
		ExtendedType:  (*descriptor.OneofOptions)(nil),
		ExtensionType: (*OneofValidator)(nil),
		Field:         65021,
		Name:          "oneof",
		Tag:           "bytes,65021,opt,name=oneof",
		Filename:      "grpc/validator/validator.proto",
	},
}

// Extension fields to descriptor.FieldOptions.
var (
	// optional FieldValidator field = 65020;
	E_Field = &file_grpc_validator_validator_proto_extTypes[0]
)

// Extension fields to descriptor.OneofOptions.
var (
	// optional OneofValidator oneof = 65021;
	E_Oneof = &file_grpc_validator_validator_proto_extTypes[1]
)

var File_grpc_validator_validator_proto protoreflect.FileDescriptor

var file_grpc_validator_validator_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xbf, 0x04, 0x0a, 0x0e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x12, 0x15, 0x0a, 0x06, 0x69,
	0x6e, 0x74, 0x5f, 0x67, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x6e, 0x74,
	0x47, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x6e, 0x74, 0x5f, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x69, 0x6e, 0x74, 0x4c, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x73, 0x67,
	0x5f, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6d,
	0x73, 0x67, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x75, 0x6d, 0x61,
	0x6e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x68,
	0x75, 0x6d, 0x61, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x6c, 0x6f,
	0x61, 0x74, 0x5f, 0x67, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x66, 0x6c, 0x6f,
	0x61, 0x74, 0x47, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x6c, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x4c, 0x74, 0x12,
	0x23, 0x0a, 0x0d, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x65, 0x70, 0x73, 0x69, 0x6c, 0x6f, 0x6e,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x45, 0x70, 0x73,
	0x69, 0x6c, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x67, 0x74,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x47, 0x74,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x6c, 0x74, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x4c, 0x74, 0x65, 0x12, 0x28,
	0x0a, 0x10, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x4e, 0x6f, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x65, 0x70, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6d, 0x69, 0x6e, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x4d, 0x69, 0x6e, 0x12, 0x2c, 0x0a, 0x12, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x10, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x4d, 0x61, 0x78, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x5f, 0x67,
	0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x47,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x5f, 0x6c, 0x74, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x4c, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x5f, 0x65, 0x71, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x45, 0x71, 0x12, 0x1c, 0x0a, 0x0a, 0x69,
	0x73, 0x5f, 0x69, 0x6e, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x69, 0x73, 0x49, 0x6e, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x75, 0x69,
	0x64, 0x5f, 0x76, 0x65, 0x72, 0x18, 0x12, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x75, 0x75, 0x69,
	0x64, 0x56, 0x65, 0x72, 0x22, 0x2c, 0x0a, 0x0e, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x3a, 0x46, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xfc, 0xfb, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x6f, 0x72, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x3a, 0x46, 0x0a, 0x05, 0x6f, 0x6e,
	0x65, 0x6f, 0x66, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xfd, 0xfb, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x4f, 0x6e, 0x65,
	0x6f, 0x66, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x05, 0x6f, 0x6e, 0x65,
	0x6f, 0x66, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x74, 0x61, 0x70, 0x32, 0x6a, 0x6f, 0x79, 0x2f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_validator_validator_proto_rawDescOnce sync.Once
	file_grpc_validator_validator_proto_rawDescData = file_grpc_validator_validator_proto_rawDesc
)

func file_grpc_validator_validator_proto_rawDescGZIP() []byte {
	file_grpc_validator_validator_proto_rawDescOnce.Do(func() {
		file_grpc_validator_validator_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_validator_validator_proto_rawDescData)
	})
	return file_grpc_validator_validator_proto_rawDescData
}

var file_grpc_validator_validator_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpc_validator_validator_proto_goTypes = []interface{}{
	(*FieldValidator)(nil),          // 0: FieldValidator
	(*OneofValidator)(nil),          // 1: OneofValidator
	(*descriptor.FieldOptions)(nil), // 2: google.protobuf.FieldOptions
	(*descriptor.OneofOptions)(nil), // 3: google.protobuf.OneofOptions
}
var file_grpc_validator_validator_proto_depIdxs = []int32{
	2, // 0: field:extendee -> google.protobuf.FieldOptions
	3, // 1: oneof:extendee -> google.protobuf.OneofOptions
	0, // 2: field:type_name -> FieldValidator
	1, // 3: oneof:type_name -> OneofValidator
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	2, // [2:4] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_validator_validator_proto_init() }
func file_grpc_validator_validator_proto_init() {
	if File_grpc_validator_validator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_validator_validator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldValidator); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_validator_validator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OneofValidator); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_validator_validator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_grpc_validator_validator_proto_goTypes,
		DependencyIndexes: file_grpc_validator_validator_proto_depIdxs,
		MessageInfos:      file_grpc_validator_validator_proto_msgTypes,
		ExtensionInfos:    file_grpc_validator_validator_proto_extTypes,
	}.Build()
	File_grpc_validator_validator_proto = out.File
	file_grpc_validator_validator_proto_rawDesc = nil
	file_grpc_validator_validator_proto_goTypes = nil
	file_grpc_validator_validator_proto_depIdxs = nil
}
