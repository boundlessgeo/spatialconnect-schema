// Code generated by protoc-gen-go.
// source: Msg.proto
// DO NOT EDIT!

/*
Package com_boundlessgeo_schema is a generated protocol buffer package.

It is generated from these files:
	Msg.proto

It has these top-level messages:
	Msg
*/
package com_boundlessgeo_schema

import proto "github.com/golang/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type Msg struct {
	Context       string `protobuf:"bytes,1,opt,name=context" json:"context,omitempty"`
	CorrelationId int64  `protobuf:"varint,2,opt,name=correlationId" json:"correlationId,omitempty"`
	To            string `protobuf:"bytes,3,opt,name=to" json:"to,omitempty"`
	Action        string `protobuf:"bytes,4,opt,name=action" json:"action,omitempty"`
	Payload       string `protobuf:"bytes,5,opt,name=payload" json:"payload,omitempty"`
	Jwt           string `protobuf:"bytes,6,opt,name=jwt" json:"jwt,omitempty"`
}

func (m *Msg) Reset()         { *m = Msg{} }
func (m *Msg) String() string { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()    {}

func init() {
}
