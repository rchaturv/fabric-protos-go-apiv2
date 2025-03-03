// Copyright the Hyperledger Fabric contributors. All rights reserved.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: ledger/rwset/rwset.proto

package rwset

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

type TxReadWriteSet_DataModel int32

const (
	TxReadWriteSet_KV TxReadWriteSet_DataModel = 0
)

// Enum value maps for TxReadWriteSet_DataModel.
var (
	TxReadWriteSet_DataModel_name = map[int32]string{
		0: "KV",
	}
	TxReadWriteSet_DataModel_value = map[string]int32{
		"KV": 0,
	}
)

func (x TxReadWriteSet_DataModel) Enum() *TxReadWriteSet_DataModel {
	p := new(TxReadWriteSet_DataModel)
	*p = x
	return p
}

func (x TxReadWriteSet_DataModel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TxReadWriteSet_DataModel) Descriptor() protoreflect.EnumDescriptor {
	return file_ledger_rwset_rwset_proto_enumTypes[0].Descriptor()
}

func (TxReadWriteSet_DataModel) Type() protoreflect.EnumType {
	return &file_ledger_rwset_rwset_proto_enumTypes[0]
}

func (x TxReadWriteSet_DataModel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TxReadWriteSet_DataModel.Descriptor instead.
func (TxReadWriteSet_DataModel) EnumDescriptor() ([]byte, []int) {
	return file_ledger_rwset_rwset_proto_rawDescGZIP(), []int{0, 0}
}

// TxReadWriteSet encapsulates a read-write set for a transaction
// DataModel specifies the enum value of the data model
// ns_rwset field specifies a list of chaincode specific read-write set (one for each chaincode)
type TxReadWriteSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataModel TxReadWriteSet_DataModel `protobuf:"varint,1,opt,name=data_model,json=dataModel,proto3,enum=rwset.TxReadWriteSet_DataModel" json:"data_model,omitempty"`
	NsRwset   []*NsReadWriteSet        `protobuf:"bytes,2,rep,name=ns_rwset,json=nsRwset,proto3" json:"ns_rwset,omitempty"`
}

func (x *TxReadWriteSet) Reset() {
	*x = TxReadWriteSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_rwset_rwset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxReadWriteSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxReadWriteSet) ProtoMessage() {}

func (x *TxReadWriteSet) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_rwset_rwset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxReadWriteSet.ProtoReflect.Descriptor instead.
func (*TxReadWriteSet) Descriptor() ([]byte, []int) {
	return file_ledger_rwset_rwset_proto_rawDescGZIP(), []int{0}
}

func (x *TxReadWriteSet) GetDataModel() TxReadWriteSet_DataModel {
	if x != nil {
		return x.DataModel
	}
	return TxReadWriteSet_KV
}

func (x *TxReadWriteSet) GetNsRwset() []*NsReadWriteSet {
	if x != nil {
		return x.NsRwset
	}
	return nil
}

