// Code generated by protoc-gen-go. DO NOT EDIT.
// source: boundlessgeo_schema/Command.proto

/*
Package schema is a generated protocol buffer package.

It is generated from these files:
	boundlessgeo_schema/Command.proto
	boundlessgeo_schema/Event.proto
	boundlessgeo_schema/Feature.proto
	boundlessgeo_schema/Metadata.proto
	boundlessgeo_schema/Msg.proto
	boundlessgeo_schema/worm.proto

It has these top-level messages:
	Command
	Event
	Feature
	Operation
	Metadata
	Msg
	Response
	Subscription
*/
package schema

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

type Command_ContextType int32

const (
	Command_DESKTOP Command_ContextType = 0
	Command_WEB     Command_ContextType = 1
	Command_MOBILE  Command_ContextType = 2
	Command_SERVICE Command_ContextType = 3
)

var Command_ContextType_name = map[int32]string{
	0: "DESKTOP",
	1: "WEB",
	2: "MOBILE",
	3: "SERVICE",
}
var Command_ContextType_value = map[string]int32{
	"DESKTOP": 0,
	"WEB":     1,
	"MOBILE":  2,
	"SERVICE": 3,
}

func (x Command_ContextType) String() string {
	return proto.EnumName(Command_ContextType_name, int32(x))
}
func (Command_ContextType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// Command from a client.
type Command struct {
	Id       string              `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Action   string              `protobuf:"bytes,2,opt,name=action" json:"action,omitempty"`
	Data     []byte              `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Context  Command_ContextType `protobuf:"varint,4,opt,name=context,enum=Command_ContextType" json:"context,omitempty"`
	Metadata *Metadata           `protobuf:"bytes,5,opt,name=metadata" json:"metadata,omitempty"`
	Created  int64               `protobuf:"varint,6,opt,name=created" json:"created,omitempty"`
}

func (m *Command) Reset()                    { *m = Command{} }
func (m *Command) String() string            { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()               {}
func (*Command) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Command) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Command) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *Command) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Command) GetContext() Command_ContextType {
	if m != nil {
		return m.Context
	}
	return Command_DESKTOP
}

func (m *Command) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Command) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func init() {
	proto.RegisterType((*Command)(nil), "Command")
	proto.RegisterEnum("Command_ContextType", Command_ContextType_name, Command_ContextType_value)
}

func init() { proto.RegisterFile("boundlessgeo_schema/Command.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x4d, 0x3b, 0x5b, 0xf7, 0x2a, 0xa3, 0x3c, 0x44, 0x83, 0xa7, 0x5a, 0x10, 0x7a, 0x31,
	0xc2, 0xbc, 0x7a, 0x6a, 0xed, 0x61, 0xe8, 0xd8, 0xc8, 0x86, 0x82, 0x17, 0x49, 0x9b, 0xa8, 0x05,
	0xd3, 0x8c, 0x35, 0x82, 0xfe, 0xdf, 0xfe, 0x01, 0xb2, 0xfe, 0x90, 0x1d, 0xbc, 0xe5, 0x7d, 0xf3,
	0x79, 0x8f, 0x0f, 0x5f, 0xb8, 0x28, 0xcc, 0x67, 0x2d, 0x3f, 0x54, 0xd3, 0xbc, 0x29, 0xf3, 0xd2,
	0x94, 0xef, 0x4a, 0x8b, 0xeb, 0xcc, 0x68, 0x2d, 0x6a, 0xc9, 0x36, 0x5b, 0x63, 0xcd, 0x79, 0xfc,
	0x1f, 0x32, 0x57, 0x56, 0x48, 0x61, 0x45, 0xc7, 0xc4, 0x3f, 0x04, 0xfc, 0x7e, 0x0b, 0x27, 0xe0,
	0x54, 0x92, 0x92, 0x88, 0x24, 0x63, 0xee, 0x54, 0x12, 0x4f, 0xc1, 0x13, 0xa5, 0xad, 0x4c, 0x4d,
	0x9d, 0x36, 0xeb, 0x27, 0x44, 0x18, 0xed, 0x2e, 0x50, 0x37, 0x22, 0xc9, 0x31, 0x6f, 0xdf, 0xc8,
	0xc0, 0x2f, 0x4d, 0x6d, 0xd5, 0x97, 0xa5, 0xa3, 0x88, 0x24, 0x93, 0xe9, 0x09, 0x1b, 0x64, 0xb2,
	0x2e, 0x5f, 0x7f, 0x6f, 0x14, 0x1f, 0x20, 0xbc, 0x84, 0x23, 0xdd, 0x9b, 0xd0, 0xc3, 0x88, 0x24,
	0xc1, 0x74, 0xcc, 0x06, 0x35, 0xfe, 0xf7, 0x85, 0x14, 0xfc, 0x72, 0xab, 0x84, 0x55, 0x92, 0x7a,
	0x11, 0x49, 0x5c, 0x3e, 0x8c, 0xf1, 0x2d, 0x04, 0x7b, 0x87, 0x31, 0x00, 0xff, 0x2e, 0x5f, 0xdd,
	0xaf, 0x17, 0xcb, 0xf0, 0x00, 0x7d, 0x70, 0x9f, 0xf2, 0x34, 0x24, 0x08, 0xe0, 0xcd, 0x17, 0xe9,
	0xec, 0x21, 0x0f, 0x9d, 0x1d, 0xb1, 0xca, 0xf9, 0xe3, 0x2c, 0xcb, 0x43, 0x37, 0xbd, 0x82, 0xb3,
	0xd2, 0x68, 0xb6, 0x5f, 0x10, 0xeb, 0x0a, 0x4a, 0xa1, 0xf7, 0x5e, 0x16, 0xaf, 0xcf, 0x5e, 0x97,
	0x15, 0x5e, 0x5b, 0xd6, 0xcd, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x21, 0x23, 0x88, 0xe2, 0x75,
	0x01, 0x00, 0x00,
}
