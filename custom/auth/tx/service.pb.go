// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: terra/tx/v1beta1/service.proto

package tx

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	tx "github.com/cosmos/cosmos-sdk/types/tx"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
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
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ComputeTaxRequest is the request type for the Service.ComputeTax
// RPC method.
type ComputeTaxRequest struct {
	// tx is the transaction to simulate.
	// Deprecated. Send raw tx bytes instead.
	Tx *tx.Tx `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"` // Deprecated: Do not use.
	// tx_bytes is the raw transaction.
	TxBytes []byte `protobuf:"bytes,2,opt,name=tx_bytes,json=txBytes,proto3" json:"tx_bytes,omitempty"`
}

func (m *ComputeTaxRequest) Reset()         { *m = ComputeTaxRequest{} }
func (m *ComputeTaxRequest) String() string { return proto.CompactTextString(m) }
func (*ComputeTaxRequest) ProtoMessage()    {}
func (*ComputeTaxRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b3c73e5d85273f4, []int{0}
}
func (m *ComputeTaxRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ComputeTaxRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ComputeTaxRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ComputeTaxRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeTaxRequest.Merge(m, src)
}
func (m *ComputeTaxRequest) XXX_Size() int {
	return m.Size()
}
func (m *ComputeTaxRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeTaxRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeTaxRequest proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *ComputeTaxRequest) GetTx() *tx.Tx {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *ComputeTaxRequest) GetTxBytes() []byte {
	if m != nil {
		return m.TxBytes
	}
	return nil
}

