// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vitepb/account_blockmeta.proto

package vitepb

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

type AccountBlockMeta struct {
	AccountId            uint64   `protobuf:"varint,1,opt,name=accountId,proto3" json:"accountId,omitempty"`
	Height               uint64   `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	ReceiveBlockHeights  []uint64 `protobuf:"varint,3,rep,packed,name=receiveBlockHeights,proto3" json:"receiveBlockHeights,omitempty"`
	RefSnapshotHeight    uint64   `protobuf:"varint,4,opt,name=refSnapshotHeight,proto3" json:"refSnapshotHeight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountBlockMeta) Reset()         { *m = AccountBlockMeta{} }
func (m *AccountBlockMeta) String() string { return proto.CompactTextString(m) }
func (*AccountBlockMeta) ProtoMessage()    {}
func (*AccountBlockMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1ba65aec86c1ab8, []int{0}
}

func (m *AccountBlockMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountBlockMeta.Unmarshal(m, b)
}
func (m *AccountBlockMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountBlockMeta.Marshal(b, m, deterministic)
}
func (m *AccountBlockMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountBlockMeta.Merge(m, src)
}
func (m *AccountBlockMeta) XXX_Size() int {
	return xxx_messageInfo_AccountBlockMeta.Size(m)
}
func (m *AccountBlockMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountBlockMeta.DiscardUnknown(m)
}

var xxx_messageInfo_AccountBlockMeta proto.InternalMessageInfo

func (m *AccountBlockMeta) GetAccountId() uint64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

func (m *AccountBlockMeta) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *AccountBlockMeta) GetReceiveBlockHeights() []uint64 {
	if m != nil {
		return m.ReceiveBlockHeights
	}
	return nil
}

func (m *AccountBlockMeta) GetRefSnapshotHeight() uint64 {
	if m != nil {
		return m.RefSnapshotHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*AccountBlockMeta)(nil), "vitepb.AccountBlockMeta")
}

func init() { proto.RegisterFile("vitepb/account_blockmeta.proto", fileDescriptor_d1ba65aec86c1ab8) }

var fileDescriptor_d1ba65aec86c1ab8 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0xcb, 0x2c, 0x49,
	0x2d, 0x48, 0xd2, 0x4f, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0x89, 0x4f, 0xca, 0xc9, 0x4f, 0xce,
	0xce, 0x4d, 0x2d, 0x49, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xc8, 0x2b, 0xad,
	0x60, 0xe4, 0x12, 0x70, 0x84, 0xa8, 0x71, 0x02, 0x29, 0xf1, 0x4d, 0x2d, 0x49, 0x14, 0x92, 0xe1,
	0xe2, 0x84, 0xea, 0xf3, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x09, 0x42, 0x08, 0x08, 0x89,
	0x71, 0xb1, 0x65, 0xa4, 0x66, 0xa6, 0x67, 0x94, 0x48, 0x30, 0x81, 0xa5, 0xa0, 0x3c, 0x21, 0x03,
	0x2e, 0xe1, 0xa2, 0xd4, 0xe4, 0xd4, 0xcc, 0xb2, 0x54, 0xb0, 0x49, 0x1e, 0x60, 0xd1, 0x62, 0x09,
	0x66, 0x05, 0x66, 0x0d, 0x96, 0x20, 0x6c, 0x52, 0x42, 0x3a, 0x5c, 0x82, 0x45, 0xa9, 0x69, 0xc1,
	0x79, 0x89, 0x05, 0xc5, 0x19, 0xf9, 0x25, 0x10, 0x51, 0x09, 0x16, 0xb0, 0xa1, 0x98, 0x12, 0x49,
	0x6c, 0x60, 0x97, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x21, 0x33, 0xcf, 0xdb, 0x00,
	0x00, 0x00,
}
