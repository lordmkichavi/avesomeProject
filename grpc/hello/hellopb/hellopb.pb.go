// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.11.4
// source: hello/hellopb/hellopb.proto

package hellopb

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

type Hello struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	Prefix    string `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
}

func (x *Hello) Reset() {
	*x = Hello{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_hellopb_hellopb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Hello) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Hello) ProtoMessage() {}

func (x *Hello) ProtoReflect() protoreflect.Message {
	mi := &file_hello_hellopb_hellopb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Hello.ProtoReflect.Descriptor instead.
func (*Hello) Descriptor() ([]byte, []int) {
	return file_hello_hellopb_hellopb_proto_rawDescGZIP(), []int{0}
}

func (x *Hello) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Hello) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hello *Hello `protobuf:"bytes,1,opt,name=hello,proto3" json:"hello,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_hellopb_hellopb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hello_hellopb_hellopb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_hello_hellopb_hellopb_proto_rawDescGZIP(), []int{1}
}

func (x *HelloRequest) GetHello() *Hello {
	if x != nil {
		return x.Hello
	}
	return nil
}

type HelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomHello string `protobuf:"bytes,1,opt,name=custom_hello,json=customHello,proto3" json:"custom_hello,omitempty"`
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_hellopb_hellopb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hello_hellopb_hellopb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_hello_hellopb_hellopb_proto_rawDescGZIP(), []int{2}
}

func (x *HelloResponse) GetCustomHello() string {
	if x != nil {
		return x.CustomHello
	}
	return ""
}

type HelloManyLanguagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hello *Hello `protobuf:"bytes,1,opt,name=hello,proto3" json:"hello,omitempty"`
}

func (x *HelloManyLanguagesRequest) Reset() {
	*x = HelloManyLanguagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_hellopb_hellopb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloManyLanguagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloManyLanguagesRequest) ProtoMessage() {}

func (x *HelloManyLanguagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hello_hellopb_hellopb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloManyLanguagesRequest.ProtoReflect.Descriptor instead.
func (*HelloManyLanguagesRequest) Descriptor() ([]byte, []int) {
	return file_hello_hellopb_hellopb_proto_rawDescGZIP(), []int{3}
}

func (x *HelloManyLanguagesRequest) GetHello() *Hello {
	if x != nil {
		return x.Hello
	}
	return nil
}

type HelloManyLanguagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HelloLanguage string `protobuf:"bytes,1,opt,name=hello_language,json=helloLanguage,proto3" json:"hello_language,omitempty"`
}

func (x *HelloManyLanguagesResponse) Reset() {
	*x = HelloManyLanguagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_hellopb_hellopb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloManyLanguagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloManyLanguagesResponse) ProtoMessage() {}

func (x *HelloManyLanguagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hello_hellopb_hellopb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloManyLanguagesResponse.ProtoReflect.Descriptor instead.
func (*HelloManyLanguagesResponse) Descriptor() ([]byte, []int) {
	return file_hello_hellopb_hellopb_proto_rawDescGZIP(), []int{4}
}

func (x *HelloManyLanguagesResponse) GetHelloLanguage() string {
	if x != nil {
		return x.HelloLanguage
	}
	return ""
}

type HellosGoodbyeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hello *Hello `protobuf:"bytes,1,opt,name=hello,proto3" json:"hello,omitempty"`
}

func (x *HellosGoodbyeRequest) Reset() {
	*x = HellosGoodbyeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_hellopb_hellopb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HellosGoodbyeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HellosGoodbyeRequest) ProtoMessage() {}

func (x *HellosGoodbyeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hello_hellopb_hellopb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HellosGoodbyeRequest.ProtoReflect.Descriptor instead.
func (*HellosGoodbyeRequest) Descriptor() ([]byte, []int) {
	return file_hello_hellopb_hellopb_proto_rawDescGZIP(), []int{5}
}

func (x *HellosGoodbyeRequest) GetHello() *Hello {
	if x != nil {
		return x.Hello
	}
	return nil
}

type HellosGoodbyeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Goodbye string `protobuf:"bytes,1,opt,name=goodbye,proto3" json:"goodbye,omitempty"`
}

func (x *HellosGoodbyeResponse) Reset() {
	*x = HellosGoodbyeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_hellopb_hellopb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HellosGoodbyeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HellosGoodbyeResponse) ProtoMessage() {}

func (x *HellosGoodbyeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hello_hellopb_hellopb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HellosGoodbyeResponse.ProtoReflect.Descriptor instead.
func (*HellosGoodbyeResponse) Descriptor() ([]byte, []int) {
	return file_hello_hellopb_hellopb_proto_rawDescGZIP(), []int{6}
}

func (x *HellosGoodbyeResponse) GetGoodbye() string {
	if x != nil {
		return x.Goodbye
	}
	return ""
}

type GoodbyeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hello *Hello `protobuf:"bytes,1,opt,name=hello,proto3" json:"hello,omitempty"`
}

