// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: common.proto

//指定包名

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ErrCode int32

const (
	ErrCode_success              ErrCode = 0
	ErrCode_unknown_error        ErrCode = 999999
	ErrCode_server_fail          ErrCode = -100001
	ErrCode_server_param_invalid ErrCode = -100002
	ErrCode_server_exception     ErrCode = -100003
	ErrCode_server_reject        ErrCode = -100004
	ErrCode_server_busy          ErrCode = -100005
	ErrCode_client_timeout       ErrCode = -100006
	ErrCode_client_exception     ErrCode = -100007
	ErrCode_client_broken        ErrCode = -100008
	ErrCode_no_device_error      ErrCode = -200009
	ErrCode_no_in_param_error    ErrCode = -200010
)

var ErrCode_name = map[int32]string{
	0:       "success",
	999999:  "unknown_error",
	-100001: "server_fail",
	-100002: "server_param_invalid",
	-100003: "server_exception",
	-100004: "server_reject",
	-100005: "server_busy",
	-100006: "client_timeout",
	-100007: "client_exception",
	-100008: "client_broken",
	-200009: "no_device_error",
	-200010: "no_in_param_error",
}

var ErrCode_value = map[string]int32{
	"success":              0,
	"unknown_error":        999999,
	"server_fail":          -100001,
	"server_param_invalid": -100002,
	"server_exception":     -100003,
	"server_reject":        -100004,
	"server_busy":          -100005,
	"client_timeout":       -100006,
	"client_exception":     -100007,
	"client_broken":        -100008,
	"no_device_error":      -200009,
	"no_in_param_error":    -200010,
}

func (x ErrCode) String() string {
	return proto.EnumName(ErrCode_name, int32(x))
}

