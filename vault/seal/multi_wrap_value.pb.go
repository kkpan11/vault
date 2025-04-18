// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: vault/seal/multi_wrap_value.proto

package seal

import (
	v2 "github.com/hashicorp/go-kms-wrapping/v2"
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

// MultiWrapValue can be used to keep track of different encryptions of a value.
type MultiWrapValue struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Generation is used to keep track of when the MultiWrapValue was generated.
	Generation uint64 `protobuf:"varint,1,opt,name=generation,proto3" json:"generation,omitempty"`
	// Slots has a BlobInfo for each key used to encrypt the value.
	Slots         []*v2.BlobInfo `protobuf:"bytes,2,rep,name=slots,proto3" json:"slots,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MultiWrapValue) Reset() {
	*x = MultiWrapValue{}
	mi := &file_vault_seal_multi_wrap_value_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MultiWrapValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiWrapValue) ProtoMessage() {}

func (x *MultiWrapValue) ProtoReflect() protoreflect.Message {
	mi := &file_vault_seal_multi_wrap_value_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiWrapValue.ProtoReflect.Descriptor instead.
func (*MultiWrapValue) Descriptor() ([]byte, []int) {
	return file_vault_seal_multi_wrap_value_proto_rawDescGZIP(), []int{0}
}

func (x *MultiWrapValue) GetGeneration() uint64 {
	if x != nil {
		return x.Generation
	}
	return 0
}

func (x *MultiWrapValue) GetSlots() []*v2.BlobInfo {
	if x != nil {
		return x.Slots
	}
	return nil
}

var File_vault_seal_multi_wrap_value_proto protoreflect.FileDescriptor

var file_vault_seal_multi_wrap_value_proto_rawDesc = string([]byte{
	0x0a, 0x21, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2f, 0x73, 0x65, 0x61, 0x6c, 0x2f, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x5f, 0x77, 0x72, 0x61, 0x70, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x73, 0x65, 0x61, 0x6c, 0x1a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e,
	0x67, 0x6f, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x77, 0x72, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x32, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7f,
	0x0a, 0x0e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x57, 0x72, 0x61, 0x70, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x4d, 0x0a, 0x05, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x37, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x67, 0x6f, 0x2e, 0x6b, 0x6d, 0x73, 0x2e, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x42, 0x6c, 0x6f, 0x62, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x42,
	0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61,
	0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2f, 0x76, 0x61,
	0x75, 0x6c, 0x74, 0x2f, 0x73, 0x65, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_vault_seal_multi_wrap_value_proto_rawDescOnce sync.Once
	file_vault_seal_multi_wrap_value_proto_rawDescData []byte
)

func file_vault_seal_multi_wrap_value_proto_rawDescGZIP() []byte {
	file_vault_seal_multi_wrap_value_proto_rawDescOnce.Do(func() {
		file_vault_seal_multi_wrap_value_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_vault_seal_multi_wrap_value_proto_rawDesc), len(file_vault_seal_multi_wrap_value_proto_rawDesc)))
	})
	return file_vault_seal_multi_wrap_value_proto_rawDescData
}

var file_vault_seal_multi_wrap_value_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_vault_seal_multi_wrap_value_proto_goTypes = []any{
	(*MultiWrapValue)(nil), // 0: seal.MultiWrapValue
	(*v2.BlobInfo)(nil),    // 1: github.com.hashicorp.go.kms.wrapping.v2.types.BlobInfo
}
var file_vault_seal_multi_wrap_value_proto_depIdxs = []int32{
	1, // 0: seal.MultiWrapValue.slots:type_name -> github.com.hashicorp.go.kms.wrapping.v2.types.BlobInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_vault_seal_multi_wrap_value_proto_init() }
func file_vault_seal_multi_wrap_value_proto_init() {
	if File_vault_seal_multi_wrap_value_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_vault_seal_multi_wrap_value_proto_rawDesc), len(file_vault_seal_multi_wrap_value_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vault_seal_multi_wrap_value_proto_goTypes,
		DependencyIndexes: file_vault_seal_multi_wrap_value_proto_depIdxs,
		MessageInfos:      file_vault_seal_multi_wrap_value_proto_msgTypes,
	}.Build()
	File_vault_seal_multi_wrap_value_proto = out.File
	file_vault_seal_multi_wrap_value_proto_goTypes = nil
	file_vault_seal_multi_wrap_value_proto_depIdxs = nil
}
