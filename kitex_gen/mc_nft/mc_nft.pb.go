// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.5.1
// source: mc_nft.proto

package mc_nft

import (
	context "context"
	credits "github.com/mc_nft/kitex_gen/credits"
	ping "github.com/mc_nft/kitex_gen/ping"
	props "github.com/mc_nft/kitex_gen/props"
	user "github.com/mc_nft/kitex_gen/user"
	virtual_currency "github.com/mc_nft/kitex_gen/virtual_currency"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_mc_nft_proto protoreflect.FileDescriptor

var file_mc_nft_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x63, 0x5f, 0x6e, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6d, 0x63, 0x5f, 0x6e, 0x66, 0x74, 0x1a, 0x0f, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73,
	0x2f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x70, 0x72, 0x6f, 0x70, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1e, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x2f, 0x64, 0x69, 0x61, 0x6d, 0x6f, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0xc8, 0x09, 0x0a, 0x0a, 0x4e, 0x66, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x2d, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x11, 0x2e, 0x70, 0x69, 0x6e, 0x67,
	0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70,
	0x69, 0x6e, 0x67, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3c, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x16, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36,
	0x0a, 0x07, 0x53, 0x6d, 0x73, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x53, 0x6d, 0x73, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x6d, 0x73, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x75,
	0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a,
	0x0a, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x17, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b,
	0x0a, 0x08, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x70, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x72,
	0x6f, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f,
	0x70, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x50, 0x0a, 0x0f, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x70, 0x73, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x2e, 0x44, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x2e, 0x44, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x70,
	0x73, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50,
	0x72, 0x6f, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72,
	0x6f, 0x70, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x2e, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x57, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x44, 0x69, 0x61, 0x6d, 0x6f, 0x6e, 0x64, 0x12, 0x23, 0x2e,
	0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x2e, 0x41, 0x64, 0x64, 0x44, 0x69, 0x61, 0x6d, 0x6f, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x24, 0x2e, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x41, 0x64, 0x64, 0x44, 0x69, 0x61, 0x6d, 0x6f, 0x6e, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x44, 0x69, 0x61, 0x6d, 0x6f, 0x6e, 0x64, 0x12, 0x27, 0x2e, 0x76, 0x69, 0x72,
	0x74, 0x75, 0x61, 0x6c, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x43, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x44, 0x69, 0x61, 0x6d, 0x6f, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x44, 0x69,
	0x61, 0x6d, 0x6f, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x69, 0x0a,
	0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x44, 0x69, 0x61, 0x6d, 0x6f, 0x6e,
	0x64, 0x12, 0x29, 0x2e, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x44, 0x69,
	0x61, 0x6d, 0x6f, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x76,
	0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x55, 0x73, 0x65, 0x72, 0x44, 0x69, 0x61, 0x6d, 0x6f, 0x6e, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x4f,
	0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x12, 0x22,
	0x2e, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x2e, 0x41, 0x64, 0x64,
	0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x12, 0x1e, 0x2e, 0x63, 0x72, 0x65, 0x64,
	0x69, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x63, 0x72, 0x65, 0x64,
	0x69, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x1e, 0x2e, 0x63,
	0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x73, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x63,
	0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x73, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x24, 0x5a,
	0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x63, 0x5f, 0x6e,
	0x66, 0x74, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x6d, 0x63, 0x5f,
	0x6e, 0x66, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_mc_nft_proto_goTypes = []interface{}{
	(*ping.PingRequest)(nil),                          // 0: ping.PingRequest
	(*user.UserLoginRequest)(nil),                     // 1: user.UserLoginRequest
	(*user.SmsSendRequest)(nil),                       // 2: user.SmsSendRequest
	(*user.AuthenticateRequest)(nil),                  // 3: user.AuthenticateRequest
	(*user.UserDetailRequest)(nil),                    // 4: user.UserDetailRequest
	(*props.AddPropsRequest)(nil),                     // 5: props.AddPropsRequest
	(*props.UpdatePropsRequest)(nil),                  // 6: props.UpdatePropsRequest
	(*props.DistributePropsRequest)(nil),              // 7: props.DistributePropsRequest
	(*props.QueryPropsRequest)(nil),                   // 8: props.QueryPropsRequest
	(*props.ConsumePropsRequest)(nil),                 // 9: props.ConsumePropsRequest
	(*virtual_currency.AddDiamondRequest)(nil),        // 10: virtual_currency.AddDiamondRequest
	(*virtual_currency.ConsumeDiamondRequest)(nil),    // 11: virtual_currency.ConsumeDiamondRequest
	(*virtual_currency.QueryUserDiamondRequest)(nil),  // 12: virtual_currency.QueryUserDiamondRequest
	(*credits.AddOrUpdateCreditsRequest)(nil),         // 13: credits.AddOrUpdateCreditsRequest
	(*credits.GetUserCreditsRequest)(nil),             // 14: credits.GetUserCreditsRequest
	(*credits.GetCreditsRankRequest)(nil),             // 15: credits.GetCreditsRankRequest
	(*ping.PingResponse)(nil),                         // 16: ping.PingResponse
	(*user.UserLoginResponse)(nil),                    // 17: user.UserLoginResponse
	(*user.SmsSendResponse)(nil),                      // 18: user.SmsSendResponse
	(*user.AuthenticateResponse)(nil),                 // 19: user.AuthenticateResponse
	(*user.UserDetailResponse)(nil),                   // 20: user.UserDetailResponse
	(*props.AddPropsResponse)(nil),                    // 21: props.AddPropsResponse
	(*props.UpdatePropsResponse)(nil),                 // 22: props.UpdatePropsResponse
	(*props.DistributePropsResponse)(nil),             // 23: props.DistributePropsResponse
	(*props.QueryPropsResponse)(nil),                  // 24: props.QueryPropsResponse
	(*props.ConsumePropsResponse)(nil),                // 25: props.ConsumePropsResponse
	(*virtual_currency.AddDiamondResponse)(nil),       // 26: virtual_currency.AddDiamondResponse
	(*virtual_currency.ConsumeDiamondResponse)(nil),   // 27: virtual_currency.ConsumeDiamondResponse
	(*virtual_currency.QueryUserDiamondResponse)(nil), // 28: virtual_currency.QueryUserDiamondResponse
	(*credits.AddOrUpdateCreditsResponse)(nil),        // 29: credits.AddOrUpdateCreditsResponse
	(*credits.GetUserCreditsResponse)(nil),            // 30: credits.GetUserCreditsResponse
	(*credits.GetCreditsRankResponse)(nil),            // 31: credits.GetCreditsRankResponse
}
var file_mc_nft_proto_depIdxs = []int32{
	0,  // 0: mc_nft.NftService.Ping:input_type -> ping.PingRequest
	1,  // 1: mc_nft.NftService.UserLogin:input_type -> user.UserLoginRequest
	2,  // 2: mc_nft.NftService.SmsSend:input_type -> user.SmsSendRequest
	3,  // 3: mc_nft.NftService.Authenticate:input_type -> user.AuthenticateRequest
	4,  // 4: mc_nft.NftService.UserDetail:input_type -> user.UserDetailRequest
	5,  // 5: mc_nft.NftService.AddProps:input_type -> props.AddPropsRequest
	6,  // 6: mc_nft.NftService.UpdateProps:input_type -> props.UpdatePropsRequest
	7,  // 7: mc_nft.NftService.DistributeProps:input_type -> props.DistributePropsRequest
	8,  // 8: mc_nft.NftService.QueryProps:input_type -> props.QueryPropsRequest
	9,  // 9: mc_nft.NftService.ConsumeProps:input_type -> props.ConsumePropsRequest
	10, // 10: mc_nft.NftService.AddDiamond:input_type -> virtual_currency.AddDiamondRequest
	11, // 11: mc_nft.NftService.ConsumeDiamond:input_type -> virtual_currency.ConsumeDiamondRequest
	12, // 12: mc_nft.NftService.QueryUserDiamond:input_type -> virtual_currency.QueryUserDiamondRequest
	13, // 13: mc_nft.NftService.AddOrUpdateCredits:input_type -> credits.AddOrUpdateCreditsRequest
	14, // 14: mc_nft.NftService.GetUserCredits:input_type -> credits.GetUserCreditsRequest
	15, // 15: mc_nft.NftService.GetCreditsRank:input_type -> credits.GetCreditsRankRequest
	16, // 16: mc_nft.NftService.Ping:output_type -> ping.PingResponse
	17, // 17: mc_nft.NftService.UserLogin:output_type -> user.UserLoginResponse
	18, // 18: mc_nft.NftService.SmsSend:output_type -> user.SmsSendResponse
	19, // 19: mc_nft.NftService.Authenticate:output_type -> user.AuthenticateResponse
	20, // 20: mc_nft.NftService.UserDetail:output_type -> user.UserDetailResponse
	21, // 21: mc_nft.NftService.AddProps:output_type -> props.AddPropsResponse
	22, // 22: mc_nft.NftService.UpdateProps:output_type -> props.UpdatePropsResponse
	23, // 23: mc_nft.NftService.DistributeProps:output_type -> props.DistributePropsResponse
	24, // 24: mc_nft.NftService.QueryProps:output_type -> props.QueryPropsResponse
	25, // 25: mc_nft.NftService.ConsumeProps:output_type -> props.ConsumePropsResponse
	26, // 26: mc_nft.NftService.AddDiamond:output_type -> virtual_currency.AddDiamondResponse
	27, // 27: mc_nft.NftService.ConsumeDiamond:output_type -> virtual_currency.ConsumeDiamondResponse
	28, // 28: mc_nft.NftService.QueryUserDiamond:output_type -> virtual_currency.QueryUserDiamondResponse
	29, // 29: mc_nft.NftService.AddOrUpdateCredits:output_type -> credits.AddOrUpdateCreditsResponse
	30, // 30: mc_nft.NftService.GetUserCredits:output_type -> credits.GetUserCreditsResponse
	31, // 31: mc_nft.NftService.GetCreditsRank:output_type -> credits.GetCreditsRankResponse
	16, // [16:32] is the sub-list for method output_type
	0,  // [0:16] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_mc_nft_proto_init() }
func file_mc_nft_proto_init() {
	if File_mc_nft_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mc_nft_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mc_nft_proto_goTypes,
		DependencyIndexes: file_mc_nft_proto_depIdxs,
	}.Build()
	File_mc_nft_proto = out.File
	file_mc_nft_proto_rawDesc = nil
	file_mc_nft_proto_goTypes = nil
	file_mc_nft_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.3.2. DO NOT EDIT.

type NftService interface {
	Ping(ctx context.Context, req *ping.PingRequest) (res *ping.PingResponse, err error)
	UserLogin(ctx context.Context, req *user.UserLoginRequest) (res *user.UserLoginResponse, err error)
	SmsSend(ctx context.Context, req *user.SmsSendRequest) (res *user.SmsSendResponse, err error)
	Authenticate(ctx context.Context, req *user.AuthenticateRequest) (res *user.AuthenticateResponse, err error)
	UserDetail(ctx context.Context, req *user.UserDetailRequest) (res *user.UserDetailResponse, err error)
	AddProps(ctx context.Context, req *props.AddPropsRequest) (res *props.AddPropsResponse, err error)
	UpdateProps(ctx context.Context, req *props.UpdatePropsRequest) (res *props.UpdatePropsResponse, err error)
	DistributeProps(ctx context.Context, req *props.DistributePropsRequest) (res *props.DistributePropsResponse, err error)
	QueryProps(ctx context.Context, req *props.QueryPropsRequest) (res *props.QueryPropsResponse, err error)
	ConsumeProps(ctx context.Context, req *props.ConsumePropsRequest) (res *props.ConsumePropsResponse, err error)
	AddDiamond(ctx context.Context, req *virtual_currency.AddDiamondRequest) (res *virtual_currency.AddDiamondResponse, err error)
	ConsumeDiamond(ctx context.Context, req *virtual_currency.ConsumeDiamondRequest) (res *virtual_currency.ConsumeDiamondResponse, err error)
	QueryUserDiamond(ctx context.Context, req *virtual_currency.QueryUserDiamondRequest) (res *virtual_currency.QueryUserDiamondResponse, err error)
	AddOrUpdateCredits(ctx context.Context, req *credits.AddOrUpdateCreditsRequest) (res *credits.AddOrUpdateCreditsResponse, err error)
	GetUserCredits(ctx context.Context, req *credits.GetUserCreditsRequest) (res *credits.GetUserCreditsResponse, err error)
	GetCreditsRank(ctx context.Context, req *credits.GetCreditsRankRequest) (res *credits.GetCreditsRankResponse, err error)
}
