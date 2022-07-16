// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.5.1
// source: props/props.proto

package props

import (
	context "context"
	base "github.com/mc_nft/kitex_gen/base"
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

type PropsInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`                         // 名称
	ImageUrl string `protobuf:"bytes,2,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"` // 图片url
	Desc     string `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`                         // 描述
}

func (x *PropsInfo) Reset() {
	*x = PropsInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_props_props_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropsInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropsInfo) ProtoMessage() {}

func (x *PropsInfo) ProtoReflect() protoreflect.Message {
	mi := &file_props_props_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropsInfo.ProtoReflect.Descriptor instead.
func (*PropsInfo) Descriptor() ([]byte, []int) {
	return file_props_props_proto_rawDescGZIP(), []int{0}
}

func (x *PropsInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PropsInfo) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *PropsInfo) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

type AddPropsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropsList []*PropsInfo `protobuf:"bytes,1,rep,name=props_list,json=propsList,proto3" json:"props_list,omitempty"`
	Base      *base.Base   `protobuf:"bytes,255,opt,name=Base,proto3" json:"Base,omitempty"`
}

func (x *AddPropsRequest) Reset() {
	*x = AddPropsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_props_props_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPropsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPropsRequest) ProtoMessage() {}

func (x *AddPropsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_props_props_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPropsRequest.ProtoReflect.Descriptor instead.
func (*AddPropsRequest) Descriptor() ([]byte, []int) {
	return file_props_props_proto_rawDescGZIP(), []int{1}
}

func (x *AddPropsRequest) GetPropsList() []*PropsInfo {
	if x != nil {
		return x.PropsList
	}
	return nil
}

func (x *AddPropsRequest) GetBase() *base.Base {
	if x != nil {
		return x.Base
	}
	return nil
}

type AddPropsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *base.BaseResp `protobuf:"bytes,255,opt,name=BaseResp,proto3" json:"BaseResp,omitempty"`
}

func (x *AddPropsResponse) Reset() {
	*x = AddPropsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_props_props_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPropsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPropsResponse) ProtoMessage() {}

func (x *AddPropsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_props_props_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPropsResponse.ProtoReflect.Descriptor instead.
func (*AddPropsResponse) Descriptor() ([]byte, []int) {
	return file_props_props_proto_rawDescGZIP(), []int{2}
}

func (x *AddPropsResponse) GetBaseResp() *base.BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

type UpdatePropsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropsId   int32      `protobuf:"varint,1,opt,name=props_id,json=propsId,proto3" json:"props_id,omitempty"`
	PropsInfo *PropsInfo `protobuf:"bytes,2,opt,name=props_info,json=propsInfo,proto3" json:"props_info,omitempty"`
	Base      *base.Base `protobuf:"bytes,255,opt,name=Base,proto3" json:"Base,omitempty"`
}

func (x *UpdatePropsRequest) Reset() {
	*x = UpdatePropsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_props_props_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePropsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePropsRequest) ProtoMessage() {}

func (x *UpdatePropsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_props_props_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePropsRequest.ProtoReflect.Descriptor instead.
func (*UpdatePropsRequest) Descriptor() ([]byte, []int) {
	return file_props_props_proto_rawDescGZIP(), []int{3}
}

func (x *UpdatePropsRequest) GetPropsId() int32 {
	if x != nil {
		return x.PropsId
	}
	return 0
}

func (x *UpdatePropsRequest) GetPropsInfo() *PropsInfo {
	if x != nil {
		return x.PropsInfo
	}
	return nil
}

func (x *UpdatePropsRequest) GetBase() *base.Base {
	if x != nil {
		return x.Base
	}
	return nil
}

type UpdatePropsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *base.BaseResp `protobuf:"bytes,255,opt,name=BaseResp,proto3" json:"BaseResp,omitempty"`
}

func (x *UpdatePropsResponse) Reset() {
	*x = UpdatePropsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_props_props_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePropsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePropsResponse) ProtoMessage() {}

func (x *UpdatePropsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_props_props_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePropsResponse.ProtoReflect.Descriptor instead.
func (*UpdatePropsResponse) Descriptor() ([]byte, []int) {
	return file_props_props_proto_rawDescGZIP(), []int{4}
}

func (x *UpdatePropsResponse) GetBaseResp() *base.BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

// 用户：
// 发放、展示、使用 props
type DistributePropsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropsItems []*DistributePropsRequest_PropsItem `protobuf:"bytes,1,rep,name=props_items,json=propsItems,proto3" json:"props_items,omitempty"`
	Uid        int64                               `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	SeqId      string                              `protobuf:"bytes,3,opt,name=seq_id,json=seqId,proto3" json:"seq_id,omitempty"`
	Base       *base.Base                          `protobuf:"bytes,255,opt,name=Base,proto3" json:"Base,omitempty"`
}

