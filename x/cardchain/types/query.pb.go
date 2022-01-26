// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cardchain/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c26c5456fa123d9f, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c26c5456fa123d9f, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryQCardRequest struct {
	CardId string `protobuf:"bytes,1,opt,name=cardId,proto3" json:"cardId,omitempty"`
}

func (m *QueryQCardRequest) Reset()         { *m = QueryQCardRequest{} }
func (m *QueryQCardRequest) String() string { return proto.CompactTextString(m) }
func (*QueryQCardRequest) ProtoMessage()    {}
func (*QueryQCardRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c26c5456fa123d9f, []int{2}
}
func (m *QueryQCardRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryQCardRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryQCardRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryQCardRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryQCardRequest.Merge(m, src)
}
func (m *QueryQCardRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryQCardRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryQCardRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryQCardRequest proto.InternalMessageInfo

func (m *QueryQCardRequest) GetCardId() string {
	if m != nil {
		return m.CardId
	}
	return ""
}

type QueryQCardResponse struct {
	Card []byte `protobuf:"bytes,1,opt,name=card,proto3" json:"card,omitempty"`
}

func (m *QueryQCardResponse) Reset()         { *m = QueryQCardResponse{} }
func (m *QueryQCardResponse) String() string { return proto.CompactTextString(m) }
func (*QueryQCardResponse) ProtoMessage()    {}
func (*QueryQCardResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c26c5456fa123d9f, []int{3}
}
func (m *QueryQCardResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryQCardResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryQCardResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryQCardResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryQCardResponse.Merge(m, src)
}
func (m *QueryQCardResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryQCardResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryQCardResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryQCardResponse proto.InternalMessageInfo

func (m *QueryQCardResponse) GetCard() []byte {
	if m != nil {
		return m.Card
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "DecentralCardGame.cardchain.cardchain.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "DecentralCardGame.cardchain.cardchain.QueryParamsResponse")
	proto.RegisterType((*QueryQCardRequest)(nil), "DecentralCardGame.cardchain.cardchain.QueryQCardRequest")
	proto.RegisterType((*QueryQCardResponse)(nil), "DecentralCardGame.cardchain.cardchain.QueryQCardResponse")
}

func init() { proto.RegisterFile("cardchain/query.proto", fileDescriptor_c26c5456fa123d9f) }

var fileDescriptor_c26c5456fa123d9f = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xbf, 0x4a, 0xfb, 0x40,
	0x1c, 0x4f, 0x4a, 0x1b, 0xf8, 0xdd, 0xcf, 0xc5, 0xb3, 0x16, 0x09, 0x12, 0x25, 0x20, 0x14, 0xa5,
	0x39, 0x5a, 0x51, 0x5a, 0x07, 0x87, 0x56, 0x10, 0x71, 0xb1, 0x19, 0x5d, 0xe4, 0x92, 0x1e, 0x69,
	0xa0, 0xc9, 0xa5, 0xb9, 0xab, 0x58, 0xc4, 0xc5, 0x27, 0x10, 0x7c, 0x17, 0x27, 0x1f, 0xa0, 0x63,
	0xc1, 0xc5, 0x49, 0xa4, 0xf5, 0x3d, 0x94, 0x5c, 0x8e, 0xb6, 0x5a, 0x87, 0xda, 0xed, 0x93, 0xbb,
	0xcf, 0xbf, 0xfb, 0x7e, 0x03, 0xd6, 0x5d, 0x1c, 0xb7, 0xdc, 0x36, 0xf6, 0x43, 0xd4, 0xed, 0x91,
	0xb8, 0x6f, 0x45, 0x31, 0xe5, 0x14, 0xee, 0x9c, 0x10, 0x97, 0x84, 0x3c, 0xc6, 0x9d, 0x06, 0x8e,
	0x5b, 0xa7, 0x38, 0x20, 0xd6, 0x84, 0x38, 0x45, 0x7a, 0xde, 0xa3, 0x1e, 0x15, 0x0a, 0x94, 0xa0,
	0x54, 0xac, 0x6f, 0x7a, 0x94, 0x7a, 0x1d, 0x82, 0x70, 0xe4, 0x23, 0x1c, 0x86, 0x94, 0x63, 0xee,
	0xd3, 0x90, 0xc9, 0xdb, 0x5d, 0x97, 0xb2, 0x80, 0x32, 0xe4, 0x60, 0x46, 0xd2, 0x4c, 0x74, 0x5d,
	0x76, 0x08, 0xc7, 0x65, 0x14, 0x61, 0xcf, 0x0f, 0x05, 0x59, 0x72, 0x0b, 0xd3, 0x76, 0x11, 0x8e,
	0x71, 0x20, 0x3d, 0xcc, 0x3c, 0x80, 0xcd, 0x44, 0x79, 0x21, 0x0e, 0x6d, 0xd2, 0xed, 0x11, 0xc6,
	0x4d, 0x07, 0xac, 0x7d, 0x3b, 0x65, 0x11, 0x0d, 0x19, 0x81, 0xe7, 0x40, 0x4b, 0xc5, 0x1b, 0xea,
	0xb6, 0x5a, 0xfc, 0x5f, 0x29, 0x59, 0x0b, 0x3d, 0xce, 0x4a, 0x6d, 0xea, 0xd9, 0xc1, 0xdb, 0x96,
	0x62, 0x4b, 0x0b, 0x73, 0x0f, 0xac, 0x8a, 0x8c, 0x66, 0x22, 0x95, 0xc1, 0xb0, 0x00, 0xb4, 0x44,
	0x76, 0xd6, 0x12, 0x09, 0xff, 0x6c, 0xf9, 0x65, 0x16, 0x65, 0x4d, 0x49, 0x96, 0x7d, 0x20, 0xc8,
	0x26, 0xf7, 0x82, 0xbb, 0x62, 0x0b, 0x5c, 0xf9, 0xcc, 0x80, 0x9c, 0xa0, 0xc2, 0x27, 0x15, 0x68,
	0x69, 0x32, 0xac, 0x2d, 0x58, 0x74, 0x7e, 0x14, 0xfa, 0xd1, 0x32, 0xd2, 0xb4, 0x9f, 0x79, 0x70,
	0xff, 0xf2, 0xf1, 0x98, 0x41, 0xb0, 0x84, 0xe6, 0x3c, 0xd0, 0x74, 0x1f, 0x3f, 0x37, 0x03, 0x9f,
	0x55, 0x90, 0x13, 0x0f, 0x85, 0xd5, 0xbf, 0x84, 0xcf, 0x0e, 0x52, 0xaf, 0x2d, 0xa1, 0x94, 0xad,
	0x8f, 0x45, 0xeb, 0x2a, 0x3c, 0x5c, 0xb0, 0x75, 0xf7, 0x2a, 0xc1, 0xe8, 0x36, 0x5d, 0xd5, 0x5d,
	0xdd, 0x1e, 0x8c, 0x0c, 0x75, 0x38, 0x32, 0xd4, 0xf7, 0x91, 0xa1, 0x3e, 0x8c, 0x0d, 0x65, 0x38,
	0x36, 0x94, 0xd7, 0xb1, 0xa1, 0x5c, 0x56, 0x3d, 0x9f, 0xb7, 0x7b, 0x8e, 0xe5, 0xd2, 0xe0, 0x17,
	0xef, 0xc6, 0xc4, 0xf1, 0x66, 0xc6, 0x9d, 0xf7, 0x23, 0xc2, 0x1c, 0x4d, 0xfc, 0xad, 0xfb, 0x5f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x35, 0x92, 0x60, 0x65, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of QCard items.
	QCard(ctx context.Context, in *QueryQCardRequest, opts ...grpc.CallOption) (*QueryQCardResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/DecentralCardGame.cardchain.cardchain.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QCard(ctx context.Context, in *QueryQCardRequest, opts ...grpc.CallOption) (*QueryQCardResponse, error) {
	out := new(QueryQCardResponse)
	err := c.cc.Invoke(ctx, "/DecentralCardGame.cardchain.cardchain.Query/QCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of QCard items.
	QCard(context.Context, *QueryQCardRequest) (*QueryQCardResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) QCard(ctx context.Context, req *QueryQCardRequest) (*QueryQCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QCard not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DecentralCardGame.cardchain.cardchain.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryQCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DecentralCardGame.cardchain.cardchain.Query/QCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QCard(ctx, req.(*QueryQCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "DecentralCardGame.cardchain.cardchain.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "QCard",
			Handler:    _Query_QCard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cardchain/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryQCardRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryQCardRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryQCardRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CardId) > 0 {
		i -= len(m.CardId)
		copy(dAtA[i:], m.CardId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.CardId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryQCardResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryQCardResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryQCardResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Card) > 0 {
		i -= len(m.Card)
		copy(dAtA[i:], m.Card)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Card)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryQCardRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CardId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryQCardResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Card)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryQCardRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryQCardRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryQCardRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CardId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CardId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryQCardResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryQCardResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryQCardResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Card", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Card = append(m.Card[:0], dAtA[iNdEx:postIndex]...)
			if m.Card == nil {
				m.Card = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
