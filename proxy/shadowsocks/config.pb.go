package shadowsocks

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v2ray_core_common_protocol "v2ray.com/core/common/protocol"
import v2ray_core_common_protocol1 "v2ray.com/core/common/protocol"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CipherType int32

const (
	CipherType_UNKNOWN       CipherType = 0
	CipherType_AES_128_CFB   CipherType = 1
	CipherType_AES_256_CFB   CipherType = 2
	CipherType_CHACHA20      CipherType = 3
	CipherType_CHACHA20_IETF CipherType = 4
)

var CipherType_name = map[int32]string{
	0: "UNKNOWN",
	1: "AES_128_CFB",
	2: "AES_256_CFB",
	3: "CHACHA20",
	4: "CHACHA20_IETF",
}
var CipherType_value = map[string]int32{
	"UNKNOWN":       0,
	"AES_128_CFB":   1,
	"AES_256_CFB":   2,
	"CHACHA20":      3,
	"CHACHA20_IETF": 4,
}

func (x CipherType) String() string {
	return proto.EnumName(CipherType_name, int32(x))
}
func (CipherType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Account_OneTimeAuth int32

const (
	Account_Auto     Account_OneTimeAuth = 0
	Account_Disabled Account_OneTimeAuth = 1
	Account_Enabled  Account_OneTimeAuth = 2
)

var Account_OneTimeAuth_name = map[int32]string{
	0: "Auto",
	1: "Disabled",
	2: "Enabled",
}
var Account_OneTimeAuth_value = map[string]int32{
	"Auto":     0,
	"Disabled": 1,
	"Enabled":  2,
}

func (x Account_OneTimeAuth) String() string {
	return proto.EnumName(Account_OneTimeAuth_name, int32(x))
}
func (Account_OneTimeAuth) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Account struct {
	Password   string              `protobuf:"bytes,1,opt,name=password" json:"password,omitempty"`
	CipherType CipherType          `protobuf:"varint,2,opt,name=cipher_type,json=cipherType,enum=v2ray.core.proxy.shadowsocks.CipherType" json:"cipher_type,omitempty"`
	Ota        Account_OneTimeAuth `protobuf:"varint,3,opt,name=ota,enum=v2ray.core.proxy.shadowsocks.Account_OneTimeAuth" json:"ota,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Account) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Account) GetCipherType() CipherType {
	if m != nil {
		return m.CipherType
	}
	return CipherType_UNKNOWN
}

func (m *Account) GetOta() Account_OneTimeAuth {
	if m != nil {
		return m.Ota
	}
	return Account_Auto
}

type ServerConfig struct {
	UdpEnabled bool                             `protobuf:"varint,1,opt,name=udp_enabled,json=udpEnabled" json:"udp_enabled,omitempty"`
	User       *v2ray_core_common_protocol.User `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
}

func (m *ServerConfig) Reset()                    { *m = ServerConfig{} }
func (m *ServerConfig) String() string            { return proto.CompactTextString(m) }
func (*ServerConfig) ProtoMessage()               {}
func (*ServerConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ServerConfig) GetUdpEnabled() bool {
	if m != nil {
		return m.UdpEnabled
	}
	return false
}

func (m *ServerConfig) GetUser() *v2ray_core_common_protocol.User {
	if m != nil {
		return m.User
	}
	return nil
}

type ClientConfig struct {
	Server []*v2ray_core_common_protocol1.ServerEndpoint `protobuf:"bytes,1,rep,name=server" json:"server,omitempty"`
}

func (m *ClientConfig) Reset()                    { *m = ClientConfig{} }
func (m *ClientConfig) String() string            { return proto.CompactTextString(m) }
func (*ClientConfig) ProtoMessage()               {}
func (*ClientConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ClientConfig) GetServer() []*v2ray_core_common_protocol1.ServerEndpoint {
	if m != nil {
		return m.Server
	}
	return nil
}

func init() {
	proto.RegisterType((*Account)(nil), "v2ray.core.proxy.shadowsocks.Account")
	proto.RegisterType((*ServerConfig)(nil), "v2ray.core.proxy.shadowsocks.ServerConfig")
	proto.RegisterType((*ClientConfig)(nil), "v2ray.core.proxy.shadowsocks.ClientConfig")
	proto.RegisterEnum("v2ray.core.proxy.shadowsocks.CipherType", CipherType_name, CipherType_value)
	proto.RegisterEnum("v2ray.core.proxy.shadowsocks.Account_OneTimeAuth", Account_OneTimeAuth_name, Account_OneTimeAuth_value)
}

func init() { proto.RegisterFile("v2ray.com/core/proxy/shadowsocks/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x51, 0xdd, 0x6e, 0xd3, 0x4c,
	0x14, 0xac, 0x93, 0xa8, 0xcd, 0x77, 0x36, 0x1f, 0x98, 0xbd, 0x8a, 0xa2, 0x4a, 0x58, 0xb9, 0x0a,
	0x95, 0x58, 0xb7, 0xcb, 0x8f, 0xb8, 0x75, 0x4c, 0xaa, 0x56, 0x48, 0x69, 0xe5, 0xa4, 0x20, 0x01,
	0x92, 0xe5, 0xae, 0x17, 0x62, 0x11, 0xfb, 0xac, 0x76, 0xed, 0x16, 0xbf, 0x12, 0x6f, 0xc6, 0x5b,
	0x20, 0xaf, 0x93, 0x36, 0xe2, 0x22, 0xdc, 0xf9, 0x9c, 0x33, 0x33, 0x9e, 0x99, 0x85, 0x97, 0x77,
	0x5c, 0x27, 0x35, 0x13, 0x98, 0xfb, 0x02, 0xb5, 0xf4, 0x95, 0xc6, 0x9f, 0xb5, 0x6f, 0x56, 0x49,
	0x8a, 0xf7, 0x06, 0xc5, 0x0f, 0xe3, 0x0b, 0x2c, 0xbe, 0x65, 0xdf, 0x99, 0xd2, 0x58, 0x22, 0x3d,
	0xde, 0xc2, 0xb5, 0x64, 0x16, 0xca, 0x76, 0xa0, 0xa3, 0x17, 0x7f, 0x89, 0x09, 0xcc, 0x73, 0x2c,
	0x7c, 0x4b, 0x15, 0xb8, 0xf6, 0x2b, 0x23, 0x75, 0x2b, 0x34, 0x3a, 0xfd, 0x07, 0xd4, 0x48, 0x7d,
	0x27, 0x75, 0x6c, 0x94, 0x14, 0x2d, 0x63, 0xfc, 0xdb, 0x81, 0xa3, 0x40, 0x08, 0xac, 0x8a, 0x92,
	0x8e, 0xa0, 0xaf, 0x12, 0x63, 0xee, 0x51, 0xa7, 0x43, 0xc7, 0x73, 0x26, 0xff, 0x45, 0x0f, 0x33,
	0xbd, 0x04, 0x22, 0x32, 0xb5, 0x92, 0x3a, 0x2e, 0x6b, 0x25, 0x87, 0x1d, 0xcf, 0x99, 0x3c, 0xe1,
	0x13, 0xb6, 0xcf, 0x38, 0x0b, 0x2d, 0x61, 0x59, 0x2b, 0x19, 0x81, 0x78, 0xf8, 0xa6, 0x21, 0x74,
	0xb1, 0x4c, 0x86, 0x5d, 0x2b, 0x71, 0xb6, 0x5f, 0x62, 0x63, 0x8d, 0x5d, 0x15, 0x72, 0x99, 0xe5,
	0x32, 0xa8, 0xca, 0x55, 0xd4, 0xb0, 0xc7, 0x1c, 0xc8, 0xce, 0x8e, 0xf6, 0xa1, 0x17, 0x54, 0x25,
	0xba, 0x07, 0x74, 0x00, 0xfd, 0xf7, 0x99, 0x49, 0x6e, 0xd7, 0x32, 0x75, 0x1d, 0x4a, 0xe0, 0x68,
	0x56, 0xb4, 0x43, 0x67, 0x2c, 0x61, 0xb0, 0xb0, 0x05, 0x84, 0xb6, 0x7c, 0xfa, 0x1c, 0x48, 0x95,
	0xaa, 0x58, 0xb6, 0x00, 0x1b, 0xb9, 0x1f, 0x41, 0x95, 0xaa, 0x0d, 0x85, 0xbe, 0x86, 0x5e, 0x53,
	0xae, 0x4d, 0x4b, 0xb8, 0xb7, 0x6b, 0xb5, 0x6d, 0x96, 0x6d, 0x9b, 0x65, 0x37, 0x46, 0xea, 0xc8,
	0xa2, 0xc7, 0x11, 0x0c, 0xc2, 0x75, 0x26, 0x8b, 0x72, 0xf3, 0x9b, 0x29, 0x1c, 0xb6, 0xbd, 0x0f,
	0x1d, 0xaf, 0x3b, 0x21, 0xfc, 0x64, 0x9f, 0x4e, 0x6b, 0x70, 0x56, 0xa4, 0x0a, 0xb3, 0xa2, 0x8c,
	0x36, 0xcc, 0x93, 0xaf, 0x00, 0x8f, 0x6d, 0x36, 0xa9, 0x6e, 0xe6, 0x1f, 0xe6, 0x57, 0x9f, 0xe6,
	0xee, 0x01, 0x7d, 0x0a, 0x24, 0x98, 0x2d, 0xe2, 0x33, 0xfe, 0x2e, 0x0e, 0xcf, 0xa7, 0xae, 0xb3,
	0x5d, 0xf0, 0x37, 0x6f, 0xed, 0xa2, 0xd3, 0x54, 0x12, 0x5e, 0x04, 0xe1, 0x45, 0xc0, 0x4f, 0xdd,
	0x2e, 0x7d, 0x06, 0xff, 0x6f, 0xa7, 0xf8, 0x72, 0xb6, 0x3c, 0x77, 0x7b, 0xd3, 0x2f, 0xe0, 0x09,
	0xcc, 0xf7, 0xbe, 0xc4, 0x94, 0xb4, 0x69, 0xae, 0x1b, 0xa3, 0x9f, 0xc9, 0xce, 0xe5, 0x57, 0xe7,
	0xf8, 0x23, 0x8f, 0x92, 0x9a, 0x85, 0x0d, 0xf1, 0xda, 0x12, 0x17, 0x8f, 0xe7, 0xdb, 0x43, 0x9b,
	0xed, 0xd5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x4e, 0xca, 0xdc, 0x14, 0x03, 0x00, 0x00,
}
