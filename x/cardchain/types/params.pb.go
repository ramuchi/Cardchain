// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cardchain/cardchain/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// Params defines the parameters for the module.
type Params struct {
	VotingRightsExpirationTime      int64                                   `protobuf:"varint,1,opt,name=votingRightsExpirationTime,proto3" json:"votingRightsExpirationTime,omitempty"`
	SetSize                         uint64                                  `protobuf:"varint,2,opt,name=setSize,proto3" json:"setSize,omitempty"`
	SetPrice                        github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,3,opt,name=setPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"setPrice"`
	ActiveSetsAmount                uint64                                  `protobuf:"varint,4,opt,name=activeSetsAmount,proto3" json:"activeSetsAmount,omitempty"`
	SetCreationFee                  github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,5,opt,name=setCreationFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"setCreationFee"`
	CollateralDeposit               github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,6,opt,name=collateralDeposit,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"collateralDeposit"`
	WinnerReward                    int64                                   `protobuf:"varint,7,opt,name=winnerReward,proto3" json:"winnerReward,omitempty"`
	HourlyFaucet                    github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,9,opt,name=hourlyFaucet,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"hourlyFaucet"`
	InflationRate                   string                                  `protobuf:"bytes,10,opt,name=inflationRate,proto3" json:"inflationRate,omitempty"`
	RaresPerPack                    uint64                                  `protobuf:"varint,11,opt,name=raresPerPack,proto3" json:"raresPerPack,omitempty"`
	CommonsPerPack                  uint64                                  `protobuf:"varint,12,opt,name=commonsPerPack,proto3" json:"commonsPerPack,omitempty"`
	UnCommonsPerPack                uint64                                  `protobuf:"varint,13,opt,name=unCommonsPerPack,proto3" json:"unCommonsPerPack,omitempty"`
	TrialPeriod                     uint64                                  `protobuf:"varint,14,opt,name=trialPeriod,proto3" json:"trialPeriod,omitempty"`
	GameVoteRatio                   int64                                   `protobuf:"varint,15,opt,name=gameVoteRatio,proto3" json:"gameVoteRatio,omitempty"`
	CardAuctionPriceReductionPeriod int64                                   `protobuf:"varint,16,opt,name=cardAuctionPriceReductionPeriod,proto3" json:"cardAuctionPriceReductionPeriod,omitempty"`
	AirDropValue                    github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,17,opt,name=airDropValue,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"airDropValue"`
	AirDropMaxBlockHeight           int64                                   `protobuf:"varint,18,opt,name=airDropMaxBlockHeight,proto3" json:"airDropMaxBlockHeight,omitempty"`
	TrialVoteReward                 github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,19,opt,name=trialVoteReward,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"trialVoteReward"`
	VotePoolFraction                int64                                   `protobuf:"varint,20,opt,name=votePoolFraction,proto3" json:"votePoolFraction,omitempty"`
	VotingRewardCap                 int64                                   `protobuf:"varint,8,opt,name=votingRewardCap,proto3" json:"votingRewardCap,omitempty"`
	MatchWorkerDelay                uint64                                  `protobuf:"varint,21,opt,name=matchWorkerDelay,proto3" json:"matchWorkerDelay,omitempty"`
	RareDropRatio                   uint64                                  `protobuf:"varint,22,opt,name=rareDropRatio,proto3" json:"rareDropRatio,omitempty"`
	ExceptionalDropRatio            uint64                                  `protobuf:"varint,23,opt,name=exceptionalDropRatio,proto3" json:"exceptionalDropRatio,omitempty"`
	UniqueDropRatio                 uint64                                  `protobuf:"varint,24,opt,name=uniqueDropRatio,proto3" json:"uniqueDropRatio,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_8843e481ee664a23, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetVotingRightsExpirationTime() int64 {
	if m != nil {
		return m.VotingRightsExpirationTime
	}
	return 0
}

func (m *Params) GetSetSize() uint64 {
	if m != nil {
		return m.SetSize
	}
	return 0
}

func (m *Params) GetActiveSetsAmount() uint64 {
	if m != nil {
		return m.ActiveSetsAmount
	}
	return 0
}

func (m *Params) GetWinnerReward() int64 {
	if m != nil {
		return m.WinnerReward
	}
	return 0
}

func (m *Params) GetInflationRate() string {
	if m != nil {
		return m.InflationRate
	}
	return ""
}

func (m *Params) GetRaresPerPack() uint64 {
	if m != nil {
		return m.RaresPerPack
	}
	return 0
}

func (m *Params) GetCommonsPerPack() uint64 {
	if m != nil {
		return m.CommonsPerPack
	}
	return 0
}

func (m *Params) GetUnCommonsPerPack() uint64 {
	if m != nil {
		return m.UnCommonsPerPack
	}
	return 0
}

func (m *Params) GetTrialPeriod() uint64 {
	if m != nil {
		return m.TrialPeriod
	}
	return 0
}

func (m *Params) GetGameVoteRatio() int64 {
	if m != nil {
		return m.GameVoteRatio
	}
	return 0
}

func (m *Params) GetCardAuctionPriceReductionPeriod() int64 {
	if m != nil {
		return m.CardAuctionPriceReductionPeriod
	}
	return 0
}

func (m *Params) GetAirDropMaxBlockHeight() int64 {
	if m != nil {
		return m.AirDropMaxBlockHeight
	}
	return 0
}

func (m *Params) GetVotePoolFraction() int64 {
	if m != nil {
		return m.VotePoolFraction
	}
	return 0
}

func (m *Params) GetVotingRewardCap() int64 {
	if m != nil {
		return m.VotingRewardCap
	}
	return 0
}

func (m *Params) GetMatchWorkerDelay() uint64 {
	if m != nil {
		return m.MatchWorkerDelay
	}
	return 0
}

func (m *Params) GetRareDropRatio() uint64 {
	if m != nil {
		return m.RareDropRatio
	}
	return 0
}

func (m *Params) GetExceptionalDropRatio() uint64 {
	if m != nil {
		return m.ExceptionalDropRatio
	}
	return 0
}

func (m *Params) GetUniqueDropRatio() uint64 {
	if m != nil {
		return m.UniqueDropRatio
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "DecentralCardGame.cardchain.cardchain.Params")
}

func init() { proto.RegisterFile("cardchain/cardchain/params.proto", fileDescriptor_8843e481ee664a23) }

var fileDescriptor_8843e481ee664a23 = []byte{
	// 651 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4d, 0x6f, 0x13, 0x3d,
	0x10, 0xce, 0xbe, 0xed, 0xdb, 0x0f, 0xf7, 0xdb, 0xb4, 0x60, 0xf5, 0x90, 0x44, 0x15, 0x1f, 0x11,
	0x12, 0x89, 0x04, 0x1c, 0x10, 0x07, 0xa4, 0x36, 0xa1, 0x54, 0x42, 0x48, 0xd1, 0x16, 0xb5, 0x02,
	0x89, 0xc3, 0xd4, 0x19, 0x12, 0x2b, 0xbb, 0xeb, 0xc5, 0xeb, 0x6d, 0x53, 0x7e, 0x05, 0x47, 0x8e,
	0xfc, 0x02, 0x7e, 0x47, 0x8f, 0x3d, 0x22, 0x0e, 0x15, 0x6a, 0xff, 0x08, 0xf2, 0x24, 0xa4, 0xd9,
	0x4d, 0x05, 0x52, 0x4e, 0x6b, 0x3f, 0xfb, 0xcc, 0xe3, 0x67, 0x66, 0xec, 0x61, 0x65, 0x09, 0xa6,
	0x25, 0x3b, 0xa0, 0xa2, 0xda, 0xf5, 0x2a, 0x06, 0x03, 0x61, 0x52, 0x8d, 0x8d, 0xb6, 0x9a, 0xdf,
	0x6b, 0xa0, 0xc4, 0xc8, 0x1a, 0x08, 0xea, 0x60, 0x5a, 0xaf, 0x20, 0xc4, 0xea, 0x90, 0x79, 0xbd,
	0xda, 0x5c, 0x6f, 0xeb, 0xb6, 0xa6, 0x88, 0x9a, 0x5b, 0xf5, 0x83, 0xb7, 0xbe, 0x33, 0x36, 0xd3,
	0x24, 0x35, 0xfe, 0x82, 0x6d, 0x1e, 0x6b, 0xab, 0xa2, 0xb6, 0xaf, 0xda, 0x1d, 0x9b, 0xbc, 0xec,
	0xc5, 0xca, 0x80, 0x55, 0x3a, 0x7a, 0xab, 0x42, 0x14, 0x5e, 0xd9, 0xab, 0x4c, 0xf9, 0x7f, 0x61,
	0x70, 0xc1, 0x66, 0x13, 0xb4, 0xfb, 0xea, 0x33, 0x8a, 0xff, 0xca, 0x5e, 0x65, 0xda, 0xff, 0xb3,
	0xe5, 0xaf, 0xd9, 0x5c, 0x82, 0xb6, 0x69, 0x94, 0x44, 0x31, 0x55, 0xf6, 0x2a, 0xf3, 0x3b, 0xb5,
	0xb3, 0x8b, 0x52, 0xe1, 0xe7, 0x45, 0xe9, 0x41, 0x5b, 0xd9, 0x4e, 0x7a, 0x54, 0x95, 0x3a, 0xac,
	0x49, 0x9d, 0x84, 0x3a, 0x19, 0x7c, 0x1e, 0x25, 0xad, 0x6e, 0xcd, 0x9e, 0xc6, 0x98, 0x54, 0xeb,
	0x5a, 0x45, 0xfe, 0x50, 0x80, 0x3f, 0x64, 0xab, 0x20, 0xad, 0x3a, 0xc6, 0x7d, 0xb4, 0xc9, 0x76,
	0xa8, 0xd3, 0xc8, 0x8a, 0x69, 0x3a, 0x6f, 0x0c, 0xe7, 0x87, 0x6c, 0x39, 0x41, 0x5b, 0x37, 0x48,
	0x2e, 0x77, 0x11, 0xc5, 0xff, 0x93, 0x1d, 0x9f, 0x93, 0xe1, 0x1f, 0xd8, 0x9a, 0xd4, 0x41, 0x00,
	0x16, 0x0d, 0x04, 0x0d, 0x8c, 0x75, 0xa2, 0xac, 0x98, 0x99, 0x4c, 0x7b, 0x5c, 0x89, 0x6f, 0xb1,
	0xc5, 0x13, 0x15, 0x45, 0x68, 0x7c, 0x3c, 0x01, 0xd3, 0x12, 0xb3, 0x54, 0xfc, 0x0c, 0xc6, 0xf7,
	0xd9, 0x62, 0x47, 0xa7, 0x26, 0x38, 0xdd, 0x85, 0x54, 0xa2, 0x15, 0xf3, 0x93, 0x9d, 0x9e, 0x11,
	0xe1, 0x77, 0xd9, 0x92, 0x8a, 0x3e, 0x06, 0x94, 0xa7, 0x0f, 0x16, 0x05, 0x73, 0xaa, 0x7e, 0x16,
	0x74, 0xf6, 0x0c, 0x18, 0x4c, 0x9a, 0x68, 0x9a, 0x20, 0xbb, 0x62, 0x81, 0xca, 0x9f, 0xc1, 0xf8,
	0x7d, 0xb6, 0x2c, 0x75, 0x18, 0xea, 0x68, 0xc8, 0x5a, 0x24, 0x56, 0x0e, 0x75, 0xed, 0x4c, 0xa3,
	0x7a, 0x96, 0xb9, 0xd4, 0x6f, 0x67, 0x1e, 0xe7, 0x65, 0xb6, 0x60, 0x8d, 0x82, 0xa0, 0x89, 0x46,
	0xe9, 0x96, 0x58, 0x26, 0xda, 0x28, 0xe4, 0xfc, 0xb7, 0x21, 0xc4, 0x03, 0x6d, 0xd1, 0x77, 0x7e,
	0xc5, 0x0a, 0x55, 0x2e, 0x0b, 0xf2, 0x3d, 0x56, 0x72, 0xef, 0x62, 0x3b, 0x95, 0x2e, 0x25, 0xba,
	0x56, 0x3e, 0xb6, 0x06, 0xbb, 0xbe, 0xf6, 0x2a, 0xc5, 0xfd, 0x8b, 0xe6, 0x9a, 0x00, 0xca, 0x34,
	0x8c, 0x8e, 0x0f, 0x20, 0x48, 0x51, 0xac, 0x4d, 0xd8, 0x84, 0x51, 0x11, 0xfe, 0x94, 0x6d, 0x0c,
	0xf6, 0x6f, 0xa0, 0xb7, 0x13, 0x68, 0xd9, 0xdd, 0x43, 0xf7, 0xe0, 0x04, 0x27, 0x53, 0x37, 0xff,
	0xe4, 0xef, 0xd8, 0x0a, 0x55, 0x82, 0xd2, 0xec, 0x5f, 0x9b, 0x5b, 0x93, 0xb9, 0xc9, 0xeb, 0xb8,
	0x1e, 0x1d, 0x6b, 0x8b, 0x4d, 0xad, 0x83, 0x5d, 0x03, 0x94, 0xbf, 0x58, 0x27, 0x2f, 0x63, 0x38,
	0xaf, 0xb0, 0x95, 0xc1, 0x8c, 0xa0, 0xd8, 0x3a, 0xc4, 0x62, 0x8e, 0xa8, 0x79, 0xd8, 0xa9, 0x86,
	0x60, 0x65, 0xe7, 0x50, 0x9b, 0x2e, 0x9a, 0x06, 0x06, 0x70, 0x2a, 0x36, 0xfa, 0x9d, 0xcf, 0xe3,
	0xae, 0xaf, 0xee, 0x76, 0xb9, 0xb4, 0xfb, 0x7d, 0xbd, 0x4d, 0xc4, 0x2c, 0xc8, 0x1f, 0xb3, 0x75,
	0xec, 0x49, 0x8c, 0x9d, 0x11, 0x08, 0xae, 0xc9, 0x77, 0x88, 0x7c, 0xe3, 0x3f, 0xe7, 0x37, 0x8d,
	0xd4, 0xa7, 0x74, 0x44, 0x5b, 0x10, 0x3d, 0x0f, 0x3f, 0x9f, 0xfe, 0xfa, 0xad, 0x54, 0xd8, 0xf1,
	0xcf, 0x2e, 0x8b, 0xde, 0xf9, 0x65, 0xd1, 0xfb, 0x75, 0x59, 0xf4, 0xbe, 0x5c, 0x15, 0x0b, 0xe7,
	0x57, 0xc5, 0xc2, 0x8f, 0xab, 0x62, 0xe1, 0xfd, 0xb3, 0x91, 0xfa, 0x8e, 0x8d, 0xe4, 0x5a, 0x7d,
	0x38, 0xbc, 0x7b, 0x23, 0x83, 0x9c, 0xaa, 0x7e, 0x34, 0x43, 0xb3, 0xf8, 0xc9, 0xef, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x7c, 0x7a, 0xa5, 0x82, 0xec, 0x05, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UniqueDropRatio != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.UniqueDropRatio))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xc0
	}
	if m.ExceptionalDropRatio != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ExceptionalDropRatio))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xb8
	}
	if m.RareDropRatio != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.RareDropRatio))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xb0
	}
	if m.MatchWorkerDelay != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MatchWorkerDelay))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xa8
	}
	if m.VotePoolFraction != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.VotePoolFraction))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0xa0
	}
	{
		size := m.TrialVoteReward.Size()
		i -= size
		if _, err := m.TrialVoteReward.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1
	i--
	dAtA[i] = 0x9a
	if m.AirDropMaxBlockHeight != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.AirDropMaxBlockHeight))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x90
	}
	{
		size := m.AirDropValue.Size()
		i -= size
		if _, err := m.AirDropValue.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1
	i--
	dAtA[i] = 0x8a
	if m.CardAuctionPriceReductionPeriod != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.CardAuctionPriceReductionPeriod))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x80
	}
	if m.GameVoteRatio != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.GameVoteRatio))
		i--
		dAtA[i] = 0x78
	}
	if m.TrialPeriod != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.TrialPeriod))
		i--
		dAtA[i] = 0x70
	}
	if m.UnCommonsPerPack != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.UnCommonsPerPack))
		i--
		dAtA[i] = 0x68
	}
	if m.CommonsPerPack != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.CommonsPerPack))
		i--
		dAtA[i] = 0x60
	}
	if m.RaresPerPack != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.RaresPerPack))
		i--
		dAtA[i] = 0x58
	}
	if len(m.InflationRate) > 0 {
		i -= len(m.InflationRate)
		copy(dAtA[i:], m.InflationRate)
		i = encodeVarintParams(dAtA, i, uint64(len(m.InflationRate)))
		i--
		dAtA[i] = 0x52
	}
	{
		size := m.HourlyFaucet.Size()
		i -= size
		if _, err := m.HourlyFaucet.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if m.VotingRewardCap != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.VotingRewardCap))
		i--
		dAtA[i] = 0x40
	}
	if m.WinnerReward != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.WinnerReward))
		i--
		dAtA[i] = 0x38
	}
	{
		size := m.CollateralDeposit.Size()
		i -= size
		if _, err := m.CollateralDeposit.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.SetCreationFee.Size()
		i -= size
		if _, err := m.SetCreationFee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.ActiveSetsAmount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.ActiveSetsAmount))
		i--
		dAtA[i] = 0x20
	}
	{
		size := m.SetPrice.Size()
		i -= size
		if _, err := m.SetPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.SetSize != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.SetSize))
		i--
		dAtA[i] = 0x10
	}
	if m.VotingRightsExpirationTime != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.VotingRightsExpirationTime))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.VotingRightsExpirationTime != 0 {
		n += 1 + sovParams(uint64(m.VotingRightsExpirationTime))
	}
	if m.SetSize != 0 {
		n += 1 + sovParams(uint64(m.SetSize))
	}
	l = m.SetPrice.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.ActiveSetsAmount != 0 {
		n += 1 + sovParams(uint64(m.ActiveSetsAmount))
	}
	l = m.SetCreationFee.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.CollateralDeposit.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.WinnerReward != 0 {
		n += 1 + sovParams(uint64(m.WinnerReward))
	}
	if m.VotingRewardCap != 0 {
		n += 1 + sovParams(uint64(m.VotingRewardCap))
	}
	l = m.HourlyFaucet.Size()
	n += 1 + l + sovParams(uint64(l))
	l = len(m.InflationRate)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.RaresPerPack != 0 {
		n += 1 + sovParams(uint64(m.RaresPerPack))
	}
	if m.CommonsPerPack != 0 {
		n += 1 + sovParams(uint64(m.CommonsPerPack))
	}
	if m.UnCommonsPerPack != 0 {
		n += 1 + sovParams(uint64(m.UnCommonsPerPack))
	}
	if m.TrialPeriod != 0 {
		n += 1 + sovParams(uint64(m.TrialPeriod))
	}
	if m.GameVoteRatio != 0 {
		n += 1 + sovParams(uint64(m.GameVoteRatio))
	}
	if m.CardAuctionPriceReductionPeriod != 0 {
		n += 2 + sovParams(uint64(m.CardAuctionPriceReductionPeriod))
	}
	l = m.AirDropValue.Size()
	n += 2 + l + sovParams(uint64(l))
	if m.AirDropMaxBlockHeight != 0 {
		n += 2 + sovParams(uint64(m.AirDropMaxBlockHeight))
	}
	l = m.TrialVoteReward.Size()
	n += 2 + l + sovParams(uint64(l))
	if m.VotePoolFraction != 0 {
		n += 2 + sovParams(uint64(m.VotePoolFraction))
	}
	if m.MatchWorkerDelay != 0 {
		n += 2 + sovParams(uint64(m.MatchWorkerDelay))
	}
	if m.RareDropRatio != 0 {
		n += 2 + sovParams(uint64(m.RareDropRatio))
	}
	if m.ExceptionalDropRatio != 0 {
		n += 2 + sovParams(uint64(m.ExceptionalDropRatio))
	}
	if m.UniqueDropRatio != 0 {
		n += 2 + sovParams(uint64(m.UniqueDropRatio))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingRightsExpirationTime", wireType)
			}
			m.VotingRightsExpirationTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VotingRightsExpirationTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SetSize", wireType)
			}
			m.SetSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SetSize |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SetPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SetPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActiveSetsAmount", wireType)
			}
			m.ActiveSetsAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ActiveSetsAmount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SetCreationFee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SetCreationFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralDeposit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CollateralDeposit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WinnerReward", wireType)
			}
			m.WinnerReward = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WinnerReward |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingRewardCap", wireType)
			}
			m.VotingRewardCap = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VotingRewardCap |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HourlyFaucet", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.HourlyFaucet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InflationRate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RaresPerPack", wireType)
			}
			m.RaresPerPack = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RaresPerPack |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommonsPerPack", wireType)
			}
			m.CommonsPerPack = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CommonsPerPack |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnCommonsPerPack", wireType)
			}
			m.UnCommonsPerPack = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnCommonsPerPack |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrialPeriod", wireType)
			}
			m.TrialPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TrialPeriod |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GameVoteRatio", wireType)
			}
			m.GameVoteRatio = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GameVoteRatio |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 16:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CardAuctionPriceReductionPeriod", wireType)
			}
			m.CardAuctionPriceReductionPeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CardAuctionPriceReductionPeriod |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AirDropValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AirDropValue.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 18:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AirDropMaxBlockHeight", wireType)
			}
			m.AirDropMaxBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AirDropMaxBlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 19:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrialVoteReward", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TrialVoteReward.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 20:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotePoolFraction", wireType)
			}
			m.VotePoolFraction = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VotePoolFraction |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 21:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MatchWorkerDelay", wireType)
			}
			m.MatchWorkerDelay = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MatchWorkerDelay |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 22:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RareDropRatio", wireType)
			}
			m.RareDropRatio = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RareDropRatio |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 23:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExceptionalDropRatio", wireType)
			}
			m.ExceptionalDropRatio = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExceptionalDropRatio |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 24:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UniqueDropRatio", wireType)
			}
			m.UniqueDropRatio = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UniqueDropRatio |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
