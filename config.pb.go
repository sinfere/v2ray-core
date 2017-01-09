package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v2ray_core_common_serial "v2ray.com/core/common/serial"
import v2ray_core_common_net "v2ray.com/core/common/net"
import v2ray_core_common_net1 "v2ray.com/core/common/net"
import v2ray_core_common_log "v2ray.com/core/common/log"
import v2ray_core_transport_internet "v2ray.com/core/transport/internet"
import v2ray_core_transport "v2ray.com/core/transport"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Configuration serialization format.
type ConfigFormat int32

const (
	ConfigFormat_Protobuf ConfigFormat = 0
	ConfigFormat_JSON     ConfigFormat = 1
)

var ConfigFormat_name = map[int32]string{
	0: "Protobuf",
	1: "JSON",
}
var ConfigFormat_value = map[string]int32{
	"Protobuf": 0,
	"JSON":     1,
}

func (x ConfigFormat) String() string {
	return proto.EnumName(ConfigFormat_name, int32(x))
}
func (ConfigFormat) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type AllocationStrategy_Type int32

const (
	// Always allocate all connection handlers.
	AllocationStrategy_Always AllocationStrategy_Type = 0
	// Randomly allocate specific range of handlers.
	AllocationStrategy_Random AllocationStrategy_Type = 1
	// External. Not supported yet.
	AllocationStrategy_External AllocationStrategy_Type = 2
)

var AllocationStrategy_Type_name = map[int32]string{
	0: "Always",
	1: "Random",
	2: "External",
}
var AllocationStrategy_Type_value = map[string]int32{
	"Always":   0,
	"Random":   1,
	"External": 2,
}

func (x AllocationStrategy_Type) String() string {
	return proto.EnumName(AllocationStrategy_Type_name, int32(x))
}
func (AllocationStrategy_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type AllocationStrategyConcurrency struct {
	Value uint32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *AllocationStrategyConcurrency) Reset()                    { *m = AllocationStrategyConcurrency{} }
func (m *AllocationStrategyConcurrency) String() string            { return proto.CompactTextString(m) }
func (*AllocationStrategyConcurrency) ProtoMessage()               {}
func (*AllocationStrategyConcurrency) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AllocationStrategyConcurrency) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type AllocationStrategyRefresh struct {
	Value uint32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *AllocationStrategyRefresh) Reset()                    { *m = AllocationStrategyRefresh{} }
func (m *AllocationStrategyRefresh) String() string            { return proto.CompactTextString(m) }
func (*AllocationStrategyRefresh) ProtoMessage()               {}
func (*AllocationStrategyRefresh) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AllocationStrategyRefresh) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type AllocationStrategy struct {
	Type AllocationStrategy_Type `protobuf:"varint,1,opt,name=type,enum=v2ray.core.AllocationStrategy_Type" json:"type,omitempty"`
	// Number of handlers (ports) running in parallel.
	// Default value is 3 if unset.
	Concurrency *AllocationStrategyConcurrency `protobuf:"bytes,2,opt,name=concurrency" json:"concurrency,omitempty"`
	// Number of minutes before a handler is regenerated.
	// Default value is 5 if unset.
	Refresh *AllocationStrategyRefresh `protobuf:"bytes,3,opt,name=refresh" json:"refresh,omitempty"`
}

func (m *AllocationStrategy) Reset()                    { *m = AllocationStrategy{} }
func (m *AllocationStrategy) String() string            { return proto.CompactTextString(m) }
func (*AllocationStrategy) ProtoMessage()               {}
func (*AllocationStrategy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AllocationStrategy) GetType() AllocationStrategy_Type {
	if m != nil {
		return m.Type
	}
	return AllocationStrategy_Always
}

func (m *AllocationStrategy) GetConcurrency() *AllocationStrategyConcurrency {
	if m != nil {
		return m.Concurrency
	}
	return nil
}

func (m *AllocationStrategy) GetRefresh() *AllocationStrategyRefresh {
	if m != nil {
		return m.Refresh
	}
	return nil
}

