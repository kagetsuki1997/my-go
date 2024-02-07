// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: pkg/protobuf/doge/doge.proto

package doge

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	uuid "my-go/pkg/protobuf/uuid"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DogeType int32

const (
	DogeType_DOGE_TYPE_UNSPECIFIED       DogeType = 0
	DogeType_DOGE_TYPE_DOGE              DogeType = 1
	DogeType_DOGE_TYPE_DOGE_CRY          DogeType = 2
	DogeType_DOGE_TYPE_DOGE_BUFFED       DogeType = 3
	DogeType_DOGE_TYPE_DOGE_KACHITORITAI DogeType = 4
)

// Enum value maps for DogeType.
var (
	DogeType_name = map[int32]string{
		0: "DOGE_TYPE_UNSPECIFIED",
		1: "DOGE_TYPE_DOGE",
		2: "DOGE_TYPE_DOGE_CRY",
		3: "DOGE_TYPE_DOGE_BUFFED",
		4: "DOGE_TYPE_DOGE_KACHITORITAI",
	}
	DogeType_value = map[string]int32{
		"DOGE_TYPE_UNSPECIFIED":       0,
		"DOGE_TYPE_DOGE":              1,
		"DOGE_TYPE_DOGE_CRY":          2,
		"DOGE_TYPE_DOGE_BUFFED":       3,
		"DOGE_TYPE_DOGE_KACHITORITAI": 4,
	}
)

func (x DogeType) Enum() *DogeType {
	p := new(DogeType)
	*p = x
	return p
}

func (x DogeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DogeType) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_protobuf_doge_doge_proto_enumTypes[0].Descriptor()
}

func (DogeType) Type() protoreflect.EnumType {
	return &file_pkg_protobuf_doge_doge_proto_enumTypes[0]
}

func (x DogeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DogeType.Descriptor instead.
func (DogeType) EnumDescriptor() ([]byte, []int) {
	return file_pkg_protobuf_doge_doge_proto_rawDescGZIP(), []int{0}
}

type Doge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              *uuid.Uuid             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	MagicNumber     int64                  `protobuf:"varint,3,opt,name=magic_number,json=magicNumber,proto3" json:"magic_number,omitempty"`
	Type            DogeType               `protobuf:"varint,4,opt,name=type,proto3,enum=doge.DogeType" json:"type,omitempty"`
	CreateTimestamp *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_timestamp,json=createTimestamp,proto3" json:"create_timestamp,omitempty"`
}

func (x *Doge) Reset() {
	*x = Doge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_protobuf_doge_doge_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Doge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Doge) ProtoMessage() {}

func (x *Doge) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_protobuf_doge_doge_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Doge.ProtoReflect.Descriptor instead.
func (*Doge) Descriptor() ([]byte, []int) {
	return file_pkg_protobuf_doge_doge_proto_rawDescGZIP(), []int{0}
}

func (x *Doge) GetId() *uuid.Uuid {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Doge) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Doge) GetMagicNumber() int64 {
	if x != nil {
		return x.MagicNumber
	}
	return 0
}

func (x *Doge) GetType() DogeType {
	if x != nil {
		return x.Type
	}
	return DogeType_DOGE_TYPE_UNSPECIFIED
}

func (x *Doge) GetCreateTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTimestamp
	}
	return nil
}

var File_pkg_protobuf_doge_doge_proto protoreflect.FileDescriptor

var file_pkg_protobuf_doge_doge_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x6f, 0x67, 0x65, 0x2f, 0x64, 0x6f, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x64, 0x6f, 0x67, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x75, 0x75, 0x69, 0x64, 0x2f, 0x75, 0x75, 0x69, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc4, 0x01, 0x0a, 0x04, 0x44, 0x6f, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x75, 0x69, 0x64, 0x2e,
	0x55, 0x75, 0x69, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x6d, 0x61, 0x67, 0x69, 0x63, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x6d, 0x61, 0x67, 0x69, 0x63, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x22, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e,
	0x64, 0x6f, 0x67, 0x65, 0x2e, 0x44, 0x6f, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x45, 0x0a, 0x10, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2a, 0x8d, 0x01, 0x0a, 0x08, 0x44,
	0x6f, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x4f, 0x47, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x4f, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x44, 0x4f, 0x47, 0x45, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x44, 0x4f, 0x47, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x44, 0x4f, 0x47, 0x45, 0x5f, 0x43, 0x52, 0x59, 0x10, 0x02, 0x12, 0x19,
	0x0a, 0x15, 0x44, 0x4f, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x4f, 0x47, 0x45,
	0x5f, 0x42, 0x55, 0x46, 0x46, 0x45, 0x44, 0x10, 0x03, 0x12, 0x1f, 0x0a, 0x1b, 0x44, 0x4f, 0x47,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x4f, 0x47, 0x45, 0x5f, 0x4b, 0x41, 0x43, 0x48,
	0x49, 0x54, 0x4f, 0x52, 0x49, 0x54, 0x41, 0x49, 0x10, 0x04, 0x42, 0x19, 0x5a, 0x17, 0x6d, 0x79,
	0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x6f, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_protobuf_doge_doge_proto_rawDescOnce sync.Once
	file_pkg_protobuf_doge_doge_proto_rawDescData = file_pkg_protobuf_doge_doge_proto_rawDesc
)

func file_pkg_protobuf_doge_doge_proto_rawDescGZIP() []byte {
	file_pkg_protobuf_doge_doge_proto_rawDescOnce.Do(func() {
		file_pkg_protobuf_doge_doge_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_protobuf_doge_doge_proto_rawDescData)
	})
	return file_pkg_protobuf_doge_doge_proto_rawDescData
}

var file_pkg_protobuf_doge_doge_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_protobuf_doge_doge_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_protobuf_doge_doge_proto_goTypes = []interface{}{
	(DogeType)(0),                 // 0: doge.DogeType
	(*Doge)(nil),                  // 1: doge.Doge
	(*uuid.Uuid)(nil),             // 2: uuid.Uuid
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_pkg_protobuf_doge_doge_proto_depIdxs = []int32{
	2, // 0: doge.Doge.id:type_name -> uuid.Uuid
	0, // 1: doge.Doge.type:type_name -> doge.DogeType
	3, // 2: doge.Doge.create_timestamp:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_protobuf_doge_doge_proto_init() }
func file_pkg_protobuf_doge_doge_proto_init() {
	if File_pkg_protobuf_doge_doge_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_protobuf_doge_doge_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Doge); i {
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
			RawDescriptor: file_pkg_protobuf_doge_doge_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_protobuf_doge_doge_proto_goTypes,
		DependencyIndexes: file_pkg_protobuf_doge_doge_proto_depIdxs,
		EnumInfos:         file_pkg_protobuf_doge_doge_proto_enumTypes,
		MessageInfos:      file_pkg_protobuf_doge_doge_proto_msgTypes,
	}.Build()
	File_pkg_protobuf_doge_doge_proto = out.File
	file_pkg_protobuf_doge_doge_proto_rawDesc = nil
	file_pkg_protobuf_doge_doge_proto_goTypes = nil
	file_pkg_protobuf_doge_doge_proto_depIdxs = nil
}
