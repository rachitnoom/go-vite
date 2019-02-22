// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vitepb/snapshot_block.proto

package vitepb

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

type SnapshotBlock struct {
	Hash                 []byte           `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	PrevHash             []byte           `protobuf:"bytes,2,opt,name=prevHash,proto3" json:"prevHash,omitempty"`
	Height               uint64           `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	PublicKey            []byte           `protobuf:"bytes,4,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	Signature            []byte           `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	Timestamp            int64            `protobuf:"varint,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	StateHash            []byte           `protobuf:"bytes,7,opt,name=stateHash,proto3" json:"stateHash,omitempty"`
	SnapshotContent      *SnapshotContent `protobuf:"bytes,8,opt,name=snapshotContent,proto3" json:"snapshotContent,omitempty"`
	Seed                 uint64           `protobuf:"varint,9,opt,name=seed,proto3" json:"seed,omitempty"`
	SeedHash             []byte           `protobuf:"bytes,10,opt,name=seedHash,proto3" json:"seedHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SnapshotBlock) Reset()         { *m = SnapshotBlock{} }
func (m *SnapshotBlock) String() string { return proto.CompactTextString(m) }
func (*SnapshotBlock) ProtoMessage()    {}
func (*SnapshotBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_snapshot_block_7f929f73b809029d, []int{0}
}
func (m *SnapshotBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnapshotBlock.Unmarshal(m, b)
}
func (m *SnapshotBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnapshotBlock.Marshal(b, m, deterministic)
}
func (dst *SnapshotBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotBlock.Merge(dst, src)
}
func (m *SnapshotBlock) XXX_Size() int {
	return xxx_messageInfo_SnapshotBlock.Size(m)
}
func (m *SnapshotBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotBlock.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotBlock proto.InternalMessageInfo

func (m *SnapshotBlock) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *SnapshotBlock) GetPrevHash() []byte {
	if m != nil {
		return m.PrevHash
	}
	return nil
}

func (m *SnapshotBlock) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *SnapshotBlock) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *SnapshotBlock) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *SnapshotBlock) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *SnapshotBlock) GetStateHash() []byte {
	if m != nil {
		return m.StateHash
	}
	return nil
}

func (m *SnapshotBlock) GetSnapshotContent() *SnapshotContent {
	if m != nil {
		return m.SnapshotContent
	}
	return nil
}

func (m *SnapshotBlock) GetSeed() uint64 {
	if m != nil {
		return m.Seed
	}
	return 0
}

func (m *SnapshotBlock) GetSeedHash() []byte {
	if m != nil {
		return m.SeedHash
	}
	return nil
}

func init() {
	proto.RegisterType((*SnapshotBlock)(nil), "vitepb.SnapshotBlock")
}

func init() {
	proto.RegisterFile("vitepb/snapshot_block.proto", fileDescriptor_snapshot_block_7f929f73b809029d)
}

var fileDescriptor_snapshot_block_7f929f73b809029d = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xd9, 0x36, 0xc6, 0x76, 0x54, 0x84, 0x3d, 0xe8, 0x52, 0x15, 0x82, 0xa7, 0x9c, 0x22,
	0xe8, 0x13, 0xa8, 0x17, 0xc1, 0x5b, 0xfa, 0x00, 0xb2, 0x89, 0x43, 0x77, 0xb1, 0xcd, 0x2e, 0xd9,
	0x69, 0xc1, 0xd7, 0xf3, 0xc9, 0x64, 0x67, 0x93, 0x96, 0xf6, 0x94, 0x99, 0xff, 0x9b, 0x9f, 0xec,
	0xff, 0xc3, 0xdd, 0xce, 0x12, 0xfa, 0xe6, 0x29, 0x74, 0xda, 0x07, 0xe3, 0xe8, 0xab, 0x59, 0xbb,
	0xf6, 0xa7, 0xf2, 0xbd, 0x23, 0x27, 0xf3, 0x04, 0x17, 0x0f, 0xa7, 0x47, 0xad, 0xeb, 0x08, 0x3b,
	0x4a, 0x67, 0x8f, 0x7f, 0x13, 0xb8, 0x5a, 0x0e, 0xe8, 0x2d, 0xda, 0xa5, 0x84, 0xcc, 0xe8, 0x60,
	0x94, 0x28, 0x44, 0x79, 0x59, 0xf3, 0x2c, 0x17, 0x30, 0xf3, 0x3d, 0xee, 0x3e, 0xa2, 0x3e, 0x61,
	0x7d, 0xbf, 0xcb, 0x1b, 0xc8, 0x0d, 0xda, 0x95, 0x21, 0x35, 0x2d, 0x44, 0x99, 0xd5, 0xc3, 0x26,
	0xef, 0x61, 0xee, 0xb7, 0xcd, 0xda, 0xb6, 0x9f, 0xf8, 0xab, 0x32, 0x36, 0x1d, 0x84, 0x48, 0x83,
	0x5d, 0x75, 0x9a, 0xb6, 0x3d, 0xaa, 0xb3, 0x44, 0xf7, 0x42, 0xa4, 0x64, 0x37, 0x18, 0x48, 0x6f,
	0xbc, 0xca, 0x0b, 0x51, 0x4e, 0xeb, 0x83, 0xc0, 0x5e, 0xd2, 0x84, 0xfc, 0x9c, 0xf3, 0xc1, 0x3b,
	0x0a, 0xf2, 0x15, 0xae, 0xc7, 0xac, 0xef, 0x29, 0xaa, 0x9a, 0x15, 0xa2, 0xbc, 0x78, 0xbe, 0xad,
	0x52, 0x15, 0xd5, 0xf2, 0x18, 0xd7, 0xa7, 0xf7, 0xb1, 0x82, 0x80, 0xf8, 0xad, 0xe6, 0x1c, 0x88,
	0xe7, 0x58, 0x41, 0xfc, 0xf2, 0x3f, 0x21, 0x55, 0x30, 0xee, 0x4d, 0xce, 0x5d, 0xbe, 0xfc, 0x07,
	0x00, 0x00, 0xff, 0xff, 0x0a, 0x83, 0x03, 0xd0, 0x91, 0x01, 0x00, 0x00,
}
