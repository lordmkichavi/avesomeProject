// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.11.4
// source: transfers/transferspb/transferspb.proto

package transferspb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	date "google.golang.org/genproto/googleapis/type/date"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type TransfersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OriginAccount   string  `protobuf:"bytes,1,opt,name=origin_account,json=originAccount,proto3" json:"origin_account,omitempty"`
	ReceiverAccount string  `protobuf:"bytes,2,opt,name=receiver_account,json=receiverAccount,proto3" json:"receiver_account,omitempty"`
	Amount          float64 `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *TransfersRequest) Reset() {
	*x = TransfersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfers_transferspb_transferspb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransfersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransfersRequest) ProtoMessage() {}

func (x *TransfersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transfers_transferspb_transferspb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransfersRequest.ProtoReflect.Descriptor instead.
func (*TransfersRequest) Descriptor() ([]byte, []int) {
	return file_transfers_transferspb_transferspb_proto_rawDescGZIP(), []int{0}
}

func (x *TransfersRequest) GetOriginAccount() string {
	if x != nil {
		return x.OriginAccount
	}
	return ""
}

func (x *TransfersRequest) GetReceiverAccount() string {
	if x != nil {
		return x.ReceiverAccount
	}
	return ""
}

func (x *TransfersRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type TransfersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperationDate *date.Date `protobuf:"bytes,1,opt,name=operation_date,json=operationDate,proto3" json:"operation_date,omitempty"`
}

func (x *TransfersResponse) Reset() {
	*x = TransfersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transfers_transferspb_transferspb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransfersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransfersResponse) ProtoMessage() {}

func (x *TransfersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_transfers_transferspb_transferspb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransfersResponse.ProtoReflect.Descriptor instead.
func (*TransfersResponse) Descriptor() ([]byte, []int) {
	return file_transfers_transferspb_transferspb_proto_rawDescGZIP(), []int{1}
}

func (x *TransfersResponse) GetOperationDate() *date.Date {
	if x != nil {
		return x.OperationDate
	}
	return nil
}

var File_transfers_transferspb_transferspb_proto protoreflect.FileDescriptor

var file_transfers_transferspb_transferspb_proto_rawDesc = []byte{
	0x0a, 0x27, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x73, 0x70, 0x62, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x73, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x73, 0x1a, 0x1b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x7c, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10,
	0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x4d, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52,
	0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x32, 0x5b,
	0x0a, 0x10, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x47, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x1b,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transfers_transferspb_transferspb_proto_rawDescOnce sync.Once
	file_transfers_transferspb_transferspb_proto_rawDescData = file_transfers_transferspb_transferspb_proto_rawDesc
)

func file_transfers_transferspb_transferspb_proto_rawDescGZIP() []byte {
	file_transfers_transferspb_transferspb_proto_rawDescOnce.Do(func() {
		file_transfers_transferspb_transferspb_proto_rawDescData = protoimpl.X.CompressGZIP(file_transfers_transferspb_transferspb_proto_rawDescData)
	})
	return file_transfers_transferspb_transferspb_proto_rawDescData
}

var file_transfers_transferspb_transferspb_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_transfers_transferspb_transferspb_proto_goTypes = []interface{}{
	(*TransfersRequest)(nil),  // 0: transfers.TransfersRequest
	(*TransfersResponse)(nil), // 1: transfers.TransfersResponse
	(*date.Date)(nil),         // 2: google.type.Date
}
var file_transfers_transferspb_transferspb_proto_depIdxs = []int32{
	2, // 0: transfers.TransfersResponse.operation_date:type_name -> google.type.Date
	0, // 1: transfers.TransfersService.Transfer:input_type -> transfers.TransfersRequest
	1, // 2: transfers.TransfersService.Transfer:output_type -> transfers.TransfersResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_transfers_transferspb_transferspb_proto_init() }
func file_transfers_transferspb_transferspb_proto_init() {
	if File_transfers_transferspb_transferspb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transfers_transferspb_transferspb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransfersRequest); i {
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
		file_transfers_transferspb_transferspb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransfersResponse); i {
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
			RawDescriptor: file_transfers_transferspb_transferspb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_transfers_transferspb_transferspb_proto_goTypes,
		DependencyIndexes: file_transfers_transferspb_transferspb_proto_depIdxs,
		MessageInfos:      file_transfers_transferspb_transferspb_proto_msgTypes,
	}.Build()
	File_transfers_transferspb_transferspb_proto = out.File
	file_transfers_transferspb_transferspb_proto_rawDesc = nil
	file_transfers_transferspb_transferspb_proto_goTypes = nil
	file_transfers_transferspb_transferspb_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TransfersServiceClient is the client API for TransfersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransfersServiceClient interface {
	Transfer(ctx context.Context, in *TransfersRequest, opts ...grpc.CallOption) (*TransfersResponse, error)
}

type transfersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransfersServiceClient(cc grpc.ClientConnInterface) TransfersServiceClient {
	return &transfersServiceClient{cc}
}

func (c *transfersServiceClient) Transfer(ctx context.Context, in *TransfersRequest, opts ...grpc.CallOption) (*TransfersResponse, error) {
	out := new(TransfersResponse)
	err := c.cc.Invoke(ctx, "/transfers.TransfersService/Transfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransfersServiceServer is the server API for TransfersService service.
type TransfersServiceServer interface {
	Transfer(context.Context, *TransfersRequest) (*TransfersResponse, error)
}

// UnimplementedTransfersServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTransfersServiceServer struct {
}

func (*UnimplementedTransfersServiceServer) Transfer(context.Context, *TransfersRequest) (*TransfersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transfer not implemented")
}

func RegisterTransfersServiceServer(s *grpc.Server, srv TransfersServiceServer) {
	s.RegisterService(&_TransfersService_serviceDesc, srv)
}

func _TransfersService_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransfersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransfersServiceServer).Transfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transfers.TransfersService/Transfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransfersServiceServer).Transfer(ctx, req.(*TransfersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TransfersService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "transfers.TransfersService",
	HandlerType: (*TransfersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Transfer",
			Handler:    _TransfersService_Transfer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transfers/transferspb/transferspb.proto",
}
