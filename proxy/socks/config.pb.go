package socks

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v2ray_core_common_net "v2ray.com/core/common/net"
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

type AuthType int32

const (
	AuthType_NO_AUTH  AuthType = 0
	AuthType_PASSWORD AuthType = 1
)

var AuthType_name = map[int32]string{
	0: "NO_AUTH",
	1: "PASSWORD",
}
var AuthType_value = map[string]int32{
	"NO_AUTH":  0,
	"PASSWORD": 1,
}

func (x AuthType) String() string {
	return proto.EnumName(AuthType_name, int32(x))
}
func (AuthType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Account struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Account) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Account) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type ServerConfig struct {
	AuthType   AuthType                          `protobuf:"varint,1,opt,name=auth_type,json=authType,enum=v2ray.core.proxy.socks.AuthType" json:"auth_type,omitempty"`
	Accounts   map[string]string                 `protobuf:"bytes,2,rep,name=accounts" json:"accounts,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Address    *v2ray_core_common_net.IPOrDomain `protobuf:"bytes,3,opt,name=address" json:"address,omitempty"`
	UdpEnabled bool                              `protobuf:"varint,4,opt,name=udp_enabled,json=udpEnabled" json:"udp_enabled,omitempty"`
	Timeout    uint32                            `protobuf:"varint,5,opt,name=timeout" json:"timeout,omitempty"`
}

func (m *ServerConfig) Reset()                    { *m = ServerConfig{} }
func (m *ServerConfig) String() string            { return proto.CompactTextString(m) }
func (*ServerConfig) ProtoMessage()               {}
func (*ServerConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ServerConfig) GetAuthType() AuthType {
	if m != nil {
		return m.AuthType
	}
	return AuthType_NO_AUTH
}

func (m *ServerConfig) GetAccounts() map[string]string {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func (m *ServerConfig) GetAddress() *v2ray_core_common_net.IPOrDomain {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *ServerConfig) GetUdpEnabled() bool {
	if m != nil {
		return m.UdpEnabled
	}
	return false
}

func (m *ServerConfig) GetTimeout() uint32 {
	if m != nil {
		return m.Timeout
	}
	return 0
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
	proto.RegisterType((*Account)(nil), "v2ray.core.proxy.socks.Account")
	proto.RegisterType((*ServerConfig)(nil), "v2ray.core.proxy.socks.ServerConfig")
	proto.RegisterType((*ClientConfig)(nil), "v2ray.core.proxy.socks.ClientConfig")
	proto.RegisterEnum("v2ray.core.proxy.socks.AuthType", AuthType_name, AuthType_value)
}

func init() { proto.RegisterFile("v2ray.com/core/proxy/socks/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x91, 0xd1, 0x6e, 0xd3, 0x3e,
	0x14, 0xc6, 0xff, 0x49, 0xff, 0x5d, 0xb3, 0xd3, 0x0e, 0x55, 0x16, 0x9a, 0xa2, 0xdc, 0x10, 0x2a,
	0x21, 0xa2, 0x5d, 0x38, 0x28, 0xdc, 0x20, 0x26, 0x90, 0xb2, 0xae, 0x12, 0xdc, 0xac, 0x95, 0x3b,
	0x40, 0xe2, 0xa6, 0xf2, 0x1c, 0xc3, 0xa2, 0x35, 0x76, 0x64, 0x3b, 0x85, 0xbc, 0x12, 0x8f, 0xc3,
	0x13, 0xa1, 0xda, 0xc9, 0x34, 0x50, 0x77, 0x77, 0x4e, 0xce, 0x77, 0xbe, 0x9c, 0xef, 0x67, 0x78,
	0xb9, 0xcb, 0x14, 0x6d, 0x31, 0x93, 0x55, 0xca, 0xa4, 0xe2, 0x69, 0xad, 0xe4, 0xcf, 0x36, 0xd5,
	0x92, 0xdd, 0xe9, 0x94, 0x49, 0xf1, 0xad, 0xfc, 0x8e, 0x6b, 0x25, 0x8d, 0x44, 0xa7, 0xbd, 0x50,
	0x71, 0x6c, 0x45, 0xd8, 0x8a, 0xa2, 0x7f, 0x0d, 0x98, 0xac, 0x2a, 0x29, 0x52, 0xc1, 0x4d, 0x4a,
	0x8b, 0x42, 0x71, 0xad, 0x9d, 0x41, 0xf4, 0xea, 0xb0, 0xd0, 0x0e, 0x99, 0xdc, 0xa6, 0x9a, 0xab,
	0x1d, 0x57, 0x1b, 0x5d, 0x73, 0xe6, 0x36, 0x66, 0x39, 0x8c, 0x72, 0xc6, 0x64, 0x23, 0x0c, 0x8a,
	0x20, 0x68, 0x34, 0x57, 0x82, 0x56, 0x3c, 0xf4, 0x62, 0x2f, 0x39, 0x26, 0xf7, 0xfd, 0x7e, 0x56,
	0x53, 0xad, 0x7f, 0x48, 0x55, 0x84, 0xbe, 0x9b, 0xf5, 0xfd, 0xec, 0xb7, 0x0f, 0x93, 0xb5, 0x35,
	0x9e, 0xdb, 0x30, 0xe8, 0x1d, 0x1c, 0xd3, 0xc6, 0xdc, 0x6e, 0x4c, 0x5b, 0x3b, 0xa7, 0x27, 0x59,
	0x8c, 0x0f, 0x47, 0xc3, 0x79, 0x63, 0x6e, 0xaf, 0xdb, 0x9a, 0x93, 0x80, 0x76, 0x15, 0xba, 0x82,
	0x80, 0xba, 0x93, 0x74, 0xe8, 0xc7, 0x83, 0x64, 0x9c, 0x65, 0x8f, 0x6d, 0x3f, 0xfc, 0x2d, 0xee,
	0x72, 0xe8, 0x85, 0x30, 0xaa, 0x25, 0xf7, 0x1e, 0xe8, 0x1c, 0x46, 0x1d, 0xa5, 0x70, 0x10, 0x7b,
	0xc9, 0x38, 0x7b, 0xfe, 0xd0, 0xce, 0x21, 0xc2, 0x82, 0x1b, 0xfc, 0x71, 0xb5, 0x54, 0x97, 0xb2,
	0xa2, 0xa5, 0x20, 0xfd, 0x06, 0x7a, 0x06, 0xe3, 0xa6, 0xa8, 0x37, 0x5c, 0xd0, 0x9b, 0x2d, 0x2f,
	0xc2, 0xff, 0x63, 0x2f, 0x09, 0x08, 0x34, 0x45, 0xbd, 0x70, 0x5f, 0x50, 0x08, 0x23, 0x53, 0x56,
	0x5c, 0x36, 0x26, 0x1c, 0xc6, 0x5e, 0x72, 0x42, 0xfa, 0x36, 0x3a, 0x87, 0x93, 0xbf, 0x4e, 0x42,
	0x53, 0x18, 0xdc, 0xf1, 0xb6, 0x63, 0xbb, 0x2f, 0xd1, 0x53, 0x18, 0xee, 0xe8, 0xb6, 0xe1, 0x1d,
	0x53, 0xd7, 0xbc, 0xf5, 0xdf, 0x78, 0x33, 0x02, 0x93, 0xf9, 0xb6, 0xe4, 0xc2, 0x74, 0x4c, 0x2f,
	0xe0, 0xc8, 0x3d, 0x5e, 0xe8, 0x59, 0x24, 0x67, 0x07, 0x32, 0xf4, 0xcf, 0xdc, 0x61, 0x59, 0x88,
	0xa2, 0x96, 0xa5, 0x30, 0xa4, 0xdb, 0x3c, 0x7b, 0x01, 0x41, 0x8f, 0x1b, 0x8d, 0x61, 0x74, 0xb5,
	0xdc, 0xe4, 0x9f, 0xae, 0x3f, 0x4c, 0xff, 0x43, 0x13, 0x08, 0x56, 0xf9, 0x7a, 0xfd, 0x65, 0x49,
	0x2e, 0xa7, 0xde, 0xc5, 0x7b, 0x88, 0x98, 0xac, 0x1e, 0x41, 0xbe, 0xf2, 0xbe, 0x0e, 0x6d, 0xf1,
	0xcb, 0x3f, 0xfd, 0x9c, 0x11, 0xda, 0xe2, 0xf9, 0x5e, 0xb1, 0xb2, 0x8a, 0xf5, 0x7e, 0x70, 0x73,
	0x64, 0xef, 0x78, 0xfd, 0x27, 0x00, 0x00, 0xff, 0xff, 0x57, 0x01, 0x54, 0x95, 0xf7, 0x02, 0x00,
	0x00,
}