func (x *DistributePropsRequest) Reset() {
	*x = DistributePropsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_props_props_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistributePropsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistributePropsRequest) ProtoMessage() {}

func (x *DistributePropsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_props_props_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistributePropsRequest.ProtoReflect.Descriptor instead.
func (*DistributePropsRequest) Descriptor() ([]byte, []int) {
	return file_props_props_proto_rawDescGZIP(), []int{5}
}

func (x *DistributePropsRequest) GetPropsItems() []*DistributePropsRequest_PropsItem {
	if x != nil {
		return x.PropsItems
	}
	return nil
}

func (x *DistributePropsRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *DistributePropsRequest) GetSeqId() string {
	if x != nil {
		return x.SeqId
	}
	return ""
}

func (x *DistributePropsRequest) GetBase() *base.Base {
	if x != nil {
		return x.Base
	}
	return nil
}

type DistributePropsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *base.BaseResp `protobuf:"bytes,255,opt,name=BaseResp,proto3" json:"BaseResp,omitempty"`
}

func (x *DistributePropsResponse) Reset() {
	*x = DistributePropsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_props_props_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistributePropsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistributePropsResponse) ProtoMessage() {}

func (x *DistributePropsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_props_props_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistributePropsResponse.ProtoReflect.Descriptor instead.
func (*DistributePropsResponse) Descriptor() ([]byte, []int) {
	return file_props_props_proto_rawDescGZIP(), []int{6}
}

func (x *DistributePropsResponse) GetBaseResp() *base.BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

// 查询道具
type QueryPropsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      int64      `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	PropsIds []int32    `protobuf:"varint,2,rep,packed,name=props_ids,json=propsIds,proto3" json:"props_ids,omitempty"`
	Limit    int32      `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset   int32      `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
	Base     *base.Base `protobuf:"bytes,255,opt,name=Base,proto3" json:"Base,omitempty"`
}

func (x *QueryPropsRequest) Reset() {
	*x = QueryPropsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_props_props_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPropsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPropsRequest) ProtoMessage() {}

func (x *QueryPropsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_props_props_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPropsRequest.ProtoReflect.Descriptor instead.
func (*QueryPropsRequest) Descriptor() ([]byte, []int) {
	return file_props_props_proto_rawDescGZIP(), []int{7}
}

func (x *QueryPropsRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *QueryPropsRequest) GetPropsIds() []int32 {
	if x != nil {
		return x.PropsIds
	}
	return nil
}

func (x *QueryPropsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *QueryPropsRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *QueryPropsRequest) GetBase() *base.Base {
	if x != nil {
		return x.Base
	}
	return nil
}

type QueryPropsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropsList []*PropsInfo   `protobuf:"bytes,1,rep,name=props_list,json=propsList,proto3" json:"props_list,omitempty"`
	PropsCnt  int32          `protobuf:"varint,2,opt,name=props_cnt,json=propsCnt,proto3" json:"props_cnt,omitempty"`
	BaseResp  *base.BaseResp `protobuf:"bytes,255,opt,name=BaseResp,proto3" json:"BaseResp,omitempty"`
}

func (x *QueryPropsResponse) Reset() {
	*x = QueryPropsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_props_props_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryPropsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryPropsResponse) ProtoMessage() {}

func (x *QueryPropsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_props_props_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryPropsResponse.ProtoReflect.Descriptor instead.
func (*QueryPropsResponse) Descriptor() ([]byte, []int) {
	return file_props_props_proto_rawDescGZIP(), []int{8}
}

func (x *QueryPropsResponse) GetPropsList() []*PropsInfo {
	if x != nil {
		return x.PropsList
	}
	return nil
}

func (x *QueryPropsResponse) GetPropsCnt() int32 {
	if x != nil {
		return x.PropsCnt
	}
	return 0
}

func (x *QueryPropsResponse) GetBaseResp() *base.BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

// 核销道具
type ConsumePropsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropsItems []*ConsumePropsRequest_PropsItem `protobuf:"bytes,1,rep,name=props_items,json=propsItems,proto3" json:"props_items,omitempty"`
	Uid        int64                            `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	RecordStr  string                           `protobuf:"bytes,3,opt,name=record_str,json=recordStr,proto3" json:"record_str,omitempty"` // 如订单id、订单号等
	Base       *base.Base                       `protobuf:"bytes,255,opt,name=Base,proto3" json:"Base,omitempty"`
}