func (x *GoodbyeRequest) Reset() {
	*x = GoodbyeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_hellopb_hellopb_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodbyeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodbyeRequest) ProtoMessage() {}

func (x *GoodbyeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hello_hellopb_hellopb_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodbyeRequest.ProtoReflect.Descriptor instead.
func (*GoodbyeRequest) Descriptor() ([]byte, []int) {
	return file_hello_hellopb_hellopb_proto_rawDescGZIP(), []int{7}
}

func (x *GoodbyeRequest) GetHello() *Hello {
	if x != nil {
		return x.Hello
	}
	return nil
}

type GoodbyeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Goodbye string `protobuf:"bytes,1,opt,name=goodbye,proto3" json:"goodbye,omitempty"`
}

func (x *GoodbyeResponse) Reset() {
	*x = GoodbyeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hello_hellopb_hellopb_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodbyeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodbyeResponse) ProtoMessage() {}

func (x *GoodbyeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hello_hellopb_hellopb_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodbyeResponse.ProtoReflect.Descriptor instead.
func (*GoodbyeResponse) Descriptor() ([]byte, []int) {
	return file_hello_hellopb_hellopb_proto_rawDescGZIP(), []int{8}
}

func (x *GoodbyeResponse) GetGoodbye() string {
	if x != nil {
		return x.Goodbye
	}
	return ""
}

var File_hello_hellopb_hellopb_proto protoreflect.FileDescriptor

var file_hello_hellopb_hellopb_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x70, 0x62, 0x2f,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x22, 0x3e, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x1d, 0x0a,
	0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x22, 0x32, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x22, 0x32, 0x0a, 0x0d, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x5f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x22, 0x3f, 0x0a, 0x19,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x4d, 0x61, 0x6e, 0x79, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x22, 0x43, 0x0a,
	0x1a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x4d, 0x61, 0x6e, 0x79, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x22, 0x3a, 0x0a, 0x14, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x73, 0x47, 0x6f, 0x6f, 0x64,
	0x62, 0x79, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x22, 0x31,
	0x0a, 0x15, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x73, 0x47, 0x6f, 0x6f, 0x64, 0x62, 0x79, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x62,
	0x79, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x62, 0x79,
	0x65, 0x22, 0x34, 0x0a, 0x0e, 0x47, 0x6f, 0x6f, 0x64, 0x62, 0x79, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x05, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x22, 0x2b, 0x0a, 0x0f, 0x47, 0x6f, 0x6f, 0x64, 0x62,
	0x79, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x6f,
	0x6f, 0x64, 0x62, 0x79, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x6f, 0x6f,
	0x64, 0x62, 0x79, 0x65, 0x32, 0xb3, 0x02, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x13,
	0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x12, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x4d, 0x61, 0x6e, 0x79, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x73, 0x12, 0x20, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x4d,
	0x61, 0x6e, 0x79, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x4d, 0x61, 0x6e, 0x79, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4e, 0x0a, 0x0d, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x73, 0x47, 0x6f, 0x6f, 0x64, 0x62, 0x79, 0x65, 0x12, 0x1b, 0x2e, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x73, 0x47, 0x6f, 0x6f, 0x64, 0x62, 0x79,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x73, 0x47, 0x6f, 0x6f, 0x64, 0x62, 0x79, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x3e, 0x0a, 0x07, 0x47, 0x6f,
	0x6f, 0x64, 0x62, 0x79, 0x65, 0x12, 0x15, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x47, 0x6f,
	0x6f, 0x64, 0x62, 0x79, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x62, 0x79, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0f, 0x5a, 0x0d, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_hello_hellopb_hellopb_proto_rawDescOnce sync.Once
	file_hello_hellopb_hellopb_proto_rawDescData = file_hello_hellopb_hellopb_proto_rawDesc
)

func file_hello_hellopb_hellopb_proto_rawDescGZIP() []byte {
	file_hello_hellopb_hellopb_proto_rawDescOnce.Do(func() {
		file_hello_hellopb_hellopb_proto_rawDescData = protoimpl.X.CompressGZIP(file_hello_hellopb_hellopb_proto_rawDescData)
	})
	return file_hello_hellopb_hellopb_proto_rawDescData
}

