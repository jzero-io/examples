// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.19.4
// source: v1/hello.proto

package hellopb

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SayHelloRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SayHelloRequest) Reset() {
	*x = SayHelloRequest{}
	mi := &file_v1_hello_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SayHelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloRequest) ProtoMessage() {}

func (x *SayHelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_hello_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloRequest.ProtoReflect.Descriptor instead.
func (*SayHelloRequest) Descriptor() ([]byte, []int) {
	return file_v1_hello_proto_rawDescGZIP(), []int{0}
}

func (x *SayHelloRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SayHelloResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SayHelloResponse) Reset() {
	*x = SayHelloResponse{}
	mi := &file_v1_hello_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SayHelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloResponse) ProtoMessage() {}

func (x *SayHelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_hello_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloResponse.ProtoReflect.Descriptor instead.
func (*SayHelloResponse) Descriptor() ([]byte, []int) {
	return file_v1_hello_proto_rawDescGZIP(), []int{1}
}

func (x *SayHelloResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_v1_hello_proto protoreflect.FileDescriptor

var file_v1_hello_proto_rawDesc = string([]byte{
	0x0a, 0x0e, 0x76, 0x31, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x07, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x70, 0x62, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x35, 0x0a, 0x0f, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x98, 0x01, 0x05,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2c, 0x0a, 0x10, 0x53, 0x61, 0x79,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x4a, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x12, 0x41, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x18, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x70, 0x62, 0x2e, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x70, 0x62,
	0x2e, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x70, 0x62, 0x2f, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_v1_hello_proto_rawDescOnce sync.Once
	file_v1_hello_proto_rawDescData []byte
)

func file_v1_hello_proto_rawDescGZIP() []byte {
	file_v1_hello_proto_rawDescOnce.Do(func() {
		file_v1_hello_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_v1_hello_proto_rawDesc), len(file_v1_hello_proto_rawDesc)))
	})
	return file_v1_hello_proto_rawDescData
}

var file_v1_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_hello_proto_goTypes = []any{
	(*SayHelloRequest)(nil),  // 0: hellopb.SayHelloRequest
	(*SayHelloResponse)(nil), // 1: hellopb.SayHelloResponse
}
var file_v1_hello_proto_depIdxs = []int32{
	0, // 0: hellopb.Hello.SayHello:input_type -> hellopb.SayHelloRequest
	1, // 1: hellopb.Hello.SayHello:output_type -> hellopb.SayHelloResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_hello_proto_init() }
func file_v1_hello_proto_init() {
	if File_v1_hello_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_v1_hello_proto_rawDesc), len(file_v1_hello_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_hello_proto_goTypes,
		DependencyIndexes: file_v1_hello_proto_depIdxs,
		MessageInfos:      file_v1_hello_proto_msgTypes,
	}.Build()
	File_v1_hello_proto = out.File
	file_v1_hello_proto_goTypes = nil
	file_v1_hello_proto_depIdxs = nil
}
