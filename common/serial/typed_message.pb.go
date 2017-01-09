package serial

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

// Serialized proto message along with its type name.
type TypedMessage struct {
	// The name of the message type, retrieved from protobuf API.
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	// Serialized proto message.
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *TypedMessage) Reset()                    { *m = TypedMessage{} }
func (m *TypedMessage) String() string            { return proto.CompactTextString(m) }
func (*TypedMessage) ProtoMessage()               {}
func (*TypedMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TypedMessage) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *TypedMessage) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*TypedMessage)(nil), "v2ray.core.common.serial.TypedMessage")
}

func init() { proto.RegisterFile("v2ray.com/core/common/serial/typed_message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x32, 0x28, 0x33, 0x2a, 0x4a,
	0xac, 0xd4, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x4f, 0xce, 0xcf, 0xcd, 0xcd,
	0xcf, 0xd3, 0x2f, 0x4e, 0x2d, 0xca, 0x4c, 0xcc, 0xd1, 0x2f, 0xa9, 0x2c, 0x48, 0x4d, 0x89, 0xcf,
	0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x80, 0xe9,
	0x28, 0x4a, 0xd5, 0x83, 0xa8, 0xd6, 0x83, 0xa8, 0x56, 0xb2, 0xe0, 0xe2, 0x09, 0x01, 0x69, 0xf0,
	0x85, 0xa8, 0x17, 0x12, 0xe2, 0x62, 0x01, 0x19, 0x20, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04,
	0x66, 0x0b, 0x89, 0x70, 0xb1, 0x96, 0x25, 0xe6, 0x94, 0xa6, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0xf0,
	0x04, 0x41, 0x38, 0x4e, 0x21, 0x5c, 0x32, 0xc9, 0xf9, 0xb9, 0x7a, 0xb8, 0x4c, 0x76, 0x12, 0x44,
	0x36, 0x37, 0x00, 0xe4, 0x8c, 0x28, 0x36, 0x88, 0xd4, 0x2a, 0x26, 0x89, 0x30, 0xa3, 0xa0, 0xc4,
	0x4a, 0x3d, 0x67, 0x90, 0x2e, 0x67, 0x88, 0xae, 0x60, 0xb0, 0x54, 0x12, 0x1b, 0xd8, 0xc1, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x74, 0x15, 0xc3, 0xe2, 0xe4, 0x00, 0x00, 0x00,
}
