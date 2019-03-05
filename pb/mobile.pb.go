// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mobile.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MobileEventType int32

const (
	MobileEventType_NODE_START     MobileEventType = 0
	MobileEventType_NODE_ONLINE    MobileEventType = 1
	MobileEventType_NODE_STOP      MobileEventType = 2
	MobileEventType_WALLET_UPDATE  MobileEventType = 10
	MobileEventType_THREAD_UPDATE  MobileEventType = 11
	MobileEventType_NOTIFICATION   MobileEventType = 12
	MobileEventType_QUERY_RESPONSE MobileEventType = 20
)

var MobileEventType_name = map[int32]string{
	0:  "NODE_START",
	1:  "NODE_ONLINE",
	2:  "NODE_STOP",
	10: "WALLET_UPDATE",
	11: "THREAD_UPDATE",
	12: "NOTIFICATION",
	20: "QUERY_RESPONSE",
}
var MobileEventType_value = map[string]int32{
	"NODE_START":     0,
	"NODE_ONLINE":    1,
	"NODE_STOP":      2,
	"WALLET_UPDATE":  10,
	"THREAD_UPDATE":  11,
	"NOTIFICATION":   12,
	"QUERY_RESPONSE": 20,
}

func (x MobileEventType) String() string {
	return proto.EnumName(MobileEventType_name, int32(x))
}
func (MobileEventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_mobile_d551be6d20e754a5, []int{0}
}

type MobileQueryEvent_Type int32

const (
	MobileQueryEvent_DATA  MobileQueryEvent_Type = 0
	MobileQueryEvent_DONE  MobileQueryEvent_Type = 1
	MobileQueryEvent_ERROR MobileQueryEvent_Type = 2
)

var MobileQueryEvent_Type_name = map[int32]string{
	0: "DATA",
	1: "DONE",
	2: "ERROR",
}
var MobileQueryEvent_Type_value = map[string]int32{
	"DATA":  0,
	"DONE":  1,
	"ERROR": 2,
}

func (x MobileQueryEvent_Type) String() string {
	return proto.EnumName(MobileQueryEvent_Type_name, int32(x))
}
func (MobileQueryEvent_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_mobile_d551be6d20e754a5, []int{3, 0}
}

type MobileWalletAccount struct {
	Seed                 string   `protobuf:"bytes,1,opt,name=seed,proto3" json:"seed,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MobileWalletAccount) Reset()         { *m = MobileWalletAccount{} }
func (m *MobileWalletAccount) String() string { return proto.CompactTextString(m) }
func (*MobileWalletAccount) ProtoMessage()    {}
func (*MobileWalletAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_mobile_d551be6d20e754a5, []int{0}
}
func (m *MobileWalletAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MobileWalletAccount.Unmarshal(m, b)
}
func (m *MobileWalletAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MobileWalletAccount.Marshal(b, m, deterministic)
}
func (dst *MobileWalletAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MobileWalletAccount.Merge(dst, src)
}
func (m *MobileWalletAccount) XXX_Size() int {
	return xxx_messageInfo_MobileWalletAccount.Size(m)
}
func (m *MobileWalletAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MobileWalletAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MobileWalletAccount proto.InternalMessageInfo

func (m *MobileWalletAccount) GetSeed() string {
	if m != nil {
		return m.Seed
	}
	return ""
}

func (m *MobileWalletAccount) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type MobilePreparedFiles struct {
	Dir                  *Directory        `protobuf:"bytes,1,opt,name=dir,proto3" json:"dir,omitempty"`
	Pin                  map[string]string `protobuf:"bytes,2,rep,name=pin,proto3" json:"pin,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MobilePreparedFiles) Reset()         { *m = MobilePreparedFiles{} }
func (m *MobilePreparedFiles) String() string { return proto.CompactTextString(m) }
func (*MobilePreparedFiles) ProtoMessage()    {}
func (*MobilePreparedFiles) Descriptor() ([]byte, []int) {
	return fileDescriptor_mobile_d551be6d20e754a5, []int{1}
}
func (m *MobilePreparedFiles) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MobilePreparedFiles.Unmarshal(m, b)
}
func (m *MobilePreparedFiles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MobilePreparedFiles.Marshal(b, m, deterministic)
}
func (dst *MobilePreparedFiles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MobilePreparedFiles.Merge(dst, src)
}
func (m *MobilePreparedFiles) XXX_Size() int {
	return xxx_messageInfo_MobilePreparedFiles.Size(m)
}
func (m *MobilePreparedFiles) XXX_DiscardUnknown() {
	xxx_messageInfo_MobilePreparedFiles.DiscardUnknown(m)
}

var xxx_messageInfo_MobilePreparedFiles proto.InternalMessageInfo

func (m *MobilePreparedFiles) GetDir() *Directory {
	if m != nil {
		return m.Dir
	}
	return nil
}