func (x *ConsumePropsRequest) Reset() {
	*x = ConsumePropsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_props_props_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumePropsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumePropsRequest) ProtoMessage() {}

func (x *ConsumePropsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_props_props_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumePropsRequest.ProtoReflect.Descriptor instead.
func (*ConsumePropsRequest) Descriptor() ([]byte, []int) {
	return file_props_props_proto_rawDescGZIP(), []int{9}
}

func (x *ConsumePropsRequest) GetPropsItems() []*ConsumePropsRequest_PropsItem {
	if x != nil {
		return x.PropsItems
	}
	return nil
}

func (x *ConsumePropsRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ConsumePropsRequest) GetRecordStr() string {
	if x != nil {
		return x.RecordStr
	}
	return ""
}

func (x *ConsumePropsRequest) GetBase() *base.Base {
	if x != nil {
		return x.Base
	}
	return nil
}

type ConsumePropsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *base.BaseResp `protobuf:"bytes,255,opt,name=BaseResp,proto3" json:"BaseResp,omitempty"`
}

func (x *ConsumePropsResponse) Reset() {
	*x = ConsumePropsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_props_props_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumePropsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumePropsResponse) ProtoMessage() {}

func (x *ConsumePropsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_props_props_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumePropsResponse.ProtoReflect.Descriptor instead.
func (*ConsumePropsResponse) Descriptor() ([]byte, []int) {
	return file_props_props_proto_rawDescGZIP(), []int{10}
}

func (x *ConsumePropsResponse) GetBaseResp() *base.BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

type DistributePropsRequest_PropsItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropsId int32 `protobuf:"varint,1,opt,name=props_id,json=propsId,proto3" json:"props_id,omitempty"`
	Amount  int32 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *DistributePropsRequest_PropsItem) Reset() {
	*x = DistributePropsRequest_PropsItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_props_props_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistributePropsRequest_PropsItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistributePropsRequest_PropsItem) ProtoMessage() {}

func (x *DistributePropsRequest_PropsItem) ProtoReflect() protoreflect.Message {
	mi := &file_props_props_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistributePropsRequest_PropsItem.ProtoReflect.Descriptor instead.
func (*DistributePropsRequest_PropsItem) Descriptor() ([]byte, []int) {
	return file_props_props_proto_rawDescGZIP(), []int{5, 0}
}

func (x *DistributePropsRequest_PropsItem) GetPropsId() int32 {
	if x != nil {
		return x.PropsId
	}
	return 0
}

func (x *DistributePropsRequest_PropsItem) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type ConsumePropsRequest_PropsItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropsId int32 `protobuf:"varint,1,opt,name=props_id,json=propsId,proto3" json:"props_id,omitempty"`
	Amount  int32 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *ConsumePropsRequest_PropsItem) Reset() {
	*x = ConsumePropsRequest_PropsItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_props_props_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumePropsRequest_PropsItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumePropsRequest_PropsItem) ProtoMessage() {}

