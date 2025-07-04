// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.19.4
// source: v1/version.proto

package versionpb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type VersionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VersionRequest) Reset() {
	*x = VersionRequest{}
	mi := &file_v1_version_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VersionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionRequest) ProtoMessage() {}

func (x *VersionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_version_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionRequest.ProtoReflect.Descriptor instead.
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return file_v1_version_proto_rawDescGZIP(), []int{0}
}

type VersionResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Version       string                 `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	GoVersion     string                 `protobuf:"bytes,2,opt,name=goVersion,proto3" json:"goVersion,omitempty"`
	Commit        string                 `protobuf:"bytes,3,opt,name=commit,proto3" json:"commit,omitempty"`
	Date          string                 `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VersionResponse) Reset() {
	*x = VersionResponse{}
	mi := &file_v1_version_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VersionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionResponse) ProtoMessage() {}

func (x *VersionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_version_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionResponse.ProtoReflect.Descriptor instead.
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return file_v1_version_proto_rawDescGZIP(), []int{1}
}

func (x *VersionResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *VersionResponse) GetGoVersion() string {
	if x != nil {
		return x.GoVersion
	}
	return ""
}

func (x *VersionResponse) GetCommit() string {
	if x != nil {
		return x.Commit
	}
	return ""
}

func (x *VersionResponse) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

var File_v1_version_proto protoreflect.FileDescriptor

const file_v1_version_proto_rawDesc = "" +
	"\n" +
	"\x10v1/version.proto\x12\tversionpb\x1a\x1cgoogle/api/annotations.proto\x1a;grpc-gateway/protoc-gen-openapiv2/options/annotations.proto\"\x10\n" +
	"\x0eVersionRequest\"u\n" +
	"\x0fVersionResponse\x12\x18\n" +
	"\aversion\x18\x01 \x01(\tR\aversion\x12\x1c\n" +
	"\tgoVersion\x18\x02 \x01(\tR\tgoVersion\x12\x16\n" +
	"\x06commit\x18\x03 \x01(\tR\x06commit\x12\x12\n" +
	"\x04date\x18\x04 \x01(\tR\x04date2d\n" +
	"\aVersion\x12Y\n" +
	"\aVersion\x12\x19.versionpb.VersionRequest\x1a\x1a.versionpb.VersionResponse\"\x17\x82\xd3\xe4\x93\x02\x11\x12\x0f/api/v1/versionB\x19\x92A\x06\x12\x042\x02v1Z\x0e./pb/versionpbb\x06proto3"

var (
	file_v1_version_proto_rawDescOnce sync.Once
	file_v1_version_proto_rawDescData []byte
)

func file_v1_version_proto_rawDescGZIP() []byte {
	file_v1_version_proto_rawDescOnce.Do(func() {
		file_v1_version_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_v1_version_proto_rawDesc), len(file_v1_version_proto_rawDesc)))
	})
	return file_v1_version_proto_rawDescData
}

var file_v1_version_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_version_proto_goTypes = []any{
	(*VersionRequest)(nil),  // 0: versionpb.VersionRequest
	(*VersionResponse)(nil), // 1: versionpb.VersionResponse
}
var file_v1_version_proto_depIdxs = []int32{
	0, // 0: versionpb.Version.Version:input_type -> versionpb.VersionRequest
	1, // 1: versionpb.Version.Version:output_type -> versionpb.VersionResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_version_proto_init() }
func file_v1_version_proto_init() {
	if File_v1_version_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_v1_version_proto_rawDesc), len(file_v1_version_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_version_proto_goTypes,
		DependencyIndexes: file_v1_version_proto_depIdxs,
		MessageInfos:      file_v1_version_proto_msgTypes,
	}.Build()
	File_v1_version_proto = out.File
	file_v1_version_proto_goTypes = nil
	file_v1_version_proto_depIdxs = nil
}
