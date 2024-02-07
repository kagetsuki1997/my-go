// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: pkg/protobuf/doge/v1/doge_service.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	common "my-go/pkg/protobuf/uuid"
	doge "my-go/pkg/protobuf/doge"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateDogeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MagicNumber int64         `protobuf:"varint,2,opt,name=magic_number,json=magicNumber,proto3" json:"magic_number,omitempty"`
	Type        doge.DogeType `protobuf:"varint,3,opt,name=type,proto3,enum=doge.DogeType" json:"type,omitempty"`
}

func (x *CreateDogeRequest) Reset() {
	*x = CreateDogeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protobuf_doge_v1_doge_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDogeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDogeRequest) ProtoMessage() {}

func (x *CreateDogeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protobuf_doge_v1_doge_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDogeRequest.ProtoReflect.Descriptor instead.
func (*CreateDogeRequest) Descriptor() ([]byte, []int) {
	return file_pkg_protobuf_doge_v1_doge_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateDogeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateDogeRequest) GetMagicNumber() int64 {
	if x != nil {
		return x.MagicNumber
	}
	return 0
}

func (x *CreateDogeRequest) GetType() doge.DogeType {
	if x != nil {
		return x.Type
	}
	return doge.DogeType(0)
}

var File_pkg_protobuf_doge_v1_doge_service_proto protoreflect.FileDescriptor

var file_pkg_protobuf_doge_v1_doge_service_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x6f, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x64, 0x6f, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x1a, 0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x75, 0x75, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x6f, 0x67, 0x65, 0x2f, 0x64, 0x6f, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x6e, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61, 0x67,
	0x69, 0x63, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x6d, 0x61, 0x67, 0x69, 0x63, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x64, 0x6f, 0x67,
	0x65, 0x2e, 0x44, 0x6f, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x32, 0x68, 0x0a, 0x0b, 0x44, 0x6f, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x23, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x67, 0x65, 0x12, 0x0c, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x1a, 0x0a, 0x2e, 0x64, 0x6f, 0x67, 0x65, 0x2e,
	0x44, 0x6f, 0x67, 0x65, 0x12, 0x34, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x6f,
	0x67, 0x65, 0x12, 0x1a, 0x2e, 0x64, 0x6f, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x44, 0x6f, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a,
	0x2e, 0x64, 0x6f, 0x67, 0x65, 0x2e, 0x44, 0x6f, 0x67, 0x65, 0x42, 0x1c, 0x5a, 0x1a, 0x6d, 0x79,
	0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x6f, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_protobuf_doge_v1_doge_service_proto_rawDescOnce sync.Once
	file_pkg_protobuf_doge_v1_doge_service_proto_rawDescData = file_pkg_protobuf_doge_v1_doge_service_proto_rawDesc
)

func file_pkg_protobuf_doge_v1_doge_service_proto_rawDescGZIP() []byte {
	file_pkg_protobuf_doge_v1_doge_service_proto_rawDescOnce.Do(func() {
		file_pkg_protobuf_doge_v1_doge_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_protobuf_doge_v1_doge_service_proto_rawDescData)
	})
	return file_pkg_protobuf_doge_v1_doge_service_proto_rawDescData
}

var file_pkg_protobuf_doge_v1_doge_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_protobuf_doge_v1_doge_service_proto_goTypes = []interface{}{
	(*CreateDogeRequest)(nil), // 0: doge.v1.CreateDogeRequest
	(doge.DogeType)(0),        // 1: doge.DogeType
	(*common.Uuid)(nil),       // 2: common.Uuid
	(*doge.Doge)(nil),         // 3: doge.Doge
}
var file_pkg_protobuf_doge_v1_doge_service_proto_depIdxs = []int32{
	1, // 0: doge.v1.CreateDogeRequest.type:type_name -> doge.DogeType
	2, // 1: doge.v1.DogeService.GetDoge:input_type -> common.Uuid
	0, // 2: doge.v1.DogeService.CreateDoge:input_type -> doge.v1.CreateDogeRequest
	3, // 3: doge.v1.DogeService.GetDoge:output_type -> doge.Doge
	3, // 4: doge.v1.DogeService.CreateDoge:output_type -> doge.Doge
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_protobuf_doge_v1_doge_service_proto_init() }
func file_pkg_protobuf_doge_v1_doge_service_proto_init() {
	if File_pkg_protobuf_doge_v1_doge_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_protobuf_doge_v1_doge_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDogeRequest); i {
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
			RawDescriptor: file_pkg_protobuf_doge_v1_doge_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_protobuf_doge_v1_doge_service_proto_goTypes,
		DependencyIndexes: file_pkg_protobuf_doge_v1_doge_service_proto_depIdxs,
		MessageInfos:      file_pkg_protobuf_doge_v1_doge_service_proto_msgTypes,
	}.Build()
	File_pkg_protobuf_doge_v1_doge_service_proto = out.File
	file_pkg_protobuf_doge_v1_doge_service_proto_rawDesc = nil
	file_pkg_protobuf_doge_v1_doge_service_proto_goTypes = nil
	file_pkg_protobuf_doge_v1_doge_service_proto_depIdxs = nil
}