func (m *MobilePreparedFiles) GetPin() map[string]string {
	if m != nil {
		return m.Pin
	}
	return nil
}

type MobileFileData struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MobileFileData) Reset()         { *m = MobileFileData{} }
func (m *MobileFileData) String() string { return proto.CompactTextString(m) }
func (*MobileFileData) ProtoMessage()    {}
func (*MobileFileData) Descriptor() ([]byte, []int) {
	return fileDescriptor_mobile_d551be6d20e754a5, []int{2}
}
func (m *MobileFileData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MobileFileData.Unmarshal(m, b)
}
func (m *MobileFileData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MobileFileData.Marshal(b, m, deterministic)
}
func (dst *MobileFileData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MobileFileData.Merge(dst, src)
}
func (m *MobileFileData) XXX_Size() int {
	return xxx_messageInfo_MobileFileData.Size(m)
}
func (m *MobileFileData) XXX_DiscardUnknown() {
	xxx_messageInfo_MobileFileData.DiscardUnknown(m)
}

var xxx_messageInfo_MobileFileData proto.InternalMessageInfo

func (m *MobileFileData) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type MobileQueryEvent struct {
	Id                   string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 MobileQueryEvent_Type `protobuf:"varint,2,opt,name=type,proto3,enum=MobileQueryEvent_Type" json:"type,omitempty"`
	Data                 *QueryResult          `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Error                *Error                `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MobileQueryEvent) Reset()         { *m = MobileQueryEvent{} }
func (m *MobileQueryEvent) String() string { return proto.CompactTextString(m) }
func (*MobileQueryEvent) ProtoMessage()    {}
func (*MobileQueryEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_mobile_d551be6d20e754a5, []int{3}
}
func (m *MobileQueryEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MobileQueryEvent.Unmarshal(m, b)
}
func (m *MobileQueryEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MobileQueryEvent.Marshal(b, m, deterministic)
}
func (dst *MobileQueryEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MobileQueryEvent.Merge(dst, src)
}
func (m *MobileQueryEvent) XXX_Size() int {
	return xxx_messageInfo_MobileQueryEvent.Size(m)
}
func (m *MobileQueryEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_MobileQueryEvent.DiscardUnknown(m)
}

var xxx_messageInfo_MobileQueryEvent proto.InternalMessageInfo

func (m *MobileQueryEvent) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MobileQueryEvent) GetType() MobileQueryEvent_Type {
	if m != nil {
		return m.Type
	}
	return MobileQueryEvent_DATA
}

func (m *MobileQueryEvent) GetData() *QueryResult {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *MobileQueryEvent) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*MobileWalletAccount)(nil), "MobileWalletAccount")
	proto.RegisterType((*MobilePreparedFiles)(nil), "MobilePreparedFiles")
	proto.RegisterMapType((map[string]string)(nil), "MobilePreparedFiles.PinEntry")
	proto.RegisterType((*MobileFileData)(nil), "MobileFileData")
	proto.RegisterType((*MobileQueryEvent)(nil), "MobileQueryEvent")
	proto.RegisterEnum("MobileEventType", MobileEventType_name, MobileEventType_value)
	proto.RegisterEnum("MobileQueryEvent_Type", MobileQueryEvent_Type_name, MobileQueryEvent_Type_value)
}

func init() { proto.RegisterFile("mobile.proto", fileDescriptor_mobile_d551be6d20e754a5) }

