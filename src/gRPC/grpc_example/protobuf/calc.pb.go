// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calc.proto

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	calc.proto
	user.proto

It has these top-level messages:
	CalcRequest
	CalcReply
	CalcRequest1
	CalcReply1
*/
package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CalcRequest struct {
	NumberA int32 `protobuf:"varint,1,opt,name=number_a,json=numberA" json:"number_a,omitempty"`
	NumberB int32 `protobuf:"varint,2,opt,name=number_b,json=numberB" json:"number_b,omitempty"`
}

func (m *CalcRequest) Reset()                    { *m = CalcRequest{} }
func (m *CalcRequest) String() string            { return proto.CompactTextString(m) }
func (*CalcRequest) ProtoMessage()               {}
func (*CalcRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CalcRequest) GetNumberA() int32 {
	if m != nil {
		return m.NumberA
	}
	return 0
}

func (m *CalcRequest) GetNumberB() int32 {
	if m != nil {
		return m.NumberB
	}
	return 0
}

type CalcReply struct {
	Result int32 `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

func (m *CalcReply) Reset()                    { *m = CalcReply{} }
func (m *CalcReply) String() string            { return proto.CompactTextString(m) }
func (*CalcReply) ProtoMessage()               {}
func (*CalcReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CalcReply) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func init() {
	proto.RegisterType((*CalcRequest)(nil), "protobuf.CalcRequest")
	proto.RegisterType((*CalcReply)(nil), "protobuf.CalcReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Calculator service

type CalculatorClient interface {
	Plus(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcReply, error)
}

type calculatorClient struct {
	cc *grpc.ClientConn
}

func NewCalculatorClient(cc *grpc.ClientConn) CalculatorClient {
	return &calculatorClient{cc}
}

func (c *calculatorClient) Plus(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcReply, error) {
	out := new(CalcReply)
	err := grpc.Invoke(ctx, "/protobuf.Calculator/Plus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Calculator service

type CalculatorServer interface {
	Plus(context.Context, *CalcRequest) (*CalcReply, error)
}

func RegisterCalculatorServer(s *grpc.Server, srv CalculatorServer) {
	s.RegisterService(&_Calculator_serviceDesc, srv)
}

func _Calculator_Plus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServer).Plus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Calculator/Plus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServer).Plus(ctx, req.(*CalcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calculator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.Calculator",
	HandlerType: (*CalculatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Plus",
			Handler:    _Calculator_Plus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calc.proto",
}

func init() { proto.RegisterFile("calc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4e, 0xcc, 0x49,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0x49, 0xa5, 0x69, 0x4a, 0xce, 0x5c,
	0xdc, 0xce, 0x89, 0x39, 0xc9, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x92, 0x5c, 0x1c,
	0x79, 0xa5, 0xb9, 0x49, 0xa9, 0x45, 0xf1, 0x89, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0xec,
	0x10, 0xbe, 0x23, 0x92, 0x54, 0x92, 0x04, 0x13, 0xb2, 0x94, 0x93, 0x92, 0x32, 0x17, 0x27, 0xc4,
	0x90, 0x82, 0x9c, 0x4a, 0x21, 0x31, 0x2e, 0xb6, 0xa2, 0xd4, 0xe2, 0xd2, 0x9c, 0x12, 0xa8, 0x01,
	0x50, 0x9e, 0x91, 0x13, 0x17, 0x17, 0x48, 0x51, 0x69, 0x4e, 0x62, 0x49, 0x7e, 0x91, 0x90, 0x09,
	0x17, 0x4b, 0x40, 0x4e, 0x69, 0xb1, 0x90, 0xa8, 0x1e, 0xcc, 0x29, 0x7a, 0x48, 0xee, 0x90, 0x12,
	0x46, 0x17, 0x2e, 0xc8, 0xa9, 0x54, 0x62, 0x48, 0x62, 0x03, 0x8b, 0x1a, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xb6, 0x03, 0x06, 0x36, 0xcc, 0x00, 0x00, 0x00,
}