// Config for an inbound connection handler.
type InboundConnectionConfig struct {
	// Protocol specific settings. Must be one of the supported protocols.
	Settings *v2ray_core_common_serial.TypedMessage `protobuf:"bytes,1,opt,name=settings" json:"settings,omitempty"`
	// Range of port number to run on. Both inclusive.
	PortRange *v2ray_core_common_net.PortRange `protobuf:"bytes,2,opt,name=port_range,json=portRange" json:"port_range,omitempty"`
	// IP address to listen on. 0.0.0.0 if unset.
	ListenOn *v2ray_core_common_net1.IPOrDomain `protobuf:"bytes,3,opt,name=listen_on,json=listenOn" json:"listen_on,omitempty"`
	// Tag of this handler.
	Tag                    string                                      `protobuf:"bytes,4,opt,name=tag" json:"tag,omitempty"`
	AllocationStrategy     *AllocationStrategy                         `protobuf:"bytes,5,opt,name=allocation_strategy,json=allocationStrategy" json:"allocation_strategy,omitempty"`
	StreamSettings         *v2ray_core_transport_internet.StreamConfig `protobuf:"bytes,6,opt,name=stream_settings,json=streamSettings" json:"stream_settings,omitempty"`
	AllowPassiveConnection bool                                        `protobuf:"varint,7,opt,name=allow_passive_connection,json=allowPassiveConnection" json:"allow_passive_connection,omitempty"`
}

func (m *InboundConnectionConfig) Reset()                    { *m = InboundConnectionConfig{} }
func (m *InboundConnectionConfig) String() string            { return proto.CompactTextString(m) }
func (*InboundConnectionConfig) ProtoMessage()               {}
func (*InboundConnectionConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *InboundConnectionConfig) GetSettings() *v2ray_core_common_serial.TypedMessage {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (m *InboundConnectionConfig) GetPortRange() *v2ray_core_common_net.PortRange {
	if m != nil {
		return m.PortRange
	}
	return nil
}

func (m *InboundConnectionConfig) GetListenOn() *v2ray_core_common_net1.IPOrDomain {
	if m != nil {
		return m.ListenOn
	}
	return nil
}

func (m *InboundConnectionConfig) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *InboundConnectionConfig) GetAllocationStrategy() *AllocationStrategy {
	if m != nil {
		return m.AllocationStrategy
	}
	return nil
}

func (m *InboundConnectionConfig) GetStreamSettings() *v2ray_core_transport_internet.StreamConfig {
	if m != nil {
		return m.StreamSettings
	}
	return nil
}

func (m *InboundConnectionConfig) GetAllowPassiveConnection() bool {
	if m != nil {
		return m.AllowPassiveConnection
	}
	return false
}

// Config for an outbound connection handler.
type OutboundConnectionConfig struct {
	Settings *v2ray_core_common_serial.TypedMessage `protobuf:"bytes,1,opt,name=settings" json:"settings,omitempty"`
	// IP address to send data through. 0.0.0.0 if unset.
	SendThrough    *v2ray_core_common_net1.IPOrDomain          `protobuf:"bytes,2,opt,name=send_through,json=sendThrough" json:"send_through,omitempty"`
	StreamSettings *v2ray_core_transport_internet.StreamConfig `protobuf:"bytes,3,opt,name=stream_settings,json=streamSettings" json:"stream_settings,omitempty"`
	ProxySettings  *v2ray_core_transport_internet.ProxyConfig  `protobuf:"bytes,5,opt,name=proxy_settings,json=proxySettings" json:"proxy_settings,omitempty"`
	Tag            string                                      `protobuf:"bytes,4,opt,name=tag" json:"tag,omitempty"`
}