var fileDescriptor_mobile_d551be6d20e754a5 = []byte{
	// 455 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x51, 0x6b, 0x9c, 0x40,
	0x18, 0x8c, 0x9e, 0x97, 0xe6, 0x3e, 0xef, 0xcc, 0x76, 0x1b, 0x8a, 0x84, 0x14, 0x0e, 0xa1, 0x10,
	0xf2, 0x60, 0xe1, 0x0a, 0xa5, 0xf4, 0xcd, 0xc6, 0x0d, 0x3d, 0xb8, 0xaa, 0xd9, 0x33, 0x84, 0xf6,
	0xe5, 0xf0, 0xce, 0xa5, 0x2c, 0x35, 0x6a, 0xd7, 0xf5, 0x8a, 0x3f, 0xa2, 0xff, 0xa0, 0x7f, 0xa3,
	0xff, 0xaf, 0xec, 0xaa, 0x2f, 0xa5, 0x6f, 0x33, 0xe3, 0xcc, 0x38, 0x7e, 0x08, 0xf3, 0xa7, 0x6a,
	0xcf, 0x0b, 0xe6, 0xd7, 0xa2, 0x92, 0xd5, 0x25, 0x1c, 0x39, 0xfb, 0x39, 0x60, 0xfb, 0x47, 0xcb,
	0x44, 0x37, 0x90, 0xc5, 0x13, 0x6b, 0x9a, 0xec, 0xdb, 0xe0, 0xf3, 0x6e, 0xe1, 0xc5, 0x67, 0x9d,
	0x7b, 0xcc, 0x8a, 0x82, 0xc9, 0xe0, 0x70, 0xa8, 0xda, 0x52, 0x62, 0x0c, 0x56, 0xc3, 0x58, 0xee,
	0x1a, 0x4b, 0xe3, 0x7a, 0x46, 0x35, 0xc6, 0x2e, 0x3c, 0xcb, 0xf2, 0x5c, 0xb0, 0xa6, 0x71, 0x4d,
	0x2d, 0x8f, 0xd4, 0xfb, 0x6d, 0x8c, 0x2d, 0x89, 0x60, 0x75, 0x26, 0x58, 0x7e, 0xc7, 0x0b, 0xd6,
	0xe0, 0x2b, 0x98, 0xe4, 0x5c, 0xe8, 0x12, 0x7b, 0x05, 0x7e, 0xc8, 0x05, 0x3b, 0xc8, 0x4a, 0x74,
	0x54, 0xc9, 0xf8, 0x0d, 0x4c, 0x6a, 0x5e, 0xba, 0xe6, 0x72, 0x72, 0x6d, 0xaf, 0x5e, 0xf9, 0xff,
	0x29, 0xf0, 0x13, 0x5e, 0x92, 0x52, 0xaa, 0x40, 0xcd, 0xcb, 0xcb, 0x77, 0x70, 0x36, 0x0a, 0x18,
	0xc1, 0xe4, 0x3b, 0xeb, 0x86, 0x7d, 0x0a, 0xe2, 0x0b, 0x98, 0x1e, 0xb3, 0xa2, 0x65, 0xc3, 0xb8,
	0x9e, 0x7c, 0x30, 0xdf, 0x1b, 0x9e, 0x07, 0x4e, 0x5f, 0xae, 0x4a, 0xc3, 0x4c, 0x66, 0x2a, 0xdd,
	0x8a, 0x62, 0x4c, 0xb7, 0xa2, 0xf0, 0xfe, 0x18, 0x80, 0x7a, 0xd3, 0xbd, 0x3a, 0x16, 0x39, 0xb2,
	0x52, 0x62, 0x07, 0x4c, 0x3e, 0xde, 0xc0, 0xe4, 0x39, 0xbe, 0x01, 0x4b, 0x76, 0x75, 0xff, 0x06,
	0x67, 0xf5, 0xd2, 0xff, 0x37, 0xe0, 0xa7, 0x5d, 0xcd, 0xa8, 0xf6, 0xe0, 0x25, 0x58, 0x79, 0x26,
	0x33, 0x77, 0xa2, 0x3f, 0x7e, 0xee, 0x6b, 0x17, 0x65, 0x4d, 0x5b, 0x48, 0xaa, 0x9f, 0xe0, 0x2b,
	0x98, 0x32, 0x21, 0x2a, 0xe1, 0x5a, 0xda, 0x72, 0xea, 0x13, 0xc5, 0x68, 0x2f, 0x7a, 0xaf, 0xc1,
	0x52, 0x6d, 0xf8, 0x0c, 0xac, 0x30, 0x48, 0x03, 0x74, 0xa2, 0x51, 0x1c, 0x11, 0x64, 0xe0, 0x19,
	0x4c, 0x09, 0xa5, 0x31, 0x45, 0xe6, 0xcd, 0x2f, 0x03, 0xce, 0xfb, 0x19, 0x7a, 0x81, 0x8e, 0x38,
	0x00, 0x51, 0x1c, 0x92, 0xdd, 0x36, 0x0d, 0x68, 0x8a, 0x4e, 0xf0, 0x39, 0xd8, 0x9a, 0xc7, 0xd1,
	0x66, 0xad, 0xf3, 0x0b, 0x98, 0x0d, 0x86, 0x38, 0x41, 0x26, 0x7e, 0x0e, 0x8b, 0xc7, 0x60, 0xb3,
	0x21, 0xe9, 0xee, 0x21, 0x09, 0x83, 0x94, 0x20, 0x50, 0x52, 0xfa, 0x89, 0x92, 0x20, 0x1c, 0x25,
	0x1b, 0x23, 0x98, 0x47, 0x71, 0xba, 0xbe, 0x5b, 0xdf, 0x06, 0xe9, 0x3a, 0x8e, 0xd0, 0x1c, 0x63,
	0x70, 0xee, 0x1f, 0x08, 0xfd, 0xb2, 0xa3, 0x64, 0x9b, 0xc4, 0xd1, 0x96, 0xa0, 0x8b, 0x8f, 0xd6,
	0x57, 0xb3, 0xde, 0xef, 0x4f, 0xf5, 0xcf, 0xf5, 0xf6, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xea,
	0xab, 0xea, 0x8b, 0x94, 0x02, 0x00, 0x00,
}