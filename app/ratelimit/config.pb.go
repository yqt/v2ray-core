package ratelimit

import (
	fmt "fmt"
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

type LimitSettings struct {
	UpRate               int64    `protobuf:"varint,1,opt,name=up_rate,json=upRate,proto3" json:"up_rate,omitempty"`
	UpCapacity           int64    `protobuf:"varint,2,opt,name=up_capacity,json=upCapacity,proto3" json:"up_capacity,omitempty"`
	DownRate             int64    `protobuf:"varint,3,opt,name=down_rate,json=downRate,proto3" json:"down_rate,omitempty"`
	DownCapacity         int64    `protobuf:"varint,4,opt,name=down_capacity,json=downCapacity,proto3" json:"down_capacity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LimitSettings) Reset()         { *m = LimitSettings{} }
func (m *LimitSettings) String() string { return proto.CompactTextString(m) }
func (*LimitSettings) ProtoMessage()    {}
func (*LimitSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{0}
}

func (m *LimitSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LimitSettings.Unmarshal(m, b)
}
func (m *LimitSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LimitSettings.Marshal(b, m, deterministic)
}
func (m *LimitSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LimitSettings.Merge(m, src)
}
func (m *LimitSettings) XXX_Size() int {
	return xxx_messageInfo_LimitSettings.Size(m)
}
func (m *LimitSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_LimitSettings.DiscardUnknown(m)
}

var xxx_messageInfo_LimitSettings proto.InternalMessageInfo

func (m *LimitSettings) GetUpRate() int64 {
	if m != nil {
		return m.UpRate
	}
	return 0
}

func (m *LimitSettings) GetUpCapacity() int64 {
	if m != nil {
		return m.UpCapacity
	}
	return 0
}

func (m *LimitSettings) GetDownRate() int64 {
	if m != nil {
		return m.DownRate
	}
	return 0
}

func (m *LimitSettings) GetDownCapacity() int64 {
	if m != nil {
		return m.DownCapacity
	}
	return 0
}

type LimitRule struct {
	UserEmail            []string       `protobuf:"bytes,1,rep,name=user_email,json=userEmail,proto3" json:"user_email,omitempty"`
	InboundTag           []string       `protobuf:"bytes,2,rep,name=inbound_tag,json=inboundTag,proto3" json:"inbound_tag,omitempty"`
	Settings             *LimitSettings `protobuf:"bytes,3,opt,name=settings,proto3" json:"settings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *LimitRule) Reset()         { *m = LimitRule{} }
func (m *LimitRule) String() string { return proto.CompactTextString(m) }
func (*LimitRule) ProtoMessage()    {}
func (*LimitRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{1}
}

func (m *LimitRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LimitRule.Unmarshal(m, b)
}
func (m *LimitRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LimitRule.Marshal(b, m, deterministic)
}
func (m *LimitRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LimitRule.Merge(m, src)
}
func (m *LimitRule) XXX_Size() int {
	return xxx_messageInfo_LimitRule.Size(m)
}
func (m *LimitRule) XXX_DiscardUnknown() {
	xxx_messageInfo_LimitRule.DiscardUnknown(m)
}

var xxx_messageInfo_LimitRule proto.InternalMessageInfo

func (m *LimitRule) GetUserEmail() []string {
	if m != nil {
		return m.UserEmail
	}
	return nil
}

func (m *LimitRule) GetInboundTag() []string {
	if m != nil {
		return m.InboundTag
	}
	return nil
}

func (m *LimitRule) GetSettings() *LimitSettings {
	if m != nil {
		return m.Settings
	}
	return nil
}

type Config struct {
	Rule                 []*LimitRule `protobuf:"bytes,1,rep,name=rule,proto3" json:"rule,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{2}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetRule() []*LimitRule {
	if m != nil {
		return m.Rule
	}
	return nil
}

func init() {
	proto.RegisterType((*LimitSettings)(nil), "v2ray.core.app.ratelimit.LimitSettings")
	proto.RegisterType((*LimitRule)(nil), "v2ray.core.app.ratelimit.LimitRule")
	proto.RegisterType((*Config)(nil), "v2ray.core.app.ratelimit.Config")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor_3eaf2c85e69e9ea4) }

var fileDescriptor_3eaf2c85e69e9ea4 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0xe5, 0xb6, 0xea, 0xdf, 0xdc, 0xb6, 0x8b, 0x97, 0xdf, 0x12, 0x20, 0xaa, 0x76, 0xa0,
	0x93, 0x87, 0x30, 0x30, 0x97, 0x08, 0x26, 0x06, 0x64, 0x10, 0x03, 0x4b, 0xe4, 0xa6, 0x26, 0xb2,
	0x94, 0xc6, 0x96, 0x63, 0x83, 0xf2, 0x0c, 0x2c, 0x3c, 0x07, 0x4f, 0x89, 0x7c, 0x49, 0x23, 0x40,
	0x42, 0x6c, 0xd6, 0x39, 0xe7, 0xbb, 0x3a, 0xf7, 0x1a, 0x66, 0x85, 0xa9, 0x9f, 0x74, 0xc9, 0xad,
	0x33, 0xde, 0x50, 0xf6, 0x9c, 0x3a, 0xd9, 0xf2, 0xc2, 0x38, 0xc5, 0xa5, 0xb5, 0xdc, 0x49, 0xaf,
	0x2a, 0xbd, 0xd7, 0x7e, 0xf9, 0x4a, 0x60, 0x7e, 0x13, 0x5f, 0x77, 0xca, 0x7b, 0x5d, 0x97, 0x0d,
	0xfd, 0x0f, 0xff, 0x82, 0xcd, 0x63, 0x82, 0x91, 0x05, 0x59, 0x0f, 0xc5, 0x38, 0x58, 0x21, 0xbd,
	0xa2, 0xa7, 0x30, 0x0d, 0x36, 0x2f, 0xa4, 0x95, 0x85, 0xf6, 0x2d, 0x1b, 0xa0, 0x09, 0xc1, 0x66,
	0x9d, 0x42, 0x8f, 0x20, 0xd9, 0x99, 0x97, 0xfa, 0x93, 0x1d, 0xa2, 0x3d, 0x89, 0x02, 0xd2, 0x2b,
	0x98, 0xa3, 0xd9, 0xf3, 0x23, 0x0c, 0xcc, 0xa2, 0x78, 0x98, 0xb0, 0x7c, 0x23, 0x90, 0x60, 0x1b,
	0x11, 0x2a, 0x45, 0x4f, 0x00, 0x42, 0xa3, 0x5c, 0xae, 0xf6, 0x52, 0x57, 0x8c, 0x2c, 0x86, 0xeb,
	0x44, 0x24, 0x51, 0xb9, 0x8a, 0x42, 0xec, 0xa3, 0xeb, 0xad, 0x09, 0xf5, 0x2e, 0xf7, 0xb2, 0x64,
	0x03, 0xf4, 0xa1, 0x93, 0xee, 0x65, 0x49, 0x33, 0x98, 0x34, 0xdd, 0x56, 0x58, 0x67, 0x9a, 0x9e,
	0xf1, 0xdf, 0x0e, 0xc1, 0xbf, 0x1d, 0x41, 0xf4, 0xe0, 0x72, 0x03, 0xe3, 0x0c, 0x4f, 0x49, 0x2f,
	0x60, 0xe4, 0x42, 0xa5, 0xb0, 0xc8, 0x34, 0x5d, 0xfd, 0x31, 0x2a, 0x6e, 0x20, 0x10, 0xb8, 0xbc,
	0x86, 0xe3, 0xc2, 0xec, 0x7f, 0xe6, 0xc5, 0x21, 0x7f, 0x4b, 0x1e, 0x93, 0x1e, 0x7e, 0x1f, 0xb0,
	0x87, 0x54, 0xc8, 0x96, 0x67, 0x31, 0xb7, 0xf9, 0x9a, 0xdb, 0x8e, 0xf1, 0x33, 0xcf, 0x3f, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xe8, 0xdc, 0x36, 0xb5, 0xdc, 0x01, 0x00, 0x00,
}
