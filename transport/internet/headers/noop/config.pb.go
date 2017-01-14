package noop

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

type Config struct {
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ConnectionConfig struct {
}

func (m *ConnectionConfig) Reset()                    { *m = ConnectionConfig{} }
func (m *ConnectionConfig) String() string            { return proto.CompactTextString(m) }
func (*ConnectionConfig) ProtoMessage()               {}
func (*ConnectionConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*Config)(nil), "v2ray.core.transport.internet.headers.noop.Config")
	proto.RegisterType((*ConnectionConfig)(nil), "v2ray.core.transport.internet.headers.noop.ConnectionConfig")
}

func init() {
	proto.RegisterFile("v2ray.com/core/transport/internet/headers/noop/config.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0xcf, 0xb1, 0xea, 0xc2, 0x40,
	0x0c, 0xc7, 0x71, 0xf8, 0xf3, 0xa7, 0xc8, 0xb9, 0x48, 0x1f, 0xa1, 0x63, 0x87, 0x1c, 0xd4, 0xd1,
	0xad, 0x5d, 0x74, 0x11, 0x11, 0x71, 0x70, 0x3b, 0xcf, 0xa8, 0x1d, 0x4c, 0x8e, 0x34, 0x08, 0x7d,
	0x25, 0x9f, 0x52, 0xae, 0xed, 0x75, 0x76, 0xfd, 0xc1, 0x27, 0x7c, 0x63, 0x36, 0xef, 0x4a, 0x5c,
	0x0f, 0x9e, 0x5f, 0xd6, 0xb3, 0xa0, 0x55, 0x71, 0xd4, 0x05, 0x16, 0xb5, 0x2d, 0x29, 0x0a, 0xa1,
	0xda, 0x27, 0xba, 0x1b, 0x4a, 0x67, 0x89, 0x39, 0x58, 0xcf, 0x74, 0x6f, 0x1f, 0x10, 0x84, 0x95,
	0xf3, 0x32, 0x61, 0x41, 0x98, 0x21, 0x24, 0x08, 0x13, 0x84, 0x08, 0x8b, 0x85, 0xc9, 0x9a, 0xc1,
	0x16, 0xb9, 0x59, 0x35, 0x4c, 0x84, 0x5e, 0x5b, 0xa6, 0x71, 0xab, 0x83, 0x89, 0x09, 0xf0, 0xfb,
	0xbd, 0x7a, 0x39, 0xca, 0x43, 0x0c, 0xb9, 0xfc, 0xc7, 0xe9, 0xf3, 0x57, 0x9e, 0xab, 0xa3, 0xeb,
	0xa1, 0x89, 0xfe, 0x34, 0xfb, 0x5d, 0xf2, 0xdb, 0xc9, 0xef, 0x99, 0xc3, 0x35, 0x1b, 0x5e, 0x58,
	0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x55, 0x14, 0x69, 0x01, 0x01, 0x00, 0x00,
}
