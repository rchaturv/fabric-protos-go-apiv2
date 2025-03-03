// Copyright the Hyperledger Fabric contributors. All rights reserved.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: peer/resources.proto

package peer

import (
	common "github.com/hyperledger/fabric-protos-go-apiv2/common"
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

// ChaincodeIdentifier identifies a piece of chaincode.  For a peer to accept invocations of
// this chaincode, the hash of the installed code must match, as must the version string
// included with the install command.
type ChaincodeIdentifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash    []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`       // The hash of the chaincode bytes
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"` // A user friendly human readable name corresponding to the ID
}

func (x *ChaincodeIdentifier) Reset() {
	*x = ChaincodeIdentifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_resources_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChaincodeIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChaincodeIdentifier) ProtoMessage() {}

func (x *ChaincodeIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_peer_resources_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChaincodeIdentifier.ProtoReflect.Descriptor instead.
func (*ChaincodeIdentifier) Descriptor() ([]byte, []int) {
	return file_peer_resources_proto_rawDescGZIP(), []int{0}
}

func (x *ChaincodeIdentifier) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *ChaincodeIdentifier) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

// ChaincodeValidation instructs the peer how transactions for this chaincode should be
// validated.  The only validation mechanism which ships with fabric today is the standard
// 'vscc' validation mechanism.  This built in validation method utilizes an endorsement policy
// which checks that a sufficient number of signatures have been included.  The 'arguement'
// field encodes any parameters required by the validation implementation.
type ChaincodeValidation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`         // Specifies which code to run to validate transactions, defaults to 'vscc'
	Argument []byte `protobuf:"bytes,2,opt,name=argument,proto3" json:"argument,omitempty"` // When 'vscc' a marshaled VSCCArgs
}

func (x *ChaincodeValidation) Reset() {
	*x = ChaincodeValidation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_resources_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChaincodeValidation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChaincodeValidation) ProtoMessage() {}

func (x *ChaincodeValidation) ProtoReflect() protoreflect.Message {
	mi := &file_peer_resources_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChaincodeValidation.ProtoReflect.Descriptor instead.
func (*ChaincodeValidation) Descriptor() ([]byte, []int) {
	return file_peer_resources_proto_rawDescGZIP(), []int{1}
}

func (x *ChaincodeValidation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChaincodeValidation) GetArgument() []byte {
	if x != nil {
		return x.Argument
	}
	return nil
}

// VSCCArgs is passed (marshaled) as a parameter to the VSCC imlementation via the
// argument field of the ChaincodeValidation message.
type VSCCArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EndorsementPolicyRef string `protobuf:"bytes,1,opt,name=endorsement_policy_ref,json=endorsementPolicyRef,proto3" json:"endorsement_policy_ref,omitempty"` // A named reference to an endorsement policy,
}

func (x *VSCCArgs) Reset() {
	*x = VSCCArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_resources_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VSCCArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VSCCArgs) ProtoMessage() {}

func (x *VSCCArgs) ProtoReflect() protoreflect.Message {
	mi := &file_peer_resources_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VSCCArgs.ProtoReflect.Descriptor instead.
func (*VSCCArgs) Descriptor() ([]byte, []int) {
	return file_peer_resources_proto_rawDescGZIP(), []int{2}
}

func (x *VSCCArgs) GetEndorsementPolicyRef() string {
	if x != nil {
		return x.EndorsementPolicyRef
	}
	return ""
}

// ChaincodeEndorsement instructs the peer how transactions should be endorsed.  The only
// endorsement mechanism which ships with the fabric today is the standard 'escc' mechanism.
// This code simply simulates the proposal to generate a RW set, then signs the result
// using the peer's local signing identity.
type ChaincodeEndorsement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // Specifies what code to run for endorsements, defaults 'escc'
}

func (x *ChaincodeEndorsement) Reset() {
	*x = ChaincodeEndorsement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_resources_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChaincodeEndorsement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChaincodeEndorsement) ProtoMessage() {}

func (x *ChaincodeEndorsement) ProtoReflect() protoreflect.Message {
	mi := &file_peer_resources_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChaincodeEndorsement.ProtoReflect.Descriptor instead.
func (*ChaincodeEndorsement) Descriptor() ([]byte, []int) {
	return file_peer_resources_proto_rawDescGZIP(), []int{3}
}

func (x *ChaincodeEndorsement) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// ConfigTree encapsulates channel and resources configuration of a channel.
// Both configurations are represented as common.Config
type ConfigTree struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChannelConfig   *common.Config `protobuf:"bytes,1,opt,name=channel_config,json=channelConfig,proto3" json:"channel_config,omitempty"`
	ResourcesConfig *common.Config `protobuf:"bytes,2,opt,name=resources_config,json=resourcesConfig,proto3" json:"resources_config,omitempty"`
}

func (x *ConfigTree) Reset() {
	*x = ConfigTree{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_resources_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigTree) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigTree) ProtoMessage() {}

func (x *ConfigTree) ProtoReflect() protoreflect.Message {
	mi := &file_peer_resources_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigTree.ProtoReflect.Descriptor instead.
func (*ConfigTree) Descriptor() ([]byte, []int) {
	return file_peer_resources_proto_rawDescGZIP(), []int{4}
}

func (x *ConfigTree) GetChannelConfig() *common.Config {
	if x != nil {
		return x.ChannelConfig
	}
	return nil
}

func (x *ConfigTree) GetResourcesConfig() *common.Config {
	if x != nil {
		return x.ResourcesConfig
	}
	return nil
}

var File_peer_resources_proto protoreflect.FileDescriptor

var file_peer_resources_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x65, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x15,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x74, 0x78, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x13, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f,
	0x64, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x45, 0x0a, 0x13, 0x43, 0x68,
	0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x22, 0x40, 0x0a, 0x08, 0x56, 0x53, 0x43, 0x43, 0x41, 0x72, 0x67, 0x73, 0x12, 0x34, 0x0a,
	0x16, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x52, 0x65, 0x66, 0x22, 0x2a, 0x0a, 0x14, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65,
	0x45, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x7e, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x54, 0x72, 0x65, 0x65, 0x12, 0x35, 0x0a,
	0x0e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x39, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42,
	0xa0, 0x01, 0x0a, 0x22, 0x6f, 0x72, 0x67, 0x2e, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x72, 0x2e, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x70, 0x65, 0x65, 0x72, 0x42, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x2f, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2d, 0x67,
	0x6f, 0x2d, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x70, 0x65, 0x65, 0x72, 0xa2, 0x02, 0x03, 0x50,
	0x58, 0x58, 0xaa, 0x02, 0x06, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0xca, 0x02, 0x06, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0xe2, 0x02, 0x12, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x06, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_peer_resources_proto_rawDescOnce sync.Once
	file_peer_resources_proto_rawDescData = file_peer_resources_proto_rawDesc
)

func file_peer_resources_proto_rawDescGZIP() []byte {
	file_peer_resources_proto_rawDescOnce.Do(func() {
		file_peer_resources_proto_rawDescData = protoimpl.X.CompressGZIP(file_peer_resources_proto_rawDescData)
	})
	return file_peer_resources_proto_rawDescData
}

var file_peer_resources_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_peer_resources_proto_goTypes = []interface{}{
	(*ChaincodeIdentifier)(nil),  // 0: protos.ChaincodeIdentifier
	(*ChaincodeValidation)(nil),  // 1: protos.ChaincodeValidation
	(*VSCCArgs)(nil),             // 2: protos.VSCCArgs
	(*ChaincodeEndorsement)(nil), // 3: protos.ChaincodeEndorsement
	(*ConfigTree)(nil),           // 4: protos.ConfigTree
	(*common.Config)(nil),        // 5: common.Config
}
var file_peer_resources_proto_depIdxs = []int32{
	5, // 0: protos.ConfigTree.channel_config:type_name -> common.Config
	5, // 1: protos.ConfigTree.resources_config:type_name -> common.Config
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_peer_resources_proto_init() }
func file_peer_resources_proto_init() {
	if File_peer_resources_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_peer_resources_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChaincodeIdentifier); i {
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
		file_peer_resources_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChaincodeValidation); i {
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
		file_peer_resources_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VSCCArgs); i {
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
		file_peer_resources_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChaincodeEndorsement); i {
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
		file_peer_resources_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigTree); i {
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
			RawDescriptor: file_peer_resources_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_peer_resources_proto_goTypes,
		DependencyIndexes: file_peer_resources_proto_depIdxs,
		MessageInfos:      file_peer_resources_proto_msgTypes,
	}.Build()
	File_peer_resources_proto = out.File
	file_peer_resources_proto_rawDesc = nil
	file_peer_resources_proto_goTypes = nil
	file_peer_resources_proto_depIdxs = nil
}
