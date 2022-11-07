// Copyright the Hyperledger Fabric contributors. All rights reserved.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: orderer/ab.proto

package orderer

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

// If BLOCK_UNTIL_READY is specified, the reply will block until the requested blocks are available,
// if FAIL_IF_NOT_READY is specified, the reply will return an error indicating that the block is not
// found.  To request that all blocks be returned indefinitely as they are created, behavior should be
// set to BLOCK_UNTIL_READY and the stop should be set to specified with a number of MAX_UINT64
type SeekInfo_SeekBehavior int32

const (
	SeekInfo_BLOCK_UNTIL_READY SeekInfo_SeekBehavior = 0
	SeekInfo_FAIL_IF_NOT_READY SeekInfo_SeekBehavior = 1
)

// Enum value maps for SeekInfo_SeekBehavior.
var (
	SeekInfo_SeekBehavior_name = map[int32]string{
		0: "BLOCK_UNTIL_READY",
		1: "FAIL_IF_NOT_READY",
	}
	SeekInfo_SeekBehavior_value = map[string]int32{
		"BLOCK_UNTIL_READY": 0,
		"FAIL_IF_NOT_READY": 1,
	}
)

func (x SeekInfo_SeekBehavior) Enum() *SeekInfo_SeekBehavior {
	p := new(SeekInfo_SeekBehavior)
	*p = x
	return p
}

func (x SeekInfo_SeekBehavior) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SeekInfo_SeekBehavior) Descriptor() protoreflect.EnumDescriptor {
	return file_orderer_ab_proto_enumTypes[0].Descriptor()
}

func (SeekInfo_SeekBehavior) Type() protoreflect.EnumType {
	return &file_orderer_ab_proto_enumTypes[0]
}

func (x SeekInfo_SeekBehavior) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SeekInfo_SeekBehavior.Descriptor instead.
func (SeekInfo_SeekBehavior) EnumDescriptor() ([]byte, []int) {
	return file_orderer_ab_proto_rawDescGZIP(), []int{6, 0}
}

// SeekErrorTolerance indicates to the server how block provider errors should be tolerated.  By default,
// if the deliver service detects a problem in the underlying block source (typically, in the orderer,
// a consenter error), it will begin to reject deliver requests.  This is to prevent a client from waiting
// for blocks from an orderer which is stuck in an errored state.  This is almost always the desired behavior
// and clients should stick with the default STRICT checking behavior.  However, in some scenarios, particularly
// when attempting to recover from a crash or other corruption, it's desirable to force an orderer to respond
// with blocks on a best effort basis, even if the backing consensus implementation is in an errored state.
// In this case, set the SeekErrorResponse to BEST_EFFORT to ignore the consenter errors.
type SeekInfo_SeekErrorResponse int32

const (
	SeekInfo_STRICT      SeekInfo_SeekErrorResponse = 0
	SeekInfo_BEST_EFFORT SeekInfo_SeekErrorResponse = 1
)

// Enum value maps for SeekInfo_SeekErrorResponse.
var (
	SeekInfo_SeekErrorResponse_name = map[int32]string{
		0: "STRICT",
		1: "BEST_EFFORT",
	}
	SeekInfo_SeekErrorResponse_value = map[string]int32{
		"STRICT":      0,
		"BEST_EFFORT": 1,
	}
)

func (x SeekInfo_SeekErrorResponse) Enum() *SeekInfo_SeekErrorResponse {
	p := new(SeekInfo_SeekErrorResponse)
	*p = x
	return p
}

func (x SeekInfo_SeekErrorResponse) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SeekInfo_SeekErrorResponse) Descriptor() protoreflect.EnumDescriptor {
	return file_orderer_ab_proto_enumTypes[1].Descriptor()
}

func (SeekInfo_SeekErrorResponse) Type() protoreflect.EnumType {
	return &file_orderer_ab_proto_enumTypes[1]
}

func (x SeekInfo_SeekErrorResponse) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SeekInfo_SeekErrorResponse.Descriptor instead.
func (SeekInfo_SeekErrorResponse) EnumDescriptor() ([]byte, []int) {
	return file_orderer_ab_proto_rawDescGZIP(), []int{6, 1}
}

