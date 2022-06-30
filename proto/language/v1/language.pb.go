// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: language/v1/language.proto

//指定包名

package language_v1

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type LanguageType int32

const (
	LANGUAGE_TYPE_UNSPECIFIED LanguageType = 0
	LANGUAGE_TYPE_EN          LanguageType = 1
	LANGUAGE_TYPE_SW          LanguageType = 2
	LANGUAGE_TYPE_AR          LanguageType = 3
	LANGUAGE_TYPE_FR          LanguageType = 4
)

var LanguageType_name = map[int32]string{
	0: "LANGUAGE_TYPE_UNSPECIFIED",
	1: "LANGUAGE_TYPE_EN",
	2: "LANGUAGE_TYPE_SW",
	3: "LANGUAGE_TYPE_AR",
	4: "LANGUAGE_TYPE_FR",
}

var LanguageType_value = map[string]int32{
	"LANGUAGE_TYPE_UNSPECIFIED": 0,
	"LANGUAGE_TYPE_EN":          1,
	"LANGUAGE_TYPE_SW":          2,
	"LANGUAGE_TYPE_AR":          3,
	"LANGUAGE_TYPE_FR":          4,
}

func (x LanguageType) String() string {
	return proto.EnumName(LanguageType_name, int32(x))
}

func (LanguageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_943923651e6e816c, []int{0}
}

func init() {
	proto.RegisterEnum("language.v1.LanguageType", LanguageType_name, LanguageType_value)
}

func init() { proto.RegisterFile("language/v1/language.proto", fileDescriptor_943923651e6e816c) }

var fileDescriptor_943923651e6e816c = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xca, 0x49, 0xcc, 0x4b,
	0x2f, 0x4d, 0x4c, 0x4f, 0xd5, 0x2f, 0x33, 0xd4, 0x87, 0xb1, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2,
	0x85, 0xb8, 0xe1, 0xfc, 0x32, 0x43, 0x29, 0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0xb0, 0xb8, 0x3e, 0x88,
	0x05, 0x51, 0xa2, 0xd5, 0xca, 0xc8, 0xc5, 0xe3, 0x03, 0x55, 0x15, 0x52, 0x59, 0x90, 0x2a, 0x24,
	0xcb, 0x25, 0xe9, 0xe3, 0xe8, 0xe7, 0x1e, 0xea, 0xe8, 0xee, 0x1a, 0x1f, 0x12, 0x19, 0xe0, 0x1a,
	0x1f, 0xea, 0x17, 0x1c, 0xe0, 0xea, 0xec, 0xe9, 0xe6, 0xe9, 0xea, 0x22, 0xc0, 0x20, 0x24, 0xc2,
	0x25, 0x80, 0x2a, 0xed, 0xea, 0x27, 0xc0, 0x88, 0x29, 0x1a, 0x1c, 0x2e, 0xc0, 0x84, 0x29, 0xea,
	0x18, 0x24, 0xc0, 0x8c, 0x29, 0xea, 0x16, 0x24, 0xc0, 0xe2, 0x14, 0x76, 0xe1, 0xa1, 0x1c, 0xc3,
	0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe3, 0xb1, 0x1c,
	0xc3, 0x81, 0xc7, 0x72, 0x8c, 0x27, 0x1e, 0xcb, 0x31, 0x46, 0x59, 0xa4, 0x67, 0x96, 0x64, 0x94,
	0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x27, 0x56, 0xa5, 0x24, 0x25, 0x82, 0x81, 0x7e, 0x69, 0x49,
	0x66, 0x8e, 0x3e, 0xc4, 0x4b, 0x48, 0xa1, 0x60, 0x0d, 0x63, 0xc7, 0x97, 0x19, 0x26, 0xb1, 0x81,
	0xa5, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x04, 0x76, 0x78, 0x27, 0x01, 0x00, 0x00,
}