func (x *ConsumePropsRequest_PropsItem) ProtoReflect() protoreflect.Message {
	mi := &file_props_props_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumePropsRequest_PropsItem.ProtoReflect.Descriptor instead.
func (*ConsumePropsRequest_PropsItem) Descriptor() ([]byte, []int) {
	return file_props_props_proto_rawDescGZIP(), []int{9, 0}
}

func (x *ConsumePropsRequest_PropsItem) GetPropsId() int32 {
	if x != nil {
		return x.PropsId
	}
	return 0
}

func (x *ConsumePropsRequest_PropsItem) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_props_props_proto protoreflect.FileDescriptor

var file_props_props_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x1a, 0x0f, 0x62, 0x61, 0x73, 0x65,
	0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x09, 0x50,
	0x72, 0x6f, 0x70, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73,
	0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x22, 0x63, 0x0a,
	0x0f, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2f, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x2e, 0x50, 0x72, 0x6f,
	0x70, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x1f, 0x0a, 0x04, 0x42, 0x61, 0x73, 0x65, 0x18, 0xff, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x04, 0x42, 0x61,
	0x73, 0x65, 0x22, 0x3f, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x18, 0xff, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08, 0x42, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x81, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x70, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x72,
	0x6f, 0x70, 0x73, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x70,
	0x73, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x70, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x04, 0x42, 0x61, 0x73, 0x65, 0x18, 0xff,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x42, 0x61, 0x73,
	0x65, 0x52, 0x04, 0x42, 0x61, 0x73, 0x65, 0x22, 0x42, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b,
	0x0a, 0x08, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x18, 0xff, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x52, 0x08, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0xec, 0x01, 0x0a, 0x16,
	0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x5f,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x72,
	0x6f, 0x70, 0x73, 0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x50, 0x72,
	0x6f, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x73,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x73, 0x65, 0x71, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x65, 0x71, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x04, 0x42, 0x61, 0x73,
	0x65, 0x18, 0xff, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x04, 0x42, 0x61, 0x73, 0x65, 0x1a, 0x3e, 0x0a, 0x09, 0x50, 0x72,
	0x6f, 0x70, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x73,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x70, 0x73,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x46, 0x0a, 0x17, 0x44, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x18, 0xff, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x91, 0x01, 0x0a, 0x11, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x70,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x72,
	0x6f, 0x70, 0x73, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x70, 0x73, 0x49, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x42, 0x61, 0x73, 0x65, 0x18, 0xff, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65,
	0x52, 0x04, 0x42, 0x61, 0x73, 0x65, 0x22, 0x8f, 0x01, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x50, 0x72, 0x6f, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x5f, 0x63, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x43, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x08, 0x42,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x18, 0xff, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0xee, 0x01, 0x0a, 0x13, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x45, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x2e, 0x43, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0a, 0x70, 0x72, 0x6f,
	0x70, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x53, 0x74, 0x72, 0x12, 0x1f, 0x0a, 0x04, 0x42, 0x61, 0x73, 0x65,
	0x18, 0xff, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x42,
	0x61, 0x73, 0x65, 0x52, 0x04, 0x42, 0x61, 0x73, 0x65, 0x1a, 0x3e, 0x0a, 0x09, 0x50, 0x72, 0x6f,
	0x70, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x43, 0x0a, 0x14, 0x43, 0x6f, 0x6e,
	0x73, 0x75, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2b, 0x0a, 0x08, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x18, 0xff, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x52, 0x08, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x42, 0x23,
	0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x63, 0x5f,
	0x6e, 0x66, 0x74, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x70, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_props_props_proto_rawDescOnce sync.Once
	file_props_props_proto_rawDescData = file_props_props_proto_rawDesc
)

func file_props_props_proto_rawDescGZIP() []byte {
	file_props_props_proto_rawDescOnce.Do(func() {
		file_props_props_proto_rawDescData = protoimpl.X.CompressGZIP(file_props_props_proto_rawDescData)
	})
	return file_props_props_proto_rawDescData
}

