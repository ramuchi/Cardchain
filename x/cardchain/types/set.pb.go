// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cardchain/cardchain/set.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type CStatus int32

const (
	CStatus_design    CStatus = 0
	CStatus_finalized CStatus = 1
	CStatus_active    CStatus = 2
	CStatus_archived  CStatus = 3
)

var CStatus_name = map[int32]string{
	0: "design",
	1: "finalized",
	2: "active",
	3: "archived",
}

var CStatus_value = map[string]int32{
	"design":    0,
	"finalized": 1,
	"active":    2,
	"archived":  3,
}

func (x CStatus) String() string {
	return proto.EnumName(CStatus_name, int32(x))
}

func (CStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4433f04964645edd, []int{0}
}

type Set struct {
	Name         string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Cards        []uint64 `protobuf:"varint,2,rep,packed,name=cards,proto3" json:"cards,omitempty"`
	Artist       string   `protobuf:"bytes,3,opt,name=artist,proto3" json:"artist,omitempty"`
	StoryWriter  string   `protobuf:"bytes,4,opt,name=storyWriter,proto3" json:"storyWriter,omitempty"`
	Contributors []string `protobuf:"bytes,5,rep,name=contributors,proto3" json:"contributors,omitempty"`
	Story        string   `protobuf:"bytes,6,opt,name=story,proto3" json:"story,omitempty"`
	ArtworkId    uint64   `protobuf:"varint,7,opt,name=artworkId,proto3" json:"artworkId,omitempty"`
	Status       CStatus  `protobuf:"varint,8,opt,name=status,proto3,enum=DecentralCardGame.cardchain.cardchain.CStatus" json:"status,omitempty"`
	TimeStamp    int64    `protobuf:"varint,9,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
}

func (m *Set) Reset()         { *m = Set{} }
func (m *Set) String() string { return proto.CompactTextString(m) }
func (*Set) ProtoMessage()    {}
func (*Set) Descriptor() ([]byte, []int) {
	return fileDescriptor_4433f04964645edd, []int{0}
}
func (m *Set) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Set) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Set.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Set) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Set.Merge(m, src)
}
func (m *Set) XXX_Size() int {
	return m.Size()
}
func (m *Set) XXX_DiscardUnknown() {
	xxx_messageInfo_Set.DiscardUnknown(m)
}

var xxx_messageInfo_Set proto.InternalMessageInfo

func (m *Set) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Set) GetCards() []uint64 {
	if m != nil {
		return m.Cards
	}
	return nil
}

func (m *Set) GetArtist() string {
	if m != nil {
		return m.Artist
	}
	return ""
}

func (m *Set) GetStoryWriter() string {
	if m != nil {
		return m.StoryWriter
	}
	return ""
}

func (m *Set) GetContributors() []string {
	if m != nil {
		return m.Contributors
	}
	return nil
}

func (m *Set) GetStory() string {
	if m != nil {
		return m.Story
	}
	return ""
}

func (m *Set) GetArtworkId() uint64 {
	if m != nil {
		return m.ArtworkId
	}
	return 0
}

func (m *Set) GetStatus() CStatus {
	if m != nil {
		return m.Status
	}
	return CStatus_design
}

func (m *Set) GetTimeStamp() int64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

type OutpSet struct {
	Name         string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Cards        []uint64 `protobuf:"varint,2,rep,packed,name=cards,proto3" json:"cards,omitempty"`
	Artist       string   `protobuf:"bytes,3,opt,name=artist,proto3" json:"artist,omitempty"`
	StoryWriter  string   `protobuf:"bytes,4,opt,name=storyWriter,proto3" json:"storyWriter,omitempty"`
	Contributors []string `protobuf:"bytes,5,rep,name=contributors,proto3" json:"contributors,omitempty"`
	Story        string   `protobuf:"bytes,6,opt,name=story,proto3" json:"story,omitempty"`
	Artwork      string   `protobuf:"bytes,7,opt,name=artwork,proto3" json:"artwork,omitempty"`
	Status       CStatus  `protobuf:"varint,8,opt,name=status,proto3,enum=DecentralCardGame.cardchain.cardchain.CStatus" json:"status,omitempty"`
	TimeStamp    int64    `protobuf:"varint,9,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
}

func (m *OutpSet) Reset()         { *m = OutpSet{} }
func (m *OutpSet) String() string { return proto.CompactTextString(m) }
func (*OutpSet) ProtoMessage()    {}
func (*OutpSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_4433f04964645edd, []int{1}
}
func (m *OutpSet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OutpSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OutpSet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OutpSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutpSet.Merge(m, src)
}
func (m *OutpSet) XXX_Size() int {
	return m.Size()
}
func (m *OutpSet) XXX_DiscardUnknown() {
	xxx_messageInfo_OutpSet.DiscardUnknown(m)
}

