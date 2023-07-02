// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: plugin/plugin.proto

package plugin

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ExitCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExitCode int32 `protobuf:"varint,1,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
}

func (x *ExitCodeResponse) Reset() {
	*x = ExitCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugin_plugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExitCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExitCodeResponse) ProtoMessage() {}

func (x *ExitCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_plugin_plugin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExitCodeResponse.ProtoReflect.Descriptor instead.
func (*ExitCodeResponse) Descriptor() ([]byte, []int) {
	return file_plugin_plugin_proto_rawDescGZIP(), []int{0}
}

func (x *ExitCodeResponse) GetExitCode() int32 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

var File_plugin_plugin_proto protoreflect.FileDescriptor

var file_plugin_plugin_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x10, 0x45, 0x78,
	0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x65, 0x78, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x65, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x32, 0xb5, 0x01, 0x0a, 0x04,
	0x4e, 0x6f, 0x64, 0x65, 0x12, 0x37, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3c, 0x0a,
	0x08, 0x45, 0x78, 0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x18, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x45, 0x78, 0x69, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x53,
	0x74, 0x6f, 0x70, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6c, 0x61, 0x73, 0x74, 0x68, 0x79, 0x70, 0x68, 0x65, 0x6e, 0x2f, 0x61, 0x76, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70,
	0x62, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plugin_plugin_proto_rawDescOnce sync.Once
	file_plugin_plugin_proto_rawDescData = file_plugin_plugin_proto_rawDesc
)

func file_plugin_plugin_proto_rawDescGZIP() []byte {
	file_plugin_plugin_proto_rawDescOnce.Do(func() {
		file_plugin_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugin_plugin_proto_rawDescData)
	})
	return file_plugin_plugin_proto_rawDescData
}

var file_plugin_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_plugin_plugin_proto_goTypes = []interface{}{
	(*ExitCodeResponse)(nil), // 0: plugin.ExitCodeResponse
	(*emptypb.Empty)(nil),    // 1: google.protobuf.Empty
}
var file_plugin_plugin_proto_depIdxs = []int32{
	1, // 0: plugin.Node.Start:input_type -> google.protobuf.Empty
	1, // 1: plugin.Node.ExitCode:input_type -> google.protobuf.Empty
	1, // 2: plugin.Node.Stop:input_type -> google.protobuf.Empty
	1, // 3: plugin.Node.Start:output_type -> google.protobuf.Empty
	0, // 4: plugin.Node.ExitCode:output_type -> plugin.ExitCodeResponse
	1, // 5: plugin.Node.Stop:output_type -> google.protobuf.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_plugin_plugin_proto_init() }
func file_plugin_plugin_proto_init() {
	if File_plugin_plugin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_plugin_plugin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExitCodeResponse); i {
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
			RawDescriptor: file_plugin_plugin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_plugin_plugin_proto_goTypes,
		DependencyIndexes: file_plugin_plugin_proto_depIdxs,
		MessageInfos:      file_plugin_plugin_proto_msgTypes,
	}.Build()
	File_plugin_plugin_proto = out.File
	file_plugin_plugin_proto_rawDesc = nil
	file_plugin_plugin_proto_goTypes = nil
	file_plugin_plugin_proto_depIdxs = nil
}