// NsReadWriteSet encapsulates the read-write set for a chaincode
type NsReadWriteSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace             string                          `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Rwset                 []byte                          `protobuf:"bytes,2,opt,name=rwset,proto3" json:"rwset,omitempty"` // Data model specific serialized proto message (e.g., kvrwset.KVRWSet for KV and Document data models)
	CollectionHashedRwset []*CollectionHashedReadWriteSet `protobuf:"bytes,3,rep,name=collection_hashed_rwset,json=collectionHashedRwset,proto3" json:"collection_hashed_rwset,omitempty"`
}

func (x *NsReadWriteSet) Reset() {
	*x = NsReadWriteSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_rwset_rwset_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NsReadWriteSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NsReadWriteSet) ProtoMessage() {}

func (x *NsReadWriteSet) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_rwset_rwset_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NsReadWriteSet.ProtoReflect.Descriptor instead.
func (*NsReadWriteSet) Descriptor() ([]byte, []int) {
	return file_ledger_rwset_rwset_proto_rawDescGZIP(), []int{1}
}

func (x *NsReadWriteSet) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *NsReadWriteSet) GetRwset() []byte {
	if x != nil {
		return x.Rwset
	}
	return nil
}

func (x *NsReadWriteSet) GetCollectionHashedRwset() []*CollectionHashedReadWriteSet {
	if x != nil {
		return x.CollectionHashedRwset
	}
	return nil
}

// CollectionHashedReadWriteSet encapsulate the hashed representation for the private read-write set for a collection
type CollectionHashedReadWriteSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CollectionName string `protobuf:"bytes,1,opt,name=collection_name,json=collectionName,proto3" json:"collection_name,omitempty"`
	HashedRwset    []byte `protobuf:"bytes,2,opt,name=hashed_rwset,json=hashedRwset,proto3" json:"hashed_rwset,omitempty"`      // Data model specific serialized proto message (e.g., kvrwset.HashedRWSet for KV and Document data models)
	PvtRwsetHash   []byte `protobuf:"bytes,3,opt,name=pvt_rwset_hash,json=pvtRwsetHash,proto3" json:"pvt_rwset_hash,omitempty"` // Hash of entire private read-write set for a specific collection. This helps in authenticating the private read-write set efficiently
}

func (x *CollectionHashedReadWriteSet) Reset() {
	*x = CollectionHashedReadWriteSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_rwset_rwset_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectionHashedReadWriteSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectionHashedReadWriteSet) ProtoMessage() {}

func (x *CollectionHashedReadWriteSet) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_rwset_rwset_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectionHashedReadWriteSet.ProtoReflect.Descriptor instead.
func (*CollectionHashedReadWriteSet) Descriptor() ([]byte, []int) {
	return file_ledger_rwset_rwset_proto_rawDescGZIP(), []int{2}
}

func (x *CollectionHashedReadWriteSet) GetCollectionName() string {
	if x != nil {
		return x.CollectionName
	}
	return ""
}

func (x *CollectionHashedReadWriteSet) GetHashedRwset() []byte {
	if x != nil {
		return x.HashedRwset
	}
	return nil
}

func (x *CollectionHashedReadWriteSet) GetPvtRwsetHash() []byte {
	if x != nil {
		return x.PvtRwsetHash
	}
	return nil
}

// TxPvtReadWriteSet encapsulate the private read-write set for a transaction
type TxPvtReadWriteSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataModel  TxReadWriteSet_DataModel `protobuf:"varint,1,opt,name=data_model,json=dataModel,proto3,enum=rwset.TxReadWriteSet_DataModel" json:"data_model,omitempty"`
	NsPvtRwset []*NsPvtReadWriteSet     `protobuf:"bytes,2,rep,name=ns_pvt_rwset,json=nsPvtRwset,proto3" json:"ns_pvt_rwset,omitempty"`
}

func (x *TxPvtReadWriteSet) Reset() {
	*x = TxPvtReadWriteSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_rwset_rwset_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TxPvtReadWriteSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TxPvtReadWriteSet) ProtoMessage() {}

func (x *TxPvtReadWriteSet) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_rwset_rwset_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TxPvtReadWriteSet.ProtoReflect.Descriptor instead.
func (*TxPvtReadWriteSet) Descriptor() ([]byte, []int) {
	return file_ledger_rwset_rwset_proto_rawDescGZIP(), []int{3}
}

func (x *TxPvtReadWriteSet) GetDataModel() TxReadWriteSet_DataModel {
	if x != nil {
		return x.DataModel
	}
	return TxReadWriteSet_KV
}

func (x *TxPvtReadWriteSet) GetNsPvtRwset() []*NsPvtReadWriteSet {
	if x != nil {
		return x.NsPvtRwset
	}
	return nil
}

// NsPvtReadWriteSet encapsulates the private read-write set for a chaincode
type NsPvtReadWriteSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace          string                       `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	CollectionPvtRwset []*CollectionPvtReadWriteSet `protobuf:"bytes,2,rep,name=collection_pvt_rwset,json=collectionPvtRwset,proto3" json:"collection_pvt_rwset,omitempty"`
}

func (x *NsPvtReadWriteSet) Reset() {
	*x = NsPvtReadWriteSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_rwset_rwset_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NsPvtReadWriteSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NsPvtReadWriteSet) ProtoMessage() {}

func (x *NsPvtReadWriteSet) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_rwset_rwset_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NsPvtReadWriteSet.ProtoReflect.Descriptor instead.
func (*NsPvtReadWriteSet) Descriptor() ([]byte, []int) {
	return file_ledger_rwset_rwset_proto_rawDescGZIP(), []int{4}
}

func (x *NsPvtReadWriteSet) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *NsPvtReadWriteSet) GetCollectionPvtRwset() []*CollectionPvtReadWriteSet {
	if x != nil {
		return x.CollectionPvtRwset
	}
	return nil
}

// CollectionPvtReadWriteSet encapsulates the private read-write set for a collection
type CollectionPvtReadWriteSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CollectionName string `protobuf:"bytes,1,opt,name=collection_name,json=collectionName,proto3" json:"collection_name,omitempty"`
	Rwset          []byte `protobuf:"bytes,2,opt,name=rwset,proto3" json:"rwset,omitempty"` // Data model specific serialized proto message (e.g., kvrwset.KVRWSet for KV and Document data models)
}

func (x *CollectionPvtReadWriteSet) Reset() {
	*x = CollectionPvtReadWriteSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ledger_rwset_rwset_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectionPvtReadWriteSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectionPvtReadWriteSet) ProtoMessage() {}