// SeekContentType indicates what type of content to deliver in response to a request. If BLOCK is specified,
// the orderer will stream blocks back to the peer. This is the default behavior. If HEADER_WITH_SIG is  specified, the
// orderer will stream only a the header and the signature, and the payload field will be set to nil. This allows
// the requester to ascertain that the respective signed block exists in the orderer (or cluster of orderers).
type SeekInfo_SeekContentType int32

const (
	SeekInfo_BLOCK           SeekInfo_SeekContentType = 0
	SeekInfo_HEADER_WITH_SIG SeekInfo_SeekContentType = 1
)

// Enum value maps for SeekInfo_SeekContentType.
var (
	SeekInfo_SeekContentType_name = map[int32]string{
		0: "BLOCK",
		1: "HEADER_WITH_SIG",
	}
	SeekInfo_SeekContentType_value = map[string]int32{
		"BLOCK":           0,
		"HEADER_WITH_SIG": 1,
	}
)

func (x SeekInfo_SeekContentType) Enum() *SeekInfo_SeekContentType {
	p := new(SeekInfo_SeekContentType)
	*p = x
	return p
}

func (x SeekInfo_SeekContentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SeekInfo_SeekContentType) Descriptor() protoreflect.EnumDescriptor {
	return file_orderer_ab_proto_enumTypes[2].Descriptor()
}

func (SeekInfo_SeekContentType) Type() protoreflect.EnumType {
	return &file_orderer_ab_proto_enumTypes[2]
}

