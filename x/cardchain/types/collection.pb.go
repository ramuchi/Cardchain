// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cardchain/cardchain/collection.proto

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
	return fileDescriptor_bfff38131151e449, []int{0}
}

type Collection struct {
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

func (m *Collection) Reset()         { *m = Collection{} }
func (m *Collection) String() string { return proto.CompactTextString(m) }
func (*Collection) ProtoMessage()    {}
func (*Collection) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfff38131151e449, []int{0}
}
func (m *Collection) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Collection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Collection.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Collection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Collection.Merge(m, src)
}
func (m *Collection) XXX_Size() int {
	return m.Size()
}
func (m *Collection) XXX_DiscardUnknown() {
	xxx_messageInfo_Collection.DiscardUnknown(m)
}

var xxx_messageInfo_Collection proto.InternalMessageInfo

func (m *Collection) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Collection) GetCards() []uint64 {
	if m != nil {
		return m.Cards
	}
	return nil
}

func (m *Collection) GetArtist() string {
	if m != nil {
		return m.Artist
	}
	return ""
}

func (m *Collection) GetStoryWriter() string {
	if m != nil {
		return m.StoryWriter
	}
	return ""
}

func (m *Collection) GetContributors() []string {
	if m != nil {
		return m.Contributors
	}
	return nil
}

func (m *Collection) GetStory() string {
	if m != nil {
		return m.Story
	}
	return ""
}

func (m *Collection) GetArtworkId() uint64 {
	if m != nil {
		return m.ArtworkId
	}
	return 0
}

func (m *Collection) GetStatus() CStatus {
	if m != nil {
		return m.Status
	}
	return CStatus_design
}

func (m *Collection) GetTimeStamp() int64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

type OutpCollection struct {
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

func (m *OutpCollection) Reset()         { *m = OutpCollection{} }
func (m *OutpCollection) String() string { return proto.CompactTextString(m) }
func (*OutpCollection) ProtoMessage()    {}
func (*OutpCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfff38131151e449, []int{1}
}
func (m *OutpCollection) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OutpCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OutpCollection.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OutpCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutpCollection.Merge(m, src)
}
func (m *OutpCollection) XXX_Size() int {
	return m.Size()
}
func (m *OutpCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_OutpCollection.DiscardUnknown(m)
}

var xxx_messageInfo_OutpCollection proto.InternalMessageInfo

func (m *OutpCollection) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OutpCollection) GetCards() []uint64 {
	if m != nil {
		return m.Cards
	}
	return nil
}

func (m *OutpCollection) GetArtist() string {
	if m != nil {
		return m.Artist
	}
	return ""
}

func (m *OutpCollection) GetStoryWriter() string {
	if m != nil {
		return m.StoryWriter
	}
	return ""
}

func (m *OutpCollection) GetContributors() []string {
	if m != nil {
		return m.Contributors
	}
	return nil
}

func (m *OutpCollection) GetStory() string {
	if m != nil {
		return m.Story
	}
	return ""
}

func (m *OutpCollection) GetArtwork() string {
	if m != nil {
		return m.Artwork
	}
	return ""
}

func (m *OutpCollection) GetStatus() CStatus {
	if m != nil {
		return m.Status
	}
	return CStatus_design
}

func (m *OutpCollection) GetTimeStamp() int64 {
	if m != nil {
		return m.TimeStamp
	}
	return 0
}

func init() {
	proto.RegisterEnum("DecentralCardGame.cardchain.cardchain.CStatus", CStatus_name, CStatus_value)
	proto.RegisterType((*Collection)(nil), "DecentralCardGame.cardchain.cardchain.Collection")
	proto.RegisterType((*OutpCollection)(nil), "DecentralCardGame.cardchain.cardchain.OutpCollection")
}

func init() {
	proto.RegisterFile("cardchain/cardchain/collection.proto", fileDescriptor_bfff38131151e449)
}

