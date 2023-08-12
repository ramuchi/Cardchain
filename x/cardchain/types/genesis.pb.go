// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cardchain/cardchain/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

// GenesisState defines the cardchain module's genesis state.
type GenesisState struct {
	Params           Params                                  `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	CardRecords      []*Card                                 `protobuf:"bytes,2,rep,name=cardRecords,proto3" json:"cardRecords,omitempty"`
	Users            []*User                                 `protobuf:"bytes,3,rep,name=users,proto3" json:"users,omitempty"`
	Addresses        []string                                `protobuf:"bytes,4,rep,name=addresses,proto3" json:"addresses,omitempty"`
	Matches          []*Match                                `protobuf:"bytes,6,rep,name=matches,proto3" json:"matches,omitempty"`
	Collections      []*Collection                           `protobuf:"bytes,7,rep,name=collections,proto3" json:"collections,omitempty"`
	SellOffers       []*SellOffer                            `protobuf:"bytes,8,rep,name=sellOffers,proto3" json:"sellOffers,omitempty"`
	Pools            []*types.Coin                           `protobuf:"bytes,9,rep,name=pools,proto3" json:"pools,omitempty"`
	CardAuctionPrice github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,11,opt,name=cardAuctionPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"cardAuctionPrice"`
	Councils         []*Council                              `protobuf:"bytes,12,rep,name=councils,proto3" json:"councils,omitempty"`
	RunningAverages  []*RunningAverage                       `protobuf:"bytes,13,rep,name=RunningAverages,proto3" json:"RunningAverages,omitempty"`
	Images           []*Image                                `protobuf:"bytes,14,rep,name=images,proto3" json:"images,omitempty"`
	Servers          []*Server                               `protobuf:"bytes,15,rep,name=Servers,proto3" json:"Servers,omitempty"`
	LastCardModified *TimeStamp                              `protobuf:"bytes,16,opt,name=lastCardModified,proto3" json:"lastCardModified,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4e78aa6e403ddd4, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetCardRecords() []*Card {
	if m != nil {
		return m.CardRecords
	}
	return nil
}

func (m *GenesisState) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *GenesisState) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func (m *GenesisState) GetMatches() []*Match {
	if m != nil {
		return m.Matches
	}
	return nil
}

func (m *GenesisState) GetCollections() []*Collection {
	if m != nil {
		return m.Collections
	}
	return nil
}

func (m *GenesisState) GetSellOffers() []*SellOffer {
	if m != nil {
		return m.SellOffers
	}
	return nil
}

func (m *GenesisState) GetPools() []*types.Coin {
	if m != nil {
		return m.Pools
	}
	return nil
}

func (m *GenesisState) GetCouncils() []*Council {
	if m != nil {
		return m.Councils
	}
	return nil
}

func (m *GenesisState) GetRunningAverages() []*RunningAverage {
	if m != nil {
		return m.RunningAverages
	}
	return nil
}

func (m *GenesisState) GetImages() []*Image {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *GenesisState) GetServers() []*Server {
	if m != nil {
		return m.Servers
	}
	return nil
}

func (m *GenesisState) GetLastCardModified() *TimeStamp {
	if m != nil {
		return m.LastCardModified
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "DecentralCardGame.cardchain.cardchain.GenesisState")
}

func init() { proto.RegisterFile("cardchain/cardchain/genesis.proto", fileDescriptor_c4e78aa6e403ddd4) }

var fileDescriptor_c4e78aa6e403ddd4 = []byte{
	// 616 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x1b, 0xb6, 0x75, 0xab, 0x3b, 0xd8, 0x64, 0x71, 0x61, 0x26, 0x94, 0x65, 0x08, 0x44,
	0x11, 0x2c, 0x61, 0x43, 0x48, 0xdc, 0xee, 0x8f, 0x98, 0x00, 0x4d, 0x4c, 0x2e, 0xdc, 0x00, 0x52,
	0xe5, 0x26, 0x6e, 0x6a, 0x91, 0xc4, 0x95, 0x4f, 0x5a, 0xc1, 0x5b, 0xf0, 0x0c, 0x3c, 0xcd, 0x2e,
	0x77, 0x89, 0xb8, 0x98, 0x50, 0xfb, 0x22, 0xc8, 0x4e, 0x9a, 0x16, 0x1a, 0x55, 0xd9, 0x55, 0xad,
	0x9e, 0xf3, 0xfd, 0xce, 0x89, 0xcf, 0x77, 0x8c, 0xf6, 0x7c, 0xa6, 0x02, 0xbf, 0xcf, 0x44, 0xe2,
	0xcd, 0x4e, 0x21, 0x4f, 0x38, 0x08, 0x70, 0x07, 0x4a, 0xa6, 0x12, 0x3f, 0x3a, 0xe5, 0x3e, 0x4f,
	0x52, 0xc5, 0xa2, 0x13, 0xa6, 0x82, 0x33, 0x16, 0x73, 0xb7, 0x48, 0x9d, 0x9d, 0x76, 0xee, 0x86,
	0x32, 0x94, 0x46, 0xe1, 0xe9, 0x53, 0x26, 0xde, 0x71, 0xca, 0xf8, 0x03, 0xa6, 0x58, 0x9c, 0xe3,
	0x77, 0xec, 0xb2, 0x0c, 0x7d, 0x5a, 0x16, 0x1f, 0x02, 0x57, 0x79, 0x7c, 0xb7, 0x2c, 0x1e, 0xb3,
	0xd4, 0xef, 0xe7, 0x09, 0x0f, 0x4b, 0x0b, 0xc8, 0x28, 0xe2, 0x7e, 0x2a, 0x64, 0xb2, 0x2c, 0x0b,
	0x78, 0x14, 0x75, 0x64, 0xaf, 0x57, 0x14, 0x7b, 0x52, 0x96, 0xa5, 0x86, 0x49, 0x22, 0x92, 0xb0,
	0xc3, 0x46, 0x5c, 0xb1, 0x90, 0xe7, 0xa9, 0x7b, 0xe5, 0x65, 0x87, 0x89, 0x2f, 0xa2, 0x65, 0xad,
	0x8b, 0x78, 0xc6, 0x70, 0xca, 0x9b, 0x52, 0xa3, 0xa2, 0x21, 0xdb, 0x97, 0x10, 0x4b, 0xf0, 0xba,
	0x0c, 0xb8, 0x37, 0x3a, 0xe8, 0xf2, 0x94, 0x1d, 0x78, 0xbe, 0x14, 0xf9, 0x67, 0x3d, 0xf8, 0xb9,
	0x81, 0x36, 0xcf, 0xb2, 0x71, 0xb6, 0x53, 0x96, 0x72, 0xfc, 0x0e, 0xd5, 0xb3, 0xeb, 0x27, 0x96,
	0x63, 0xb5, 0x9a, 0x87, 0xfb, 0x6e, 0xa5, 0xf1, 0xba, 0x17, 0x46, 0x74, 0xbc, 0x7a, 0x79, 0xbd,
	0x5b, 0xa3, 0x39, 0x02, 0x9f, 0xa3, 0xa6, 0xce, 0xa0, 0xdc, 0x97, 0x2a, 0x00, 0x72, 0xcb, 0x59,
	0x69, 0x35, 0x0f, 0x9f, 0x56, 0x24, 0xea, 0x20, 0x9d, 0xd7, 0xe3, 0x23, 0xb4, 0xa6, 0x07, 0x0b,
	0x64, 0xe5, 0x46, 0xa0, 0x8f, 0xc0, 0x15, 0xcd, 0x94, 0xf8, 0x3e, 0x6a, 0xb0, 0x20, 0x50, 0x1c,
	0x80, 0x03, 0x59, 0x75, 0x56, 0x5a, 0x0d, 0x3a, 0xfb, 0x03, 0xbf, 0x46, 0xeb, 0xc6, 0x19, 0x1c,
	0x48, 0xdd, 0x94, 0x78, 0x56, 0xb1, 0xc4, 0xb9, 0x56, 0xd1, 0xa9, 0x18, 0xb7, 0x51, 0x73, 0x66,
	0x20, 0x20, 0xeb, 0x86, 0x75, 0x50, 0xf5, 0xbb, 0x0b, 0x25, 0x9d, 0xa7, 0xe0, 0x0b, 0x84, 0xb4,
	0xdf, 0xde, 0x6b, 0xbb, 0x01, 0xd9, 0x30, 0xcc, 0xe7, 0x15, 0x99, 0xed, 0xa9, 0x90, 0xce, 0x31,
	0xb0, 0x87, 0xd6, 0x06, 0x52, 0x46, 0x40, 0x1a, 0x06, 0x76, 0xcf, 0xcd, 0xcc, 0xe2, 0x6a, 0xb3,
	0xb8, 0xb9, 0x59, 0xdc, 0x13, 0x29, 0x12, 0x9a, 0xe5, 0xe1, 0xcf, 0x68, 0x5b, 0x33, 0x8f, 0x86,
	0xa6, 0xa5, 0x0b, 0x25, 0x7c, 0x4e, 0x9a, 0x8e, 0xd5, 0x6a, 0x1c, 0x7b, 0x7a, 0xee, 0xbf, 0xaf,
	0x77, 0x1f, 0x87, 0x22, 0xed, 0x0f, 0xbb, 0xae, 0x2f, 0x63, 0x2f, 0xb7, 0x5e, 0xf6, 0xb3, 0x0f,
	0xc1, 0x57, 0x2f, 0xfd, 0x3e, 0xe0, 0x90, 0x11, 0x17, 0x40, 0xf8, 0x2d, 0xda, 0xc8, 0xed, 0x0f,
	0x64, 0xd3, 0x34, 0xe4, 0x56, 0xbe, 0x31, 0x23, 0xa3, 0x85, 0x1e, 0x77, 0xd0, 0x16, 0xcd, 0xb6,
	0xee, 0x28, 0x5b, 0x3a, 0x20, 0xb7, 0x0d, 0xf2, 0x65, 0x45, 0xe4, 0xbf, 0x6a, 0xfa, 0x3f, 0x0d,
	0x9f, 0xa2, 0xba, 0x59, 0x44, 0x20, 0x77, 0x6e, 0x64, 0x94, 0x37, 0x5a, 0x44, 0x73, 0x2d, 0x3e,
	0x43, 0xeb, 0x6d, 0xb3, 0xad, 0x40, 0xb6, 0x0c, 0x66, 0xbf, 0xf2, 0x3c, 0xb5, 0x8a, 0x4e, 0xd5,
	0xf8, 0x0b, 0xda, 0x8e, 0x18, 0xa4, 0x5a, 0x73, 0x2e, 0x03, 0xd1, 0x13, 0x3c, 0x20, 0xdb, 0x66,
	0x7f, 0xab, 0x3a, 0xe4, 0x83, 0x88, 0x79, 0x3b, 0x65, 0xf1, 0x80, 0x2e, 0x90, 0x8e, 0xe9, 0xe5,
	0xd8, 0xb6, 0xae, 0xc6, 0xb6, 0xf5, 0x67, 0x6c, 0x5b, 0x3f, 0x26, 0x76, 0xed, 0x6a, 0x62, 0xd7,
	0x7e, 0x4d, 0xec, 0xda, 0xa7, 0x57, 0x73, 0xe3, 0x5e, 0xa8, 0xe3, 0x9d, 0x14, 0x6f, 0xd2, 0xb7,
	0xb9, 0xf7, 0xc9, 0x98, 0xa0, 0x5b, 0x37, 0xef, 0xcf, 0x8b, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xeb, 0xf4, 0x84, 0x3d, 0x61, 0x06, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastCardModified != nil {
		{
			size, err := m.LastCardModified.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x82
	}
	if len(m.Servers) > 0 {
		for iNdEx := len(m.Servers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Servers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x7a
		}
	}
	if len(m.Images) > 0 {
		for iNdEx := len(m.Images) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Images[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x72
		}
	}
	if len(m.RunningAverages) > 0 {
		for iNdEx := len(m.RunningAverages) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RunningAverages[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x6a
		}
	}
	if len(m.Councils) > 0 {
		for iNdEx := len(m.Councils) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Councils[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x62
		}
	}
	{
		size := m.CardAuctionPrice.Size()
		i -= size
		if _, err := m.CardAuctionPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	if len(m.Pools) > 0 {
		for iNdEx := len(m.Pools) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Pools[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x4a
		}
	}
	if len(m.SellOffers) > 0 {
		for iNdEx := len(m.SellOffers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SellOffers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.Collections) > 0 {
		for iNdEx := len(m.Collections) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Collections[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.Matches) > 0 {
		for iNdEx := len(m.Matches) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Matches[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Addresses) > 0 {
		for iNdEx := len(m.Addresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Addresses[iNdEx])
			copy(dAtA[i:], m.Addresses[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.Addresses[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Users) > 0 {
		for iNdEx := len(m.Users) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Users[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.CardRecords) > 0 {
		for iNdEx := len(m.CardRecords) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CardRecords[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.CardRecords) > 0 {
		for _, e := range m.CardRecords {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Users) > 0 {
		for _, e := range m.Users {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Addresses) > 0 {
		for _, s := range m.Addresses {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Matches) > 0 {
		for _, e := range m.Matches {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Collections) > 0 {
		for _, e := range m.Collections {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.SellOffers) > 0 {
		for _, e := range m.SellOffers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Pools) > 0 {
		for _, e := range m.Pools {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.CardAuctionPrice.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Councils) > 0 {
		for _, e := range m.Councils {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.RunningAverages) > 0 {
		for _, e := range m.RunningAverages {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Images) > 0 {
		for _, e := range m.Images {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Servers) > 0 {
		for _, e := range m.Servers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.LastCardModified != nil {
		l = m.LastCardModified.Size()
		n += 2 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CardRecords", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CardRecords = append(m.CardRecords, &Card{})
			if err := m.CardRecords[len(m.CardRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Users", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Users = append(m.Users, &User{})
			if err := m.Users[len(m.Users)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addresses = append(m.Addresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Matches", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Matches = append(m.Matches, &Match{})
			if err := m.Matches[len(m.Matches)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Collections", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Collections = append(m.Collections, &Collection{})
			if err := m.Collections[len(m.Collections)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SellOffers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SellOffers = append(m.SellOffers, &SellOffer{})
			if err := m.SellOffers[len(m.SellOffers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pools", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pools = append(m.Pools, &types.Coin{})
			if err := m.Pools[len(m.Pools)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CardAuctionPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CardAuctionPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Councils", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Councils = append(m.Councils, &Council{})
			if err := m.Councils[len(m.Councils)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RunningAverages", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RunningAverages = append(m.RunningAverages, &RunningAverage{})
			if err := m.RunningAverages[len(m.RunningAverages)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Images", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Images = append(m.Images, &Image{})
			if err := m.Images[len(m.Images)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Servers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Servers = append(m.Servers, &Server{})
			if err := m.Servers[len(m.Servers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastCardModified", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LastCardModified == nil {
				m.LastCardModified = &TimeStamp{}
			}
			if err := m.LastCardModified.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