func (m *OutboundConnectionConfig) Reset()                    { *m = OutboundConnectionConfig{} }
func (m *OutboundConnectionConfig) String() string            { return proto.CompactTextString(m) }
func (*OutboundConnectionConfig) ProtoMessage()               {}
func (*OutboundConnectionConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *OutboundConnectionConfig) GetSettings() *v2ray_core_common_serial.TypedMessage {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (m *OutboundConnectionConfig) GetSendThrough() *v2ray_core_common_net1.IPOrDomain {
	if m != nil {
		return m.SendThrough
	}
	return nil
}

func (m *OutboundConnectionConfig) GetStreamSettings() *v2ray_core_transport_internet.StreamConfig {
	if m != nil {
		return m.StreamSettings
	}
	return nil
}

func (m *OutboundConnectionConfig) GetProxySettings() *v2ray_core_transport_internet.ProxyConfig {
	if m != nil {
		return m.ProxySettings
	}
	return nil
}

func (m *OutboundConnectionConfig) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type Config struct {
	// Inbound handler configurations. Must have at least one item.
	Inbound []*InboundConnectionConfig `protobuf:"bytes,1,rep,name=inbound" json:"inbound,omitempty"`
	// Outbound handler configurations. Must have at least one item. The first item is used as default for routing.
	Outbound []*OutboundConnectionConfig   `protobuf:"bytes,2,rep,name=outbound" json:"outbound,omitempty"`
	Log      *v2ray_core_common_log.Config `protobuf:"bytes,3,opt,name=log" json:"log,omitempty"`
	// App configuration. Must be one in the app directory.
	App       []*v2ray_core_common_serial.TypedMessage `protobuf:"bytes,4,rep,name=app" json:"app,omitempty"`
	Transport *v2ray_core_transport.Config             `protobuf:"bytes,5,opt,name=transport" json:"transport,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Config) GetInbound() []*InboundConnectionConfig {
	if m != nil {
		return m.Inbound
	}
	return nil
}

func (m *Config) GetOutbound() []*OutboundConnectionConfig {
	if m != nil {
		return m.Outbound
	}
	return nil
}

func (m *Config) GetLog() *v2ray_core_common_log.Config {
	if m != nil {
		return m.Log
	}
	return nil
}

func (m *Config) GetApp() []*v2ray_core_common_serial.TypedMessage {
	if m != nil {
		return m.App
	}
	return nil
}

func (m *Config) GetTransport() *v2ray_core_transport.Config {
	if m != nil {
		return m.Transport
	}
	return nil
}

func init() {
	proto.RegisterType((*AllocationStrategyConcurrency)(nil), "v2ray.core.AllocationStrategyConcurrency")
	proto.RegisterType((*AllocationStrategyRefresh)(nil), "v2ray.core.AllocationStrategyRefresh")
	proto.RegisterType((*AllocationStrategy)(nil), "v2ray.core.AllocationStrategy")
	proto.RegisterType((*InboundConnectionConfig)(nil), "v2ray.core.InboundConnectionConfig")
	proto.RegisterType((*OutboundConnectionConfig)(nil), "v2ray.core.OutboundConnectionConfig")
	proto.RegisterType((*Config)(nil), "v2ray.core.Config")
	proto.RegisterEnum("v2ray.core.ConfigFormat", ConfigFormat_name, ConfigFormat_value)
	proto.RegisterEnum("v2ray.core.AllocationStrategy_Type", AllocationStrategy_Type_name, AllocationStrategy_Type_value)
}

func init() { proto.RegisterFile("v2ray.com/core/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 745 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x95, 0xd1, 0x6e, 0xd3, 0x3a,
	0x1c, 0xc6, 0x97, 0xb6, 0xeb, 0xda, 0x7f, 0xb7, 0x9e, 0xca, 0xe7, 0xe8, 0x9c, 0x9c, 0xc1, 0x50,
	0x29, 0xdb, 0x28, 0x03, 0xa5, 0xa3, 0x08, 0x31, 0x21, 0xc1, 0xd8, 0x3a, 0x90, 0x06, 0x82, 0x16,
	0x77, 0xe2, 0x82, 0x9b, 0xc8, 0x4b, 0xbd, 0x2c, 0x52, 0x62, 0x47, 0xb6, 0xbb, 0x2d, 0x8f, 0xc0,
	0xab, 0x70, 0xc5, 0xab, 0xf0, 0x04, 0xbc, 0x0a, 0x72, 0x92, 0xa6, 0x1d, 0x6d, 0xb7, 0x49, 0x88,
	0xbb, 0x34, 0xf9, 0x7e, 0x9f, 0x9d, 0xef, 0xfb, 0xc7, 0x85, 0x5b, 0x67, 0x6d, 0x41, 0x22, 0xcb,
	0xe1, 0x41, 0xcb, 0xe1, 0x82, 0xb6, 0x1c, 0xce, 0x4e, 0x3c, 0xd7, 0x0a, 0x05, 0x57, 0x1c, 0xc1,
	0xe8, 0xa1, 0xa0, 0xab, 0xdb, 0x53, 0xc2, 0x20, 0xe0, 0xac, 0x25, 0xa9, 0xf0, 0x88, 0xdf, 0x52,
	0x51, 0x48, 0x07, 0x76, 0x40, 0xa5, 0x24, 0x2e, 0x4d, 0xe8, 0xd5, 0xf5, 0xd9, 0x04, 0xa3, 0xaa,
	0x15, 0x72, 0xa1, 0x52, 0xd5, 0xfd, 0xf9, 0x2a, 0x32, 0x18, 0x08, 0x2a, 0x65, 0x2a, 0xdc, 0x9c,
	0x2d, 0xf4, 0xb9, 0x7b, 0x69, 0xd3, 0xab, 0xd6, 0x2f, 0x3a, 0x25, 0x08, 0x93, 0x7a, 0xc1, 0x96,
	0xc7, 0x14, 0x15, 0xda, 0xf8, 0x92, 0x7e, 0x63, 0xae, 0x7e, 0x52, 0xd6, 0x78, 0x0a, 0x6b, 0x7b,
	0xbe, 0xcf, 0x1d, 0xa2, 0x3c, 0xce, 0xfa, 0x4a, 0x10, 0x45, 0xdd, 0xa8, 0xc3, 0x99, 0x33, 0x14,
	0x82, 0x32, 0x27, 0x42, 0xff, 0xc0, 0xe2, 0x19, 0xf1, 0x87, 0xd4, 0x34, 0xea, 0x46, 0x73, 0x05,
	0x27, 0x3f, 0x1a, 0x8f, 0xe1, 0xff, 0x69, 0x0c, 0xd3, 0x13, 0x41, 0xe5, 0xe9, 0x1c, 0xe4, 0x4b,
	0x0e, 0xd0, 0x34, 0x83, 0x9e, 0x41, 0x41, 0xa7, 0x1c, 0x6b, 0xab, 0xed, 0x7b, 0xd6, 0xb8, 0x1b,
	0x6b, 0x5a, 0x6d, 0x1d, 0x45, 0x21, 0xc5, 0x31, 0x80, 0xde, 0x41, 0xc5, 0x19, 0xef, 0xd3, 0xcc,
	0xd5, 0x8d, 0x66, 0xa5, 0xfd, 0xe0, 0x6a, 0x7e, 0xe2, 0xc5, 0xf0, 0x24, 0x8d, 0x76, 0x61, 0x49,
	0x24, 0xbb, 0x37, 0xf3, 0xb1, 0xd1, 0xc6, 0xd5, 0x46, 0xe9, 0xab, 0xe2, 0x11, 0xd5, 0x78, 0x04,
	0x05, 0xbd, 0x37, 0x04, 0x50, 0xdc, 0xf3, 0xcf, 0x49, 0x24, 0x6b, 0x0b, 0xfa, 0x1a, 0x13, 0x36,
	0xe0, 0x41, 0xcd, 0x40, 0xcb, 0x50, 0x7a, 0x7d, 0xa1, 0x7b, 0x22, 0x7e, 0x2d, 0xd7, 0xf8, 0x9e,
	0x87, 0xff, 0x0e, 0xd9, 0x31, 0x1f, 0xb2, 0x41, 0x87, 0x33, 0x46, 0x1d, 0xed, 0xdd, 0x89, 0x7b,
	0x41, 0xfb, 0x50, 0x92, 0x54, 0x29, 0x8f, 0xb9, 0x32, 0x0e, 0xa5, 0xd2, 0xde, 0x9c, 0xdc, 0x4b,
	0x32, 0x1f, 0x56, 0x32, 0xa0, 0x71, 0x1e, 0x83, 0xf7, 0xc9, 0x7c, 0xe2, 0x8c, 0x43, 0xbb, 0x00,
	0xba, 0x6a, 0x5b, 0x10, 0xe6, 0xd2, 0x34, 0x9a, 0xfa, 0x0c, 0x17, 0x46, 0x95, 0xd5, 0xe3, 0x42,
	0x61, 0xad, 0xc3, 0xe5, 0x70, 0x74, 0x89, 0x5e, 0x42, 0xd9, 0xf7, 0xa4, 0xa2, 0xcc, 0xe6, 0x2c,
	0x4d, 0xe4, 0xee, 0x1c, 0xfe, 0xb0, 0xd7, 0x15, 0x07, 0x3c, 0x20, 0x1e, 0xc3, 0xa5, 0x84, 0xe9,
	0x32, 0x54, 0x83, 0xbc, 0x22, 0xae, 0x59, 0xa8, 0x1b, 0xcd, 0x32, 0xd6, 0x97, 0xa8, 0x0b, 0x7f,
	0x93, 0x2c, 0x46, 0x5b, 0xa6, 0x39, 0x9a, 0x8b, 0xb1, 0xf7, 0x9d, 0x6b, 0xd2, 0x46, 0x64, 0x7a,
	0x70, 0x8e, 0xe0, 0x2f, 0xa9, 0x04, 0x25, 0x81, 0x9d, 0xc5, 0x55, 0x8c, 0xcd, 0x1e, 0x4e, 0x9a,
	0x65, 0x63, 0x6f, 0x8d, 0x3e, 0x13, 0xab, 0x1f, 0x53, 0x49, 0xda, 0xb8, 0x9a, 0x78, 0xf4, 0x47,
	0xc9, 0xed, 0x80, 0xa9, 0xd7, 0x3a, 0xb7, 0x43, 0x22, 0xa5, 0x77, 0x46, 0x6d, 0x27, 0xeb, 0xc7,
	0x5c, 0xaa, 0x1b, 0xcd, 0x12, 0xfe, 0x37, 0x7e, 0xde, 0x4b, 0x1e, 0x8f, 0xdb, 0x6b, 0xfc, 0xc8,
	0x81, 0xd9, 0x1d, 0xaa, 0x3f, 0x57, 0xea, 0x01, 0x2c, 0x4b, 0xca, 0x06, 0xb6, 0x3a, 0x15, 0x7c,
	0xe8, 0x9e, 0xa6, 0xb5, 0xde, 0xa0, 0x96, 0x8a, 0xc6, 0x8e, 0x12, 0x6a, 0x56, 0x6c, 0xf9, 0xdf,
	0x8f, 0xed, 0x23, 0x54, 0x43, 0xc1, 0x2f, 0xa2, 0xb1, 0x69, 0x52, 0xec, 0xd6, 0x35, 0xa6, 0x3d,
	0x0d, 0xa5, 0x9e, 0x2b, 0xb1, 0x43, 0x66, 0x39, 0x35, 0x42, 0x8d, 0x6f, 0x39, 0x28, 0xa6, 0x79,
	0xbe, 0x80, 0x25, 0x2f, 0xf9, 0x7e, 0x4c, 0xa3, 0x9e, 0x6f, 0x56, 0x2e, 0x1f, 0x1c, 0x73, 0x3e,
	0x2d, 0x3c, 0x62, 0xd0, 0x2b, 0x28, 0xf1, 0xb4, 0x2a, 0x33, 0x17, 0xf3, 0xeb, 0x93, 0xfc, 0xbc,
	0x1a, 0x71, 0x46, 0xa1, 0x16, 0xe4, 0x7d, 0xee, 0xa6, 0xd1, 0xad, 0xcd, 0xe8, 0xc0, 0xe7, 0xae,
	0x95, 0x52, 0x5a, 0x89, 0x76, 0x20, 0x4f, 0xc2, 0xd0, 0x2c, 0xc4, 0xab, 0xdd, 0xb4, 0x7c, 0x8d,
	0xa0, 0xe7, 0x50, 0xce, 0x92, 0x4b, 0x63, 0xbd, 0x3d, 0x3b, 0xd6, 0x74, 0xbd, 0xb1, 0x7c, 0x6b,
	0x13, 0x96, 0x93, 0x9b, 0x6f, 0xb8, 0x08, 0x88, 0xd2, 0xc7, 0x50, 0x4f, 0x9f, 0xfb, 0xc7, 0xc3,
	0x93, 0xda, 0x02, 0x2a, 0x41, 0xe1, 0x6d, 0xbf, 0xfb, 0xa1, 0x66, 0xec, 0x6f, 0x43, 0xd5, 0xe1,
	0xc1, 0x84, 0xeb, 0x7e, 0x25, 0xe1, 0x62, 0xf5, 0xe7, 0x82, 0xbe, 0xf5, 0x35, 0x07, 0x9f, 0xda,
	0x98, 0x44, 0x56, 0x87, 0x0b, 0x7a, 0x5c, 0x8c, 0xff, 0x3f, 0x9e, 0xfc, 0x0c, 0x00, 0x00, 0xff,
	0xff, 0xff, 0x44, 0xf3, 0xc8, 0x6a, 0x07, 0x00, 0x00,
}