var file_props_props_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_props_props_proto_goTypes = []interface{}{
	(*PropsInfo)(nil),                        // 0: props.PropsInfo
	(*AddPropsRequest)(nil),                  // 1: props.AddPropsRequest
	(*AddPropsResponse)(nil),                 // 2: props.AddPropsResponse
	(*UpdatePropsRequest)(nil),               // 3: props.UpdatePropsRequest
	(*UpdatePropsResponse)(nil),              // 4: props.UpdatePropsResponse
	(*DistributePropsRequest)(nil),           // 5: props.DistributePropsRequest
	(*DistributePropsResponse)(nil),          // 6: props.DistributePropsResponse
	(*QueryPropsRequest)(nil),                // 7: props.QueryPropsRequest
	(*QueryPropsResponse)(nil),               // 8: props.QueryPropsResponse
	(*ConsumePropsRequest)(nil),              // 9: props.ConsumePropsRequest
	(*ConsumePropsResponse)(nil),             // 10: props.ConsumePropsResponse
	(*DistributePropsRequest_PropsItem)(nil), // 11: props.DistributePropsRequest.PropsItem
	(*ConsumePropsRequest_PropsItem)(nil),    // 12: props.ConsumePropsRequest.PropsItem
	(*base.Base)(nil),                        // 13: base.Base
	(*base.BaseResp)(nil),                    // 14: base.BaseResp
}
var file_props_props_proto_depIdxs = []int32{
	0,  // 0: props.AddPropsRequest.props_list:type_name -> props.PropsInfo
	13, // 1: props.AddPropsRequest.Base:type_name -> base.Base
	14, // 2: props.AddPropsResponse.BaseResp:type_name -> base.BaseResp
	0,  // 3: props.UpdatePropsRequest.props_info:type_name -> props.PropsInfo
	13, // 4: props.UpdatePropsRequest.Base:type_name -> base.Base
	14, // 5: props.UpdatePropsResponse.BaseResp:type_name -> base.BaseResp
	11, // 6: props.DistributePropsRequest.props_items:type_name -> props.DistributePropsRequest.PropsItem
	13, // 7: props.DistributePropsRequest.Base:type_name -> base.Base
	14, // 8: props.DistributePropsResponse.BaseResp:type_name -> base.BaseResp
	13, // 9: props.QueryPropsRequest.Base:type_name -> base.Base
	0,  // 10: props.QueryPropsResponse.props_list:type_name -> props.PropsInfo
	14, // 11: props.QueryPropsResponse.BaseResp:type_name -> base.BaseResp
	12, // 12: props.ConsumePropsRequest.props_items:type_name -> props.ConsumePropsRequest.PropsItem
	13, // 13: props.ConsumePropsRequest.Base:type_name -> base.Base
	14, // 14: props.ConsumePropsResponse.BaseResp:type_name -> base.BaseResp
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_props_props_proto_init() }
func file_props_props_proto_init() {
	if File_props_props_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_props_props_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropsInfo); i {
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
		file_props_props_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPropsRequest); i {
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
		file_props_props_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPropsResponse); i {
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
		file_props_props_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePropsRequest); i {
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
		file_props_props_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePropsResponse); i {
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
		file_props_props_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistributePropsRequest); i {
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
		file_props_props_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistributePropsResponse); i {
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
		file_props_props_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPropsRequest); i {
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
		file_props_props_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryPropsResponse); i {
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
		file_props_props_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumePropsRequest); i {
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
		file_props_props_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumePropsResponse); i {
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
		file_props_props_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistributePropsRequest_PropsItem); i {
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
		file_props_props_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumePropsRequest_PropsItem); i {
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
			RawDescriptor: file_props_props_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_props_props_proto_goTypes,
		DependencyIndexes: file_props_props_proto_depIdxs,
		MessageInfos:      file_props_props_proto_msgTypes,
	}.Build()
	File_props_props_proto = out.File
	file_props_props_proto_rawDesc = nil
	file_props_props_proto_goTypes = nil
	file_props_props_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.3.2. DO NOT EDIT.