func (ErrCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

type ErrSubCode int32

const (
	ErrSubCode_err_sub_code_unknown                 ErrSubCode = 0
	ErrSubCode_err_sub_code_book_service            ErrSubCode = 101
	ErrSubCode_err_sub_code_go_book_service         ErrSubCode = 200
	ErrSubCode_err_sub_code_agreement_service       ErrSubCode = 201
	ErrSubCode_err_sub_code_book_store_service      ErrSubCode = 202
	ErrSubCode_err_sub_code_go_device_service       ErrSubCode = 203
	ErrSubCode_err_sub_code_task_service            ErrSubCode = 204
	ErrSubCode_err_sub_code_account_service         ErrSubCode = 205
	ErrSubCode_err_sub_code_writing_contest_service ErrSubCode = 206
	ErrSubCode_err_sub_code_go_user_service         ErrSubCode = 207
	ErrSubCode_err_sub_code_balance_service         ErrSubCode = 208
	ErrSubCode_err_sub_code_track_service           ErrSubCode = 209
)

var ErrSubCode_name = map[int32]string{
	0:   "err_sub_code_unknown",
	101: "err_sub_code_book_service",
	200: "err_sub_code_go_book_service",
	201: "err_sub_code_agreement_service",
	202: "err_sub_code_book_store_service",
	203: "err_sub_code_go_device_service",
	204: "err_sub_code_task_service",
	205: "err_sub_code_account_service",
	206: "err_sub_code_writing_contest_service",
	207: "err_sub_code_go_user_service",
	208: "err_sub_code_balance_service",
	209: "err_sub_code_track_service",
}

var ErrSubCode_value = map[string]int32{
	"err_sub_code_unknown":                 0,
	"err_sub_code_book_service":            101,
	"err_sub_code_go_book_service":         200,
	"err_sub_code_agreement_service":       201,
	"err_sub_code_book_store_service":      202,
	"err_sub_code_go_device_service":       203,
	"err_sub_code_task_service":            204,
	"err_sub_code_account_service":         205,
	"err_sub_code_writing_contest_service": 206,
	"err_sub_code_go_user_service":         207,
	"err_sub_code_balance_service":         208,
	"err_sub_code_track_service":           209,
}

func (x ErrSubCode) String() string {
	return proto.EnumName(ErrSubCode_name, int32(x))
}

func (ErrSubCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

type AppIdType int32

const (
	AppIdType_APP_ID_UNKNOWN    AppIdType = 0
	AppIdType_APP_ID_LIGHTHOUSE AppIdType = 100
)

var AppIdType_name = map[int32]string{
	0:   "APP_ID_UNKNOWN",
	100: "APP_ID_LIGHTHOUSE",
}

var AppIdType_value = map[string]int32{
	"APP_ID_UNKNOWN":    0,
	"APP_ID_LIGHTHOUSE": 100,
}

func (x AppIdType) String() string {
	return proto.EnumName(AppIdType_name, int32(x))
}

func (AppIdType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

type LanguageType int32

const (
	LanguageType_LAN_ZH LanguageType = 0
	LanguageType_LAN_EN LanguageType = 1
	LanguageType_LAN_ES LanguageType = 2
	LanguageType_LAN_IN LanguageType = 3
	LanguageType_LAN_TL LanguageType = 4
	LanguageType_LAN_MS LanguageType = 5
	LanguageType_LAN_KO LanguageType = 6
)

var LanguageType_name = map[int32]string{
	0: "LAN_ZH",
	1: "LAN_EN",
	2: "LAN_ES",
	3: "LAN_IN",
	4: "LAN_TL",
	5: "LAN_MS",
	6: "LAN_KO",
}

var LanguageType_value = map[string]int32{
	"LAN_ZH": 0,
	"LAN_EN": 1,
	"LAN_ES": 2,
	"LAN_IN": 3,
	"LAN_TL": 4,
	"LAN_MS": 5,
	"LAN_KO": 6,
}

func (x LanguageType) String() string {
	return proto.EnumName(LanguageType_name, int32(x))
}

func (LanguageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{3}
}

type ClientType int32

const (
	ClientType_CLIENT_TYPE_UNKNOWN ClientType = 0
	ClientType_CLIENT_TYPE_ANDROID ClientType = 1
	ClientType_CLIENT_TYPE_IOS     ClientType = 5
)

var ClientType_name = map[int32]string{
	0: "CLIENT_TYPE_UNKNOWN",
	1: "CLIENT_TYPE_ANDROID",
	5: "CLIENT_TYPE_IOS",
}

var ClientType_value = map[string]int32{
	"CLIENT_TYPE_UNKNOWN": 0,
	"CLIENT_TYPE_ANDROID": 1,
	"CLIENT_TYPE_IOS":     5,
}

func (x ClientType) String() string {
	return proto.EnumName(ClientType_name, int32(x))
}

func (ClientType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{4}
}

type ChannelType int32

const (
	ChannelType_ChannelType_Unknown ChannelType = 0
	ChannelType_sdk_phoenix         ChannelType = 2
)

var ChannelType_name = map[int32]string{
	0: "ChannelType_Unknown",
	2: "sdk_phoenix",
}

var ChannelType_value = map[string]int32{
	"ChannelType_Unknown": 0,
	"sdk_phoenix":         2,
}

func (x ChannelType) String() string {
	return proto.EnumName(ChannelType_name, int32(x))
}

func (ChannelType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{5}
}

type SDKType int32

const (
	SDKType_SDK_UNKNOWN SDKType = 0
	SDKType_SDK_FICOOL  SDKType = 1
	SDKType_SDK_PHOENIX SDKType = 2
)

var SDKType_name = map[int32]string{
	0: "SDK_UNKNOWN",
	1: "SDK_FICOOL",
	2: "SDK_PHOENIX",
}

var SDKType_value = map[string]int32{
	"SDK_UNKNOWN": 0,
	"SDK_FICOOL":  1,
	"SDK_PHOENIX": 2,
}

func (x SDKType) String() string {
	return proto.EnumName(SDKType_name, int32(x))
}

func (SDKType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{6}
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
	return fileDescriptor_555bd8c177793206, []int{0}
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

func init() {
	proto.RegisterEnum("common.ErrCode", ErrCode_name, ErrCode_value)
	proto.RegisterEnum("common.ErrSubCode", ErrSubCode_name, ErrSubCode_value)
	proto.RegisterEnum("common.AppIdType", AppIdType_name, AppIdType_value)
	proto.RegisterEnum("common.LanguageType", LanguageType_name, LanguageType_value)
	proto.RegisterEnum("common.ClientType", ClientType_name, ClientType_value)
	proto.RegisterEnum("common.ChannelType", ChannelType_name, ChannelType_value)
	proto.RegisterEnum("common.SDKType", SDKType_name, SDKType_value)
	proto.RegisterType((*OutParam)(nil), "common.OutParam")
	proto.RegisterMapType((map[string]string)(nil), "common.OutParam.MetadataEntry")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 808 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x94, 0x4d, 0x6f, 0xe3, 0x44,
	0x18, 0xc7, 0x33, 0x49, 0x93, 0xb4, 0x4f, 0xb6, 0xed, 0xec, 0xb4, 0x40, 0xb6, 0xec, 0x66, 0xbb,
	0xb0, 0x87, 0x92, 0x43, 0x2b, 0x81, 0x04, 0x68, 0x2b, 0x24, 0x4a, 0x12, 0xa8, 0xd5, 0xac, 0x1d,
	0x25, 0xa9, 0x80, 0xbd, 0x8c, 0xc6, 0xce, 0xe0, 0x9a, 0x24, 0x33, 0xd1, 0x78, 0xdc, 0xdd, 0x70,
	0xe6, 0x43, 0xf0, 0x25, 0xb8, 0xc2, 0x91, 0xeb, 0xf2, 0xbe, 0x88, 0x0b, 0xef, 0xa0, 0x02, 0x5f,
	0x00, 0x4e, 0x9c, 0x40, 0x76, 0xfc, 0x12, 0xd3, 0xf5, 0xc5, 0xff, 0xe7, 0xf9, 0xfd, 0xed, 0xe7,
	0x3f, 0xa3, 0xd1, 0xc0, 0x15, 0x47, 0x4e, 0xa7, 0x52, 0xec, 0xcf, 0x94, 0xd4, 0x92, 0x54, 0x16,
	0xd5, 0x33, 0x7f, 0x20, 0x58, 0xb5, 0x02, 0xdd, 0x63, 0x8a, 0x4d, 0x09, 0x81, 0x15, 0x47, 0x8e,
	0x78, 0x1d, 0xed, 0xa2, 0xbd, 0x72, 0x3f, 0xd2, 0xe4, 0x1a, 0xac, 0xfa, 0x81, 0x4d, 0xa3, 0x7e,
	0x31, 0xea, 0x57, 0xfd, 0xc0, 0x6e, 0x85, 0xa8, 0x0e, 0xd5, 0x29, 0xf7, 0x7d, 0xe6, 0xf2, 0x7a,
	0x69, 0x17, 0xed, 0xad, 0xf5, 0x93, 0x92, 0x3c, 0x09, 0x15, 0xc5, 0x99, 0x2f, 0x45, 0x7d, 0x25,
	0x02, 0x71, 0x45, 0xee, 0xc0, 0xea, 0x94, 0x6b, 0x36, 0x62, 0x9a, 0xd5, 0xcb, 0xbb, 0xa5, 0xbd,
	0xda, 0xf3, 0x8d, 0xfd, 0x38, 0x56, 0x12, 0x62, 0xff, 0x6e, 0x6c, 0xe8, 0x08, 0xad, 0xe6, 0xfd,
	0xd4, 0xbf, 0x73, 0x08, 0xeb, 0x39, 0x44, 0x30, 0x94, 0xc6, 0x7c, 0x1e, 0x85, 0x5d, 0xeb, 0x87,
	0x92, 0x6c, 0x43, 0xf9, 0x9c, 0x4d, 0x82, 0x45, 0xd0, 0xb5, 0xfe, 0xa2, 0xb8, 0x53, 0x7c, 0x19,
	0x35, 0xbf, 0x2d, 0x42, 0xb5, 0xa3, 0x54, 0x14, 0xbb, 0x06, 0x55, 0x3f, 0x70, 0x1c, 0xee, 0xfb,
	0xb8, 0x40, 0xb6, 0x60, 0x3d, 0x10, 0x63, 0x21, 0xef, 0x0b, 0xca, 0x95, 0x92, 0x0a, 0x7f, 0xf2,
	0xfe, 0x2b, 0xa4, 0x0e, 0x35, 0x9f, 0xab, 0x73, 0xae, 0xe8, 0x3b, 0xcc, 0x9b, 0xe0, 0x5f, 0xff,
	0xfa, 0xe7, 0xdf, 0xe8, 0x41, 0xe4, 0x16, 0x6c, 0xc7, 0x64, 0x16, 0x86, 0xa5, 0x9e, 0x38, 0x67,
	0x13, 0x6f, 0x84, 0x7f, 0xc9, 0x2c, 0x37, 0x00, 0xc7, 0x16, 0xfe, 0xc0, 0xe1, 0x33, 0xed, 0x49,
	0x81, 0x7f, 0xce, 0xf0, 0x0e, 0xac, 0xc7, 0x58, 0xf1, 0x77, 0xb9, 0xa3, 0xf1, 0x4f, 0x19, 0xcb,
	0xe6, 0xda, 0x81, 0x3f, 0xc7, 0x3f, 0x66, 0xe4, 0x69, 0xd8, 0x70, 0x26, 0x1e, 0x17, 0x9a, 0x6a,
	0x6f, 0xca, 0x65, 0xa0, 0xf1, 0x0f, 0xb9, 0x89, 0x31, 0xcc, 0x26, 0x7e, 0x9f, 0x9b, 0x18, 0x63,
	0x5b, 0xc9, 0x31, 0x17, 0xf8, 0xbb, 0x8c, 0x5d, 0x87, 0x4d, 0x21, 0xe9, 0x88, 0x9f, 0x7b, 0x0e,
	0x8f, 0x37, 0xe0, 0xe3, 0x3f, 0xff, 0x4e, 0x68, 0x03, 0xae, 0x0a, 0x49, 0x3d, 0x11, 0x2f, 0x76,
	0xc1, 0x3f, 0x4a, 0x79, 0xf3, 0xc3, 0x12, 0x40, 0x47, 0xa9, 0x41, 0x7a, 0x1e, 0xb6, 0xb9, 0x52,
	0x34, 0x39, 0x2e, 0x34, 0xde, 0x58, 0x5c, 0x20, 0x37, 0xe0, 0x5a, 0x8e, 0xd8, 0x52, 0x8e, 0x69,
	0xb8, 0x54, 0xcf, 0xe1, 0x98, 0x93, 0x5b, 0x70, 0x3d, 0x87, 0x5d, 0x99, 0x77, 0x3c, 0x44, 0xe4,
	0x59, 0x68, 0xe4, 0x2c, 0xcc, 0x55, 0x9c, 0x4f, 0xc3, 0x45, 0x25, 0xa6, 0x4f, 0x11, 0xb9, 0x0d,
	0x37, 0x1f, 0x33, 0x46, 0x4b, 0xc5, 0x53, 0xd7, 0x67, 0x97, 0x7f, 0xe5, 0xa6, 0x1b, 0x90, 0x98,
	0x3e, 0x0f, 0x97, 0x9e, 0x4f, 0xac, 0x99, 0x9f, 0xe5, 0xf9, 0x02, 0x5d, 0x8a, 0xcc, 0x1c, 0x47,
	0x06, 0x4b, 0x69, 0xbe, 0x44, 0xe4, 0x39, 0xb8, 0x9d, 0xb3, 0xdc, 0x57, 0x9e, 0xf6, 0x84, 0x4b,
	0x1d, 0x29, 0x34, 0xf7, 0x33, 0xeb, 0x57, 0xe8, 0x71, 0x1b, 0x10, 0xf8, 0x5c, 0xa5, 0x96, 0xaf,
	0x2f, 0x5b, 0x6c, 0x36, 0x61, 0x62, 0x29, 0xf3, 0x23, 0x44, 0x6e, 0xc2, 0x4e, 0x3e, 0xb3, 0x62,
	0x4e, 0x16, 0xfa, 0x1b, 0xd4, 0x7c, 0x11, 0xd6, 0x8e, 0x66, 0x33, 0x63, 0x34, 0x9c, 0xcf, 0x38,
	0x21, 0xb0, 0x71, 0xd4, 0xeb, 0x51, 0xa3, 0x4d, 0x4f, 0xcd, 0x13, 0xd3, 0x7a, 0xd3, 0xc4, 0x05,
	0xf2, 0x04, 0x5c, 0x8d, 0x7b, 0x5d, 0xe3, 0x8d, 0xe3, 0xe1, 0xb1, 0x75, 0x3a, 0xe8, 0xe0, 0x51,
	0xd3, 0x86, 0x2b, 0x5d, 0x26, 0xdc, 0x80, 0xb9, 0x3c, 0xfa, 0x14, 0xa0, 0xd2, 0x3d, 0x32, 0xe9,
	0xbd, 0x63, 0x5c, 0x48, 0x74, 0xc7, 0xc4, 0x28, 0xd5, 0x03, 0x5c, 0x4c, 0xb4, 0x61, 0xe2, 0x52,
	0xa2, 0x87, 0x5d, 0xbc, 0x92, 0xe8, 0xbb, 0x03, 0x5c, 0x4e, 0xf4, 0x89, 0x85, 0x2b, 0xcd, 0x01,
	0x40, 0x2b, 0x3a, 0xa5, 0xd1, 0x84, 0xa7, 0x60, 0xab, 0xd5, 0x35, 0x3a, 0xe6, 0x90, 0x0e, 0xdf,
	0xee, 0x75, 0x96, 0x12, 0xfe, 0x0f, 0x1c, 0x99, 0xed, 0xbe, 0x65, 0xb4, 0x31, 0x22, 0x5b, 0xb0,
	0xb9, 0x0c, 0x0c, 0x6b, 0x80, 0xcb, 0xcd, 0x97, 0xa0, 0xd6, 0x3a, 0x63, 0x42, 0xf0, 0x49, 0xfa,
	0xd7, 0xac, 0xa4, 0xa7, 0xe9, 0xf9, 0xdc, 0x84, 0x9a, 0x3f, 0x1a, 0xd3, 0xd9, 0x99, 0xe4, 0xc2,
	0x7b, 0x80, 0x8b, 0xcd, 0x43, 0xa8, 0x0e, 0xda, 0x27, 0xd1, 0x47, 0x9b, 0x50, 0x1b, 0xb4, 0x4f,
	0x96, 0x22, 0x6c, 0x00, 0x84, 0x8d, 0xd7, 0x8d, 0x96, 0x65, 0x75, 0x31, 0x4a, 0x0c, 0xbd, 0x63,
	0xab, 0x63, 0x1a, 0x6f, 0xe1, 0xe2, 0x6b, 0xaf, 0x3e, 0xbc, 0x68, 0xa0, 0x47, 0x17, 0x0d, 0xf4,
	0xdb, 0x45, 0x03, 0x7d, 0xf0, 0x7b, 0xa3, 0x70, 0x6f, 0xdf, 0xf5, 0xf4, 0x59, 0x60, 0x87, 0x77,
	0xdd, 0x01, 0x7b, 0x6f, 0x64, 0xb3, 0xe8, 0x39, 0x08, 0xb4, 0x37, 0x39, 0x88, 0xee, 0xe3, 0x83,
	0xc5, 0x2d, 0x78, 0xb8, 0x78, 0xd9, 0x95, 0xa8, 0xf9, 0xc2, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x16, 0xa9, 0x98, 0xa8, 0xb4, 0x05, 0x00, 0x00,
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
