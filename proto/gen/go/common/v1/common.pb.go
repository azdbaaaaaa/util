// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: common/v1/common.proto

//指定包名

package common_v1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	anypb "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ErrCode int32

const (
	ErrCode_ERR_CODE_UNSPECIFIED            ErrCode = 0
	ErrCode_ERR_CODE_SERVER_UNKNOWN         ErrCode = 999999
	ErrCode_ERR_CODE_SERVER_FAILED          ErrCode = -100001
	ErrCode_ERR_CODE_SERVER_PARAM_INVALID   ErrCode = -100002
	ErrCode_ERR_CODE_SERVER_EXCEPTION       ErrCode = -100003
	ErrCode_ERR_CODE_SERVER_REJECTED        ErrCode = -100004
	ErrCode_ERR_CODE_SERVER_BUSY            ErrCode = -100005
	ErrCode_ERR_CODE_CLIENT_TIMEOUT         ErrCode = -100006
	ErrCode_ERR_CODE_CLIENT_EXCEPTION       ErrCode = -100007
	ErrCode_ERR_CODE_CLIENT_BROKEN          ErrCode = -100008
	ErrCode_ERR_CODE_CONTEXT_NO_DEVICE      ErrCode = -200000
	ErrCode_ERR_CODE_CONTEXT_NO_INPARAM     ErrCode = -200001
	ErrCode_ERR_CODE_RECORD_NOT_FOUND       ErrCode = -200002
	ErrCode_ERR_CODE_CONTEXT_DEVICE_INVALID ErrCode = -200003
	ErrCode_ERR_CODE_CONTEXT_NO_USER        ErrCode = -200004
)

var ErrCode_name = map[int32]string{
	0:       "ERR_CODE_UNSPECIFIED",
	999999:  "ERR_CODE_SERVER_UNKNOWN",
	-100001: "ERR_CODE_SERVER_FAILED",
	-100002: "ERR_CODE_SERVER_PARAM_INVALID",
	-100003: "ERR_CODE_SERVER_EXCEPTION",
	-100004: "ERR_CODE_SERVER_REJECTED",
	-100005: "ERR_CODE_SERVER_BUSY",
	-100006: "ERR_CODE_CLIENT_TIMEOUT",
	-100007: "ERR_CODE_CLIENT_EXCEPTION",
	-100008: "ERR_CODE_CLIENT_BROKEN",
	-200000: "ERR_CODE_CONTEXT_NO_DEVICE",
	-200001: "ERR_CODE_CONTEXT_NO_INPARAM",
	-200002: "ERR_CODE_RECORD_NOT_FOUND",
	-200003: "ERR_CODE_CONTEXT_DEVICE_INVALID",
	-200004: "ERR_CODE_CONTEXT_NO_USER",
}

var ErrCode_value = map[string]int32{
	"ERR_CODE_UNSPECIFIED":            0,
	"ERR_CODE_SERVER_UNKNOWN":         999999,
	"ERR_CODE_SERVER_FAILED":          -100001,
	"ERR_CODE_SERVER_PARAM_INVALID":   -100002,
	"ERR_CODE_SERVER_EXCEPTION":       -100003,
	"ERR_CODE_SERVER_REJECTED":        -100004,
	"ERR_CODE_SERVER_BUSY":            -100005,
	"ERR_CODE_CLIENT_TIMEOUT":         -100006,
	"ERR_CODE_CLIENT_EXCEPTION":       -100007,
	"ERR_CODE_CLIENT_BROKEN":          -100008,
	"ERR_CODE_CONTEXT_NO_DEVICE":      -200000,
	"ERR_CODE_CONTEXT_NO_INPARAM":     -200001,
	"ERR_CODE_RECORD_NOT_FOUND":       -200002,
	"ERR_CODE_CONTEXT_DEVICE_INVALID": -200003,
	"ERR_CODE_CONTEXT_NO_USER":        -200004,
}

func (x ErrCode) String() string {
	return proto.EnumName(ErrCode_name, int32(x))
}

func (ErrCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_92d5df4519b8f2e3, []int{0}
}

type ErrSubCode int32

const (
	ErrSubCode_ERR_SUB_CODE_UNSPECIFIED  ErrSubCode = 0
	ErrSubCode_ERR_SUB_CODE_BOOK_SERVICE ErrSubCode = 101
)

