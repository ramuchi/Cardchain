// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cardchain/cardchain/collection_proposal.proto

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

type CollectionProposal struct {
	Title        string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description  string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	CollectionId uint64 `protobuf:"varint,3,opt,name=collectionId,proto3" json:"collectionId,omitempty"`
}

func (m *CollectionProposal) Reset()         { *m = CollectionProposal{} }
func (m *CollectionProposal) String() string { return proto.CompactTextString(m) }
func (*CollectionProposal) ProtoMessage()    {}
func (*CollectionProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1f16944bc7ec925, []int{0}
}
func (m *CollectionProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CollectionProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CollectionProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CollectionProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionProposal.Merge(m, src)
}
func (m *CollectionProposal) XXX_Size() int {
	return m.Size()
}
func (m *CollectionProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionProposal.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionProposal proto.InternalMessageInfo

func (m *CollectionProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CollectionProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CollectionProposal) GetCollectionId() uint64 {
	if m != nil {
		return m.CollectionId
	}
	return 0
}

func init() {
	proto.RegisterType((*CollectionProposal)(nil), "DecentralCardGame.cardchain.cardchain.CollectionProposal")
}

func init() {
	proto.RegisterFile("cardchain/cardchain/collection_proposal.proto", fileDescriptor_d1f16944bc7ec925)
}

var fileDescriptor_d1f16944bc7ec925 = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4d, 0x4e, 0x2c, 0x4a,
	0x49, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x47, 0x62, 0xe5, 0xe7, 0xe4, 0xa4, 0x26, 0x97, 0x64, 0xe6,
	0xe7, 0xc5, 0x17, 0x14, 0xe5, 0x17, 0xe4, 0x17, 0x27, 0xe6, 0xe8, 0x15, 0x14, 0xe5, 0x97, 0xe4,
	0x0b, 0xa9, 0xba, 0xa4, 0x26, 0xa7, 0xe6, 0x95, 0x14, 0x25, 0xe6, 0x38, 0x27, 0x16, 0xa5, 0xb8,
	0x27, 0xe6, 0xa6, 0xea, 0xc1, 0xb5, 0x21, 0x58, 0x4a, 0x05, 0x5c, 0x42, 0xce, 0x70, 0x33, 0x02,
	0xa0, 0x46, 0x08, 0x89, 0x70, 0xb1, 0x96, 0x64, 0x96, 0xe4, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a,
	0x70, 0x06, 0x41, 0x38, 0x42, 0x0a, 0x5c, 0xdc, 0x29, 0xa9, 0xc5, 0xc9, 0x45, 0x99, 0x05, 0x20,
	0xc5, 0x12, 0x4c, 0x60, 0x39, 0x64, 0x21, 0x21, 0x25, 0x2e, 0x1e, 0x84, 0x8b, 0x3c, 0x53, 0x24,
	0x98, 0x15, 0x18, 0x35, 0x58, 0x82, 0x50, 0xc4, 0x9c, 0x82, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0,
	0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8,
	0xf1, 0x58, 0x8e, 0x21, 0xca, 0x22, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57,
	0x1f, 0xc3, 0xf5, 0xfa, 0xce, 0x70, 0x4f, 0x57, 0x20, 0x05, 0x40, 0x49, 0x65, 0x41, 0x6a, 0x71,
	0x12, 0x1b, 0xd8, 0xcf, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x28, 0x9a, 0x21, 0x24,
	0x01, 0x00, 0x00,
}

func (m *CollectionProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CollectionProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CollectionProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CollectionId != 0 {
		i = encodeVarintCollectionProposal(dAtA, i, uint64(m.CollectionId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintCollectionProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintCollectionProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCollectionProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovCollectionProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CollectionProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovCollectionProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovCollectionProposal(uint64(l))
	}
	if m.CollectionId != 0 {
		n += 1 + sovCollectionProposal(uint64(m.CollectionId))
	}
	return n
}

func sovCollectionProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCollectionProposal(x uint64) (n int) {
	return sovCollectionProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CollectionProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCollectionProposal
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
			return fmt.Errorf("proto: CollectionProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CollectionProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectionProposal
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
				return ErrInvalidLengthCollectionProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollectionProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectionProposal
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
				return ErrInvalidLengthCollectionProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCollectionProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollectionId", wireType)
			}
			m.CollectionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCollectionProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CollectionId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCollectionProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCollectionProposal
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
func skipCollectionProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCollectionProposal
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
					return 0, ErrIntOverflowCollectionProposal
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
					return 0, ErrIntOverflowCollectionProposal
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
				return 0, ErrInvalidLengthCollectionProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCollectionProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCollectionProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCollectionProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCollectionProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCollectionProposal = fmt.Errorf("proto: unexpected end of group")
)
