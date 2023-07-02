// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: teleporter/message.proto

package teleporter

import (
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

type SignRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceChainId      []byte `protobuf:"bytes,1,opt,name=source_chain_id,json=sourceChainId,proto3" json:"source_chain_id,omitempty"`
	DestinationChainId []byte `protobuf:"bytes,2,opt,name=destination_chain_id,json=destinationChainId,proto3" json:"destination_chain_id,omitempty"`
	Payload            []byte `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *SignRequest) Reset() {
	*x = SignRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleporter_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignRequest) ProtoMessage() {}

func (x *SignRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleporter_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignRequest.ProtoReflect.Descriptor instead.
func (*SignRequest) Descriptor() ([]byte, []int) {
	return file_teleporter_message_proto_rawDescGZIP(), []int{0}
}

func (x *SignRequest) GetSourceChainId() []byte {
	if x != nil {
		return x.SourceChainId
	}
	return nil
}

func (x *SignRequest) GetDestinationChainId() []byte {
	if x != nil {
		return x.DestinationChainId
	}
	return nil
}

func (x *SignRequest) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type SignResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *SignResponse) Reset() {
	*x = SignResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleporter_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignResponse) ProtoMessage() {}

func (x *SignResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleporter_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignResponse.ProtoReflect.Descriptor instead.
func (*SignResponse) Descriptor() ([]byte, []int) {
	return file_teleporter_message_proto_rawDescGZIP(), []int{1}
}

func (x *SignResponse) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

var File_teleporter_message_proto protoreflect.FileDescriptor

var file_teleporter_message_proto_rawDesc = []byte{
	0x0a, 0x18, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x22, 0x81, 0x01, 0x0a, 0x0b, 0x53, 0x69, 0x67, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x30,
	0x0a, 0x14, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x12, 0x64, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x2c, 0x0a, 0x0c, 0x53, 0x69,
	0x67, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x32, 0x43, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e,
	0x65, 0x72, 0x12, 0x39, 0x0a, 0x04, 0x53, 0x69, 0x67, 0x6e, 0x12, 0x17, 0x2e, 0x74, 0x65, 0x6c,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x37, 0x5a,
	0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x61, 0x73, 0x74,
	0x68, 0x79, 0x70, 0x68, 0x65, 0x6e, 0x2f, 0x61, 0x76, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x68, 0x65,
	0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleporter_message_proto_rawDescOnce sync.Once
	file_teleporter_message_proto_rawDescData = file_teleporter_message_proto_rawDesc
)

func file_teleporter_message_proto_rawDescGZIP() []byte {
	file_teleporter_message_proto_rawDescOnce.Do(func() {
		file_teleporter_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleporter_message_proto_rawDescData)
	})
	return file_teleporter_message_proto_rawDescData
}

var file_teleporter_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_teleporter_message_proto_goTypes = []interface{}{
	(*SignRequest)(nil),  // 0: teleporter.SignRequest
	(*SignResponse)(nil), // 1: teleporter.SignResponse
}
var file_teleporter_message_proto_depIdxs = []int32{
	0, // 0: teleporter.Signer.Sign:input_type -> teleporter.SignRequest
	1, // 1: teleporter.Signer.Sign:output_type -> teleporter.SignResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_teleporter_message_proto_init() }
func file_teleporter_message_proto_init() {
	if File_teleporter_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_teleporter_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignRequest); i {
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
		file_teleporter_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignResponse); i {
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
			RawDescriptor: file_teleporter_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teleporter_message_proto_goTypes,
		DependencyIndexes: file_teleporter_message_proto_depIdxs,
		MessageInfos:      file_teleporter_message_proto_msgTypes,
	}.Build()
	File_teleporter_message_proto = out.File
	file_teleporter_message_proto_rawDesc = nil
	file_teleporter_message_proto_goTypes = nil
	file_teleporter_message_proto_depIdxs = nil
}
