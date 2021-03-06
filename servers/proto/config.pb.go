// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/google/cloudprober/servers/proto/config.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import proto1 "github.com/google/cloudprober/servers/http/proto"
import proto2 "github.com/google/cloudprober/servers/udp/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ServerDef_Type int32

const (
	ServerDef_HTTP ServerDef_Type = 0
	ServerDef_UDP  ServerDef_Type = 1
)

var ServerDef_Type_name = map[int32]string{
	0: "HTTP",
	1: "UDP",
}
var ServerDef_Type_value = map[string]int32{
	"HTTP": 0,
	"UDP":  1,
}

func (x ServerDef_Type) Enum() *ServerDef_Type {
	p := new(ServerDef_Type)
	*p = x
	return p
}
func (x ServerDef_Type) String() string {
	return proto.EnumName(ServerDef_Type_name, int32(x))
}
func (x *ServerDef_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ServerDef_Type_value, data, "ServerDef_Type")
	if err != nil {
		return err
	}
	*x = ServerDef_Type(value)
	return nil
}
func (ServerDef_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_config_32ba5fc994ff6f3e, []int{0, 0}
}

type ServerDef struct {
	Type *ServerDef_Type `protobuf:"varint,1,req,name=type,enum=cloudprober.servers.ServerDef_Type" json:"type,omitempty"`
	// Types that are valid to be assigned to Server:
	//	*ServerDef_HttpServer
	//	*ServerDef_UdpServer
	Server               isServerDef_Server `protobuf_oneof:"server"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ServerDef) Reset()         { *m = ServerDef{} }
func (m *ServerDef) String() string { return proto.CompactTextString(m) }
func (*ServerDef) ProtoMessage()    {}
func (*ServerDef) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_32ba5fc994ff6f3e, []int{0}
}
func (m *ServerDef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerDef.Unmarshal(m, b)
}
func (m *ServerDef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerDef.Marshal(b, m, deterministic)
}
func (dst *ServerDef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerDef.Merge(dst, src)
}
func (m *ServerDef) XXX_Size() int {
	return xxx_messageInfo_ServerDef.Size(m)
}
func (m *ServerDef) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerDef.DiscardUnknown(m)
}

var xxx_messageInfo_ServerDef proto.InternalMessageInfo

func (m *ServerDef) GetType() ServerDef_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ServerDef_HTTP
}

type isServerDef_Server interface {
	isServerDef_Server()
}

type ServerDef_HttpServer struct {
	HttpServer *proto1.ServerConf `protobuf:"bytes,2,opt,name=http_server,json=httpServer,oneof"`
}

type ServerDef_UdpServer struct {
	UdpServer *proto2.ServerConf `protobuf:"bytes,3,opt,name=udp_server,json=udpServer,oneof"`
}

func (*ServerDef_HttpServer) isServerDef_Server() {}

func (*ServerDef_UdpServer) isServerDef_Server() {}

func (m *ServerDef) GetServer() isServerDef_Server {
	if m != nil {
		return m.Server
	}
	return nil
}

func (m *ServerDef) GetHttpServer() *proto1.ServerConf {
	if x, ok := m.GetServer().(*ServerDef_HttpServer); ok {
		return x.HttpServer
	}
	return nil
}

func (m *ServerDef) GetUdpServer() *proto2.ServerConf {
	if x, ok := m.GetServer().(*ServerDef_UdpServer); ok {
		return x.UdpServer
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ServerDef) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ServerDef_OneofMarshaler, _ServerDef_OneofUnmarshaler, _ServerDef_OneofSizer, []interface{}{
		(*ServerDef_HttpServer)(nil),
		(*ServerDef_UdpServer)(nil),
	}
}

func _ServerDef_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ServerDef)
	// server
	switch x := m.Server.(type) {
	case *ServerDef_HttpServer:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.HttpServer); err != nil {
			return err
		}
	case *ServerDef_UdpServer:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.UdpServer); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ServerDef.Server has unexpected type %T", x)
	}
	return nil
}

func _ServerDef_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ServerDef)
	switch tag {
	case 2: // server.http_server
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(proto1.ServerConf)
		err := b.DecodeMessage(msg)
		m.Server = &ServerDef_HttpServer{msg}
		return true, err
	case 3: // server.udp_server
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(proto2.ServerConf)
		err := b.DecodeMessage(msg)
		m.Server = &ServerDef_UdpServer{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ServerDef_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ServerDef)
	// server
	switch x := m.Server.(type) {
	case *ServerDef_HttpServer:
		s := proto.Size(x.HttpServer)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ServerDef_UdpServer:
		s := proto.Size(x.UdpServer)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*ServerDef)(nil), "cloudprober.servers.ServerDef")
	proto.RegisterEnum("cloudprober.servers.ServerDef_Type", ServerDef_Type_name, ServerDef_Type_value)
}

func init() {
	proto.RegisterFile("github.com/google/cloudprober/servers/proto/config.proto", fileDescriptor_config_32ba5fc994ff6f3e)
}

var fileDescriptor_config_32ba5fc994ff6f3e = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x48, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0xce,
	0xc9, 0x2f, 0x4d, 0x29, 0x28, 0xca, 0x4f, 0x4a, 0x2d, 0xd2, 0x2f, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d,
	0x2a, 0xd6, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x03,
	0x73, 0x84, 0x84, 0x91, 0xd4, 0xe9, 0x41, 0xd5, 0x49, 0xd9, 0x12, 0x67, 0x5c, 0x46, 0x49, 0x49,
	0x01, 0x16, 0x33, 0xa5, 0x6c, 0x88, 0xd3, 0x5e, 0x9a, 0x82, 0x4d, 0xb7, 0xd2, 0x0f, 0x46, 0x2e,
	0xce, 0x60, 0xb0, 0x12, 0x97, 0xd4, 0x34, 0x21, 0x73, 0x2e, 0x96, 0x92, 0xca, 0x82, 0x54, 0x09,
	0x46, 0x05, 0x26, 0x0d, 0x3e, 0x23, 0x65, 0x3d, 0x2c, 0xce, 0xd5, 0x83, 0xab, 0xd6, 0x0b, 0xa9,
	0x2c, 0x48, 0x0d, 0x02, 0x6b, 0x10, 0x72, 0xe7, 0xe2, 0x06, 0xb9, 0x2f, 0x1e, 0xa2, 0x48, 0x82,
	0x49, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x05, 0xab, 0x7e, 0x90, 0x3a, 0xa8, 0x21, 0xce, 0xf9, 0x79,
	0x69, 0x1e, 0x0c, 0x41, 0x5c, 0x20, 0x21, 0x88, 0x88, 0x90, 0x0b, 0x17, 0x57, 0x69, 0x0a, 0xdc,
	0x1c, 0x66, 0xb0, 0x39, 0xd8, 0xdd, 0x51, 0x9a, 0x82, 0x66, 0x0c, 0x67, 0x69, 0x0a, 0xd4, 0x14,
	0x25, 0x49, 0x2e, 0x16, 0x90, 0xe3, 0x84, 0x38, 0xb8, 0x58, 0x3c, 0x42, 0x42, 0x02, 0x04, 0x18,
	0x84, 0xd8, 0xb9, 0x98, 0x43, 0x5d, 0x02, 0x04, 0x18, 0x9d, 0x38, 0xb8, 0xd8, 0x20, 0x26, 0x00,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x63, 0x10, 0x34, 0xc7, 0x01, 0x00, 0x00,
}
