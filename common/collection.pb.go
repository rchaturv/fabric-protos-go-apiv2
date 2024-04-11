// Copyright the Hyperledger Fabric contributors. All rights reserved.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: common/collection.proto

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

// CollectionConfigPackage represents an array of CollectionConfig
// messages; the extra struct is required because repeated oneof is
// forbidden by the protobuf syntax
//
// Deprecated: Marked as deprecated in common/collection.proto.
type CollectionConfigPackage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config []*CollectionConfig `protobuf:"bytes,1,rep,name=config,proto3" json:"config,omitempty"`
}

func (x *CollectionConfigPackage) Reset() {
	*x = CollectionConfigPackage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_collection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectionConfigPackage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectionConfigPackage) ProtoMessage() {}

func (x *CollectionConfigPackage) ProtoReflect() protoreflect.Message {
	mi := &file_common_collection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectionConfigPackage.ProtoReflect.Descriptor instead.
func (*CollectionConfigPackage) Descriptor() ([]byte, []int) {
	return file_common_collection_proto_rawDescGZIP(), []int{0}
}

func (x *CollectionConfigPackage) GetConfig() []*CollectionConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

// CollectionConfig defines the configuration of a collection object;
// it currently contains a single, static type.
// Dynamic collections are deferred.
//
// Deprecated: Marked as deprecated in common/collection.proto.
type CollectionConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//
	//	*CollectionConfig_StaticCollectionConfig
	Payload isCollectionConfig_Payload `protobuf_oneof:"payload"`
}

func (x *CollectionConfig) Reset() {
	*x = CollectionConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_collection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectionConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectionConfig) ProtoMessage() {}

func (x *CollectionConfig) ProtoReflect() protoreflect.Message {
	mi := &file_common_collection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectionConfig.ProtoReflect.Descriptor instead.
func (*CollectionConfig) Descriptor() ([]byte, []int) {
	return file_common_collection_proto_rawDescGZIP(), []int{1}
}

func (m *CollectionConfig) GetPayload() isCollectionConfig_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *CollectionConfig) GetStaticCollectionConfig() *StaticCollectionConfig {
	if x, ok := x.GetPayload().(*CollectionConfig_StaticCollectionConfig); ok {
		return x.StaticCollectionConfig
	}
	return nil
}

type isCollectionConfig_Payload interface {
	isCollectionConfig_Payload()
}

type CollectionConfig_StaticCollectionConfig struct {
	StaticCollectionConfig *StaticCollectionConfig `protobuf:"bytes,1,opt,name=static_collection_config,json=staticCollectionConfig,proto3,oneof"`
}

func (*CollectionConfig_StaticCollectionConfig) isCollectionConfig_Payload() {}

// StaticCollectionConfig constitutes the configuration parameters of a
// static collection object. Static collections are collections that are
// known at chaincode instantiation time, and that cannot be changed.
// Dynamic collections are deferred.
//
// Deprecated: Marked as deprecated in common/collection.proto.
type StaticCollectionConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the name of the collection inside the denoted chaincode
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// a reference to a policy residing / managed in the config block
	// to define which orgs have access to this collection’s private data
	MemberOrgsPolicy *CollectionPolicyConfig `protobuf:"bytes,2,opt,name=member_orgs_policy,json=memberOrgsPolicy,proto3" json:"member_orgs_policy,omitempty"`
	// The minimum number of peers private data will be sent to upon
	// endorsement. The endorsement would fail if dissemination to at least
	// this number of peers is not achieved.
	RequiredPeerCount int32 `protobuf:"varint,3,opt,name=required_peer_count,json=requiredPeerCount,proto3" json:"required_peer_count,omitempty"`
	// The maximum number of peers that private data will be sent to
	// upon endorsement. This number has to be bigger than required_peer_count.
	MaximumPeerCount int32 `protobuf:"varint,4,opt,name=maximum_peer_count,json=maximumPeerCount,proto3" json:"maximum_peer_count,omitempty"`
	// The number of blocks after which the collection data expires.
	// For instance if the value is set to 10, a key last modified by block number 100
	// will be purged at block number 111. A zero value is treated same as MaxUint64
	BlockToLive uint64 `protobuf:"varint,5,opt,name=block_to_live,json=blockToLive,proto3" json:"block_to_live,omitempty"`
	// The member only read access denotes whether only collection member clients
	// can read the private data (if set to true), or even non members can
	// read the data (if set to false, for example if you want to implement more granular
	// access logic in the chaincode)
	MemberOnlyRead bool `protobuf:"varint,6,opt,name=member_only_read,json=memberOnlyRead,proto3" json:"member_only_read,omitempty"`
	// The member only write access denotes whether only collection member clients
	// can write the private data (if set to true), or even non members can
	// write the data (if set to false, for example if you want to implement more granular
	// access logic in the chaincode)
	MemberOnlyWrite bool `protobuf:"varint,7,opt,name=member_only_write,json=memberOnlyWrite,proto3" json:"member_only_write,omitempty"`
	// a reference to a policy residing / managed in the config block
	// to define the endorsement policy for this collection
	EndorsementPolicy *ApplicationPolicy `protobuf:"bytes,8,opt,name=endorsement_policy,json=endorsementPolicy,proto3" json:"endorsement_policy,omitempty"`
}