func (x SeekInfo_SeekContentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SeekInfo_SeekContentType.Descriptor instead.
func (SeekInfo_SeekContentType) EnumDescriptor() ([]byte, []int) {
	return file_orderer_ab_proto_rawDescGZIP(), []int{6, 2}
}

type BroadcastResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Status code, which may be used to programatically respond to success/failure
	Status common.Status `protobuf:"varint,1,opt,name=status,proto3,enum=common.Status" json:"status,omitempty"`
	// Info string which may contain additional information about the status returned
	Info string `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *BroadcastResponse) Reset() {
	*x = BroadcastResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_ab_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BroadcastResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BroadcastResponse) ProtoMessage() {}

func (x *BroadcastResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_ab_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BroadcastResponse.ProtoReflect.Descriptor instead.
func (*BroadcastResponse) Descriptor() ([]byte, []int) {
	return file_orderer_ab_proto_rawDescGZIP(), []int{0}
}

func (x *BroadcastResponse) GetStatus() common.Status {
	if x != nil {
		return x.Status
	}
	return common.Status(0)
}

func (x *BroadcastResponse) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

type SeekNewest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SeekNewest) Reset() {
	*x = SeekNewest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_ab_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeekNewest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeekNewest) ProtoMessage() {}

func (x *SeekNewest) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_ab_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeekNewest.ProtoReflect.Descriptor instead.
func (*SeekNewest) Descriptor() ([]byte, []int) {
	return file_orderer_ab_proto_rawDescGZIP(), []int{1}
}

type SeekOldest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SeekOldest) Reset() {
	*x = SeekOldest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_ab_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeekOldest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeekOldest) ProtoMessage() {}

func (x *SeekOldest) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_ab_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeekOldest.ProtoReflect.Descriptor instead.
func (*SeekOldest) Descriptor() ([]byte, []int) {
	return file_orderer_ab_proto_rawDescGZIP(), []int{2}
}

type SeekSpecified struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number uint64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *SeekSpecified) Reset() {
	*x = SeekSpecified{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_ab_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeekSpecified) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeekSpecified) ProtoMessage() {}

func (x *SeekSpecified) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_ab_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeekSpecified.ProtoReflect.Descriptor instead.
func (*SeekSpecified) Descriptor() ([]byte, []int) {
	return file_orderer_ab_proto_rawDescGZIP(), []int{3}
}

func (x *SeekSpecified) GetNumber() uint64 {
	if x != nil {
		return x.Number
	}
	return 0
}

// SeekNextCommit refers to the next block that will be committed
type SeekNextCommit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SeekNextCommit) Reset() {
	*x = SeekNextCommit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_ab_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeekNextCommit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeekNextCommit) ProtoMessage() {}

func (x *SeekNextCommit) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_ab_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeekNextCommit.ProtoReflect.Descriptor instead.
func (*SeekNextCommit) Descriptor() ([]byte, []int) {
	return file_orderer_ab_proto_rawDescGZIP(), []int{4}
}

type SeekPosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//	*SeekPosition_Newest
	//	*SeekPosition_Oldest
	//	*SeekPosition_Specified
	//	*SeekPosition_NextCommit
	Type isSeekPosition_Type `protobuf_oneof:"Type"`
}

func (x *SeekPosition) Reset() {
	*x = SeekPosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_ab_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeekPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeekPosition) ProtoMessage() {}

func (x *SeekPosition) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_ab_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeekPosition.ProtoReflect.Descriptor instead.
func (*SeekPosition) Descriptor() ([]byte, []int) {
	return file_orderer_ab_proto_rawDescGZIP(), []int{5}
}

func (m *SeekPosition) GetType() isSeekPosition_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *SeekPosition) GetNewest() *SeekNewest {
	if x, ok := x.GetType().(*SeekPosition_Newest); ok {
		return x.Newest
	}
	return nil
}

func (x *SeekPosition) GetOldest() *SeekOldest {
	if x, ok := x.GetType().(*SeekPosition_Oldest); ok {
		return x.Oldest
	}
	return nil
}

func (x *SeekPosition) GetSpecified() *SeekSpecified {
	if x, ok := x.GetType().(*SeekPosition_Specified); ok {
		return x.Specified
	}
	return nil
}

func (x *SeekPosition) GetNextCommit() *SeekNextCommit {
	if x, ok := x.GetType().(*SeekPosition_NextCommit); ok {
		return x.NextCommit
	}
	return nil
}

type isSeekPosition_Type interface {
	isSeekPosition_Type()
}

type SeekPosition_Newest struct {
	Newest *SeekNewest `protobuf:"bytes,1,opt,name=newest,proto3,oneof"`
}

type SeekPosition_Oldest struct {
	Oldest *SeekOldest `protobuf:"bytes,2,opt,name=oldest,proto3,oneof"`
}

type SeekPosition_Specified struct {
	Specified *SeekSpecified `protobuf:"bytes,3,opt,name=specified,proto3,oneof"`
}

type SeekPosition_NextCommit struct {
	NextCommit *SeekNextCommit `protobuf:"bytes,4,opt,name=next_commit,json=nextCommit,proto3,oneof"`
}

func (*SeekPosition_Newest) isSeekPosition_Type() {}

func (*SeekPosition_Oldest) isSeekPosition_Type() {}

func (*SeekPosition_Specified) isSeekPosition_Type() {}

func (*SeekPosition_NextCommit) isSeekPosition_Type() {}

// SeekInfo specifies the range of requested blocks to return
// If the start position is not found, an error is immediately returned
// Otherwise, blocks are returned until a missing block is encountered, then behavior is dictated
// by the SeekBehavior specified.
type SeekInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start         *SeekPosition              `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`                                                                               // The position to start the deliver from
	Stop          *SeekPosition              `protobuf:"bytes,2,opt,name=stop,proto3" json:"stop,omitempty"`                                                                                 // The position to stop the deliver
	Behavior      SeekInfo_SeekBehavior      `protobuf:"varint,3,opt,name=behavior,proto3,enum=orderer.SeekInfo_SeekBehavior" json:"behavior,omitempty"`                                     // The behavior when a missing block is encountered
	ErrorResponse SeekInfo_SeekErrorResponse `protobuf:"varint,4,opt,name=error_response,json=errorResponse,proto3,enum=orderer.SeekInfo_SeekErrorResponse" json:"error_response,omitempty"` // How to respond to errors reported to the deliver service
	ContentType   SeekInfo_SeekContentType   `protobuf:"varint,5,opt,name=content_type,json=contentType,proto3,enum=orderer.SeekInfo_SeekContentType" json:"content_type,omitempty"`         // Defines what type of content to deliver in response to a request
}

func (x *SeekInfo) Reset() {
	*x = SeekInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_ab_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeekInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeekInfo) ProtoMessage() {}

func (x *SeekInfo) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_ab_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeekInfo.ProtoReflect.Descriptor instead.
func (*SeekInfo) Descriptor() ([]byte, []int) {
	return file_orderer_ab_proto_rawDescGZIP(), []int{6}
}

func (x *SeekInfo) GetStart() *SeekPosition {
	if x != nil {
		return x.Start
	}
	return nil
}

func (x *SeekInfo) GetStop() *SeekPosition {
	if x != nil {
		return x.Stop
	}
	return nil
}

func (x *SeekInfo) GetBehavior() SeekInfo_SeekBehavior {
	if x != nil {
		return x.Behavior
	}
	return SeekInfo_BLOCK_UNTIL_READY
}

func (x *SeekInfo) GetErrorResponse() SeekInfo_SeekErrorResponse {
	if x != nil {
		return x.ErrorResponse
	}
	return SeekInfo_STRICT
}

func (x *SeekInfo) GetContentType() SeekInfo_SeekContentType {
	if x != nil {
		return x.ContentType
	}
	return SeekInfo_BLOCK
}

type DeliverResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//	*DeliverResponse_Status
	//	*DeliverResponse_Block
	Type isDeliverResponse_Type `protobuf_oneof:"Type"`
}

func (x *DeliverResponse) Reset() {
	*x = DeliverResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_ab_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliverResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliverResponse) ProtoMessage() {}

func (x *DeliverResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_ab_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliverResponse.ProtoReflect.Descriptor instead.
func (*DeliverResponse) Descriptor() ([]byte, []int) {
	return file_orderer_ab_proto_rawDescGZIP(), []int{7}
}

func (m *DeliverResponse) GetType() isDeliverResponse_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *DeliverResponse) GetStatus() common.Status {
	if x, ok := x.GetType().(*DeliverResponse_Status); ok {
		return x.Status
	}
	return common.Status(0)
}

func (x *DeliverResponse) GetBlock() *common.Block {
	if x, ok := x.GetType().(*DeliverResponse_Block); ok {
		return x.Block
	}
	return nil
}

type isDeliverResponse_Type interface {
	isDeliverResponse_Type()
}

type DeliverResponse_Status struct {
	Status common.Status `protobuf:"varint,1,opt,name=status,proto3,enum=common.Status,oneof"`
}

type DeliverResponse_Block struct {
	Block *common.Block `protobuf:"bytes,2,opt,name=block,proto3,oneof"`
}

func (*DeliverResponse_Status) isDeliverResponse_Type() {}

func (*DeliverResponse_Block) isDeliverResponse_Type() {}

var File_orderer_ab_proto protoreflect.FileDescriptor

var file_orderer_ab_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2f, 0x61, 0x62, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x1a, 0x13, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x4f, 0x0a, 0x11, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x22, 0x0c, 0x0a, 0x0a, 0x53, 0x65, 0x65, 0x6b, 0x4e, 0x65, 0x77, 0x65, 0x73, 0x74, 0x22,
	0x0c, 0x0a, 0x0a, 0x53, 0x65, 0x65, 0x6b, 0x4f, 0x6c, 0x64, 0x65, 0x73, 0x74, 0x22, 0x27, 0x0a,
	0x0d, 0x53, 0x65, 0x65, 0x6b, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x10, 0x0a, 0x0e, 0x53, 0x65, 0x65, 0x6b, 0x4e, 0x65,
	0x78, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x22, 0xe8, 0x01, 0x0a, 0x0c, 0x53, 0x65, 0x65,
	0x6b, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x06, 0x6e, 0x65, 0x77,
	0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x65, 0x6b, 0x4e, 0x65, 0x77, 0x65, 0x73, 0x74, 0x48, 0x00,
	0x52, 0x06, 0x6e, 0x65, 0x77, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x06, 0x6f, 0x6c, 0x64, 0x65,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x65, 0x72, 0x2e, 0x53, 0x65, 0x65, 0x6b, 0x4f, 0x6c, 0x64, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52,
	0x06, 0x6f, 0x6c, 0x64, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x09, 0x73, 0x70, 0x65, 0x63, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x65, 0x6b, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x65, 0x64, 0x48, 0x00, 0x52, 0x09, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12,
	0x3a, 0x0a, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2e, 0x53,
	0x65, 0x65, 0x6b, 0x4e, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x48, 0x00, 0x52,
	0x0a, 0x6e, 0x65, 0x78, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x22, 0xd3, 0x03, 0x0a, 0x08, 0x53, 0x65, 0x65, 0x6b, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x2b, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x65, 0x6b, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x29, 0x0a,
	0x04, 0x73, 0x74, 0x6f, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x65, 0x6b, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x04, 0x73, 0x74, 0x6f, 0x70, 0x12, 0x3a, 0x0a, 0x08, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x65, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x65,
	0x65, 0x6b, 0x42, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x52, 0x08, 0x62, 0x65, 0x68, 0x61,
	0x76, 0x69, 0x6f, 0x72, 0x12, 0x4a, 0x0a, 0x0e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x65, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x53, 0x65, 0x65, 0x6b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x44, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72,
	0x2e, 0x53, 0x65, 0x65, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x65, 0x65, 0x6b, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x3c, 0x0a, 0x0c, 0x53, 0x65, 0x65, 0x6b, 0x42, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x12, 0x15, 0x0a, 0x11, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x5f,
	0x55, 0x4e, 0x54, 0x49, 0x4c, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x00, 0x12, 0x15, 0x0a,
	0x11, 0x46, 0x41, 0x49, 0x4c, 0x5f, 0x49, 0x46, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x52, 0x45, 0x41,
	0x44, 0x59, 0x10, 0x01, 0x22, 0x30, 0x0a, 0x11, 0x53, 0x65, 0x65, 0x6b, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52,
	0x49, 0x43, 0x54, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x42, 0x45, 0x53, 0x54, 0x5f, 0x45, 0x46,
	0x46, 0x4f, 0x52, 0x54, 0x10, 0x01, 0x22, 0x31, 0x0a, 0x0f, 0x53, 0x65, 0x65, 0x6b, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x4c, 0x4f,
	0x43, 0x4b, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x48, 0x45, 0x41, 0x44, 0x45, 0x52, 0x5f, 0x57,
	0x49, 0x54, 0x48, 0x5f, 0x53, 0x49, 0x47, 0x10, 0x01, 0x22, 0x6a, 0x0a, 0x0f, 0x44, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x00, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x00, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x06, 0x0a,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x32, 0x8b, 0x01, 0x0a, 0x0f, 0x41, 0x74, 0x6f, 0x6d, 0x69, 0x63,
	0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x09, 0x42, 0x72, 0x6f,
	0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x1a, 0x1a, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x65, 0x72, 0x2e, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x39, 0x0a, 0x07, 0x44, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6e, 0x76,
	0x65, 0x6c, 0x6f, 0x70, 0x65, 0x1a, 0x18, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2e,
	0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28,
	0x01, 0x30, 0x01, 0x42, 0xa3, 0x01, 0x0a, 0x25, 0x6f, 0x72, 0x67, 0x2e, 0x68, 0x79, 0x70, 0x65,
	0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x42, 0x07, 0x41,
	0x62, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x2f, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2d, 0x67,
	0x6f, 0x2d, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0xa2,
	0x02, 0x03, 0x4f, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0xca,
	0x02, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0xe2, 0x02, 0x13, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_orderer_ab_proto_rawDescOnce sync.Once
	file_orderer_ab_proto_rawDescData = file_orderer_ab_proto_rawDesc
)

func file_orderer_ab_proto_rawDescGZIP() []byte {
	file_orderer_ab_proto_rawDescOnce.Do(func() {
		file_orderer_ab_proto_rawDescData = protoimpl.X.CompressGZIP(file_orderer_ab_proto_rawDescData)
	})
	return file_orderer_ab_proto_rawDescData
}

var file_orderer_ab_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_orderer_ab_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_orderer_ab_proto_goTypes = []interface{}{
	(SeekInfo_SeekBehavior)(0),      // 0: orderer.SeekInfo.SeekBehavior
	(SeekInfo_SeekErrorResponse)(0), // 1: orderer.SeekInfo.SeekErrorResponse
	(SeekInfo_SeekContentType)(0),   // 2: orderer.SeekInfo.SeekContentType
	(*BroadcastResponse)(nil),       // 3: orderer.BroadcastResponse
	(*SeekNewest)(nil),              // 4: orderer.SeekNewest
	(*SeekOldest)(nil),              // 5: orderer.SeekOldest
	(*SeekSpecified)(nil),           // 6: orderer.SeekSpecified
	(*SeekNextCommit)(nil),          // 7: orderer.SeekNextCommit
	(*SeekPosition)(nil),            // 8: orderer.SeekPosition
	(*SeekInfo)(nil),                // 9: orderer.SeekInfo
	(*DeliverResponse)(nil),         // 10: orderer.DeliverResponse
	(common.Status)(0),              // 11: common.Status
	(*common.Block)(nil),            // 12: common.Block
	(*common.Envelope)(nil),         // 13: common.Envelope
}
var file_orderer_ab_proto_depIdxs = []int32{
	11, // 0: orderer.BroadcastResponse.status:type_name -> common.Status
	4,  // 1: orderer.SeekPosition.newest:type_name -> orderer.SeekNewest
	5,  // 2: orderer.SeekPosition.oldest:type_name -> orderer.SeekOldest
	6,  // 3: orderer.SeekPosition.specified:type_name -> orderer.SeekSpecified
	7,  // 4: orderer.SeekPosition.next_commit:type_name -> orderer.SeekNextCommit
	8,  // 5: orderer.SeekInfo.start:type_name -> orderer.SeekPosition
	8,  // 6: orderer.SeekInfo.stop:type_name -> orderer.SeekPosition
	0,  // 7: orderer.SeekInfo.behavior:type_name -> orderer.SeekInfo.SeekBehavior
	1,  // 8: orderer.SeekInfo.error_response:type_name -> orderer.SeekInfo.SeekErrorResponse
	2,  // 9: orderer.SeekInfo.content_type:type_name -> orderer.SeekInfo.SeekContentType
	11, // 10: orderer.DeliverResponse.status:type_name -> common.Status
	12, // 11: orderer.DeliverResponse.block:type_name -> common.Block
	13, // 12: orderer.AtomicBroadcast.Broadcast:input_type -> common.Envelope
	13, // 13: orderer.AtomicBroadcast.Deliver:input_type -> common.Envelope
	3,  // 14: orderer.AtomicBroadcast.Broadcast:output_type -> orderer.BroadcastResponse
	10, // 15: orderer.AtomicBroadcast.Deliver:output_type -> orderer.DeliverResponse
	14, // [14:16] is the sub-list for method output_type
	12, // [12:14] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_orderer_ab_proto_init() }
func file_orderer_ab_proto_init() {
	if File_orderer_ab_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orderer_ab_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BroadcastResponse); i {
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
		file_orderer_ab_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeekNewest); i {
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
		file_orderer_ab_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeekOldest); i {
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
		file_orderer_ab_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeekSpecified); i {
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
		file_orderer_ab_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeekNextCommit); i {
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
		file_orderer_ab_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeekPosition); i {
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
		file_orderer_ab_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeekInfo); i {
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
		file_orderer_ab_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliverResponse); i {
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
	file_orderer_ab_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*SeekPosition_Newest)(nil),
		(*SeekPosition_Oldest)(nil),
		(*SeekPosition_Specified)(nil),
		(*SeekPosition_NextCommit)(nil),
	}
	file_orderer_ab_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*DeliverResponse_Status)(nil),
		(*DeliverResponse_Block)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_orderer_ab_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orderer_ab_proto_goTypes,
		DependencyIndexes: file_orderer_ab_proto_depIdxs,
		EnumInfos:         file_orderer_ab_proto_enumTypes,
		MessageInfos:      file_orderer_ab_proto_msgTypes,
	}.Build()
	File_orderer_ab_proto = out.File
	file_orderer_ab_proto_rawDesc = nil
	file_orderer_ab_proto_goTypes = nil
	file_orderer_ab_proto_depIdxs = nil
}
