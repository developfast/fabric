// Copyright the Hyperledger Fabric contributors. All rights reserved.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: common/ledger.proto

package common

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

// Contains information about the blockchain ledger such as height, current
// block hash, and previous block hash.
type BlockchainInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Height            uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	CurrentBlockHash  []byte `protobuf:"bytes,2,opt,name=currentBlockHash,proto3" json:"currentBlockHash,omitempty"`
	PreviousBlockHash []byte `protobuf:"bytes,3,opt,name=previousBlockHash,proto3" json:"previousBlockHash,omitempty"`
	// Specifies bootstrapping snapshot info if the channel is bootstrapped from a snapshot.
	// It is nil if the channel is not bootstrapped from a snapshot.
	BootstrappingSnapshotInfo *BootstrappingSnapshotInfo `protobuf:"bytes,4,opt,name=bootstrappingSnapshotInfo,proto3" json:"bootstrappingSnapshotInfo,omitempty"`
}

func (x *BlockchainInfo) Reset() {
	*x = BlockchainInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_ledger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockchainInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockchainInfo) ProtoMessage() {}

func (x *BlockchainInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_ledger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockchainInfo.ProtoReflect.Descriptor instead.
func (*BlockchainInfo) Descriptor() ([]byte, []int) {
	return file_common_ledger_proto_rawDescGZIP(), []int{0}
}

func (x *BlockchainInfo) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *BlockchainInfo) GetCurrentBlockHash() []byte {
	if x != nil {
		return x.CurrentBlockHash
	}
	return nil
}

func (x *BlockchainInfo) GetPreviousBlockHash() []byte {
	if x != nil {
		return x.PreviousBlockHash
	}
	return nil
}

func (x *BlockchainInfo) GetBootstrappingSnapshotInfo() *BootstrappingSnapshotInfo {
	if x != nil {
		return x.BootstrappingSnapshotInfo
	}
	return nil
}

// Contains information for the bootstrapping snapshot.
type BootstrappingSnapshotInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastBlockInSnapshot uint64 `protobuf:"varint,1,opt,name=lastBlockInSnapshot,proto3" json:"lastBlockInSnapshot,omitempty"`
}

func (x *BootstrappingSnapshotInfo) Reset() {
	*x = BootstrappingSnapshotInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_ledger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BootstrappingSnapshotInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BootstrappingSnapshotInfo) ProtoMessage() {}

func (x *BootstrappingSnapshotInfo) ProtoReflect() protoreflect.Message {
	mi := &file_common_ledger_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BootstrappingSnapshotInfo.ProtoReflect.Descriptor instead.
func (*BootstrappingSnapshotInfo) Descriptor() ([]byte, []int) {
	return file_common_ledger_proto_rawDescGZIP(), []int{1}
}

func (x *BootstrappingSnapshotInfo) GetLastBlockInSnapshot() uint64 {
	if x != nil {
		return x.LastBlockInSnapshot
	}
	return 0
}

var File_common_ledger_proto protoreflect.FileDescriptor

var file_common_ledger_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0xe3, 0x01,
	0x0a, 0x0e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x10, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x11, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x5f, 0x0a, 0x19, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42,
	0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x19, 0x62, 0x6f, 0x6f, 0x74, 0x73, 0x74,
	0x72, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x4d, 0x0a, 0x19, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x30, 0x0a, 0x13, 0x6c, 0x61, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x53,
	0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x13, 0x6c,
	0x61, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68,
	0x6f, 0x74, 0x42, 0xa1, 0x01, 0x0a, 0x24, 0x6f, 0x72, 0x67, 0x2e, 0x68, 0x79, 0x70, 0x65, 0x72,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x0b, 0x4c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x2f, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2d, 0x67, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xca,
	0x02, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xe2, 0x02, 0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x06,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_ledger_proto_rawDescOnce sync.Once
	file_common_ledger_proto_rawDescData = file_common_ledger_proto_rawDesc
)

func file_common_ledger_proto_rawDescGZIP() []byte {
	file_common_ledger_proto_rawDescOnce.Do(func() {
		file_common_ledger_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_ledger_proto_rawDescData)
	})
	return file_common_ledger_proto_rawDescData
}

var file_common_ledger_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_ledger_proto_goTypes = []any{
	(*BlockchainInfo)(nil),            // 0: common.BlockchainInfo
	(*BootstrappingSnapshotInfo)(nil), // 1: common.BootstrappingSnapshotInfo
}
var file_common_ledger_proto_depIdxs = []int32{
	1, // 0: common.BlockchainInfo.bootstrappingSnapshotInfo:type_name -> common.BootstrappingSnapshotInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_common_ledger_proto_init() }
func file_common_ledger_proto_init() {
	if File_common_ledger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_ledger_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*BlockchainInfo); i {
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
		file_common_ledger_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*BootstrappingSnapshotInfo); i {
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
			RawDescriptor: file_common_ledger_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_ledger_proto_goTypes,
		DependencyIndexes: file_common_ledger_proto_depIdxs,
		MessageInfos:      file_common_ledger_proto_msgTypes,
	}.Build()
	File_common_ledger_proto = out.File
	file_common_ledger_proto_rawDesc = nil
	file_common_ledger_proto_goTypes = nil
	file_common_ledger_proto_depIdxs = nil
}