func (x *StaticCollectionConfig) Reset() {
	*x = StaticCollectionConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_collection_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaticCollectionConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaticCollectionConfig) ProtoMessage() {}

func (x *StaticCollectionConfig) ProtoReflect() protoreflect.Message {
	mi := &file_common_collection_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaticCollectionConfig.ProtoReflect.Descriptor instead.
func (*StaticCollectionConfig) Descriptor() ([]byte, []int) {
	return file_common_collection_proto_rawDescGZIP(), []int{2}
}

func (x *StaticCollectionConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StaticCollectionConfig) GetMemberOrgsPolicy() *CollectionPolicyConfig {
	if x != nil {
		return x.MemberOrgsPolicy
	}
	return nil
}

func (x *StaticCollectionConfig) GetRequiredPeerCount() int32 {
	if x != nil {
		return x.RequiredPeerCount
	}
	return 0
}

func (x *StaticCollectionConfig) GetMaximumPeerCount() int32 {
	if x != nil {
		return x.MaximumPeerCount
	}
	return 0
}

func (x *StaticCollectionConfig) GetBlockToLive() uint64 {
	if x != nil {
		return x.BlockToLive
	}
	return 0
}

func (x *StaticCollectionConfig) GetMemberOnlyRead() bool {
	if x != nil {
		return x.MemberOnlyRead
	}
	return false
}

func (x *StaticCollectionConfig) GetMemberOnlyWrite() bool {
	if x != nil {
		return x.MemberOnlyWrite
	}
	return false
}

func (x *StaticCollectionConfig) GetEndorsementPolicy() *ApplicationPolicy {
	if x != nil {
		return x.EndorsementPolicy
	}
	return nil
}

// Collection policy configuration. Initially, the configuration can only
// contain a SignaturePolicy. In the future, the SignaturePolicy may be a
// more general Policy. Instead of containing the actual policy, the
// configuration may in the future contain a string reference to a policy.
//
// Deprecated: Marked as deprecated in common/collection.proto.
type CollectionPolicyConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//
	//	*CollectionPolicyConfig_SignaturePolicy
	Payload isCollectionPolicyConfig_Payload `protobuf_oneof:"payload"`
}

func (x *CollectionPolicyConfig) Reset() {
	*x = CollectionPolicyConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_collection_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectionPolicyConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectionPolicyConfig) ProtoMessage() {}

func (x *CollectionPolicyConfig) ProtoReflect() protoreflect.Message {
	mi := &file_common_collection_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectionPolicyConfig.ProtoReflect.Descriptor instead.
func (*CollectionPolicyConfig) Descriptor() ([]byte, []int) {
	return file_common_collection_proto_rawDescGZIP(), []int{3}
}

func (m *CollectionPolicyConfig) GetPayload() isCollectionPolicyConfig_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *CollectionPolicyConfig) GetSignaturePolicy() *SignaturePolicyEnvelope {
	if x, ok := x.GetPayload().(*CollectionPolicyConfig_SignaturePolicy); ok {
		return x.SignaturePolicy
	}
	return nil
}

type isCollectionPolicyConfig_Payload interface {
	isCollectionPolicyConfig_Payload()
}

type CollectionPolicyConfig_SignaturePolicy struct {
	// Initially, only a signature policy is supported.
	SignaturePolicy *SignaturePolicyEnvelope `protobuf:"bytes,1,opt,name=signature_policy,json=signaturePolicy,proto3,oneof"`
}

