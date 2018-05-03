// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Command.proto

package schema

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import Metadata "."

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

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
func (Command_ContextType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

// Command from a client.
type Command struct {
	Id       string              `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Action   string              `protobuf:"bytes,2,opt,name=action" json:"action,omitempty"`
	Data     []byte              `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Context  Command_ContextType `protobuf:"varint,4,opt,name=context,enum=Command_ContextType" json:"context,omitempty"`
	Metadata *Metadata.Metadata  `protobuf:"bytes,5,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *Command) Reset()                    { *m = Command{} }
func (m *Command) String() string            { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()               {}
func (*Command) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

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

func (m *Command) GetMetadata() *Metadata.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*Command)(nil), "Command")
	proto.RegisterEnum("Command_ContextType", Command_ContextType_name, Command_ContextType_value)
}

func init() { proto.RegisterFile("Command.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x75, 0xce, 0xcf, 0xcd,
	0x4d, 0xcc, 0x4b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0xe2, 0xf3, 0x4d, 0x2d, 0x49, 0x4c,
	0x49, 0x2c, 0x49, 0x84, 0xf0, 0x95, 0x6e, 0x31, 0x72, 0xb1, 0x43, 0x55, 0x08, 0xf1, 0x71, 0x31,
	0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31, 0x65, 0xa6, 0x08, 0x89, 0x71, 0xb1,
	0x25, 0x26, 0x97, 0x64, 0xe6, 0xe7, 0x49, 0x30, 0x81, 0xc5, 0xa0, 0x3c, 0x21, 0x21, 0x2e, 0x16,
	0x90, 0x09, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x60, 0xb6, 0x90, 0x1e, 0x17, 0x7b, 0x72,
	0x7e, 0x5e, 0x49, 0x6a, 0x45, 0x89, 0x04, 0x8b, 0x02, 0xa3, 0x06, 0x9f, 0x91, 0x88, 0x1e, 0xcc,
	0x62, 0x67, 0x88, 0x78, 0x48, 0x65, 0x41, 0x6a, 0x10, 0x4c, 0x91, 0x90, 0x2a, 0x17, 0x47, 0x2e,
	0xd4, 0x25, 0x12, 0xac, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x9c, 0x7a, 0x30, 0xa7, 0x05, 0xc1, 0xa5,
	0x94, 0x6c, 0xb8, 0xb8, 0x91, 0xb4, 0x0b, 0x71, 0x73, 0xb1, 0xbb, 0xb8, 0x06, 0x7b, 0x87, 0xf8,
	0x07, 0x08, 0x30, 0x08, 0xb1, 0x73, 0x31, 0x87, 0xbb, 0x3a, 0x09, 0x30, 0x0a, 0x71, 0x71, 0xb1,
	0xf9, 0xfa, 0x3b, 0x79, 0xfa, 0xb8, 0x0a, 0x30, 0x81, 0x54, 0x04, 0xbb, 0x06, 0x85, 0x79, 0x3a,
	0xbb, 0x0a, 0x30, 0x3b, 0x71, 0x44, 0xb1, 0x15, 0x27, 0x67, 0xa4, 0xe6, 0x26, 0x26, 0xb1, 0x81,
	0x7d, 0x6b, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x4a, 0xb7, 0x64, 0x45, 0x0e, 0x01, 0x00, 0x00,
}