var xxx_messageInfo_OutpSet proto.InternalMessageInfo

func (m *OutpSet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OutpSet) GetCards() []uint64 {
	if m != nil {
		return m.Cards
	}
	return nil
}

func (m *OutpSet) GetArtist() string {
	if m != nil {
		return m.Artist
	}
	return ""
}

func (m *OutpSet) GetStoryWriter() string {
	if m != nil {
		return m.StoryWriter
	}
	return ""
}

func (m *OutpSet) GetContributors() []string {
	if m != nil {
		return m.Contributors
	}
	return nil
}

func (m *OutpSet) GetStory() string {
	if m != nil {
		return m.Story
	}
	return ""
}

func (m *OutpSet) GetArtwork() string {
	if m != nil {
		return m.Artwork
	}
	return ""
}

func (m *OutpSet) GetStatus() CStatus {
	if m != nil {
		return m.Status
	}
	return CStatus_design
}

func (m *OutpSet) GetTimeStamp() int64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

func init() {
	proto.RegisterEnum("DecentralCardGame.cardchain.cardchain.CStatus", CStatus_name, CStatus_value)
	proto.RegisterType((*Set)(nil), "DecentralCardGame.cardchain.cardchain.Set")
	proto.RegisterType((*OutpSet)(nil), "DecentralCardGame.cardchain.cardchain.OutpSet")
}

func init() { proto.RegisterFile("cardchain/cardchain/set.proto", fileDescriptor_4433f04964645edd) }

var fileDescriptor_4433f04964645edd = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x53, 0xcd, 0x6a, 0xdb, 0x40,
	0x18, 0xd4, 0x5a, 0xb2, 0x64, 0x6d, 0xdd, 0x22, 0x96, 0x52, 0xf6, 0xd0, 0x0a, 0x61, 0x28, 0x88,
	0x1e, 0x64, 0x68, 0x2f, 0x3d, 0xf5, 0x50, 0x97, 0x96, 0x9e, 0x0a, 0xeb, 0x43, 0x21, 0xb7, 0xb5,
	0xb4, 0xb1, 0x97, 0x58, 0x3f, 0xec, 0x7e, 0x72, 0xe2, 0x3c, 0x45, 0xde, 0x22, 0xaf, 0x92, 0xa3,
	0x8f, 0x39, 0x06, 0xfb, 0x45, 0x82, 0xd6, 0x8a, 0xe5, 0x90, 0x4b, 0x4e, 0x21, 0xb7, 0x6f, 0xe6,
	0x9b, 0xe1, 0xdb, 0x1d, 0x18, 0xfc, 0x29, 0xe5, 0x2a, 0x4b, 0x17, 0x5c, 0x16, 0xe3, 0x6e, 0xd2,
	0x02, 0x92, 0x4a, 0x95, 0x50, 0x92, 0xcf, 0xbf, 0x44, 0x2a, 0x0a, 0x50, 0x7c, 0x39, 0xe1, 0x2a,
	0xfb, 0xc3, 0x73, 0x91, 0x1c, 0x64, 0xdd, 0x34, 0xba, 0xee, 0x61, 0x7b, 0x2a, 0x80, 0x10, 0xec,
	0x14, 0x3c, 0x17, 0x14, 0x45, 0x28, 0xf6, 0x99, 0x99, 0xc9, 0x7b, 0xdc, 0x6f, 0x84, 0x9a, 0xf6,
	0x22, 0x3b, 0x76, 0xd8, 0x1e, 0x90, 0x0f, 0xd8, 0xe5, 0x0a, 0xa4, 0x06, 0x6a, 0x1b, 0x6d, 0x8b,
	0x48, 0x84, 0xdf, 0x68, 0x28, 0xd5, 0xfa, 0xbf, 0x92, 0x20, 0x14, 0x75, 0xcc, 0xf2, 0x98, 0x22,
	0x23, 0x3c, 0x4c, 0xcb, 0x02, 0x94, 0x9c, 0xd5, 0x50, 0x2a, 0x4d, 0xfb, 0x91, 0x1d, 0xfb, 0xec,
	0x11, 0xd7, 0xdc, 0x34, 0x16, 0xea, 0x1a, 0xff, 0x1e, 0x90, 0x8f, 0xd8, 0xe7, 0x0a, 0xce, 0x4b,
	0x75, 0xf6, 0x37, 0xa3, 0x5e, 0x84, 0x62, 0x87, 0x75, 0x04, 0xf9, 0x8d, 0x5d, 0x0d, 0x1c, 0x6a,
	0x4d, 0x07, 0x11, 0x8a, 0xdf, 0x7d, 0x4d, 0x92, 0x67, 0xfd, 0x3d, 0x99, 0x4c, 0x8d, 0x8b, 0xb5,
	0xee, 0xe6, 0x0a, 0xc8, 0x5c, 0x4c, 0x81, 0xe7, 0x15, 0xf5, 0x23, 0x14, 0xdb, 0xac, 0x23, 0x9a,
	0xa4, 0xbc, 0x7f, 0x35, 0x54, 0xaf, 0x3f, 0x2d, 0x8a, 0xbd, 0x36, 0x1c, 0x93, 0x95, 0xcf, 0x1e,
	0xe0, 0xcb, 0x24, 0xf5, 0xe5, 0x07, 0xf6, 0x5a, 0x03, 0xc1, 0xd8, 0xcd, 0x84, 0x96, 0xf3, 0x22,
	0xb0, 0xc8, 0x5b, 0xec, 0x9f, 0xca, 0x82, 0x2f, 0xe5, 0xa5, 0xc8, 0x02, 0xd4, 0xac, 0x78, 0x0a,
	0x72, 0x25, 0x82, 0x1e, 0x19, 0xe2, 0x01, 0x57, 0xe9, 0x42, 0xae, 0x44, 0x16, 0xd8, 0x3f, 0xd9,
	0xcd, 0x36, 0x44, 0x9b, 0x6d, 0x88, 0xee, 0xb6, 0x21, 0xba, 0xda, 0x85, 0xd6, 0x66, 0x17, 0x5a,
	0xb7, 0xbb, 0xd0, 0x3a, 0xf9, 0x3e, 0x97, 0xb0, 0xa8, 0x67, 0x49, 0x5a, 0xe6, 0xe3, 0x27, 0x2f,
	0x1f, 0x4f, 0x0e, 0x35, 0xb8, 0x38, 0xaa, 0x04, 0xac, 0x2b, 0xa1, 0x67, 0xae, 0x69, 0xc5, 0xb7,
	0xfb, 0x00, 0x00, 0x00, 0xff, 0xff, 0x98, 0x0f, 0xc5, 0x59, 0x36, 0x03, 0x00, 0x00,
}