var file_hello_hellopb_hellopb_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_hello_hellopb_hellopb_proto_goTypes = []interface{}{
	(*Hello)(nil),                      // 0: hello.Hello
	(*HelloRequest)(nil),               // 1: hello.HelloRequest
	(*HelloResponse)(nil),              // 2: hello.HelloResponse
	(*HelloManyLanguagesRequest)(nil),  // 3: hello.HelloManyLanguagesRequest
	(*HelloManyLanguagesResponse)(nil), // 4: hello.HelloManyLanguagesResponse
	(*HellosGoodbyeRequest)(nil),       // 5: hello.HellosGoodbyeRequest
	(*HellosGoodbyeResponse)(nil),      // 6: hello.HellosGoodbyeResponse
	(*GoodbyeRequest)(nil),             // 7: hello.GoodbyeRequest
	(*GoodbyeResponse)(nil),            // 8: hello.GoodbyeResponse
}
var file_hello_hellopb_hellopb_proto_depIdxs = []int32{
	0, // 0: hello.HelloRequest.hello:type_name -> hello.Hello
	0, // 1: hello.HelloManyLanguagesRequest.hello:type_name -> hello.Hello
	0, // 2: hello.HellosGoodbyeRequest.hello:type_name -> hello.Hello
	0, // 3: hello.GoodbyeRequest.hello:type_name -> hello.Hello
	1, // 4: hello.HelloService.Hello:input_type -> hello.HelloRequest
	3, // 5: hello.HelloService.HelloManyLanguages:input_type -> hello.HelloManyLanguagesRequest
	5, // 6: hello.HelloService.HellosGoodbye:input_type -> hello.HellosGoodbyeRequest
	7, // 7: hello.HelloService.Goodbye:input_type -> hello.GoodbyeRequest
	2, // 8: hello.HelloService.Hello:output_type -> hello.HelloResponse
	4, // 9: hello.HelloService.HelloManyLanguages:output_type -> hello.HelloManyLanguagesResponse
	6, // 10: hello.HelloService.HellosGoodbye:output_type -> hello.HellosGoodbyeResponse
	8, // 11: hello.HelloService.Goodbye:output_type -> hello.GoodbyeResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_hello_hellopb_hellopb_proto_init() }
func file_hello_hellopb_hellopb_proto_init() {
	if File_hello_hellopb_hellopb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hello_hellopb_hellopb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Hello); i {
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
		file_hello_hellopb_hellopb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_hello_hellopb_hellopb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloResponse); i {
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
		file_hello_hellopb_hellopb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloManyLanguagesRequest); i {
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
		file_hello_hellopb_hellopb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloManyLanguagesResponse); i {
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
		file_hello_hellopb_hellopb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HellosGoodbyeRequest); i {
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
		file_hello_hellopb_hellopb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HellosGoodbyeResponse); i {
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
		file_hello_hellopb_hellopb_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodbyeRequest); i {
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
		file_hello_hellopb_hellopb_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodbyeResponse); i {
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
			RawDescriptor: file_hello_hellopb_hellopb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hello_hellopb_hellopb_proto_goTypes,
		DependencyIndexes: file_hello_hellopb_hellopb_proto_depIdxs,
		MessageInfos:      file_hello_hellopb_hellopb_proto_msgTypes,
	}.Build()
	File_hello_hellopb_hellopb_proto = out.File
	file_hello_hellopb_hellopb_proto_rawDesc = nil
	file_hello_hellopb_hellopb_proto_goTypes = nil
	file_hello_hellopb_hellopb_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloServiceClient interface {
	// Unary
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	// Server streaming
	// The service return hello/greeting in different languages
	HelloManyLanguages(ctx context.Context, in *HelloManyLanguagesRequest, opts ...grpc.CallOption) (HelloService_HelloManyLanguagesClient, error)
	// Client streaming
	// Send many hellos and response with one goodbye for all people
	HellosGoodbye(ctx context.Context, opts ...grpc.CallOption) (HelloService_HellosGoodbyeClient, error)
	// bidirectional streaming
	// It will send many hellos and the the server will response a goodbye by each one of them
	Goodbye(ctx context.Context, opts ...grpc.CallOption) (HelloService_GoodbyeClient, error)
}

type helloServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloServiceClient(cc grpc.ClientConnInterface) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/hello.HelloService/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) HelloManyLanguages(ctx context.Context, in *HelloManyLanguagesRequest, opts ...grpc.CallOption) (HelloService_HelloManyLanguagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HelloService_serviceDesc.Streams[0], "/hello.HelloService/HelloManyLanguages", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceHelloManyLanguagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HelloService_HelloManyLanguagesClient interface {
	Recv() (*HelloManyLanguagesResponse, error)
	grpc.ClientStream
}

type helloServiceHelloManyLanguagesClient struct {
	grpc.ClientStream
}

func (x *helloServiceHelloManyLanguagesClient) Recv() (*HelloManyLanguagesResponse, error) {
	m := new(HelloManyLanguagesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloServiceClient) HellosGoodbye(ctx context.Context, opts ...grpc.CallOption) (HelloService_HellosGoodbyeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HelloService_serviceDesc.Streams[1], "/hello.HelloService/HellosGoodbye", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceHellosGoodbyeClient{stream}
	return x, nil
}

type HelloService_HellosGoodbyeClient interface {
	Send(*HellosGoodbyeRequest) error
	CloseAndRecv() (*HellosGoodbyeResponse, error)
	grpc.ClientStream
}

type helloServiceHellosGoodbyeClient struct {
	grpc.ClientStream
}

func (x *helloServiceHellosGoodbyeClient) Send(m *HellosGoodbyeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloServiceHellosGoodbyeClient) CloseAndRecv() (*HellosGoodbyeResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(HellosGoodbyeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloServiceClient) Goodbye(ctx context.Context, opts ...grpc.CallOption) (HelloService_GoodbyeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HelloService_serviceDesc.Streams[2], "/hello.HelloService/Goodbye", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceGoodbyeClient{stream}
	return x, nil
}

type HelloService_GoodbyeClient interface {
	Send(*GoodbyeRequest) error
	Recv() (*GoodbyeResponse, error)
	grpc.ClientStream
}

type helloServiceGoodbyeClient struct {
	grpc.ClientStream
}

func (x *helloServiceGoodbyeClient) Send(m *GoodbyeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloServiceGoodbyeClient) Recv() (*GoodbyeResponse, error) {
	m := new(GoodbyeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloServiceServer is the server API for HelloService service.
type HelloServiceServer interface {
	// Unary
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	// Server streaming
	// The service return hello/greeting in different languages
	HelloManyLanguages(*HelloManyLanguagesRequest, HelloService_HelloManyLanguagesServer) error
	// Client streaming
	// Send many hellos and response with one goodbye for all people
	HellosGoodbye(HelloService_HellosGoodbyeServer) error
	// bidirectional streaming
	// It will send many hellos and the the server will response a goodbye by each one of them
	Goodbye(HelloService_GoodbyeServer) error
}

// UnimplementedHelloServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHelloServiceServer struct {
}

func (*UnimplementedHelloServiceServer) Hello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (*UnimplementedHelloServiceServer) HelloManyLanguages(*HelloManyLanguagesRequest, HelloService_HelloManyLanguagesServer) error {
	return status.Errorf(codes.Unimplemented, "method HelloManyLanguages not implemented")
}
func (*UnimplementedHelloServiceServer) HellosGoodbye(HelloService_HellosGoodbyeServer) error {
	return status.Errorf(codes.Unimplemented, "method HellosGoodbye not implemented")
}
func (*UnimplementedHelloServiceServer) Goodbye(HelloService_GoodbyeServer) error {
	return status.Errorf(codes.Unimplemented, "method Goodbye not implemented")
}

func RegisterHelloServiceServer(s *grpc.Server, srv HelloServiceServer) {
	s.RegisterService(&_HelloService_serviceDesc, srv)
}

func _HelloService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.HelloService/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_HelloManyLanguages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloManyLanguagesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HelloServiceServer).HelloManyLanguages(m, &helloServiceHelloManyLanguagesServer{stream})
}

type HelloService_HelloManyLanguagesServer interface {
	Send(*HelloManyLanguagesResponse) error
	grpc.ServerStream
}

type helloServiceHelloManyLanguagesServer struct {
	grpc.ServerStream
}

func (x *helloServiceHelloManyLanguagesServer) Send(m *HelloManyLanguagesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _HelloService_HellosGoodbye_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServiceServer).HellosGoodbye(&helloServiceHellosGoodbyeServer{stream})
}

type HelloService_HellosGoodbyeServer interface {
	SendAndClose(*HellosGoodbyeResponse) error
	Recv() (*HellosGoodbyeRequest, error)
	grpc.ServerStream
}

type helloServiceHellosGoodbyeServer struct {
	grpc.ServerStream
}

func (x *helloServiceHellosGoodbyeServer) SendAndClose(m *HellosGoodbyeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloServiceHellosGoodbyeServer) Recv() (*HellosGoodbyeRequest, error) {
	m := new(HellosGoodbyeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _HelloService_Goodbye_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServiceServer).Goodbye(&helloServiceGoodbyeServer{stream})
}

type HelloService_GoodbyeServer interface {
	Send(*GoodbyeResponse) error
	Recv() (*GoodbyeRequest, error)
	grpc.ServerStream
}

type helloServiceGoodbyeServer struct {
	grpc.ServerStream
}

func (x *helloServiceGoodbyeServer) Send(m *GoodbyeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloServiceGoodbyeServer) Recv() (*GoodbyeRequest, error) {
	m := new(GoodbyeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _HelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _HelloService_Hello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "HelloManyLanguages",
			Handler:       _HelloService_HelloManyLanguages_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "HellosGoodbye",
			Handler:       _HelloService_HellosGoodbye_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Goodbye",
			Handler:       _HelloService_Goodbye_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "hello/hellopb/hellopb.proto",
}