var fileDescriptor_bfff38131151e449 = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x93, 0x31, 0x8b, 0xd4, 0x40,
	0x14, 0xc7, 0x33, 0x9b, 0x5c, 0xf6, 0xf2, 0x3c, 0x8f, 0x30, 0x88, 0x4c, 0x21, 0x21, 0x1c, 0x0a,
	0xc1, 0x22, 0x0b, 0xda, 0x58, 0x59, 0x18, 0x51, 0xac, 0x84, 0xb9, 0x42, 0xb0, 0x9b, 0x9d, 0x8c,
	0xb7, 0x83, 0x49, 0x26, 0xcc, 0xbc, 0x9c, 0x9e, 0x9f, 0xc2, 0x4f, 0xe2, 0xe7, 0xb0, 0xdc, 0xd2,
	0x52, 0x76, 0xbf, 0x88, 0x64, 0x36, 0xee, 0xae, 0xd8, 0x58, 0x09, 0xd7, 0xbd, 0xff, 0x7f, 0xde,
	0x9f, 0x37, 0xfc, 0xe0, 0x0f, 0x0f, 0xa5, 0xb0, 0xb5, 0x5c, 0x09, 0xdd, 0x2d, 0x8e, 0x26, 0xd3,
	0x34, 0x4a, 0xa2, 0x36, 0x5d, 0xd9, 0x5b, 0x83, 0x86, 0x3e, 0x7a, 0xa9, 0xa4, 0xea, 0xd0, 0x8a,
	0xa6, 0x12, 0xb6, 0x7e, 0x2d, 0x5a, 0x55, 0xee, 0xb7, 0x0f, 0xd3, 0xc5, 0xb7, 0x19, 0x40, 0xb5,
	0xcf, 0x52, 0x0a, 0x51, 0x27, 0x5a, 0xc5, 0x48, 0x4e, 0x8a, 0x84, 0xfb, 0x99, 0xde, 0x83, 0x93,
	0x71, 0xdf, 0xb1, 0x59, 0x1e, 0x16, 0x11, 0xdf, 0x09, 0x7a, 0x1f, 0x62, 0x61, 0x51, 0x3b, 0x64,
	0xa1, 0xdf, 0x9d, 0x14, 0xcd, 0xe1, 0x8e, 0x43, 0x63, 0x6f, 0xde, 0x59, 0x8d, 0xca, 0xb2, 0xc8,
	0x3f, 0x1e, 0x5b, 0xf4, 0x02, 0xce, 0xa4, 0xe9, 0xd0, 0xea, 0xe5, 0x80, 0xc6, 0x3a, 0x76, 0x92,
	0x87, 0x45, 0xc2, 0xff, 0xf0, 0xc6, 0x9b, 0x3e, 0xc2, 0x62, 0x9f, 0xdf, 0x09, 0xfa, 0x00, 0x12,
	0x61, 0xf1, 0x93, 0xb1, 0x1f, 0xdf, 0xd4, 0x6c, 0x9e, 0x93, 0x22, 0xe2, 0x07, 0x83, 0xbe, 0x82,
	0xd8, 0xa1, 0xc0, 0xc1, 0xb1, 0xd3, 0x9c, 0x14, 0xe7, 0x4f, 0xca, 0xf2, 0x9f, 0x10, 0x94, 0xd5,
	0xa5, 0x4f, 0xf1, 0x29, 0x3d, 0x5e, 0x41, 0xdd, 0xaa, 0x4b, 0x14, 0x6d, 0xcf, 0x92, 0x9c, 0x14,
	0x21, 0x3f, 0x18, 0x23, 0xb0, 0xf3, 0xb7, 0x03, 0xf6, 0xb7, 0x06, 0x1a, 0x83, 0xf9, 0xc4, 0xc8,
	0x23, 0x4b, 0xf8, 0x6f, 0xf9, 0x7f, 0x80, 0x3d, 0x7e, 0x0e, 0xf3, 0x29, 0x40, 0x01, 0xe2, 0x5a,
	0x39, 0x7d, 0xd5, 0xa5, 0x01, 0xbd, 0x0b, 0xc9, 0x07, 0xdd, 0x89, 0x46, 0x7f, 0x51, 0x75, 0x4a,
	0xc6, 0x27, 0x21, 0x51, 0x5f, 0xab, 0x74, 0x46, 0xcf, 0xe0, 0x54, 0x58, 0xb9, 0xd2, 0xd7, 0xaa,
	0x4e, 0xc3, 0x17, 0xfc, 0xfb, 0x26, 0x23, 0xeb, 0x4d, 0x46, 0x7e, 0x6e, 0x32, 0xf2, 0x75, 0x9b,
	0x05, 0xeb, 0x6d, 0x16, 0xfc, 0xd8, 0x66, 0xc1, 0xfb, 0x67, 0x57, 0x1a, 0x57, 0xc3, 0xb2, 0x94,
	0xa6, 0x5d, 0xfc, 0xf5, 0xf3, 0x45, 0xb5, 0xef, 0xc6, 0xe7, 0xa3, 0x9e, 0xe0, 0x4d, 0xaf, 0xdc,
	0x32, 0xf6, 0x1d, 0x79, 0xfa, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xb9, 0xa6, 0x10, 0x41, 0x4b, 0x03,
	0x00, 0x00,
}