var ErrSubCode_name = map[int32]string{
	0:   "ERR_SUB_CODE_UNSPECIFIED",
	101: "ERR_SUB_CODE_BOOK_SERVICE",
}

var ErrSubCode_value = map[string]int32{
	"ERR_SUB_CODE_UNSPECIFIED":  0,
	"ERR_SUB_CODE_BOOK_SERVICE": 101,
}

func (x ErrSubCode) String() string {
	return proto.EnumName(ErrSubCode_name, int32(x))
}

func (ErrSubCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_92d5df4519b8f2e3, []int{1}
}

type AppIdType int32

const (
	AppIdType_APP_ID_TYPE_UNSPECIFIED AppIdType = 0
	AppIdType_APP_ID_TYPE_LIGHTHOUSE  AppIdType = 100
)

var AppIdType_name = map[int32]string{
	0:   "APP_ID_TYPE_UNSPECIFIED",
	100: "APP_ID_TYPE_LIGHTHOUSE",
}

var AppIdType_value = map[string]int32{
	"APP_ID_TYPE_UNSPECIFIED": 0,
	"APP_ID_TYPE_LIGHTHOUSE":  100,
}

func (x AppIdType) String() string {
	return proto.EnumName(AppIdType_name, int32(x))
}

func (AppIdType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_92d5df4519b8f2e3, []int{2}
}

type AreaIdType int32

const (
	AreaIdType_AREA_ID_TYPE_UNSPECIFIED AreaIdType = 0
	AreaIdType_AREA_ID_TYPE_ANDROID     AreaIdType = 30
	AreaIdType_AREA_ID_TYPE_IOS         AreaIdType = 40
	AreaIdType_AREA_ID_TYPE_WEB         AreaIdType = 99
	AreaIdType_AREA_ID_TYPE_H5          AreaIdType = 98
)

var AreaIdType_name = map[int32]string{
	0:  "AREA_ID_TYPE_UNSPECIFIED",
	30: "AREA_ID_TYPE_ANDROID",
	40: "AREA_ID_TYPE_IOS",
	99: "AREA_ID_TYPE_WEB",
	98: "AREA_ID_TYPE_H5",
}

var AreaIdType_value = map[string]int32{
	"AREA_ID_TYPE_UNSPECIFIED": 0,
	"AREA_ID_TYPE_ANDROID":     30,
	"AREA_ID_TYPE_IOS":         40,
	"AREA_ID_TYPE_WEB":         99,
	"AREA_ID_TYPE_H5":          98,
}

func (x AreaIdType) String() string {
	return proto.EnumName(AreaIdType_name, int32(x))
}

func (AreaIdType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_92d5df4519b8f2e3, []int{3}
}

type LanguageType int32

const (
	LanguageType_LANGUAGE_TYPE_UNSPECIFIED LanguageType = 0
	LanguageType_LANGUAGE_TYPE_EN          LanguageType = 1
	LanguageType_LANGUAGE_TYPE_SW          LanguageType = 2
	LanguageType_LANGUAGE_TYPE_AR          LanguageType = 3
	LanguageType_LANGUAGE_TYPE_FR          LanguageType = 4
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
	return fileDescriptor_92d5df4519b8f2e3, []int{4}
}

type ClientType int32

const (
	ClientType_CLIENT_TYPE_UNSPECIFIED ClientType = 0
	ClientType_CLIENT_TYPE_ANDROID     ClientType = 1
	ClientType_CLIENT_TYPE_IOS         ClientType = 5
)

var ClientType_name = map[int32]string{
	0: "CLIENT_TYPE_UNSPECIFIED",
	1: "CLIENT_TYPE_ANDROID",
	5: "CLIENT_TYPE_IOS",
}

var ClientType_value = map[string]int32{
	"CLIENT_TYPE_UNSPECIFIED": 0,
	"CLIENT_TYPE_ANDROID":     1,
	"CLIENT_TYPE_IOS":         5,
}

func (x ClientType) String() string {
	return proto.EnumName(ClientType_name, int32(x))
}

func (ClientType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_92d5df4519b8f2e3, []int{5}
}

