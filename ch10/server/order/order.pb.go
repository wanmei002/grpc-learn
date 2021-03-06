// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/order.proto

package order

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type OrderInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Items                []string `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	Desc                 string   `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderInfo) Reset()         { *m = OrderInfo{} }
func (m *OrderInfo) String() string { return proto.CompactTextString(m) }
func (*OrderInfo) ProtoMessage()    {}
func (*OrderInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f65b0626cc3aada8, []int{0}
}

func (m *OrderInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderInfo.Unmarshal(m, b)
}
func (m *OrderInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderInfo.Marshal(b, m, deterministic)
}
func (m *OrderInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderInfo.Merge(m, src)
}
func (m *OrderInfo) XXX_Size() int {
	return xxx_messageInfo_OrderInfo.Size(m)
}
func (m *OrderInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OrderInfo proto.InternalMessageInfo

func (m *OrderInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OrderInfo) GetItems() []string {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *OrderInfo) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

type Res struct {
	Code                 int32    `protobuf:"varint,4,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,5,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Res) Reset()         { *m = Res{} }
func (m *Res) String() string { return proto.CompactTextString(m) }
func (*Res) ProtoMessage()    {}
func (*Res) Descriptor() ([]byte, []int) {
	return fileDescriptor_f65b0626cc3aada8, []int{1}
}

func (m *Res) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Res.Unmarshal(m, b)
}
func (m *Res) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Res.Marshal(b, m, deterministic)
}
func (m *Res) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Res.Merge(m, src)
}
func (m *Res) XXX_Size() int {
	return xxx_messageInfo_Res.Size(m)
}
func (m *Res) XXX_DiscardUnknown() {
	xxx_messageInfo_Res.DiscardUnknown(m)
}

var xxx_messageInfo_Res proto.InternalMessageInfo

func (m *Res) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Res) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*OrderInfo)(nil), "order.OrderInfo")
	proto.RegisterType((*Res)(nil), "order.Res")
}

func init() { proto.RegisterFile("proto/order.proto", fileDescriptor_f65b0626cc3aada8) }

var fileDescriptor_f65b0626cc3aada8 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0xb1, 0xcb, 0xc2, 0x30,
	0x10, 0xc5, 0xe9, 0x97, 0xe6, 0xc3, 0xde, 0x54, 0x8f, 0x0e, 0xc1, 0xa9, 0x74, 0x0a, 0x08, 0x15,
	0x75, 0x74, 0xd2, 0xad, 0x93, 0x90, 0xd1, 0x4d, 0x9b, 0x53, 0x1c, 0xd2, 0x48, 0x52, 0xfc, 0xfb,
	0x25, 0x17, 0x71, 0x39, 0x7e, 0xef, 0xee, 0xc1, 0xbd, 0x07, 0xcb, 0x57, 0xf0, 0xb3, 0xdf, 0xf8,
	0x60, 0x29, 0xf4, 0xcc, 0x28, 0x59, 0x74, 0x03, 0x54, 0xe7, 0x04, 0xc3, 0x74, 0xf7, 0x88, 0x50,
	0x4e, 0x57, 0x47, 0xaa, 0x68, 0x0b, 0x5d, 0x19, 0x66, 0x6c, 0x40, 0x3e, 0x67, 0x72, 0x51, 0xfd,
	0xb5, 0x42, 0x57, 0x26, 0x8b, 0xe4, 0xb4, 0x14, 0x47, 0x25, 0xb2, 0x33, 0x71, 0xb7, 0x06, 0x61,
	0x88, 0x4f, 0xa3, 0xb7, 0xa4, 0xca, 0xb6, 0xd0, 0xd2, 0x30, 0x63, 0x0d, 0xc2, 0xc5, 0x87, 0x92,
	0xec, 0x4e, 0xb8, 0xdb, 0x82, 0xe4, 0xbf, 0xa8, 0x61, 0x71, 0xb4, 0x36, 0x73, 0xdd, 0xe7, 0x84,
	0xbf, 0x44, 0x2b, 0xf8, 0x6e, 0x0c, 0xc5, 0x53, 0x73, 0xc1, 0x48, 0xe1, 0x4d, 0x21, 0xf7, 0x38,
	0xf0, 0xbc, 0xfd, 0x73, 0x9d, 0xfd, 0x27, 0x00, 0x00, 0xff, 0xff, 0xe9, 0xb4, 0x67, 0x91, 0xe3,
	0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OrderClient is the client API for Order service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderClient interface {
	AddOrder(ctx context.Context, in *OrderInfo, opts ...grpc.CallOption) (*Res, error)
}

type orderClient struct {
	cc *grpc.ClientConn
}

func NewOrderClient(cc *grpc.ClientConn) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) AddOrder(ctx context.Context, in *OrderInfo, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := c.cc.Invoke(ctx, "/order.Order/AddOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServer is the server API for Order service.
type OrderServer interface {
	AddOrder(context.Context, *OrderInfo) (*Res, error)
}

// UnimplementedOrderServer can be embedded to have forward compatible implementations.
type UnimplementedOrderServer struct {
}

func (*UnimplementedOrderServer) AddOrder(ctx context.Context, req *OrderInfo) (*Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOrder not implemented")
}

func RegisterOrderServer(s *grpc.Server, srv OrderServer) {
	s.RegisterService(&_Order_serviceDesc, srv)
}

func _Order_AddOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).AddOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.Order/AddOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).AddOrder(ctx, req.(*OrderInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Order_serviceDesc = grpc.ServiceDesc{
	ServiceName: "order.Order",
	HandlerType: (*OrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddOrder",
			Handler:    _Order_AddOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/order.proto",
}