func (*CollectionPolicyConfig_SignaturePolicy) isCollectionPolicyConfig_Payload() {}

var File_common_collection_proto protoreflect.FileDescriptor

var file_common_collection_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x1a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x17, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x02, 0x18, 0x01, 0x22, 0x7d, 0x0a, 0x10, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x5a, 0x0a,
	0x18, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48,
	0x00, 0x52, 0x16, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x02, 0x18, 0x01, 0x42, 0x09, 0x0a,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xa0, 0x03, 0x0a, 0x16, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x63, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4c, 0x0a, 0x12, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x5f, 0x6f, 0x72, 0x67, 0x73, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x10, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x72, 0x67, 0x73, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x2e, 0x0a, 0x13, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x11, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x50, 0x65, 0x65, 0x72,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d,
	0x5f, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x10, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x50, 0x65, 0x65, 0x72, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x6f, 0x5f,
	0x6c, 0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x54, 0x6f, 0x4c, 0x69, 0x76, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x6e, 0x6c, 0x79, 0x52, 0x65, 0x61,
	0x64, 0x12, 0x2a, 0x0a, 0x11, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x6f, 0x6e, 0x6c, 0x79,
	0x5f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x4f, 0x6e, 0x6c, 0x79, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x48, 0x0a,
	0x12, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x52, 0x11, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x3a, 0x02, 0x18, 0x01, 0x22, 0x75, 0x0a, 0x16, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4c, 0x0a, 0x10, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65,
	0x48, 0x00, 0x52, 0x0f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x3a, 0x02, 0x18, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x42, 0xa5, 0x01, 0x0a, 0x24, 0x6f, 0x72, 0x67, 0x2e, 0x68, 0x79, 0x70, 0x65, 0x72,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x0f, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x34,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2d, 0x67, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0xca, 0x02, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xe2, 0x02, 0x12, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_common_collection_proto_rawDescOnce sync.Once
	file_common_collection_proto_rawDescData = file_common_collection_proto_rawDesc
)

func file_common_collection_proto_rawDescGZIP() []byte {
	file_common_collection_proto_rawDescOnce.Do(func() {
		file_common_collection_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_collection_proto_rawDescData)
	})
	return file_common_collection_proto_rawDescData
}

var file_common_collection_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_common_collection_proto_goTypes = []interface{}{
	(*CollectionConfigPackage)(nil), // 0: common.CollectionConfigPackage
	(*CollectionConfig)(nil),        // 1: common.CollectionConfig
	(*StaticCollectionConfig)(nil),  // 2: common.StaticCollectionConfig
	(*CollectionPolicyConfig)(nil),  // 3: common.CollectionPolicyConfig
	(*ApplicationPolicy)(nil),       // 4: common.ApplicationPolicy
	(*SignaturePolicyEnvelope)(nil), // 5: common.SignaturePolicyEnvelope
}
var file_common_collection_proto_depIdxs = []int32{
	1, // 0: common.CollectionConfigPackage.config:type_name -> common.CollectionConfig
	2, // 1: common.CollectionConfig.static_collection_config:type_name -> common.StaticCollectionConfig
	3, // 2: common.StaticCollectionConfig.member_orgs_policy:type_name -> common.CollectionPolicyConfig
	4, // 3: common.StaticCollectionConfig.endorsement_policy:type_name -> common.ApplicationPolicy
	5, // 4: common.CollectionPolicyConfig.signature_policy:type_name -> common.SignaturePolicyEnvelope
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_common_collection_proto_init() }
func file_common_collection_proto_init() {
	if File_common_collection_proto != nil {
		return
	}
	file_common_policies_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_common_collection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectionConfigPackage); i {
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
		file_common_collection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectionConfig); i {
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
		file_common_collection_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaticCollectionConfig); i {
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
		file_common_collection_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectionPolicyConfig); i {
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
	file_common_collection_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*CollectionConfig_StaticCollectionConfig)(nil),
	}
	file_common_collection_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*CollectionPolicyConfig_SignaturePolicy)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_collection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_collection_proto_goTypes,
		DependencyIndexes: file_common_collection_proto_depIdxs,
		MessageInfos:      file_common_collection_proto_msgTypes,
	}.Build()
	File_common_collection_proto = out.File
	file_common_collection_proto_rawDesc = nil
	file_common_collection_proto_goTypes = nil
	file_common_collection_proto_depIdxs = nil
}
