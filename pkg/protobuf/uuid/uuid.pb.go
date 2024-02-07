// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: pkg/protobuf/uuid/uuid.proto

package uuid

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

type Uuid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HighBits uint64 `protobuf:"fixed64,1,opt,name=high_bits,json=highBits,proto3" json:"high_bits,omitempty"`
	LowBits  uint64 `protobuf:"fixed64,2,opt,name=low_bits,json=lowBits,proto3" json:"low_bits,omitempty"`
}

func (x *Uuid) Reset() {
	*x = Uuid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protobuf_uuid_uuid_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Uuid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Uuid) ProtoMessage() {}

func (x *Uuid) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protobuf_uuid_uuid_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Uuid.ProtoReflect.Descriptor instead.
func (*Uuid) Descriptor() ([]byte, []int) {
	return file_pkg_protobuf_uuid_uuid_proto_rawDescGZIP(), []int{0}
}

func (x *Uuid) GetHighBits() uint64 {
	if x != nil {
		return x.HighBits
	}
	return 0
}

func (x *Uuid) GetLowBits() uint64 {
	if x != nil {
		return x.LowBits
	}
	return 0
}

var File_pkg_protobuf_uuid_uuid_proto protoreflect.FileDescriptor

var file_pkg_protobuf_uuid_uuid_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x75,
	0x75, 0x69, 0x64, 0x2f, 0x75, 0x75, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x22, 0x3e, 0x0a, 0x04, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x68, 0x69, 0x67, 0x68, 0x5f, 0x62, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x06, 0x52,
	0x08, 0x68, 0x69, 0x67, 0x68, 0x42, 0x69, 0x74, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x77,
	0x5f, 0x62, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x06, 0x52, 0x07, 0x6c, 0x6f, 0x77,
	0x42, 0x69, 0x74, 0x73, 0x42, 0x19, 0x5a, 0x17, 0x6d, 0x79, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x75, 0x75, 0x69, 0x64, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_protobuf_uuid_uuid_proto_rawDescOnce sync.Once
	file_pkg_protobuf_uuid_uuid_proto_rawDescData = file_pkg_protobuf_uuid_uuid_proto_rawDesc
)

func file_pkg_protobuf_uuid_uuid_proto_rawDescGZIP() []byte {
	file_pkg_protobuf_uuid_uuid_proto_rawDescOnce.Do(func() {
		file_pkg_protobuf_uuid_uuid_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_protobuf_uuid_uuid_proto_rawDescData)
	})
	return file_pkg_protobuf_uuid_uuid_proto_rawDescData
}

var file_pkg_protobuf_uuid_uuid_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_protobuf_uuid_uuid_proto_goTypes = []interface{}{
	(*Uuid)(nil), // 0: uuid.Uuid
}
var file_pkg_protobuf_uuid_uuid_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_protobuf_uuid_uuid_proto_init() }
func file_pkg_protobuf_uuid_uuid_proto_init() {
	if File_pkg_protobuf_uuid_uuid_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_protobuf_uuid_uuid_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Uuid); i {
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
			RawDescriptor: file_pkg_protobuf_uuid_uuid_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_protobuf_uuid_uuid_proto_goTypes,
		DependencyIndexes: file_pkg_protobuf_uuid_uuid_proto_depIdxs,
		MessageInfos:      file_pkg_protobuf_uuid_uuid_proto_msgTypes,
	}.Build()
	File_pkg_protobuf_uuid_uuid_proto = out.File
	file_pkg_protobuf_uuid_uuid_proto_rawDesc = nil
	file_pkg_protobuf_uuid_uuid_proto_goTypes = nil
	file_pkg_protobuf_uuid_uuid_proto_depIdxs = nil
}