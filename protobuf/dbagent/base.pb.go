// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base.proto

/*
Package dbagent is a generated protocol buffer package.

It is generated from these files:
	base.proto
	db_agent_server.proto
	test_user.proto

It has these top-level messages:
	Nil
	BoolValue
	IntValue
	StringValue
	KV
	StringList
	Int64List
	TestUser
	ListUser
*/
package dbagent

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

type Nil struct {
}

func (m *Nil) Reset()                    { *m = Nil{} }
func (m *Nil) String() string            { return proto.CompactTextString(m) }
func (*Nil) ProtoMessage()               {}
func (*Nil) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type BoolValue struct {
	Value bool `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *BoolValue) Reset()                    { *m = BoolValue{} }
func (m *BoolValue) String() string            { return proto.CompactTextString(m) }
func (*BoolValue) ProtoMessage()               {}
func (*BoolValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BoolValue) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

type IntValue struct {
	Value int64 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *IntValue) Reset()                    { *m = IntValue{} }
func (m *IntValue) String() string            { return proto.CompactTextString(m) }
func (*IntValue) ProtoMessage()               {}
func (*IntValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *IntValue) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type StringValue struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *StringValue) Reset()                    { *m = StringValue{} }
func (m *StringValue) String() string            { return proto.CompactTextString(m) }
func (*StringValue) ProtoMessage()               {}
func (*StringValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *StringValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type KV struct {
	Field string `protobuf:"bytes,1,opt,name=field" json:"field,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *KV) Reset()                    { *m = KV{} }
func (m *KV) String() string            { return proto.CompactTextString(m) }
func (*KV) ProtoMessage()               {}
func (*KV) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *KV) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *KV) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type StringList struct {
	List []string `protobuf:"bytes,1,rep,name=list" json:"list,omitempty"`
}

func (m *StringList) Reset()                    { *m = StringList{} }
func (m *StringList) String() string            { return proto.CompactTextString(m) }
func (*StringList) ProtoMessage()               {}
func (*StringList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *StringList) GetList() []string {
	if m != nil {
		return m.List
	}
	return nil
}

type Int64List struct {
	List []int64 `protobuf:"varint,1,rep,packed,name=list" json:"list,omitempty"`
}

func (m *Int64List) Reset()                    { *m = Int64List{} }
func (m *Int64List) String() string            { return proto.CompactTextString(m) }
func (*Int64List) ProtoMessage()               {}
func (*Int64List) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Int64List) GetList() []int64 {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*Nil)(nil), "dbagent.Nil")
	proto.RegisterType((*BoolValue)(nil), "dbagent.BoolValue")
	proto.RegisterType((*IntValue)(nil), "dbagent.IntValue")
	proto.RegisterType((*StringValue)(nil), "dbagent.StringValue")
	proto.RegisterType((*KV)(nil), "dbagent.KV")
	proto.RegisterType((*StringList)(nil), "dbagent.StringList")
	proto.RegisterType((*Int64List)(nil), "dbagent.Int64List")
}

func init() { proto.RegisterFile("base.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4a, 0x2c, 0x4e,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0x49, 0x4a, 0x4c, 0x4f, 0xcd, 0x2b, 0x51,
	0x62, 0xe5, 0x62, 0xf6, 0xcb, 0xcc, 0x51, 0x52, 0xe4, 0xe2, 0x74, 0xca, 0xcf, 0xcf, 0x09, 0x4b,
	0xcc, 0x29, 0x4d, 0x15, 0x12, 0xe1, 0x62, 0x2d, 0x03, 0x31, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38,
	0x82, 0x20, 0x1c, 0x25, 0x05, 0x2e, 0x0e, 0xcf, 0xbc, 0x12, 0x2c, 0x2a, 0x98, 0x61, 0x2a, 0x94,
	0xb9, 0xb8, 0x83, 0x4b, 0x8a, 0x32, 0xf3, 0xd2, 0xb1, 0x28, 0xe2, 0x84, 0x29, 0x32, 0xe0, 0x62,
	0xf2, 0x0e, 0x03, 0xc9, 0xa5, 0x65, 0xa6, 0xe6, 0xa4, 0xc0, 0xe4, 0xc0, 0x1c, 0x84, 0x0e, 0x26,
	0x64, 0x1d, 0x0a, 0x5c, 0x5c, 0x10, 0x63, 0x7d, 0x32, 0x8b, 0x4b, 0x84, 0x84, 0xb8, 0x58, 0x72,
	0x32, 0x8b, 0x4b, 0x24, 0x18, 0x15, 0x98, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x25, 0x79, 0x2e, 0x4e,
	0xcf, 0xbc, 0x12, 0x33, 0x13, 0x0c, 0x05, 0xcc, 0x10, 0x05, 0x49, 0x6c, 0x60, 0x5f, 0x1b, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x08, 0x85, 0x7a, 0xcd, 0x03, 0x01, 0x00, 0x00,
}