package tls

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

type Certificate struct {
	// TLS certificate in x509 format.
	Certificate []byte `protobuf:"bytes,1,opt,name=Certificate,proto3" json:"Certificate,omitempty"`
	// TLS key in x509 format.
	Key []byte `protobuf:"bytes,2,opt,name=Key,proto3" json:"Key,omitempty"`
}

func (m *Certificate) Reset()                    { *m = Certificate{} }
func (m *Certificate) String() string            { return proto.CompactTextString(m) }
func (*Certificate) ProtoMessage()               {}
func (*Certificate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Certificate) GetCertificate() []byte {
	if m != nil {
		return m.Certificate
	}
	return nil
}

func (m *Certificate) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

type Config struct {
	// Whether or not to allow self-signed certificates.
	AllowInsecure bool `protobuf:"varint,1,opt,name=allow_insecure,json=allowInsecure" json:"allow_insecure,omitempty"`
	// List of certificates to be served on server.
	Certificate []*Certificate `protobuf:"bytes,2,rep,name=certificate" json:"certificate,omitempty"`
	// Override server name.
	ServerName string `protobuf:"bytes,3,opt,name=server_name,json=serverName" json:"server_name,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Config) GetAllowInsecure() bool {
	if m != nil {
		return m.AllowInsecure
	}
	return false
}

func (m *Config) GetCertificate() []*Certificate {
	if m != nil {
		return m.Certificate
	}
	return nil
}

func (m *Config) GetServerName() string {
	if m != nil {
		return m.ServerName
	}
	return ""
}

func init() {
	proto.RegisterType((*Certificate)(nil), "v2ray.core.transport.internet.tls.Certificate")
	proto.RegisterType((*Config)(nil), "v2ray.core.transport.internet.tls.Config")
}

func init() { proto.RegisterFile("v2ray.com/core/transport/internet/tls/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x49, 0x17, 0x8a, 0x66, 0x55, 0x24, 0xa7, 0xbd, 0xb9, 0x2d, 0x14, 0xf6, 0x94, 0xc0,
	0xfa, 0x0b, 0x74, 0x4f, 0x45, 0x90, 0x12, 0x8a, 0x07, 0x2f, 0x25, 0x86, 0xa9, 0x04, 0xb2, 0x49,
	0x99, 0x8c, 0x95, 0xfd, 0x3b, 0x1e, 0xfd, 0x95, 0xd2, 0xad, 0x5b, 0xb6, 0xa7, 0xde, 0x92, 0x37,
	0xdf, 0xbc, 0xf7, 0x18, 0x5e, 0xef, 0x6b, 0x34, 0x9d, 0xb4, 0xb1, 0x55, 0x36, 0x22, 0x28, 0x42,
	0x13, 0xd2, 0x2e, 0x22, 0x29, 0x17, 0x08, 0x30, 0x00, 0x29, 0xf2, 0x49, 0xd9, 0x18, 0xb6, 0xee,
	0x53, 0xee, 0x30, 0x52, 0x14, 0xb3, 0x61, 0x07, 0x41, 0x9e, 0x78, 0x39, 0xf0, 0x92, 0x7c, 0x9a,
	0x3f, 0xf1, 0xbc, 0x01, 0x24, 0xb7, 0x75, 0xd6, 0x10, 0x88, 0xf2, 0xec, 0x5b, 0xb0, 0x92, 0x55,
	0x37, 0xfa, 0x8c, 0xb8, 0xe7, 0xd9, 0x0b, 0x74, 0xc5, 0xa4, 0x9f, 0x1c, 0x9e, 0xf3, 0x1f, 0xc6,
	0xa7, 0x4d, 0x1f, 0x2b, 0x16, 0xfc, 0xce, 0x78, 0x1f, 0xbf, 0x37, 0x2e, 0x24, 0xb0, 0x5f, 0x78,
	0x74, 0xb8, 0xd2, 0xb7, 0xbd, 0xba, 0xfc, 0x17, 0xc5, 0x8a, 0xe7, 0x76, 0x94, 0x32, 0x29, 0xb3,
	0x2a, 0xaf, 0xa5, 0xbc, 0xd8, 0x56, 0x8e, 0x8a, 0xe8, 0xb1, 0x85, 0x78, 0xe0, 0x79, 0x02, 0xdc,
	0x03, 0x6e, 0x82, 0x69, 0xa1, 0xc8, 0x4a, 0x56, 0x5d, 0x6b, 0x7e, 0x94, 0x5e, 0x4d, 0x0b, 0xcf,
	0x9a, 0x2f, 0x6c, 0x6c, 0x2f, 0x47, 0xac, 0xd8, 0x7b, 0x46, 0x3e, 0xfd, 0x4e, 0x66, 0x6f, 0xb5,
	0x36, 0x9d, 0x6c, 0x0e, 0xe8, 0xfa, 0x84, 0x2e, 0x07, 0x74, 0xed, 0xd3, 0xc7, 0xb4, 0xbf, 0xf2,
	0xe3, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x77, 0x0b, 0x8b, 0x74, 0x9b, 0x01, 0x00, 0x00,
}
