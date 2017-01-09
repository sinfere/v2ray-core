package tcp

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v2ray_core_common_serial "v2ray.com/core/common/serial"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ConnectionReuse struct {
	Enable bool `protobuf:"varint,1,opt,name=enable" json:"enable,omitempty"`
}

func (m *ConnectionReuse) Reset()                    { *m = ConnectionReuse{} }
func (m *ConnectionReuse) String() string            { return proto.CompactTextString(m) }
func (*ConnectionReuse) ProtoMessage()               {}
func (*ConnectionReuse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ConnectionReuse) GetEnable() bool {
	if m != nil {
		return m.Enable
	}
	return false
}

type Config struct {
	ConnectionReuse *ConnectionReuse                       `protobuf:"bytes,1,opt,name=connection_reuse,json=connectionReuse" json:"connection_reuse,omitempty"`
	HeaderSettings  *v2ray_core_common_serial.TypedMessage `protobuf:"bytes,2,opt,name=header_settings,json=headerSettings" json:"header_settings,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Config) GetConnectionReuse() *ConnectionReuse {
	if m != nil {
		return m.ConnectionReuse
	}
	return nil
}

func (m *Config) GetHeaderSettings() *v2ray_core_common_serial.TypedMessage {
	if m != nil {
		return m.HeaderSettings
	}
	return nil
}

func init() {
	proto.RegisterType((*ConnectionReuse)(nil), "v2ray.core.transport.internet.tcp.ConnectionReuse")
	proto.RegisterType((*Config)(nil), "v2ray.core.transport.internet.tcp.Config")
}

func init() { proto.RegisterFile("v2ray.com/core/transport/internet/tcp/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x69, 0x85, 0x22, 0x5b, 0x30, 0x92, 0x83, 0x14, 0x4f, 0xb6, 0xa0, 0xe8, 0x65, 0x57,
	0xe2, 0x1b, 0x34, 0x27, 0x0f, 0xa2, 0xc4, 0xe0, 0x41, 0x28, 0x61, 0x3b, 0x1d, 0x63, 0xa0, 0xd9,
	0x59, 0x76, 0x47, 0x21, 0xaf, 0xe4, 0x13, 0xf8, 0x78, 0x92, 0x6c, 0x13, 0x24, 0x97, 0x1e, 0x17,
	0xfe, 0xef, 0xdf, 0xff, 0x1b, 0x91, 0x7c, 0x27, 0x4e, 0x37, 0x12, 0xa8, 0x56, 0x40, 0x0e, 0x15,
	0x3b, 0x6d, 0xbc, 0x25, 0xc7, 0xaa, 0x32, 0x8c, 0xce, 0x20, 0x2b, 0x06, 0xab, 0x80, 0xcc, 0x47,
	0x55, 0x4a, 0xeb, 0x88, 0x29, 0x5e, 0xf6, 0x8c, 0x43, 0x39, 0xe4, 0x65, 0x9f, 0x97, 0x0c, 0xf6,
	0xf2, 0x7e, 0x54, 0x0b, 0x54, 0xd7, 0x64, 0x94, 0x47, 0x57, 0xe9, 0xbd, 0xe2, 0xc6, 0xe2, 0xae,
	0xa8, 0xd1, 0x7b, 0x5d, 0x62, 0x28, 0x5d, 0xdd, 0x89, 0x28, 0x25, 0x63, 0x10, 0xb8, 0x22, 0x93,
	0xe1, 0x97, 0xc7, 0xf8, 0x42, 0xcc, 0xd0, 0xe8, 0xed, 0x1e, 0x17, 0x93, 0xab, 0xc9, 0xed, 0x69,
	0x76, 0x78, 0xad, 0x7e, 0x27, 0x62, 0x96, 0x76, 0x83, 0xe2, 0x8d, 0x38, 0x87, 0x81, 0x2a, 0x5c,
	0x8b, 0x75, 0xe1, 0x79, 0x92, 0xc8, 0xa3, 0x2b, 0xe5, 0xe8, 0xc3, 0x2c, 0x82, 0xd1, 0x82, 0x67,
	0x11, 0x7d, 0xa2, 0xde, 0xa1, 0x2b, 0x3c, 0x32, 0x57, 0xa6, 0xf4, 0x8b, 0x69, 0xd7, 0x7e, 0xf3,
	0xbf, 0x3d, 0xc8, 0xc9, 0x20, 0x27, 0xf3, 0x56, 0xee, 0x29, 0xb8, 0x65, 0x67, 0x01, 0x7f, 0x3d,
	0xd0, 0xeb, 0x8d, 0xb8, 0x06, 0xaa, 0x8f, 0x4f, 0x5b, 0xcf, 0x83, 0xe0, 0x4b, 0x7b, 0x9b, 0xf7,
	0x13, 0x06, 0xfb, 0x33, 0x5d, 0xbe, 0x25, 0x99, 0x6e, 0x64, 0xda, 0x52, 0xf9, 0x40, 0x3d, 0xf6,
	0x54, 0x0e, 0x76, 0x3b, 0xeb, 0x6e, 0xf9, 0xf0, 0x17, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x15, 0xa3,
	0xb6, 0xd6, 0x01, 0x00, 0x00,
}