type OutParam struct {
	Code                 int32             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	SubCode              int32             `protobuf:"varint,2,opt,name=sub_code,json=subCode,proto3" json:"sub_code,omitempty"`
	Message              string            `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Reason               string            `protobuf:"bytes,4,opt,name=reason,proto3" json:"reason,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,5,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *OutParam) Reset()         { *m = OutParam{} }
func (m *OutParam) String() string { return proto.CompactTextString(m) }
func (*OutParam) ProtoMessage()    {}
func (*OutParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_92d5df4519b8f2e3, []int{0}
}
func (m *OutParam) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OutParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OutParam.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OutParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutParam.Merge(m, src)
}
func (m *OutParam) XXX_Size() int {
	return m.Size()
}
func (m *OutParam) XXX_DiscardUnknown() {
	xxx_messageInfo_OutParam.DiscardUnknown(m)
}

var xxx_messageInfo_OutParam proto.InternalMessageInfo

func (m *OutParam) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *OutParam) GetSubCode() int32 {
	if m != nil {
		return m.SubCode
	}
	return 0
}

func (m *OutParam) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *OutParam) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *OutParam) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type CommonResponse struct {
	Code                 int32             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	SubCode              int32             `protobuf:"varint,2,opt,name=sub_code,json=subCode,proto3" json:"sub_code,omitempty"`
	Message              string            `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Reason               string            `protobuf:"bytes,4,opt,name=reason,proto3" json:"reason,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,5,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Data                 *anypb.Any        `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CommonResponse) Reset()         { *m = CommonResponse{} }
func (m *CommonResponse) String() string { return proto.CompactTextString(m) }
func (*CommonResponse) ProtoMessage()    {}
func (*CommonResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_92d5df4519b8f2e3, []int{1}
}
func (m *CommonResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CommonResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CommonResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CommonResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonResponse.Merge(m, src)
}
func (m *CommonResponse) XXX_Size() int {
	return m.Size()
}
func (m *CommonResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CommonResponse proto.InternalMessageInfo

func (m *CommonResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *CommonResponse) GetSubCode() int32 {
	if m != nil {
		return m.SubCode
	}
	return 0
}

func (m *CommonResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CommonResponse) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *CommonResponse) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *CommonResponse) GetData() *anypb.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("common.v1.ErrCode", ErrCode_name, ErrCode_value)
	proto.RegisterEnum("common.v1.ErrSubCode", ErrSubCode_name, ErrSubCode_value)
	proto.RegisterEnum("common.v1.AppIdType", AppIdType_name, AppIdType_value)
	proto.RegisterEnum("common.v1.AreaIdType", AreaIdType_name, AreaIdType_value)
	proto.RegisterEnum("common.v1.LanguageType", LanguageType_name, LanguageType_value)
	proto.RegisterEnum("common.v1.ClientType", ClientType_name, ClientType_value)
	proto.RegisterType((*OutParam)(nil), "common.v1.OutParam")
	proto.RegisterMapType((map[string]string)(nil), "common.v1.OutParam.MetadataEntry")
	proto.RegisterType((*CommonResponse)(nil), "common.v1.CommonResponse")
	proto.RegisterMapType((map[string]string)(nil), "common.v1.CommonResponse.MetadataEntry")
}

func init() { proto.RegisterFile("common/v1/common.proto", fileDescriptor_92d5df4519b8f2e3) }

