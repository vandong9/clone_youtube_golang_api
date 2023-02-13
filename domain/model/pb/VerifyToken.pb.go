// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.14.0
// source: domain/model/proto/VerifyToken.proto

package model

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

type VerifyToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *VerifyToken) Reset() {
	*x = VerifyToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_model_proto_VerifyToken_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyToken) ProtoMessage() {}

func (x *VerifyToken) ProtoReflect() protoreflect.Message {
	mi := &file_domain_model_proto_VerifyToken_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyToken.ProtoReflect.Descriptor instead.
func (*VerifyToken) Descriptor() ([]byte, []int) {
	return file_domain_model_proto_VerifyToken_proto_rawDescGZIP(), []int{0}
}

func (x *VerifyToken) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *VerifyToken) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_domain_model_proto_VerifyToken_proto protoreflect.FileDescriptor

var file_domain_model_proto_VerifyToken_proto_rawDesc = []byte{
	0x0a, 0x24, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x3c, 0x0a,
	0x0b, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x03, 0x5a, 0x01, 0x2e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_domain_model_proto_VerifyToken_proto_rawDescOnce sync.Once
	file_domain_model_proto_VerifyToken_proto_rawDescData = file_domain_model_proto_VerifyToken_proto_rawDesc
)

func file_domain_model_proto_VerifyToken_proto_rawDescGZIP() []byte {
	file_domain_model_proto_VerifyToken_proto_rawDescOnce.Do(func() {
		file_domain_model_proto_VerifyToken_proto_rawDescData = protoimpl.X.CompressGZIP(file_domain_model_proto_VerifyToken_proto_rawDescData)
	})
	return file_domain_model_proto_VerifyToken_proto_rawDescData
}

var file_domain_model_proto_VerifyToken_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_domain_model_proto_VerifyToken_proto_goTypes = []interface{}{
	(*VerifyToken)(nil), // 0: model.VerifyToken
}
var file_domain_model_proto_VerifyToken_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_domain_model_proto_VerifyToken_proto_init() }
func file_domain_model_proto_VerifyToken_proto_init() {
	if File_domain_model_proto_VerifyToken_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_domain_model_proto_VerifyToken_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyToken); i {
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
			RawDescriptor: file_domain_model_proto_VerifyToken_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_domain_model_proto_VerifyToken_proto_goTypes,
		DependencyIndexes: file_domain_model_proto_VerifyToken_proto_depIdxs,
		MessageInfos:      file_domain_model_proto_VerifyToken_proto_msgTypes,
	}.Build()
	File_domain_model_proto_VerifyToken_proto = out.File
	file_domain_model_proto_VerifyToken_proto_rawDesc = nil
	file_domain_model_proto_VerifyToken_proto_goTypes = nil
	file_domain_model_proto_VerifyToken_proto_depIdxs = nil
}