// ComputeTaxResponse is the response type for the Service.ComputeTax
// RPC method.
type ComputeTaxResponse struct {
	// amount is the amount of coins to be paid as a fee
	TaxAmount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=tax_amount,json=taxAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"tax_amount"`
}

func (m *ComputeTaxResponse) Reset()         { *m = ComputeTaxResponse{} }
func (m *ComputeTaxResponse) String() string { return proto.CompactTextString(m) }
func (*ComputeTaxResponse) ProtoMessage()    {}
func (*ComputeTaxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b3c73e5d85273f4, []int{1}
}
func (m *ComputeTaxResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ComputeTaxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ComputeTaxResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ComputeTaxResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeTaxResponse.Merge(m, src)
}
func (m *ComputeTaxResponse) XXX_Size() int {
	return m.Size()
}
func (m *ComputeTaxResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeTaxResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeTaxResponse proto.InternalMessageInfo

func (m *ComputeTaxResponse) GetTaxAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.TaxAmount
	}
	return nil
}

func init() {
	proto.RegisterType((*ComputeTaxRequest)(nil), "terra.tx.v1beta1.ComputeTaxRequest")
	golang_proto.RegisterType((*ComputeTaxRequest)(nil), "terra.tx.v1beta1.ComputeTaxRequest")
	proto.RegisterType((*ComputeTaxResponse)(nil), "terra.tx.v1beta1.ComputeTaxResponse")
	golang_proto.RegisterType((*ComputeTaxResponse)(nil), "terra.tx.v1beta1.ComputeTaxResponse")
}

func init() { proto.RegisterFile("terra/tx/v1beta1/service.proto", fileDescriptor_0b3c73e5d85273f4) }
func init() {
	golang_proto.RegisterFile("terra/tx/v1beta1/service.proto", fileDescriptor_0b3c73e5d85273f4)
}

var fileDescriptor_0b3c73e5d85273f4 = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x3f, 0x8f, 0x94, 0x40,
	0x18, 0xc6, 0x19, 0x4c, 0x3c, 0x9d, 0xb3, 0x50, 0xa2, 0x09, 0x47, 0x74, 0x6e, 0x83, 0x16, 0x78,
	0xc9, 0xcd, 0x78, 0x6b, 0x67, 0x27, 0xd7, 0x58, 0xe3, 0x35, 0xda, 0x6c, 0x06, 0x9c, 0xb0, 0xa8,
	0xcc, 0x8b, 0xcc, 0xcb, 0x66, 0xb6, 0x53, 0x7b, 0x13, 0x13, 0xbf, 0x85, 0x9f, 0xc2, 0x72, 0xcb,
	0x4d, 0x6c, 0xac, 0xd4, 0x2c, 0x7e, 0x10, 0xb3, 0x80, 0x2b, 0x71, 0x13, 0x2b, 0x98, 0xfc, 0xde,
	0x3f, 0xcf, 0xf3, 0xbc, 0x94, 0xa1, 0xaa, 0x6b, 0x29, 0xd0, 0x8a, 0xc5, 0x59, 0xaa, 0x50, 0x9e,
	0x09, 0xa3, 0xea, 0x45, 0x91, 0x29, 0x5e, 0xd5, 0x80, 0xe0, 0x5d, 0xef, 0x38, 0x47, 0xcb, 0x07,
	0x1e, 0xdc, 0xcc, 0x21, 0x87, 0x0e, 0x8a, 0xed, 0x5f, 0x5f, 0x17, 0xdc, 0xce, 0x01, 0xf2, 0xd7,
	0x4a, 0xc8, 0xaa, 0x10, 0x52, 0x6b, 0x40, 0x89, 0x05, 0x68, 0x33, 0x50, 0x96, 0x81, 0x29, 0xc1,
	0x88, 0x54, 0x1a, 0xb5, 0x5b, 0x94, 0x41, 0xa1, 0x07, 0x1e, 0x0c, 0x7c, 0x24, 0x03, 0x6d, 0xcf,
	0xc2, 0x67, 0xf4, 0xc6, 0x39, 0x94, 0x55, 0x83, 0xea, 0x42, 0xda, 0x44, 0xbd, 0x69, 0x94, 0x41,
	0xef, 0x3e, 0x75, 0xd1, 0xfa, 0x64, 0x42, 0xa2, 0xc3, 0xe9, 0x2d, 0xde, 0x77, 0x8f, 0x44, 0xf2,
	0x0b, 0x1b, 0xbb, 0x3e, 0x49, 0x5c, 0xb4, 0xde, 0x11, 0xbd, 0x82, 0x76, 0x96, 0x2e, 0x51, 0x19,
	0xdf, 0x9d, 0x90, 0xe8, 0x5a, 0x72, 0x80, 0x36, 0xde, 0x3e, 0xc3, 0xb7, 0x84, 0x7a, 0xe3, 0xd9,
	0xa6, 0x02, 0x6d, 0x94, 0xf7, 0x92, 0x52, 0x94, 0x76, 0x26, 0x4b, 0x68, 0x34, 0xfa, 0x64, 0x72,
	0x29, 0x3a, 0x9c, 0x1e, 0xfd, 0x59, 0xb2, 0xb5, 0xb0, 0x5b, 0x73, 0x0e, 0x85, 0x8e, 0x1f, 0xac,
	0xbe, 0x1f, 0x3b, 0x9f, 0x7f, 0x1c, 0x47, 0x79, 0x81, 0xf3, 0x26, 0xe5, 0x19, 0x94, 0x62, 0xf0,
	0xd3, 0x7f, 0x4e, 0xcd, 0x8b, 0x57, 0x02, 0x97, 0x95, 0x32, 0x5d, 0x83, 0x49, 0xae, 0xa2, 0xb4,
	0x8f, 0xbb, 0xe9, 0xd3, 0x0f, 0x84, 0x1e, 0x3c, 0xed, 0x13, 0xf7, 0xde, 0x11, 0x4a, 0xff, 0xca,
	0xf1, 0xee, 0xf2, 0x7f, 0xb3, 0xe7, 0x7b, 0x41, 0x04, 0xf7, 0xfe, 0x5f, 0xd4, 0x3b, 0x0a, 0xa3,
	0xf7, 0x5f, 0x7f, 0x7d, 0x72, 0xc3, 0xf0, 0x8e, 0xd8, 0x3b, 0x77, 0xd6, 0x57, 0xcf, 0x50, 0xda,
	0x47, 0xe4, 0x24, 0x7e, 0xb2, 0xda, 0x30, 0xb2, 0xde, 0x30, 0xf2, 0x73, 0xc3, 0xc8, 0xc7, 0x96,
	0x39, 0x5f, 0x5a, 0x46, 0xd6, 0x2d, 0x73, 0xbe, 0xb5, 0xcc, 0x79, 0x7e, 0x32, 0xb2, 0xd8, 0x4d,
	0x3a, 0x2d, 0x41, 0xab, 0xa5, 0xc8, 0xa0, 0x56, 0x22, 0x6b, 0x0c, 0x42, 0x29, 0x64, 0x83, 0x73,
	0x81, 0x36, 0xbd, 0xdc, 0x9d, 0xef, 0xe1, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcc, 0xeb, 0xe4,
	0x8f, 0x62, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	// EstimateFee simulates executing a transaction for estimating gas usage.
	ComputeTax(ctx context.Context, in *ComputeTaxRequest, opts ...grpc.CallOption) (*ComputeTaxResponse, error)
}

type serviceClient struct {
	cc grpc1.ClientConn
}

func NewServiceClient(cc grpc1.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) ComputeTax(ctx context.Context, in *ComputeTaxRequest, opts ...grpc.CallOption) (*ComputeTaxResponse, error) {
	out := new(ComputeTaxResponse)
	err := c.cc.Invoke(ctx, "/terra.tx.v1beta1.Service/ComputeTax", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	// EstimateFee simulates executing a transaction for estimating gas usage.
	ComputeTax(context.Context, *ComputeTaxRequest) (*ComputeTaxResponse, error)
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) ComputeTax(ctx context.Context, req *ComputeTaxRequest) (*ComputeTaxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComputeTax not implemented")
}

func RegisterServiceServer(s grpc1.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_ComputeTax_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComputeTaxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ComputeTax(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/terra.tx.v1beta1.Service/ComputeTax",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ComputeTax(ctx, req.(*ComputeTaxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "terra.tx.v1beta1.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ComputeTax",
			Handler:    _Service_ComputeTax_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "terra/tx/v1beta1/service.proto",
}

func (m *ComputeTaxRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ComputeTaxRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ComputeTaxRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TxBytes) > 0 {
		i -= len(m.TxBytes)
		copy(dAtA[i:], m.TxBytes)
		i = encodeVarintService(dAtA, i, uint64(len(m.TxBytes)))
		i--
		dAtA[i] = 0x12
	}
	if m.Tx != nil {
		{
			size, err := m.Tx.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ComputeTaxResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ComputeTaxResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ComputeTaxResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TaxAmount) > 0 {
		for iNdEx := len(m.TaxAmount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TaxAmount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintService(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintService(dAtA []byte, offset int, v uint64) int {
	offset -= sovService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ComputeTaxRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Tx != nil {
		l = m.Tx.Size()
		n += 1 + l + sovService(uint64(l))
	}
	l = len(m.TxBytes)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func (m *ComputeTaxResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.TaxAmount) > 0 {
		for _, e := range m.TaxAmount {
			l = e.Size()
			n += 1 + l + sovService(uint64(l))
		}
	}
	return n
}

func sovService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozService(x uint64) (n int) {
	return sovService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ComputeTaxRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: ComputeTaxRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ComputeTaxRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Tx == nil {
				m.Tx = &tx.Tx{}
			}
			if err := m.Tx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxBytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxBytes = append(m.TxBytes[:0], dAtA[iNdEx:postIndex]...)
			if m.TxBytes == nil {
				m.TxBytes = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
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
func (m *ComputeTaxResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: ComputeTaxResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ComputeTaxResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaxAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TaxAmount = append(m.TaxAmount, types.Coin{})
			if err := m.TaxAmount[len(m.TaxAmount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
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
func skipService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowService
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
					return 0, ErrIntOverflowService
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
					return 0, ErrIntOverflowService
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
				return 0, ErrInvalidLengthService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupService = fmt.Errorf("proto: unexpected end of group")
)