var fileDescriptor_92d5df4519b8f2e3 = []byte{
	// 808 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x95, 0xcd, 0x6f, 0xe3, 0x44,
	0x18, 0xc6, 0x71, 0x9b, 0x7e, 0xe4, 0x5d, 0x3e, 0x46, 0xde, 0xa8, 0x75, 0xba, 0xdb, 0x92, 0x5d,
	0x3e, 0x36, 0x8a, 0x50, 0xac, 0x2e, 0x42, 0x42, 0x54, 0x7b, 0x70, 0xec, 0x49, 0x3b, 0x34, 0x9d,
	0x89, 0xc6, 0x76, 0xb3, 0xcb, 0xc5, 0x72, 0x1a, 0x63, 0x2a, 0x12, 0x3b, 0x72, 0xec, 0x48, 0xe1,
	0x0c, 0x7f, 0x06, 0xff, 0x09, 0xb0, 0x1c, 0xf7, 0xc8, 0x91, 0x6f, 0x50, 0x25, 0xc4, 0x1d, 0x4e,
	0x9c, 0x40, 0xb1, 0x9d, 0xd8, 0xf9, 0xd8, 0x13, 0x5a, 0x9f, 0x66, 0xde, 0xe7, 0xe7, 0xd7, 0xcf,
	0xf3, 0x8e, 0x13, 0xc3, 0xde, 0x95, 0x3f, 0x18, 0xf8, 0x9e, 0x3c, 0x3e, 0x96, 0x93, 0x55, 0x7d,
	0x18, 0xf8, 0xa1, 0x2f, 0x16, 0xd3, 0xdd, 0xf8, 0xf8, 0xa0, 0xec, 0xfa, 0xbe, 0xdb, 0x77, 0xe4,
	0x58, 0xe8, 0x46, 0x1f, 0xcb, 0xb6, 0x37, 0x49, 0xa8, 0xfb, 0x7f, 0x0a, 0xb0, 0xcb, 0xa2, 0xb0,
	0x6d, 0x07, 0xf6, 0x40, 0x14, 0xa1, 0x70, 0xe5, 0xf7, 0x1c, 0x49, 0xa8, 0x08, 0xd5, 0x2d, 0x1e,
	0xaf, 0xc5, 0x32, 0xec, 0x8e, 0xa2, 0xae, 0x15, 0xd7, 0x37, 0xe2, 0xfa, 0xce, 0x28, 0xea, 0xaa,
	0x53, 0x49, 0x82, 0x9d, 0x81, 0x33, 0x1a, 0xd9, 0xae, 0x23, 0x6d, 0x56, 0x84, 0x6a, 0x91, 0xcf,
	0xb6, 0xe2, 0x1e, 0x6c, 0x07, 0x8e, 0x3d, 0xf2, 0x3d, 0xa9, 0x10, 0x0b, 0xe9, 0x4e, 0x7c, 0x04,
	0xbb, 0x03, 0x27, 0xb4, 0x7b, 0x76, 0x68, 0x4b, 0x5b, 0x95, 0xcd, 0xea, 0xad, 0x87, 0xf7, 0xea,
	0x73, 0x9b, 0xf5, 0x99, 0x8f, 0xfa, 0x45, 0xca, 0x60, 0x2f, 0x0c, 0x26, 0x7c, 0x7e, 0xcb, 0xc1,
	0x09, 0xbc, 0xb2, 0x20, 0x89, 0x08, 0x36, 0x3f, 0x75, 0x26, 0xb1, 0xdf, 0x22, 0x9f, 0x2e, 0xc5,
	0x12, 0x6c, 0x8d, 0xed, 0x7e, 0x94, 0x78, 0x2d, 0xf2, 0x64, 0xf3, 0xc1, 0xc6, 0xfb, 0xc2, 0xfd,
	0x2f, 0x37, 0xe0, 0x55, 0x35, 0x7e, 0x16, 0x77, 0x46, 0x43, 0xdf, 0x1b, 0x39, 0x2f, 0x3e, 0xaf,
	0xba, 0x92, 0xf7, 0x41, 0x2e, 0xef, 0xa2, 0x9b, 0xe7, 0xa5, 0x16, 0xab, 0x50, 0x88, 0x1b, 0x6c,
	0x57, 0x84, 0xea, 0xad, 0x87, 0xa5, 0x7a, 0x72, 0x98, 0xf5, 0xd9, 0x61, 0xd6, 0x15, 0x6f, 0xc2,
	0x0b, 0xff, 0x7b, 0x3e, 0xb5, 0x67, 0x05, 0xd8, 0xc1, 0x41, 0x90, 0x26, 0x2d, 0x61, 0xce, 0x2d,
	0x95, 0x69, 0xd8, 0x32, 0xa9, 0xde, 0xc6, 0x2a, 0x69, 0x12, 0xac, 0xa1, 0x97, 0xc4, 0x43, 0xd8,
	0x9f, 0x2b, 0x3a, 0xe6, 0x97, 0x98, 0x5b, 0x26, 0x3d, 0xa7, 0xac, 0x43, 0xd1, 0xd3, 0xcf, 0x1f,
	0x89, 0x6f, 0xc0, 0xde, 0xb2, 0xdc, 0x54, 0x48, 0x0b, 0x6b, 0xe8, 0xb7, 0xbf, 0xfe, 0xf9, 0x37,
	0xbe, 0x04, 0xb1, 0x06, 0x87, 0xcb, 0x50, 0x5b, 0xe1, 0xca, 0x85, 0x45, 0xe8, 0xa5, 0xd2, 0x22,
	0x1a, 0xfa, 0x35, 0x63, 0xdf, 0x86, 0xf2, 0x32, 0x8b, 0x1f, 0xab, 0xb8, 0x6d, 0x10, 0x46, 0xd1,
	0x2f, 0x19, 0xf7, 0x16, 0x48, 0xcb, 0x1c, 0xc7, 0x1f, 0x62, 0xd5, 0xc0, 0x1a, 0xfa, 0x39, 0xc3,
	0xee, 0xe5, 0x82, 0xa5, 0x58, 0xc3, 0xd4, 0x9f, 0xa0, 0x9f, 0x32, 0xe4, 0xcd, 0x5c, 0x42, 0xb5,
	0x45, 0x30, 0x35, 0x2c, 0x83, 0x5c, 0x60, 0x66, 0x1a, 0xe8, 0xc7, 0xf5, 0xbe, 0x52, 0x2a, 0xf3,
	0xf5, 0x43, 0xc6, 0xe5, 0x07, 0x92, 0x72, 0x0d, 0xce, 0xce, 0x31, 0x45, 0xdf, 0x67, 0xd0, 0x03,
	0x38, 0xc8, 0x20, 0x46, 0x0d, 0xfc, 0xd8, 0xb0, 0x28, 0xb3, 0x34, 0x7c, 0x49, 0x54, 0x8c, 0xbe,
	0xfd, 0xe3, 0xef, 0x19, 0x58, 0x85, 0x3b, 0xeb, 0x40, 0x42, 0xe3, 0xf9, 0xa1, 0xa7, 0x19, 0x99,
	0xf7, 0xc7, 0xb1, 0xca, 0xb8, 0x66, 0x51, 0x66, 0x58, 0x4d, 0x66, 0x52, 0x0d, 0x7d, 0x93, 0x71,
	0xef, 0xc0, 0xeb, 0x2b, 0x1d, 0x93, 0xe7, 0xce, 0x4f, 0xe3, 0xeb, 0x8c, 0xce, 0x4f, 0x39, 0xf7,
	0x7c, 0x53, 0xc7, 0x1c, 0x7d, 0x35, 0xc7, 0x6a, 0x04, 0x00, 0x07, 0x81, 0x9e, 0xfe, 0x6c, 0xee,
	0x26, 0x37, 0xe9, 0x66, 0x63, 0xfd, 0x0b, 0x55, 0x5e, 0x50, 0x1b, 0x8c, 0x9d, 0xc7, 0x47, 0x33,
	0x8d, 0xee, 0xd4, 0x34, 0x28, 0x2a, 0xc3, 0x21, 0xe9, 0x19, 0x93, 0xa1, 0x23, 0xde, 0x81, 0x7d,
	0xa5, 0xdd, 0xb6, 0x88, 0x66, 0x19, 0x4f, 0xda, 0xcb, 0x8d, 0x0e, 0x60, 0x2f, 0x2f, 0xb6, 0xc8,
	0xe9, 0x99, 0x71, 0xc6, 0x4c, 0x1d, 0xa3, 0x5e, 0xed, 0x0b, 0x01, 0x40, 0x09, 0x1c, 0x3b, 0xed,
	0x73, 0x17, 0x24, 0x85, 0x63, 0xe5, 0x39, 0x8d, 0x24, 0x28, 0x2d, 0xa8, 0x0a, 0xd5, 0x38, 0x23,
	0x1a, 0x3a, 0x12, 0x4b, 0x80, 0x16, 0x14, 0xc2, 0x74, 0x54, 0x5d, 0xa9, 0x76, 0x70, 0x03, 0x5d,
	0x89, 0xb7, 0xe1, 0xb5, 0x85, 0xea, 0xd9, 0x7b, 0xa8, 0x3b, 0xf5, 0xf1, 0x72, 0xcb, 0xf6, 0xdc,
	0xc8, 0x76, 0x9d, 0xd8, 0xc9, 0x21, 0x94, 0x5b, 0x0a, 0x3d, 0x35, 0x95, 0x53, 0xbc, 0xce, 0x4a,
	0x09, 0xd0, 0xa2, 0x8c, 0x29, 0x12, 0x56, 0xab, 0x7a, 0x07, 0x6d, 0xac, 0x56, 0x15, 0x8e, 0x36,
	0x57, 0xab, 0x4d, 0x8e, 0x0a, 0xb5, 0x0e, 0x80, 0xda, 0xbf, 0x76, 0xbc, 0x70, 0x36, 0xd6, 0xd9,
	0x8b, 0xbe, 0x6a, 0x61, 0x1f, 0x6e, 0xe7, 0xc5, 0xd9, 0x30, 0x84, 0x69, 0xc0, 0xbc, 0x30, 0x9d,
	0xc5, 0x56, 0x83, 0x3c, 0xbb, 0x39, 0x12, 0xbe, 0xbb, 0x39, 0x12, 0x7e, 0xbf, 0x39, 0x12, 0x3e,
	0x3a, 0x71, 0xaf, 0xc3, 0x4f, 0xa2, 0xee, 0xf4, 0x2f, 0x4f, 0xb6, 0x3f, 0xeb, 0x75, 0xed, 0xf8,
	0x92, 0xa3, 0xf0, 0xba, 0x9f, 0x7c, 0x8b, 0x64, 0xd7, 0xf1, 0x64, 0xd7, 0x97, 0xe7, 0x9f, 0xb0,
	0x93, 0x64, 0x65, 0x8d, 0x8f, 0xbb, 0xdb, 0x31, 0xf1, 0xee, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x19, 0x91, 0x9e, 0xf7, 0xe0, 0x06, 0x00, 0x00,
}