func (x *CollectionPvtReadWriteSet) ProtoReflect() protoreflect.Message {
	mi := &file_ledger_rwset_rwset_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectionPvtReadWriteSet.ProtoReflect.Descriptor instead.
func (*CollectionPvtReadWriteSet) Descriptor() ([]byte, []int) {
	return file_ledger_rwset_rwset_proto_rawDescGZIP(), []int{5}
}

func (x *CollectionPvtReadWriteSet) GetCollectionName() string {
	if x != nil {
		return x.CollectionName
	}
	return ""
}

func (x *CollectionPvtReadWriteSet) GetRwset() []byte {
	if x != nil {
		return x.Rwset
	}
	return nil
}

var File_ledger_rwset_rwset_proto protoreflect.FileDescriptor

var file_ledger_rwset_rwset_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x72, 0x77, 0x73, 0x65, 0x74, 0x2f, 0x72,
	0x77, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x72, 0x77, 0x73, 0x65,
	0x74, 0x22, 0x97, 0x01, 0x0a, 0x0e, 0x54, 0x78, 0x52, 0x65, 0x61, 0x64, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x53, 0x65, 0x74, 0x12, 0x3e, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x72, 0x77, 0x73, 0x65, 0x74,
	0x2e, 0x54, 0x78, 0x52, 0x65, 0x61, 0x64, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65, 0x74, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x12, 0x30, 0x0a, 0x08, 0x6e, 0x73, 0x5f, 0x72, 0x77, 0x73, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72, 0x77, 0x73, 0x65, 0x74, 0x2e, 0x4e,
	0x73, 0x52, 0x65, 0x61, 0x64, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65, 0x74, 0x52, 0x07, 0x6e,
	0x73, 0x52, 0x77, 0x73, 0x65, 0x74, 0x22, 0x13, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x12, 0x06, 0x0a, 0x02, 0x4b, 0x56, 0x10, 0x00, 0x22, 0xa1, 0x01, 0x0a, 0x0e,
	0x4e, 0x73, 0x52, 0x65, 0x61, 0x64, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x72, 0x77, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x72, 0x77, 0x73,
	0x65, 0x74, 0x12, 0x5b, 0x0a, 0x17, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x72, 0x77, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x72, 0x77, 0x73, 0x65, 0x74, 0x2e, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x65, 0x64, 0x52, 0x65, 0x61, 0x64,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65, 0x74, 0x52, 0x15, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x65, 0x64, 0x52, 0x77, 0x73, 0x65, 0x74, 0x22,
	0x90, 0x01, 0x0a, 0x1c, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61,
	0x73, 0x68, 0x65, 0x64, 0x52, 0x65, 0x61, 0x64, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65, 0x74,
	0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x61, 0x73,
	0x68, 0x65, 0x64, 0x5f, 0x72, 0x77, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0b, 0x68, 0x61, 0x73, 0x68, 0x65, 0x64, 0x52, 0x77, 0x73, 0x65, 0x74, 0x12, 0x24, 0x0a, 0x0e,
	0x70, 0x76, 0x74, 0x5f, 0x72, 0x77, 0x73, 0x65, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x70, 0x76, 0x74, 0x52, 0x77, 0x73, 0x65, 0x74, 0x48, 0x61,
	0x73, 0x68, 0x22, 0x8f, 0x01, 0x0a, 0x11, 0x54, 0x78, 0x50, 0x76, 0x74, 0x52, 0x65, 0x61, 0x64,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65, 0x74, 0x12, 0x3e, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x72,
	0x77, 0x73, 0x65, 0x74, 0x2e, 0x54, 0x78, 0x52, 0x65, 0x61, 0x64, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x53, 0x65, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x09, 0x64,
	0x61, 0x74, 0x61, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x3a, 0x0a, 0x0c, 0x6e, 0x73, 0x5f, 0x70,
	0x76, 0x74, 0x5f, 0x72, 0x77, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x72, 0x77, 0x73, 0x65, 0x74, 0x2e, 0x4e, 0x73, 0x50, 0x76, 0x74, 0x52, 0x65, 0x61, 0x64,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65, 0x74, 0x52, 0x0a, 0x6e, 0x73, 0x50, 0x76, 0x74, 0x52,
	0x77, 0x73, 0x65, 0x74, 0x22, 0x85, 0x01, 0x0a, 0x11, 0x4e, 0x73, 0x50, 0x76, 0x74, 0x52, 0x65,
	0x61, 0x64, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x14, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x76, 0x74, 0x5f, 0x72, 0x77, 0x73, 0x65, 0x74,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x72, 0x77, 0x73, 0x65, 0x74, 0x2e, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x76, 0x74, 0x52, 0x65, 0x61, 0x64,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65, 0x74, 0x52, 0x12, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x76, 0x74, 0x52, 0x77, 0x73, 0x65, 0x74, 0x22, 0x5a, 0x0a, 0x19,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x76, 0x74, 0x52, 0x65, 0x61,
	0x64, 0x57, 0x72, 0x69, 0x74, 0x65, 0x53, 0x65, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x77, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x72, 0x77, 0x73, 0x65, 0x74, 0x42, 0xa8, 0x01, 0x0a, 0x2a, 0x6f, 0x72, 0x67,
	0x2e, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x61, 0x62,
	0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2e, 0x72, 0x77, 0x73, 0x65, 0x74, 0x42, 0x0a, 0x52, 0x57, 0x53, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x66, 0x61,
	0x62, 0x72, 0x69, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2d, 0x67, 0x6f, 0x2d, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x72, 0x77, 0x73, 0x65,
	0x74, 0xa2, 0x02, 0x03, 0x52, 0x58, 0x58, 0xaa, 0x02, 0x05, 0x52, 0x77, 0x73, 0x65, 0x74, 0xca,
	0x02, 0x05, 0x52, 0x77, 0x73, 0x65, 0x74, 0xe2, 0x02, 0x11, 0x52, 0x77, 0x73, 0x65, 0x74, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x05, 0x52, 0x77,
	0x73, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ledger_rwset_rwset_proto_rawDescOnce sync.Once
	file_ledger_rwset_rwset_proto_rawDescData = file_ledger_rwset_rwset_proto_rawDesc
)

