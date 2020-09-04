// Code generated by protoc-gen-go. DO NOT EDIT.
// source: action.proto

package PlayStructures

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Action struct {
	Action               string   `protobuf:"bytes,1,opt,name=Action,proto3" json:"Action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Action) Reset()         { *m = Action{} }
func (m *Action) String() string { return proto.CompactTextString(m) }
func (*Action) ProtoMessage()    {}
func (*Action) Descriptor() ([]byte, []int) {
	return fileDescriptor_59885c909ad4dfd3, []int{0}
}

func (m *Action) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Action.Unmarshal(m, b)
}
func (m *Action) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Action.Marshal(b, m, deterministic)
}
func (m *Action) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Action.Merge(m, src)
}
func (m *Action) XXX_Size() int {
	return xxx_messageInfo_Action.Size(m)
}
func (m *Action) XXX_DiscardUnknown() {
	xxx_messageInfo_Action.DiscardUnknown(m)
}

var xxx_messageInfo_Action proto.InternalMessageInfo

func (m *Action) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func init() {
	proto.RegisterType((*Action)(nil), "PlayStructures.Action")
}

func init() { proto.RegisterFile("action.proto", fileDescriptor_59885c909ad4dfd3) }

var fileDescriptor_59885c909ad4dfd3 = []byte{
	// 78 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x4c, 0x2e, 0xc9,
	0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x0b, 0xc8, 0x49, 0xac, 0x0c, 0x2e,
	0x29, 0x2a, 0x4d, 0x2e, 0x29, 0x2d, 0x4a, 0x2d, 0x56, 0x52, 0xe0, 0x62, 0x73, 0x04, 0xcb, 0x0b,
	0x89, 0xc1, 0x58, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x50, 0x5e, 0x12, 0x1b, 0x58, 0xa3,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xaf, 0xcf, 0x87, 0x6f, 0x48, 0x00, 0x00, 0x00,
}