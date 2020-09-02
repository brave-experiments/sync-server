// Copyright 2020 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// If you change or add any fields in this file, update proto_visitors.h and
// potentially proto_enum_conversions.{h, cc}.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.2
// source: sync_invalidations_payload.proto

package sync_pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// The payload coming from the server for all invalidated data types.
type SyncInvalidationsPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This field contains information about each data type which needs to be
	// updated.
	DataTypeInvalidations []*SyncInvalidationsPayload_DataTypeInvalidation `protobuf:"bytes,1,rep,name=data_type_invalidations,json=dataTypeInvalidations" json:"data_type_invalidations,omitempty"`
	// Opaque field, which has to be provided as part of resulting GetUpdates
	// back to the server.
	Hint []byte `protobuf:"bytes,2,opt,name=hint" json:"hint,omitempty"`
}

func (x *SyncInvalidationsPayload) Reset() {
	*x = SyncInvalidationsPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_invalidations_payload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncInvalidationsPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncInvalidationsPayload) ProtoMessage() {}

func (x *SyncInvalidationsPayload) ProtoReflect() protoreflect.Message {
	mi := &file_sync_invalidations_payload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncInvalidationsPayload.ProtoReflect.Descriptor instead.
func (*SyncInvalidationsPayload) Descriptor() ([]byte, []int) {
	return file_sync_invalidations_payload_proto_rawDescGZIP(), []int{0}
}

func (x *SyncInvalidationsPayload) GetDataTypeInvalidations() []*SyncInvalidationsPayload_DataTypeInvalidation {
	if x != nil {
		return x.DataTypeInvalidations
	}
	return nil
}

func (x *SyncInvalidationsPayload) GetHint() []byte {
	if x != nil {
		return x.Hint
	}
	return nil
}

type SyncInvalidationsPayload_DataTypeInvalidation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The legitimate values of this field correspond to the protobuf field
	// numbers of all EntitySpecifics fields supported by the server (see
	// components/sync/protocol/sync.proto).
	DataTypeId *int32 `protobuf:"varint,1,opt,name=data_type_id,json=dataTypeId" json:"data_type_id,omitempty"`
}

func (x *SyncInvalidationsPayload_DataTypeInvalidation) Reset() {
	*x = SyncInvalidationsPayload_DataTypeInvalidation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_invalidations_payload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncInvalidationsPayload_DataTypeInvalidation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncInvalidationsPayload_DataTypeInvalidation) ProtoMessage() {}

func (x *SyncInvalidationsPayload_DataTypeInvalidation) ProtoReflect() protoreflect.Message {
	mi := &file_sync_invalidations_payload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncInvalidationsPayload_DataTypeInvalidation.ProtoReflect.Descriptor instead.
func (*SyncInvalidationsPayload_DataTypeInvalidation) Descriptor() ([]byte, []int) {
	return file_sync_invalidations_payload_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SyncInvalidationsPayload_DataTypeInvalidation) GetDataTypeId() int32 {
	if x != nil && x.DataTypeId != nil {
		return *x.DataTypeId
	}
	return 0
}

var File_sync_invalidations_payload_proto protoreflect.FileDescriptor

var file_sync_invalidations_payload_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x22, 0xd8, 0x01, 0x0a, 0x18,
	0x53, 0x79, 0x6e, 0x63, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x6e, 0x0a, 0x17, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x73, 0x79, 0x6e, 0x63,
	0x5f, 0x70, 0x62, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x15, 0x64, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x69, 0x6e, 0x74, 0x1a, 0x38, 0x0a, 0x14,
	0x44, 0x61, 0x74, 0x61, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x42, 0x2b, 0x0a, 0x25, 0x6f, 0x72, 0x67, 0x2e, 0x63, 0x68,
	0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x48,
	0x03, 0x50, 0x01,
}

var (
	file_sync_invalidations_payload_proto_rawDescOnce sync.Once
	file_sync_invalidations_payload_proto_rawDescData = file_sync_invalidations_payload_proto_rawDesc
)

func file_sync_invalidations_payload_proto_rawDescGZIP() []byte {
	file_sync_invalidations_payload_proto_rawDescOnce.Do(func() {
		file_sync_invalidations_payload_proto_rawDescData = protoimpl.X.CompressGZIP(file_sync_invalidations_payload_proto_rawDescData)
	})
	return file_sync_invalidations_payload_proto_rawDescData
}

var file_sync_invalidations_payload_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sync_invalidations_payload_proto_goTypes = []interface{}{
	(*SyncInvalidationsPayload)(nil),                      // 0: sync_pb.SyncInvalidationsPayload
	(*SyncInvalidationsPayload_DataTypeInvalidation)(nil), // 1: sync_pb.SyncInvalidationsPayload.DataTypeInvalidation
}
var file_sync_invalidations_payload_proto_depIdxs = []int32{
	1, // 0: sync_pb.SyncInvalidationsPayload.data_type_invalidations:type_name -> sync_pb.SyncInvalidationsPayload.DataTypeInvalidation
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sync_invalidations_payload_proto_init() }
func file_sync_invalidations_payload_proto_init() {
	if File_sync_invalidations_payload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sync_invalidations_payload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncInvalidationsPayload); i {
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
		file_sync_invalidations_payload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncInvalidationsPayload_DataTypeInvalidation); i {
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
			RawDescriptor: file_sync_invalidations_payload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sync_invalidations_payload_proto_goTypes,
		DependencyIndexes: file_sync_invalidations_payload_proto_depIdxs,
		MessageInfos:      file_sync_invalidations_payload_proto_msgTypes,
	}.Build()
	File_sync_invalidations_payload_proto = out.File
	file_sync_invalidations_payload_proto_rawDesc = nil
	file_sync_invalidations_payload_proto_goTypes = nil
	file_sync_invalidations_payload_proto_depIdxs = nil
}