func file_ledger_rwset_rwset_proto_rawDescGZIP() []byte {
	file_ledger_rwset_rwset_proto_rawDescOnce.Do(func() {
		file_ledger_rwset_rwset_proto_rawDescData = protoimpl.X.CompressGZIP(file_ledger_rwset_rwset_proto_rawDescData)
	})
	return file_ledger_rwset_rwset_proto_rawDescData
}

var file_ledger_rwset_rwset_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ledger_rwset_rwset_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_ledger_rwset_rwset_proto_goTypes = []interface{}{
	(TxReadWriteSet_DataModel)(0),        // 0: rwset.TxReadWriteSet.DataModel
	(*TxReadWriteSet)(nil),               // 1: rwset.TxReadWriteSet
	(*NsReadWriteSet)(nil),               // 2: rwset.NsReadWriteSet
	(*CollectionHashedReadWriteSet)(nil), // 3: rwset.CollectionHashedReadWriteSet
	(*TxPvtReadWriteSet)(nil),            // 4: rwset.TxPvtReadWriteSet
	(*NsPvtReadWriteSet)(nil),            // 5: rwset.NsPvtReadWriteSet
	(*CollectionPvtReadWriteSet)(nil),    // 6: rwset.CollectionPvtReadWriteSet
}
var file_ledger_rwset_rwset_proto_depIdxs = []int32{
	0, // 0: rwset.TxReadWriteSet.data_model:type_name -> rwset.TxReadWriteSet.DataModel
	2, // 1: rwset.TxReadWriteSet.ns_rwset:type_name -> rwset.NsReadWriteSet
	3, // 2: rwset.NsReadWriteSet.collection_hashed_rwset:type_name -> rwset.CollectionHashedReadWriteSet
	0, // 3: rwset.TxPvtReadWriteSet.data_model:type_name -> rwset.TxReadWriteSet.DataModel
	5, // 4: rwset.TxPvtReadWriteSet.ns_pvt_rwset:type_name -> rwset.NsPvtReadWriteSet
	6, // 5: rwset.NsPvtReadWriteSet.collection_pvt_rwset:type_name -> rwset.CollectionPvtReadWriteSet
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_ledger_rwset_rwset_proto_init() }
func file_ledger_rwset_rwset_proto_init() {
	if File_ledger_rwset_rwset_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ledger_rwset_rwset_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxReadWriteSet); i {
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
		file_ledger_rwset_rwset_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NsReadWriteSet); i {
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
		file_ledger_rwset_rwset_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectionHashedReadWriteSet); i {
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
		file_ledger_rwset_rwset_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TxPvtReadWriteSet); i {
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
		file_ledger_rwset_rwset_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NsPvtReadWriteSet); i {
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
		file_ledger_rwset_rwset_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectionPvtReadWriteSet); i {
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
			RawDescriptor: file_ledger_rwset_rwset_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ledger_rwset_rwset_proto_goTypes,
		DependencyIndexes: file_ledger_rwset_rwset_proto_depIdxs,
		EnumInfos:         file_ledger_rwset_rwset_proto_enumTypes,
		MessageInfos:      file_ledger_rwset_rwset_proto_msgTypes,
	}.Build()
	File_ledger_rwset_rwset_proto = out.File
	file_ledger_rwset_rwset_proto_rawDesc = nil
	file_ledger_rwset_rwset_proto_goTypes = nil
	file_ledger_rwset_rwset_proto_depIdxs = nil
}