func (m *Set) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Set) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Set) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TimeStamp != 0 {
		i = encodeVarintSet(dAtA, i, uint64(m.TimeStamp))
		i--
		dAtA[i] = 0x48
	}
	if m.Status != 0 {
		i = encodeVarintSet(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x40
	}
	if m.ArtworkId != 0 {
		i = encodeVarintSet(dAtA, i, uint64(m.ArtworkId))
		i--
		dAtA[i] = 0x38
	}
	if len(m.Story) > 0 {
		i -= len(m.Story)
		copy(dAtA[i:], m.Story)
		i = encodeVarintSet(dAtA, i, uint64(len(m.Story)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Contributors) > 0 {
		for iNdEx := len(m.Contributors) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Contributors[iNdEx])
			copy(dAtA[i:], m.Contributors[iNdEx])
			i = encodeVarintSet(dAtA, i, uint64(len(m.Contributors[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.StoryWriter) > 0 {
		i -= len(m.StoryWriter)
		copy(dAtA[i:], m.StoryWriter)
		i = encodeVarintSet(dAtA, i, uint64(len(m.StoryWriter)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Artist) > 0 {
		i -= len(m.Artist)
		copy(dAtA[i:], m.Artist)
		i = encodeVarintSet(dAtA, i, uint64(len(m.Artist)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Cards) > 0 {
		dAtA2 := make([]byte, len(m.Cards)*10)
		var j1 int
		for _, num := range m.Cards {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintSet(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintSet(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OutpSet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OutpSet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OutpSet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TimeStamp != 0 {
		i = encodeVarintSet(dAtA, i, uint64(m.TimeStamp))
		i--
		dAtA[i] = 0x48
	}
	if m.Status != 0 {
		i = encodeVarintSet(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x40
	}
	if len(m.Artwork) > 0 {
		i -= len(m.Artwork)
		copy(dAtA[i:], m.Artwork)
		i = encodeVarintSet(dAtA, i, uint64(len(m.Artwork)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Story) > 0 {
		i -= len(m.Story)
		copy(dAtA[i:], m.Story)
		i = encodeVarintSet(dAtA, i, uint64(len(m.Story)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Contributors) > 0 {
		for iNdEx := len(m.Contributors) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Contributors[iNdEx])
			copy(dAtA[i:], m.Contributors[iNdEx])
			i = encodeVarintSet(dAtA, i, uint64(len(m.Contributors[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.StoryWriter) > 0 {
		i -= len(m.StoryWriter)
		copy(dAtA[i:], m.StoryWriter)
		i = encodeVarintSet(dAtA, i, uint64(len(m.StoryWriter)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Artist) > 0 {
		i -= len(m.Artist)
		copy(dAtA[i:], m.Artist)
		i = encodeVarintSet(dAtA, i, uint64(len(m.Artist)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Cards) > 0 {
		dAtA4 := make([]byte, len(m.Cards)*10)
		var j3 int
		for _, num := range m.Cards {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		i -= j3
		copy(dAtA[i:], dAtA4[:j3])
		i = encodeVarintSet(dAtA, i, uint64(j3))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintSet(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSet(dAtA []byte, offset int, v uint64) int {
	offset -= sovSet(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Set) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSet(uint64(l))
	}
	if len(m.Cards) > 0 {
		l = 0
		for _, e := range m.Cards {
			l += sovSet(uint64(e))
		}
		n += 1 + sovSet(uint64(l)) + l
	}
	l = len(m.Artist)
	if l > 0 {
		n += 1 + l + sovSet(uint64(l))
	}
	l = len(m.StoryWriter)
	if l > 0 {
		n += 1 + l + sovSet(uint64(l))
	}
	if len(m.Contributors) > 0 {
		for _, s := range m.Contributors {
			l = len(s)
			n += 1 + l + sovSet(uint64(l))
		}
	}
	l = len(m.Story)
	if l > 0 {
		n += 1 + l + sovSet(uint64(l))
	}
	if m.ArtworkId != 0 {
		n += 1 + sovSet(uint64(m.ArtworkId))
	}
	if m.Status != 0 {
		n += 1 + sovSet(uint64(m.Status))
	}
	if m.TimeStamp != 0 {
		n += 1 + sovSet(uint64(m.TimeStamp))
	}
	return n
}

func (m *OutpSet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSet(uint64(l))
	}
	if len(m.Cards) > 0 {
		l = 0
		for _, e := range m.Cards {
			l += sovSet(uint64(e))
		}
		n += 1 + sovSet(uint64(l)) + l
	}
	l = len(m.Artist)
	if l > 0 {
		n += 1 + l + sovSet(uint64(l))
	}
	l = len(m.StoryWriter)
	if l > 0 {
		n += 1 + l + sovSet(uint64(l))
	}
	if len(m.Contributors) > 0 {
		for _, s := range m.Contributors {
			l = len(s)
			n += 1 + l + sovSet(uint64(l))
		}
	}
	l = len(m.Story)
	if l > 0 {
		n += 1 + l + sovSet(uint64(l))
	}
	l = len(m.Artwork)
	if l > 0 {
		n += 1 + l + sovSet(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovSet(uint64(m.Status))
	}
	if m.TimeStamp != 0 {
		n += 1 + sovSet(uint64(m.TimeStamp))
	}
	return n
}

func sovSet(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSet(x uint64) (n int) {
	return sovSet(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Set) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSet
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Set: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Set: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowSet
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Cards = append(m.Cards, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowSet
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthSet
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthSet
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Cards) == 0 {
					m.Cards = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowSet
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Cards = append(m.Cards, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Cards", wireType)
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Artist", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Artist = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoryWriter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StoryWriter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contributors", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contributors = append(m.Contributors, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Story", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Story = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ArtworkId", wireType)
			}
			m.ArtworkId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ArtworkId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= CStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeStamp", wireType)
			}
			m.TimeStamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeStamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSet
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *OutpSet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSet
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OutpSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OutpSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowSet
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Cards = append(m.Cards, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowSet
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthSet
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthSet
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Cards) == 0 {
					m.Cards = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowSet
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Cards = append(m.Cards, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Cards", wireType)
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Artist", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Artist = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoryWriter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StoryWriter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contributors", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contributors = append(m.Contributors, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Story", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Story = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Artwork", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Artwork = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= CStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeStamp", wireType)
			}
			m.TimeStamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeStamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSet
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSet(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSet
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSet
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSet
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthSet
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSet
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSet
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSet        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSet          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSet = fmt.Errorf("proto: unexpected end of group")
)
