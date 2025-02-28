// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.2
// source: pkg/proto/key.proto

package proto

import (
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

type GetPrivateKeyRequestDto struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPrivateKeyRequestDto) Reset() {
	*x = GetPrivateKeyRequestDto{}
	mi := &file_pkg_proto_key_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPrivateKeyRequestDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPrivateKeyRequestDto) ProtoMessage() {}

func (x *GetPrivateKeyRequestDto) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_key_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPrivateKeyRequestDto.ProtoReflect.Descriptor instead.
func (*GetPrivateKeyRequestDto) Descriptor() ([]byte, []int) {
	return file_pkg_proto_key_proto_rawDescGZIP(), []int{0}
}

type GetPrivateKeyResponseDto struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PrivateKey    string                 `protobuf:"bytes,1,opt,name=privateKey,proto3" json:"privateKey,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPrivateKeyResponseDto) Reset() {
	*x = GetPrivateKeyResponseDto{}
	mi := &file_pkg_proto_key_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPrivateKeyResponseDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPrivateKeyResponseDto) ProtoMessage() {}

func (x *GetPrivateKeyResponseDto) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_key_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPrivateKeyResponseDto.ProtoReflect.Descriptor instead.
func (*GetPrivateKeyResponseDto) Descriptor() ([]byte, []int) {
	return file_pkg_proto_key_proto_rawDescGZIP(), []int{1}
}

func (x *GetPrivateKeyResponseDto) GetPrivateKey() string {
	if x != nil {
		return x.PrivateKey
	}
	return ""
}

type GetPublicKeyRequestDto struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPublicKeyRequestDto) Reset() {
	*x = GetPublicKeyRequestDto{}
	mi := &file_pkg_proto_key_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPublicKeyRequestDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPublicKeyRequestDto) ProtoMessage() {}

func (x *GetPublicKeyRequestDto) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_key_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPublicKeyRequestDto.ProtoReflect.Descriptor instead.
func (*GetPublicKeyRequestDto) Descriptor() ([]byte, []int) {
	return file_pkg_proto_key_proto_rawDescGZIP(), []int{2}
}

type GetPublicKeyResponseDto struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PublicKey     string                 `protobuf:"bytes,1,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPublicKeyResponseDto) Reset() {
	*x = GetPublicKeyResponseDto{}
	mi := &file_pkg_proto_key_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPublicKeyResponseDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPublicKeyResponseDto) ProtoMessage() {}

func (x *GetPublicKeyResponseDto) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_key_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPublicKeyResponseDto.ProtoReflect.Descriptor instead.
func (*GetPublicKeyResponseDto) Descriptor() ([]byte, []int) {
	return file_pkg_proto_key_proto_rawDescGZIP(), []int{3}
}

func (x *GetPublicKeyResponseDto) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

var File_pkg_proto_key_proto protoreflect.FileDescriptor

var file_pkg_proto_key_proto_rawDesc = string([]byte{
	0x0a, 0x13, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6b, 0x65, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x19, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x74, 0x6f,
	0x22, 0x3a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x74, 0x6f, 0x12, 0x1e, 0x0a, 0x0a,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x22, 0x18, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x44, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x74,
	0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x32,
	0x99, 0x01, 0x0a, 0x0e, 0x4b, 0x65, 0x79, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x44, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x4b, 0x65, 0x79, 0x12, 0x18, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x74, 0x6f, 0x1a, 0x19, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x74, 0x6f, 0x12, 0x41, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x17, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x74,
	0x6f, 0x1a, 0x18, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x74, 0x6f, 0x42, 0x0d, 0x5a, 0x0b, 0x2e,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_pkg_proto_key_proto_rawDescOnce sync.Once
	file_pkg_proto_key_proto_rawDescData []byte
)

func file_pkg_proto_key_proto_rawDescGZIP() []byte {
	file_pkg_proto_key_proto_rawDescOnce.Do(func() {
		file_pkg_proto_key_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_pkg_proto_key_proto_rawDesc), len(file_pkg_proto_key_proto_rawDesc)))
	})
	return file_pkg_proto_key_proto_rawDescData
}

var file_pkg_proto_key_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_proto_key_proto_goTypes = []any{
	(*GetPrivateKeyRequestDto)(nil),  // 0: GetPrivateKeyRequestDto
	(*GetPrivateKeyResponseDto)(nil), // 1: GetPrivateKeyResponseDto
	(*GetPublicKeyRequestDto)(nil),   // 2: GetPublicKeyRequestDto
	(*GetPublicKeyResponseDto)(nil),  // 3: GetPublicKeyResponseDto
}
var file_pkg_proto_key_proto_depIdxs = []int32{
	0, // 0: KeyGrpcService.GetPrivateKey:input_type -> GetPrivateKeyRequestDto
	2, // 1: KeyGrpcService.GetPublicKey:input_type -> GetPublicKeyRequestDto
	1, // 2: KeyGrpcService.GetPrivateKey:output_type -> GetPrivateKeyResponseDto
	3, // 3: KeyGrpcService.GetPublicKey:output_type -> GetPublicKeyResponseDto
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_proto_key_proto_init() }
func file_pkg_proto_key_proto_init() {
	if File_pkg_proto_key_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_pkg_proto_key_proto_rawDesc), len(file_pkg_proto_key_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_key_proto_goTypes,
		DependencyIndexes: file_pkg_proto_key_proto_depIdxs,
		MessageInfos:      file_pkg_proto_key_proto_msgTypes,
	}.Build()
	File_pkg_proto_key_proto = out.File
	file_pkg_proto_key_proto_goTypes = nil
	file_pkg_proto_key_proto_depIdxs = nil
}