func (m *OutParam) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OutParam) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OutParam) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Metadata) > 0 {
		for k := range m.Metadata {
			v := m.Metadata[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintCommon(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintCommon(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintCommon(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Reason) > 0 {
		i -= len(m.Reason)
		copy(dAtA[i:], m.Reason)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Reason)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x1a
	}
	if m.SubCode != 0 {
		i = encodeVarintCommon(dAtA, i, uint64(m.SubCode))
		i--
		dAtA[i] = 0x10
	}
	if m.Code != 0 {
		i = encodeVarintCommon(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CommonResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CommonResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CommonResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCommon(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.Metadata) > 0 {
		for k := range m.Metadata {
			v := m.Metadata[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintCommon(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintCommon(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintCommon(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Reason) > 0 {
		i -= len(m.Reason)
		copy(dAtA[i:], m.Reason)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Reason)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x1a
	}
	if m.SubCode != 0 {
		i = encodeVarintCommon(dAtA, i, uint64(m.SubCode))
		i--
		dAtA[i] = 0x10
	}
	if m.Code != 0 {
		i = encodeVarintCommon(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCommon(dAtA []byte, offset int, v uint64) int {
	offset -= sovCommon(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OutParam) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovCommon(uint64(m.Code))
	}
	if m.SubCode != 0 {
		n += 1 + sovCommon(uint64(m.SubCode))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.Reason)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	if len(m.Metadata) > 0 {
		for k, v := range m.Metadata {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovCommon(uint64(len(k))) + 1 + len(v) + sovCommon(uint64(len(v)))
			n += mapEntrySize + 1 + sovCommon(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CommonResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovCommon(uint64(m.Code))
	}
	if m.SubCode != 0 {
		n += 1 + sovCommon(uint64(m.SubCode))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.Reason)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	if len(m.Metadata) > 0 {
		for k, v := range m.Metadata {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovCommon(uint64(len(k))) + 1 + len(v) + sovCommon(uint64(len(v)))
			n += mapEntrySize + 1 + sovCommon(uint64(mapEntrySize))
		}
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovCommon(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCommon(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCommon(x uint64) (n int) {
	return sovCommon(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OutParam) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OutParam: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OutParam: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubCode", wireType)
			}
			m.SubCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SubCode |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCommon
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCommon
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthCommon
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthCommon
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCommon
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthCommon
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthCommon
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipCommon(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthCommon
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Metadata[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CommonResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CommonResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CommonResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubCode", wireType)
			}
			m.SubCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SubCode |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCommon
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCommon
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthCommon
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthCommon
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCommon
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthCommon
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthCommon
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipCommon(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthCommon
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Metadata[mapkey] = mapvalue
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &anypb.Any{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCommon(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommon
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthCommon
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCommon
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCommon
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCommon        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommon          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCommon = fmt.Errorf("proto: unexpected end of group")
)
