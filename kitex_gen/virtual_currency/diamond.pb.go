// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.5.1
// source: virtual_currency/diamond.proto

package virtual_currency

import (
	context "context"
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

type DiamondSource int32

const (
	DiamondSource_DiamondSource_Unknown     DiamondSource = 0
	DiamondSource_DiamondSource_Promotion   DiamondSource = 1 // 推销拉新
	DiamondSource_DiamondSource_Consumption DiamondSource = 2 // 消费
)

// Enum value maps for DiamondSource.
var (
	DiamondSource_name = map[int32]string{
		0: "DiamondSource_Unknown",
		1: "DiamondSource_Promotion",
		2: "DiamondSource_Consumption",
	}
	DiamondSource_value = map[string]int32{
		"DiamondSource_Unknown":     0,
		"DiamondSource_Promotion":   1,
		"DiamondSource_Consumption": 2,
	}
)

func (x DiamondSource) Enum() *DiamondSource {
	p := new(DiamondSource)
	*p = x
	return p
}

func (x DiamondSource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DiamondSource) Descriptor() protoreflect.EnumDescriptor {
	return file_virtual_currency_diamond_proto_enumTypes[0].Descriptor()
}

func (DiamondSource) Type() protoreflect.EnumType {
	return &file_virtual_currency_diamond_proto_enumTypes[0]
}

func (x DiamondSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DiamondSource.Descriptor instead.
func (DiamondSource) EnumDescriptor() ([]byte, []int) {
	return file_virtual_currency_diamond_proto_rawDescGZIP(), []int{0}
}

type AddDiamondRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid    int64         `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Source DiamondSource `protobuf:"varint,2,opt,name=source,proto3,enum=virtual_currency.DiamondSource" json:"source,omitempty"` // 获取渠道
	Amount int32         `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`                                     // 数量
}

func (x *AddDiamondRequest) Reset() {
	*x = AddDiamondRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_virtual_currency_diamond_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDiamondRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDiamondRequest) ProtoMessage() {}

func (x *AddDiamondRequest) ProtoReflect() protoreflect.Message {
	mi := &file_virtual_currency_diamond_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDiamondRequest.ProtoReflect.Descriptor instead.
func (*AddDiamondRequest) Descriptor() ([]byte, []int) {
	return file_virtual_currency_diamond_proto_rawDescGZIP(), []int{0}
}

func (x *AddDiamondRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *AddDiamondRequest) GetSource() DiamondSource {
	if x != nil {
		return x.Source
	}
	return DiamondSource_DiamondSource_Unknown
}

func (x *AddDiamondRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type AddDiamondResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddDiamondResponse) Reset() {
	*x = AddDiamondResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_virtual_currency_diamond_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDiamondResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDiamondResponse) ProtoMessage() {}

func (x *AddDiamondResponse) ProtoReflect() protoreflect.Message {
	mi := &file_virtual_currency_diamond_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDiamondResponse.ProtoReflect.Descriptor instead.
func (*AddDiamondResponse) Descriptor() ([]byte, []int) {
	return file_virtual_currency_diamond_proto_rawDescGZIP(), []int{1}
}

// 消费钻石
type ConsumeDiamondRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid    int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Amount int32 `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"` // 数量
}

func (x *ConsumeDiamondRequest) Reset() {
	*x = ConsumeDiamondRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_virtual_currency_diamond_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumeDiamondRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumeDiamondRequest) ProtoMessage() {}

func (x *ConsumeDiamondRequest) ProtoReflect() protoreflect.Message {
	mi := &file_virtual_currency_diamond_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumeDiamondRequest.ProtoReflect.Descriptor instead.
func (*ConsumeDiamondRequest) Descriptor() ([]byte, []int) {
	return file_virtual_currency_diamond_proto_rawDescGZIP(), []int{2}
}

func (x *ConsumeDiamondRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ConsumeDiamondRequest) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type ConsumeDiamondResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ConsumeDiamondResponse) Reset() {
	*x = ConsumeDiamondResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_virtual_currency_diamond_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsumeDiamondResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumeDiamondResponse) ProtoMessage() {}

func (x *ConsumeDiamondResponse) ProtoReflect() protoreflect.Message {
	mi := &file_virtual_currency_diamond_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumeDiamondResponse.ProtoReflect.Descriptor instead.
func (*ConsumeDiamondResponse) Descriptor() ([]byte, []int) {
	return file_virtual_currency_diamond_proto_rawDescGZIP(), []int{3}
}

// 查询钻石
type QueryUserDiamondRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *QueryUserDiamondRequest) Reset() {
	*x = QueryUserDiamondRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_virtual_currency_diamond_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryUserDiamondRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryUserDiamondRequest) ProtoMessage() {}

func (x *QueryUserDiamondRequest) ProtoReflect() protoreflect.Message {
	mi := &file_virtual_currency_diamond_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryUserDiamondRequest.ProtoReflect.Descriptor instead.
func (*QueryUserDiamondRequest) Descriptor() ([]byte, []int) {
	return file_virtual_currency_diamond_proto_rawDescGZIP(), []int{4}
}

func (x *QueryUserDiamondRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type QueryUserDiamondResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount int32 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *QueryUserDiamondResponse) Reset() {
	*x = QueryUserDiamondResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_virtual_currency_diamond_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryUserDiamondResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryUserDiamondResponse) ProtoMessage() {}

func (x *QueryUserDiamondResponse) ProtoReflect() protoreflect.Message {
	mi := &file_virtual_currency_diamond_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryUserDiamondResponse.ProtoReflect.Descriptor instead.
func (*QueryUserDiamondResponse) Descriptor() ([]byte, []int) {
	return file_virtual_currency_diamond_proto_rawDescGZIP(), []int{5}
}

func (x *QueryUserDiamondResponse) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_virtual_currency_diamond_proto protoreflect.FileDescriptor

var file_virtual_currency_diamond_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x2f, 0x64, 0x69, 0x61, 0x6d, 0x6f, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x22, 0x76, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x44, 0x69, 0x61, 0x6d, 0x6f, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x76, 0x69, 0x72, 0x74,
	0x75, 0x61, 0x6c, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x44, 0x69, 0x61,
	0x6d, 0x6f, 0x6e, 0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x41, 0x64,
	0x64, 0x44, 0x69, 0x61, 0x6d, 0x6f, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x41, 0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x44, 0x69, 0x61, 0x6d, 0x6f,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x18, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x44, 0x69,
	0x61, 0x6d, 0x6f, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x0a,
	0x17, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x44, 0x69, 0x61, 0x6d, 0x6f, 0x6e,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x18, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x44, 0x69, 0x61, 0x6d, 0x6f, 0x6e, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x2a, 0x66,
	0x0a, 0x0d, 0x44, 0x69, 0x61, 0x6d, 0x6f, 0x6e, 0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x19, 0x0a, 0x15, 0x44, 0x69, 0x61, 0x6d, 0x6f, 0x6e, 0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x44, 0x69,
	0x61, 0x6d, 0x6f, 0x6e, 0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x50, 0x72, 0x6f, 0x6d,
	0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x01, 0x12, 0x1d, 0x0a, 0x19, 0x44, 0x69, 0x61, 0x6d, 0x6f,
	0x6e, 0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x10, 0x02, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x63, 0x5f, 0x6e, 0x66, 0x74, 0x2f, 0x6b, 0x69, 0x74, 0x65,
	0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_virtual_currency_diamond_proto_rawDescOnce sync.Once
	file_virtual_currency_diamond_proto_rawDescData = file_virtual_currency_diamond_proto_rawDesc
)

func file_virtual_currency_diamond_proto_rawDescGZIP() []byte {
	file_virtual_currency_diamond_proto_rawDescOnce.Do(func() {
		file_virtual_currency_diamond_proto_rawDescData = protoimpl.X.CompressGZIP(file_virtual_currency_diamond_proto_rawDescData)
	})
	return file_virtual_currency_diamond_proto_rawDescData
}

var file_virtual_currency_diamond_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_virtual_currency_diamond_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_virtual_currency_diamond_proto_goTypes = []interface{}{
	(DiamondSource)(0),               // 0: virtual_currency.DiamondSource
	(*AddDiamondRequest)(nil),        // 1: virtual_currency.AddDiamondRequest
	(*AddDiamondResponse)(nil),       // 2: virtual_currency.AddDiamondResponse
	(*ConsumeDiamondRequest)(nil),    // 3: virtual_currency.ConsumeDiamondRequest
	(*ConsumeDiamondResponse)(nil),   // 4: virtual_currency.ConsumeDiamondResponse
	(*QueryUserDiamondRequest)(nil),  // 5: virtual_currency.QueryUserDiamondRequest
	(*QueryUserDiamondResponse)(nil), // 6: virtual_currency.QueryUserDiamondResponse
}
var file_virtual_currency_diamond_proto_depIdxs = []int32{
	0, // 0: virtual_currency.AddDiamondRequest.source:type_name -> virtual_currency.DiamondSource
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_virtual_currency_diamond_proto_init() }
func file_virtual_currency_diamond_proto_init() {
	if File_virtual_currency_diamond_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_virtual_currency_diamond_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDiamondRequest); i {
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
		file_virtual_currency_diamond_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDiamondResponse); i {
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
		file_virtual_currency_diamond_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumeDiamondRequest); i {
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
		file_virtual_currency_diamond_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsumeDiamondResponse); i {
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
		file_virtual_currency_diamond_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryUserDiamondRequest); i {
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
		file_virtual_currency_diamond_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryUserDiamondResponse); i {
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
			RawDescriptor: file_virtual_currency_diamond_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_virtual_currency_diamond_proto_goTypes,
		DependencyIndexes: file_virtual_currency_diamond_proto_depIdxs,
		EnumInfos:         file_virtual_currency_diamond_proto_enumTypes,
		MessageInfos:      file_virtual_currency_diamond_proto_msgTypes,
	}.Build()
	File_virtual_currency_diamond_proto = out.File
	file_virtual_currency_diamond_proto_rawDesc = nil
	file_virtual_currency_diamond_proto_goTypes = nil
	file_virtual_currency_diamond_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.3.2. DO NOT EDIT.