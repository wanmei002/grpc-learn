// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.12.0
// source: proto/product.proto

package product

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ProductId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,8,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *ProductId) Reset() {
	*x = ProductId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductId) ProtoMessage() {}

func (x *ProductId) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductId.ProtoReflect.Descriptor instead.
func (*ProductId) Descriptor() ([]byte, []int) {
	return file_proto_product_proto_rawDescGZIP(), []int{0}
}

func (x *ProductId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CommentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName    string `protobuf:"bytes,9,opt,name=userName,proto3" json:"userName,omitempty"`
	UserComment string `protobuf:"bytes,10,opt,name=userComment,proto3" json:"userComment,omitempty"`
}

func (x *CommentInfo) Reset() {
	*x = CommentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentInfo) ProtoMessage() {}

func (x *CommentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentInfo.ProtoReflect.Descriptor instead.
func (*CommentInfo) Descriptor() ([]byte, []int) {
	return file_proto_product_proto_rawDescGZIP(), []int{1}
}

func (x *CommentInfo) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *CommentInfo) GetUserComment() string {
	if x != nil {
		return x.UserComment
	}
	return ""
}

type OrderList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Order `protobuf:"bytes,7,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *OrderList) Reset() {
	*x = OrderList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderList) ProtoMessage() {}

func (x *OrderList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderList.ProtoReflect.Descriptor instead.
func (*OrderList) Descriptor() ([]byte, []int) {
	return file_proto_product_proto_rawDescGZIP(), []int{2}
}

func (x *OrderList) GetItems() []*Order {
	if x != nil {
		return x.Items
	}
	return nil
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Items []string `protobuf:"bytes,2,rep,name=Items,proto3" json:"Items,omitempty"`
	Desc  string   `protobuf:"bytes,3,opt,name=Desc,proto3" json:"Desc,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_proto_product_proto_rawDescGZIP(), []int{3}
}

func (x *Order) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Order) GetItems() []string {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *Order) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

type RespRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,4,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg  string `protobuf:"bytes,5,opt,name=msg,proto3" json:"msg,omitempty"`
	Data string `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RespRes) Reset() {
	*x = RespRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespRes) ProtoMessage() {}

func (x *RespRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespRes.ProtoReflect.Descriptor instead.
func (*RespRes) Descriptor() ([]byte, []int) {
	return file_proto_product_proto_rawDescGZIP(), []int{4}
}

func (x *RespRes) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RespRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *RespRes) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_proto_product_proto protoreflect.FileDescriptor

var file_proto_product_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x1b, 0x0a, 0x09, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x22, 0x4b, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x22, 0x2c, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x22, 0x45, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x44, 0x65, 0x73, 0x63, 0x22, 0x43, 0x0a, 0x07, 0x52, 0x65, 0x73, 0x70,
	0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x63, 0x0a,
	0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x24, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x1a, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x52, 0x65, 0x73, 0x12, 0x32,
	0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x12, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x1a,
	0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x30, 0x01, 0x42, 0x18, 0x5a, 0x16, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x3b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_product_proto_rawDescOnce sync.Once
	file_proto_product_proto_rawDescData = file_proto_product_proto_rawDesc
)

func file_proto_product_proto_rawDescGZIP() []byte {
	file_proto_product_proto_rawDescOnce.Do(func() {
		file_proto_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_product_proto_rawDescData)
	})
	return file_proto_product_proto_rawDescData
}

var file_proto_product_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_product_proto_goTypes = []interface{}{
	(*ProductId)(nil),   // 0: pb.ProductId
	(*CommentInfo)(nil), // 1: pb.CommentInfo
	(*OrderList)(nil),   // 2: pb.OrderList
	(*Order)(nil),       // 3: pb.Order
	(*RespRes)(nil),     // 4: pb.RespRes
}
var file_proto_product_proto_depIdxs = []int32{
	3, // 0: pb.OrderList.items:type_name -> pb.Order
	3, // 1: pb.Product.AddProduct:input_type -> pb.Order
	0, // 2: pb.Product.CommentProduct:input_type -> pb.ProductId
	4, // 3: pb.Product.AddProduct:output_type -> pb.RespRes
	1, // 4: pb.Product.CommentProduct:output_type -> pb.CommentInfo
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_product_proto_init() }
func file_proto_product_proto_init() {
	if File_proto_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_product_proto_goTypes,
		DependencyIndexes: file_proto_product_proto_depIdxs,
		MessageInfos:      file_proto_product_proto_msgTypes,
	}.Build()
	File_proto_product_proto = out.File
	file_proto_product_proto_rawDesc = nil
	file_proto_product_proto_goTypes = nil
	file_proto_product_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProductClient is the client API for Product service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductClient interface {
	// 一元
	AddProduct(ctx context.Context, in *Order, opts ...grpc.CallOption) (*RespRes, error)
	// 服务端流数据的接收
	CommentProduct(ctx context.Context, in *ProductId, opts ...grpc.CallOption) (Product_CommentProductClient, error)
}

type productClient struct {
	cc grpc.ClientConnInterface
}

func NewProductClient(cc grpc.ClientConnInterface) ProductClient {
	return &productClient{cc}
}

func (c *productClient) AddProduct(ctx context.Context, in *Order, opts ...grpc.CallOption) (*RespRes, error) {
	out := new(RespRes)
	err := c.cc.Invoke(ctx, "/pb.Product/AddProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productClient) CommentProduct(ctx context.Context, in *ProductId, opts ...grpc.CallOption) (Product_CommentProductClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Product_serviceDesc.Streams[0], "/pb.Product/CommentProduct", opts...)
	if err != nil {
		return nil, err
	}
	x := &productCommentProductClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Product_CommentProductClient interface {
	Recv() (*CommentInfo, error)
	grpc.ClientStream
}

type productCommentProductClient struct {
	grpc.ClientStream
}

func (x *productCommentProductClient) Recv() (*CommentInfo, error) {
	m := new(CommentInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProductServer is the server API for Product service.
type ProductServer interface {
	// 一元
	AddProduct(context.Context, *Order) (*RespRes, error)
	// 服务端流数据的接收
	CommentProduct(*ProductId, Product_CommentProductServer) error
}

// UnimplementedProductServer can be embedded to have forward compatible implementations.
type UnimplementedProductServer struct {
}

func (*UnimplementedProductServer) AddProduct(context.Context, *Order) (*RespRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProduct not implemented")
}
func (*UnimplementedProductServer) CommentProduct(*ProductId, Product_CommentProductServer) error {
	return status.Errorf(codes.Unimplemented, "method CommentProduct not implemented")
}

func RegisterProductServer(s *grpc.Server, srv ProductServer) {
	s.RegisterService(&_Product_serviceDesc, srv)
}

func _Product_AddProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServer).AddProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Product/AddProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServer).AddProduct(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _Product_CommentProduct_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ProductId)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProductServer).CommentProduct(m, &productCommentProductServer{stream})
}

type Product_CommentProductServer interface {
	Send(*CommentInfo) error
	grpc.ServerStream
}

type productCommentProductServer struct {
	grpc.ServerStream
}

func (x *productCommentProductServer) Send(m *CommentInfo) error {
	return x.ServerStream.SendMsg(m)
}

var _Product_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Product",
	HandlerType: (*ProductServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddProduct",
			Handler:    _Product_AddProduct_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CommentProduct",
			Handler:       _Product_CommentProduct_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/product.proto",
}