func (m *Collection) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Collection) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Collection) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TimeStamp != 0 {
		i = encodeVarintCollection(dAtA, i, uint64(m.TimeStamp))
		i--
		dAtA[i] = 0x48
	}
	if m.Status != 0 {
		i = encodeVarintCollection(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x40
	}
	if m.ArtworkId != 0 {
		i = encodeVarintCollection(dAtA, i, uint64(m.ArtworkId))
		i--
		dAtA[i] = 0x38
	}
	if len(m.Story) > 0 {
		i -= len(m.Story)
		copy(dAtA[i:], m.Story)
		i = encodeVarintCollection(dAtA, i, uint64(len(m.Story)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Contributors) > 0 {
		for iNdEx := len(m.Contributors) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Contributors[iNdEx])
			copy(dAtA[i:], m.Contributors[iNdEx])
			i = encodeVarintCollection(dAtA, i, uint64(len(m.Contributors[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.StoryWriter) > 0 {
		i -= len(m.StoryWriter)
		copy(dAtA[i:], m.StoryWriter)
		i = encodeVarintCollection(dAtA, i, uint64(len(m.StoryWriter)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Artist) > 0 {
		i -= len(m.Artist)
		copy(dAtA[i:], m.Artist)
		i = encodeVarintCollection(dAtA, i, uint64(len(m.Artist)))
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
		i = encodeVarintCollection(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintCollection(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *OutpCollection) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OutpCollection) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OutpCollection) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TimeStamp != 0 {
		i = encodeVarintCollection(dAtA, i, uint64(m.TimeStamp))
		i--
		dAtA[i] = 0x48
	}
	if m.Status != 0 {
		i = encodeVarintCollection(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x40
	}
	if len(m.Artwork) > 0 {
		i -= len(m.Artwork)
		copy(dAtA[i:], m.Artwork)
		i = encodeVarintCollection(dAtA, i, uint64(len(m.Artwork)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Story) > 0 {
		i -= len(m.Story)
		copy(dAtA[i:], m.Story)
		i = encodeVarintCollection(dAtA, i, uint64(len(m.Story)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Contributors) > 0 {
		for iNdEx := len(m.Contributors) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Contributors[iNdEx])
			copy(dAtA[i:], m.Contributors[iNdEx])
			i = encodeVarintCollection(dAtA, i, uint64(len(m.Contributors[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.StoryWriter) > 0 {
		i -= len(m.StoryWriter)
		copy(dAtA[i:], m.StoryWriter)
		i = encodeVarintCollection(dAtA, i, uint64(len(m.StoryWriter)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Artist) > 0 {
		i -= len(m.Artist)
		copy(dAtA[i:], m.Artist)
		i = encodeVarintCollection(dAtA, i, uint64(len(m.Artist)))
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
		i = encodeVarintCollection(dAtA, i, uint64(j3))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintCollection(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCollection(dAtA []byte, offset int, v uint64) int {
	offset -= sovCollection(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Collection) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovCollection(uint64(l))
	}
	if len(m.Cards) > 0 {
		l = 0
		for _, e := range m.Cards {
			l += sovCollection(uint64(e))
		}
		n += 1 + sovCollection(uint64(l)) + l
	}
	l = len(m.Artist)
	if l > 0 {
		n += 1 + l + sovCollection(uint64(l))
	}
	l = len(m.StoryWriter)
	if l > 0 {
		n += 1 + l + sovCollection(uint64(l))
	}
	if len(m.Contributors) > 0 {
		for _, s := range m.Contributors {
			l = len(s)
			n += 1 + l + sovCollection(uint64(l))
		}
	}
	l = len(m.Story)
	if l > 0 {
		n += 1 + l + sovCollection(uint64(l))
	}
	if m.ArtworkId != 0 {
		n += 1 + sovCollection(uint64(m.ArtworkId))
	}
	if m.Status != 0 {
		n += 1 + sovCollection(uint64(m.Status))
	}
	if m.TimeStamp != 0 {
		n += 1 + sovCollection(uint64(m.TimeStamp))
	}
	return n
}

func (m *OutpCollection) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovCollection(uint64(l))
	}
	if len(m.Cards) > 0 {
		l = 0
		for _, e := range m.Cards {
			l += sovCollection(uint64(e))
		}
		n += 1 + sovCollection(uint64(l)) + l
	}
	l = len(m.Artist)
	if l > 0 {
		n += 1 + l + sovCollection(uint64(l))
	}
	l = len(m.StoryWriter)
	if l > 0 {
		n += 1 + l + sovCollection(uint64(l))
	}
	if len(m.Contributors) > 0 {
		for _, s := range m.Contributors {
			l = len(s)
			n += 1 + l + sovCollection(uint64(l))
		}
	}
	l = len(m.Story)
	if l > 0 {
		n += 1 + l + sovCollection(uint64(l))
	}
	l = len(m.Artwork)
	if l > 0 {
		n += 1 + l + sovCollection(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovCollection(uint64(m.Status))
	}
	if m.TimeStamp != 0 {
		n += 1 + sovCollection(uint64(m.TimeStamp))
	}
	return n
}

func sovCollection(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCollection(x uint64) (n int) {
	return sovCollection(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Collection) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCollection
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
			return fmt.Errorf("proto: Collection: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Collection: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
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
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
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
						return ErrIntOverflowCollection
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
						return ErrIntOverflowCollection
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
					return ErrInvalidLengthCollection
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthCollection
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
							return ErrIntOverflowCollection
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
					return ErrIntOverflowCollection
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
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
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
					return ErrIntOverflowCollection
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
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
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
					return ErrIntOverflowCollection
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
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
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
					return ErrIntOverflowCollection
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
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
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
					return ErrIntOverflowCollection
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
					return ErrIntOverflowCollection
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
					return ErrIntOverflowCollection
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
			skippy, err := skipCollection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCollection
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
func (m *OutpCollection) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCollection
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
			return fmt.Errorf("proto: OutpCollection: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OutpCollection: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollection
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
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
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
						return ErrIntOverflowCollection
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
						return ErrIntOverflowCollection
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
					return ErrInvalidLengthCollection
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthCollection
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
							return ErrIntOverflowCollection
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
					return ErrIntOverflowCollection
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
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
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
					return ErrIntOverflowCollection
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
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
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
					return ErrIntOverflowCollection
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
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
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
					return ErrIntOverflowCollection
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
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
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
					return ErrIntOverflowCollection
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
				return ErrInvalidLengthCollection
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollection
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
					return ErrIntOverflowCollection
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
					return ErrIntOverflowCollection
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
			skippy, err := skipCollection(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCollection
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
func skipCollection(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCollection
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
					return 0, ErrIntOverflowCollection
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
					return 0, ErrIntOverflowCollection
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
				return 0, ErrInvalidLengthCollection
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCollection
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCollection
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCollection        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCollection          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCollection = fmt.Errorf("proto: unexpected end of group")
)
