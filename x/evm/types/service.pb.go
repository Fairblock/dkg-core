// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: evm/v1beta1/service.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("evm/v1beta1/service.proto", fileDescriptor_e674b0c159b5e0b5) }
func init() { golang_proto.RegisterFile("evm/v1beta1/service.proto", fileDescriptor_e674b0c159b5e0b5) }

var fileDescriptor_e674b0c159b5e0b5 = []byte{
	// 812 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x96, 0x4f, 0x6f, 0xd3, 0x48,
	0x18, 0xc6, 0x3b, 0xab, 0x55, 0xb5, 0x9a, 0x5d, 0xad, 0xba, 0xde, 0x52, 0xda, 0x50, 0x4c, 0xeb,
	0x26, 0x29, 0x4d, 0xeb, 0xb8, 0x7f, 0x6e, 0xdc, 0x68, 0x7a, 0xe3, 0xaf, 0x68, 0xc5, 0x81, 0x0b,
	0x72, 0x9c, 0xa9, 0x3b, 0x24, 0x99, 0x31, 0x9e, 0x49, 0x9a, 0x08, 0x21, 0x01, 0x17, 0x24, 0x0e,
	0x08, 0xc1, 0x85, 0x33, 0x48, 0x7c, 0x00, 0x3e, 0x01, 0x47, 0x8e, 0x45, 0x5c, 0x38, 0xa2, 0x86,
	0x0f, 0x82, 0x66, 0x3c, 0xd3, 0xda, 0x8e, 0xe3, 0xe4, 0xd6, 0xce, 0xfb, 0xcc, 0xfb, 0xfc, 0x3c,
	0xef, 0x63, 0x4f, 0xe0, 0x02, 0xea, 0xb6, 0x9d, 0xee, 0x56, 0x1d, 0x71, 0x77, 0xcb, 0x61, 0x28,
	0xec, 0x62, 0x0f, 0x55, 0x83, 0x90, 0x72, 0x6a, 0xfc, 0x8d, 0xba, 0xed, 0xaa, 0x2a, 0x15, 0x66,
	0x7d, 0xea, 0x53, 0xb9, 0xee, 0x88, 0xbf, 0x22, 0x49, 0x61, 0xd1, 0xa7, 0xd4, 0x6f, 0x21, 0xc7,
	0x0d, 0xb0, 0xe3, 0x12, 0x42, 0xb9, 0xcb, 0x31, 0x25, 0x4c, 0x55, 0x67, 0xe3, 0xbd, 0x79, 0x2f,
	0x5a, 0xdd, 0xfe, 0x66, 0x40, 0x78, 0x8b, 0xf9, 0xfb, 0x91, 0x97, 0xf1, 0x08, 0xfe, 0x79, 0x13,
	0x93, 0xa6, 0x31, 0x5f, 0x8d, 0xd9, 0x55, 0xc5, 0xd2, 0x3d, 0xf4, 0xb8, 0x83, 0x18, 0x2f, 0x2c,
	0x64, 0x54, 0x58, 0x40, 0x09, 0x43, 0x96, 0xfd, 0xe2, 0xfb, 0xaf, 0x77, 0x7f, 0xac, 0x5a, 0x96,
	0xe3, 0xf6, 0x50, 0xcb, 0x0d, 0x1d, 0xe1, 0xd8, 0xc2, 0xa4, 0xe9, 0x3c, 0x09, 0x91, 0x87, 0x03,
	0x8c, 0x08, 0x7f, 0xe8, 0x1d, 0xb9, 0x98, 0x3c, 0xbd, 0x06, 0x2a, 0x46, 0x1f, 0xfe, 0x53, 0xa3,
	0xe4, 0x10, 0x87, 0xed, 0x9a, 0x58, 0x33, 0x96, 0x12, 0x9d, 0xe3, 0x25, 0xed, 0xbd, 0x9c, 0xa3,
	0x50, 0x0c, 0x45, 0xc9, 0x60, 0x5a, 0x0b, 0x71, 0x06, 0x2f, 0x52, 0xda, 0xd2, 0x5b, 0x58, 0x3f,
	0x03, 0x67, 0xde, 0x07, 0xb4, 0x89, 0x46, 0x78, 0xcb, 0x52, 0xae, 0xb7, 0x52, 0x28, 0xef, 0x75,
	0xe9, 0x5d, 0xb2, 0x96, 0xb2, 0xbc, 0x51, 0xe8, 0x6d, 0x6f, 0xda, 0x0d, 0x14, 0xb4, 0x68, 0x5f,
	0x20, 0xbc, 0x04, 0xf0, 0x5f, 0xd5, 0x65, 0x0f, 0x05, 0x94, 0x61, 0x6e, 0x58, 0x59, 0x16, 0xaa,
	0xa8, 0x31, 0x56, 0x72, 0x35, 0x0a, 0x64, 0x43, 0x82, 0x94, 0xad, 0xe5, 0x5c, 0x10, 0xb1, 0x45,
	0x90, 0xbc, 0x07, 0xd0, 0xd0, 0xcf, 0x13, 0xba, 0x84, 0x1d, 0xa2, 0xf0, 0x06, 0xea, 0x1b, 0xe5,
	0xcc, 0x07, 0x3e, 0x17, 0x68, 0xa2, 0xd5, 0xb1, 0x3a, 0x45, 0xb5, 0x25, 0xa9, 0xd6, 0xad, 0x72,
	0x16, 0x15, 0x57, 0x1b, 0x6c, 0x7a, 0x4c, 0x50, 0xc8, 0x8e, 0x70, 0x20, 0xd0, 0x5e, 0x01, 0x38,
	0x73, 0x9f, 0x72, 0x94, 0xc8, 0x49, 0x31, 0x61, 0x98, 0x2e, 0x6b, 0xac, 0xd2, 0x18, 0x95, 0x82,
	0x5a, 0x93, 0x50, 0x2b, 0x96, 0x19, 0x87, 0xea, 0x52, 0x8e, 0xec, 0xa1, 0xd0, 0xbc, 0x05, 0xd0,
	0x88, 0xf5, 0xd1, 0x53, 0x2b, 0x8f, 0x32, 0x4a, 0x4d, 0x6e, 0x75, 0xac, 0x2e, 0x2f, 0x46, 0x09,
	0xa4, 0xd8, 0xf0, 0x52, 0x27, 0x14, 0xa5, 0x79, 0xe4, 0x09, 0x25, 0x12, 0x5d, 0x1a, 0xa3, 0x9a,
	0xf8, 0x84, 0xb8, 0xd0, 0x0b, 0x98, 0x8f, 0x00, 0xce, 0xc5, 0xfb, 0xc4, 0xd2, 0x54, 0x19, 0x69,
	0x36, 0x9c, 0xa8, 0xf5, 0x89, 0xb4, 0x0a, 0x6f, 0x53, 0xe2, 0x55, 0xac, 0xd2, 0x68, 0x3c, 0x1d,
	0xad, 0x26, 0x92, 0x6f, 0xde, 0x6b, 0x00, 0xff, 0xab, 0x85, 0xc8, 0xe5, 0x68, 0x4f, 0xbe, 0x8d,
	0xd1, 0x99, 0x25, 0x4f, 0x63, 0xa8, 0xae, 0xd9, 0xca, 0xe3, 0x64, 0x0a, 0xab, 0x22, 0xb1, 0x8a,
	0xd6, 0x95, 0x44, 0xd8, 0xa5, 0x5c, 0x7d, 0x04, 0xce, 0x8f, 0xed, 0x39, 0x80, 0x33, 0x51, 0xa7,
	0xdd, 0x4e, 0x48, 0x64, 0x1f, 0x96, 0x9a, 0x61, 0xba, 0x9c, 0x3d, 0xc3, 0x61, 0x95, 0xa2, 0x59,
	0x92, 0x34, 0x05, 0xeb, 0x42, 0x9c, 0x86, 0x61, 0x9f, 0xd8, 0xf5, 0x4e, 0x28, 0x19, 0x5c, 0x38,
	0xbd, 0x8f, 0x7d, 0x72, 0xd0, 0x33, 0x0a, 0x89, 0x96, 0xd1, 0xa2, 0xb6, 0xbb, 0x94, 0x59, 0x53,
	0x26, 0xa6, 0x34, 0x99, 0xb7, 0xfe, 0x1f, 0x32, 0xe1, 0x3d, 0x61, 0xf1, 0x01, 0xc0, 0xb9, 0x88,
	0xf0, 0x2e, 0x22, 0x0d, 0x4c, 0x7c, 0x3d, 0x4e, 0x96, 0x4a, 0x47, 0xb6, 0x28, 0x3b, 0x1d, 0xa3,
	0xb4, 0x8a, 0xc9, 0x91, 0x4c, 0x6b, 0x56, 0x31, 0x63, 0x0c, 0x41, 0xb4, 0xe9, 0x2c, 0x1f, 0x4c,
	0x40, 0x7e, 0x02, 0xf0, 0x62, 0xd4, 0x53, 0x37, 0xbb, 0xa3, 0x3f, 0x48, 0x46, 0x96, 0xf3, 0x90,
	0x4a, 0x63, 0x6e, 0x4c, 0x26, 0xce, 0x4b, 0xb1, 0xe2, 0xcc, 0xfe, 0x34, 0x7e, 0x06, 0xb0, 0x90,
	0xea, 0x1a, 0xa0, 0xd0, 0xe5, 0x34, 0x62, 0xad, 0xe6, 0xd9, 0xc7, 0x84, 0x1a, 0xd7, 0x99, 0x58,
	0xaf, 0x88, 0x77, 0x24, 0xb1, 0x6d, 0x5d, 0xcd, 0x25, 0x8e, 0xed, 0x54, 0x57, 0xbe, 0x08, 0x4d,
	0x8d, 0xb6, 0xdb, 0x2e, 0x69, 0xb0, 0xd4, 0xb5, 0x1b, 0x2f, 0x65, 0x5f, 0xbb, 0x49, 0x45, 0xde,
	0x95, 0x2f, 0x73, 0xe7, 0x29, 0xa9, 0xb0, 0xc6, 0xf0, 0xaf, 0xeb, 0x8d, 0x46, 0x74, 0x83, 0x2c,
	0x26, 0x9a, 0xea, 0x65, 0x6d, 0x79, 0x79, 0x44, 0x35, 0xef, 0x5d, 0x72, 0x1b, 0x8d, 0xb3, 0x8b,
	0x62, 0xf7, 0xf6, 0xd7, 0x53, 0x13, 0x9c, 0x9c, 0x9a, 0xe0, 0xe7, 0xa9, 0x09, 0xde, 0x0c, 0xcc,
	0xa9, 0x2f, 0x03, 0x13, 0x9c, 0x0c, 0xcc, 0xa9, 0x1f, 0x03, 0x73, 0xea, 0xc1, 0xa6, 0x8f, 0xf9,
	0x51, 0xa7, 0x5e, 0xf5, 0x68, 0x5b, 0x75, 0x20, 0x88, 0x1f, 0xd3, 0xb0, 0xa9, 0xfe, 0xb3, 0x3d,
	0x1a, 0x22, 0xa7, 0x27, 0xdb, 0xf2, 0x7e, 0x80, 0x58, 0x7d, 0x5a, 0xfe, 0x54, 0xdb, 0xf9, 0x1d,
	0x00, 0x00, 0xff, 0xff, 0x10, 0x4e, 0x4c, 0xba, 0x1e, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgServiceClient is the client API for MsgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgServiceClient interface {
	Link(ctx context.Context, in *LinkRequest, opts ...grpc.CallOption) (*LinkResponse, error)
	ConfirmChain(ctx context.Context, in *ConfirmChainRequest, opts ...grpc.CallOption) (*ConfirmChainResponse, error)
	ConfirmToken(ctx context.Context, in *ConfirmTokenRequest, opts ...grpc.CallOption) (*ConfirmTokenResponse, error)
	ConfirmDeposit(ctx context.Context, in *ConfirmDepositRequest, opts ...grpc.CallOption) (*ConfirmDepositResponse, error)
	ConfirmTransferKey(ctx context.Context, in *ConfirmTransferKeyRequest, opts ...grpc.CallOption) (*ConfirmTransferKeyResponse, error)
	VoteConfirmChain(ctx context.Context, in *VoteConfirmChainRequest, opts ...grpc.CallOption) (*VoteConfirmChainResponse, error)
	VoteConfirmDeposit(ctx context.Context, in *VoteConfirmDepositRequest, opts ...grpc.CallOption) (*VoteConfirmDepositResponse, error)
	VoteConfirmToken(ctx context.Context, in *VoteConfirmTokenRequest, opts ...grpc.CallOption) (*VoteConfirmTokenResponse, error)
	VoteConfirmTransferKey(ctx context.Context, in *VoteConfirmTransferKeyRequest, opts ...grpc.CallOption) (*VoteConfirmTransferKeyResponse, error)
	CreateDeployToken(ctx context.Context, in *CreateDeployTokenRequest, opts ...grpc.CallOption) (*CreateDeployTokenResponse, error)
	CreateBurnTokens(ctx context.Context, in *CreateBurnTokensRequest, opts ...grpc.CallOption) (*CreateBurnTokensResponse, error)
	SignTx(ctx context.Context, in *SignTxRequest, opts ...grpc.CallOption) (*SignTxResponse, error)
	CreatePendingTransfers(ctx context.Context, in *CreatePendingTransfersRequest, opts ...grpc.CallOption) (*CreatePendingTransfersResponse, error)
	CreateTransferOwnership(ctx context.Context, in *CreateTransferOwnershipRequest, opts ...grpc.CallOption) (*CreateTransferOwnershipResponse, error)
	CreateTransferOperatorship(ctx context.Context, in *CreateTransferOperatorshipRequest, opts ...grpc.CallOption) (*CreateTransferOperatorshipResponse, error)
	SignCommands(ctx context.Context, in *SignCommandsRequest, opts ...grpc.CallOption) (*SignCommandsResponse, error)
	AddChain(ctx context.Context, in *AddChainRequest, opts ...grpc.CallOption) (*AddChainResponse, error)
}

type msgServiceClient struct {
	cc grpc1.ClientConn
}

func NewMsgServiceClient(cc grpc1.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) Link(ctx context.Context, in *LinkRequest, opts ...grpc.CallOption) (*LinkResponse, error) {
	out := new(LinkResponse)
	err := c.cc.Invoke(ctx, "/evm.v1beta1.MsgService/Link", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) ConfirmChain(ctx context.Context, in *ConfirmChainRequest, opts ...grpc.CallOption) (*ConfirmChainResponse, error) {
	out := new(ConfirmChainResponse)
	err := c.cc.Invoke(ctx, "/evm.v1beta1.MsgService/ConfirmChain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) ConfirmToken(ctx context.Context, in *ConfirmTokenRequest, opts ...grpc.CallOption) (*ConfirmTokenResponse, error) {
	out := new(ConfirmTokenResponse)
	err := c.cc.Invoke(ctx, "/evm.v1beta1.MsgService/ConfirmToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) ConfirmDeposit(ctx context.Context, in *ConfirmDepositRequest, opts ...grpc.CallOption) (*ConfirmDepositResponse, error) {
	out := new(ConfirmDepositResponse)
	err := c.cc.Invoke(ctx, "/evm.v1beta1.MsgService/ConfirmDeposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) ConfirmTransferKey(ctx context.Context, in *ConfirmTransferKeyRequest, opts ...grpc.CallOption) (*ConfirmTransferKeyResponse, error) {
	out := new(ConfirmTransferKeyResponse)
	err := c.cc.Invoke(ctx, "/evm.v1beta1.MsgService/ConfirmTransferKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) VoteConfirmChain(ctx context.Context, in *VoteConfirmChainRequest, opts ...grpc.CallOption) (*VoteConfirmChainResponse, error) {
	out := new(VoteConfirmChainResponse)
	err := c.cc.Invoke(ctx, "/evm.v1beta1.MsgService/VoteConfirmChain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) VoteConfirmDeposit(ctx context.Context, in *VoteConfirmDepositRequest, opts ...grpc.CallOption) (*VoteConfirmDepositResponse, error) {
	out := new(VoteConfirmDepositResponse)
	err := c.cc.Invoke(ctx, "/evm.v1beta1.MsgService/VoteConfirmDeposit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) VoteConfirmToken(ctx context.Context, in *VoteConfirmTokenRequest, opts ...grpc.CallOption) (*VoteConfirmTokenResponse, error) {
	out := new(VoteConfirmTokenResponse)
	err := c.cc.Invoke(ctx, "/evm.v1beta1.MsgService/VoteConfirmToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) VoteConfirmTransferKey(ctx context.Context, in *VoteConfirmTransferKeyRequest, opts ...grpc.CallOption) (*VoteConfirmTransferKeyResponse, error) {
	out := new(VoteConfirmTransferKeyResponse)
	err := c.cc.Invoke(ctx, "/evm.v1beta1.MsgService/VoteConfirmTransferKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) CreateDeployToken(ctx context.Context, in *CreateDeployTokenRequest, opts ...grpc.CallOption) (*CreateDeployTokenResponse, error) {
	out := new(CreateDeployTokenResponse)
	err := c.cc.Invoke(ctx, "/evm.v1beta1.MsgService/CreateDeployToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) CreateBurnTokens(ctx context.Context, in *CreateBurnTokensRequest, opts ...grpc.CallOption) (*CreateBurnTokensResponse, error) {
	out := new(CreateBurnTokensResponse)
	err := c.cc.Invoke(ctx, "/evm.v1beta1.MsgService/CreateBurnTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) SignTx(ctx context.Context, in *SignTxRequest, opts ...grpc.CallOption) (*SignTxResponse, error) {
	out := new(SignTxResponse)
	err := c.cc.Invoke(ctx, "/evm.v1beta1.MsgService/SignTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) CreatePendingTransfers(ctx context.Context, in *CreatePendingTransfersRequest, opts ...grpc.CallOption) (*CreatePendingTransfersResponse, error) {
	out := new(CreatePendingTransfersResponse)
	err := c.cc.Invoke(ctx, "/evm.v1beta1.MsgService/CreatePendingTransfers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) CreateTransferOwnership(ctx context.Context, in *CreateTransferOwnershipRequest, opts ...grpc.CallOption) (*CreateTransferOwnershipResponse, error) {
	out := new(CreateTransferOwnershipResponse)
	err := c.cc.Invoke(ctx, "/evm.v1beta1.MsgService/CreateTransferOwnership", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) CreateTransferOperatorship(ctx context.Context, in *CreateTransferOperatorshipRequest, opts ...grpc.CallOption) (*CreateTransferOperatorshipResponse, error) {
	out := new(CreateTransferOperatorshipResponse)
	err := c.cc.Invoke(ctx, "/evm.v1beta1.MsgService/CreateTransferOperatorship", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) SignCommands(ctx context.Context, in *SignCommandsRequest, opts ...grpc.CallOption) (*SignCommandsResponse, error) {
	out := new(SignCommandsResponse)
	err := c.cc.Invoke(ctx, "/evm.v1beta1.MsgService/SignCommands", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) AddChain(ctx context.Context, in *AddChainRequest, opts ...grpc.CallOption) (*AddChainResponse, error) {
	out := new(AddChainResponse)
	err := c.cc.Invoke(ctx, "/evm.v1beta1.MsgService/AddChain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
	Link(context.Context, *LinkRequest) (*LinkResponse, error)
	ConfirmChain(context.Context, *ConfirmChainRequest) (*ConfirmChainResponse, error)
	ConfirmToken(context.Context, *ConfirmTokenRequest) (*ConfirmTokenResponse, error)
	ConfirmDeposit(context.Context, *ConfirmDepositRequest) (*ConfirmDepositResponse, error)
	ConfirmTransferKey(context.Context, *ConfirmTransferKeyRequest) (*ConfirmTransferKeyResponse, error)
	VoteConfirmChain(context.Context, *VoteConfirmChainRequest) (*VoteConfirmChainResponse, error)
	VoteConfirmDeposit(context.Context, *VoteConfirmDepositRequest) (*VoteConfirmDepositResponse, error)
	VoteConfirmToken(context.Context, *VoteConfirmTokenRequest) (*VoteConfirmTokenResponse, error)
	VoteConfirmTransferKey(context.Context, *VoteConfirmTransferKeyRequest) (*VoteConfirmTransferKeyResponse, error)
	CreateDeployToken(context.Context, *CreateDeployTokenRequest) (*CreateDeployTokenResponse, error)
	CreateBurnTokens(context.Context, *CreateBurnTokensRequest) (*CreateBurnTokensResponse, error)
	SignTx(context.Context, *SignTxRequest) (*SignTxResponse, error)
	CreatePendingTransfers(context.Context, *CreatePendingTransfersRequest) (*CreatePendingTransfersResponse, error)
	CreateTransferOwnership(context.Context, *CreateTransferOwnershipRequest) (*CreateTransferOwnershipResponse, error)
	CreateTransferOperatorship(context.Context, *CreateTransferOperatorshipRequest) (*CreateTransferOperatorshipResponse, error)
	SignCommands(context.Context, *SignCommandsRequest) (*SignCommandsResponse, error)
	AddChain(context.Context, *AddChainRequest) (*AddChainResponse, error)
}

// UnimplementedMsgServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func (*UnimplementedMsgServiceServer) Link(ctx context.Context, req *LinkRequest) (*LinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Link not implemented")
}
func (*UnimplementedMsgServiceServer) ConfirmChain(ctx context.Context, req *ConfirmChainRequest) (*ConfirmChainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmChain not implemented")
}
func (*UnimplementedMsgServiceServer) ConfirmToken(ctx context.Context, req *ConfirmTokenRequest) (*ConfirmTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmToken not implemented")
}
func (*UnimplementedMsgServiceServer) ConfirmDeposit(ctx context.Context, req *ConfirmDepositRequest) (*ConfirmDepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmDeposit not implemented")
}
func (*UnimplementedMsgServiceServer) ConfirmTransferKey(ctx context.Context, req *ConfirmTransferKeyRequest) (*ConfirmTransferKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmTransferKey not implemented")
}
func (*UnimplementedMsgServiceServer) VoteConfirmChain(ctx context.Context, req *VoteConfirmChainRequest) (*VoteConfirmChainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VoteConfirmChain not implemented")
}
func (*UnimplementedMsgServiceServer) VoteConfirmDeposit(ctx context.Context, req *VoteConfirmDepositRequest) (*VoteConfirmDepositResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VoteConfirmDeposit not implemented")
}
func (*UnimplementedMsgServiceServer) VoteConfirmToken(ctx context.Context, req *VoteConfirmTokenRequest) (*VoteConfirmTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VoteConfirmToken not implemented")
}
func (*UnimplementedMsgServiceServer) VoteConfirmTransferKey(ctx context.Context, req *VoteConfirmTransferKeyRequest) (*VoteConfirmTransferKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VoteConfirmTransferKey not implemented")
}
func (*UnimplementedMsgServiceServer) CreateDeployToken(ctx context.Context, req *CreateDeployTokenRequest) (*CreateDeployTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDeployToken not implemented")
}
func (*UnimplementedMsgServiceServer) CreateBurnTokens(ctx context.Context, req *CreateBurnTokensRequest) (*CreateBurnTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBurnTokens not implemented")
}
func (*UnimplementedMsgServiceServer) SignTx(ctx context.Context, req *SignTxRequest) (*SignTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignTx not implemented")
}
func (*UnimplementedMsgServiceServer) CreatePendingTransfers(ctx context.Context, req *CreatePendingTransfersRequest) (*CreatePendingTransfersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePendingTransfers not implemented")
}
func (*UnimplementedMsgServiceServer) CreateTransferOwnership(ctx context.Context, req *CreateTransferOwnershipRequest) (*CreateTransferOwnershipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransferOwnership not implemented")
}
func (*UnimplementedMsgServiceServer) CreateTransferOperatorship(ctx context.Context, req *CreateTransferOperatorshipRequest) (*CreateTransferOperatorshipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransferOperatorship not implemented")
}
func (*UnimplementedMsgServiceServer) SignCommands(ctx context.Context, req *SignCommandsRequest) (*SignCommandsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignCommands not implemented")
}
func (*UnimplementedMsgServiceServer) AddChain(ctx context.Context, req *AddChainRequest) (*AddChainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddChain not implemented")
}

func RegisterMsgServiceServer(s grpc1.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

func _MsgService_Link_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).Link(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evm.v1beta1.MsgService/Link",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).Link(ctx, req.(*LinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_ConfirmChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmChainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).ConfirmChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evm.v1beta1.MsgService/ConfirmChain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).ConfirmChain(ctx, req.(*ConfirmChainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_ConfirmToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).ConfirmToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evm.v1beta1.MsgService/ConfirmToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).ConfirmToken(ctx, req.(*ConfirmTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_ConfirmDeposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmDepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).ConfirmDeposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evm.v1beta1.MsgService/ConfirmDeposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).ConfirmDeposit(ctx, req.(*ConfirmDepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_ConfirmTransferKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmTransferKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).ConfirmTransferKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evm.v1beta1.MsgService/ConfirmTransferKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).ConfirmTransferKey(ctx, req.(*ConfirmTransferKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_VoteConfirmChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteConfirmChainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).VoteConfirmChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evm.v1beta1.MsgService/VoteConfirmChain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).VoteConfirmChain(ctx, req.(*VoteConfirmChainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_VoteConfirmDeposit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteConfirmDepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).VoteConfirmDeposit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evm.v1beta1.MsgService/VoteConfirmDeposit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).VoteConfirmDeposit(ctx, req.(*VoteConfirmDepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_VoteConfirmToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteConfirmTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).VoteConfirmToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evm.v1beta1.MsgService/VoteConfirmToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).VoteConfirmToken(ctx, req.(*VoteConfirmTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_VoteConfirmTransferKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteConfirmTransferKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).VoteConfirmTransferKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evm.v1beta1.MsgService/VoteConfirmTransferKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).VoteConfirmTransferKey(ctx, req.(*VoteConfirmTransferKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_CreateDeployToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeployTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).CreateDeployToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evm.v1beta1.MsgService/CreateDeployToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).CreateDeployToken(ctx, req.(*CreateDeployTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_CreateBurnTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBurnTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).CreateBurnTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evm.v1beta1.MsgService/CreateBurnTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).CreateBurnTokens(ctx, req.(*CreateBurnTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_SignTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).SignTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evm.v1beta1.MsgService/SignTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).SignTx(ctx, req.(*SignTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_CreatePendingTransfers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePendingTransfersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).CreatePendingTransfers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evm.v1beta1.MsgService/CreatePendingTransfers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).CreatePendingTransfers(ctx, req.(*CreatePendingTransfersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_CreateTransferOwnership_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransferOwnershipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).CreateTransferOwnership(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evm.v1beta1.MsgService/CreateTransferOwnership",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).CreateTransferOwnership(ctx, req.(*CreateTransferOwnershipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_CreateTransferOperatorship_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransferOperatorshipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).CreateTransferOperatorship(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evm.v1beta1.MsgService/CreateTransferOperatorship",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).CreateTransferOperatorship(ctx, req.(*CreateTransferOperatorshipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_SignCommands_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignCommandsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).SignCommands(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evm.v1beta1.MsgService/SignCommands",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).SignCommands(ctx, req.(*SignCommandsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_AddChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddChainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).AddChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evm.v1beta1.MsgService/AddChain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).AddChain(ctx, req.(*AddChainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "evm.v1beta1.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Link",
			Handler:    _MsgService_Link_Handler,
		},
		{
			MethodName: "ConfirmChain",
			Handler:    _MsgService_ConfirmChain_Handler,
		},
		{
			MethodName: "ConfirmToken",
			Handler:    _MsgService_ConfirmToken_Handler,
		},
		{
			MethodName: "ConfirmDeposit",
			Handler:    _MsgService_ConfirmDeposit_Handler,
		},
		{
			MethodName: "ConfirmTransferKey",
			Handler:    _MsgService_ConfirmTransferKey_Handler,
		},
		{
			MethodName: "VoteConfirmChain",
			Handler:    _MsgService_VoteConfirmChain_Handler,
		},
		{
			MethodName: "VoteConfirmDeposit",
			Handler:    _MsgService_VoteConfirmDeposit_Handler,
		},
		{
			MethodName: "VoteConfirmToken",
			Handler:    _MsgService_VoteConfirmToken_Handler,
		},
		{
			MethodName: "VoteConfirmTransferKey",
			Handler:    _MsgService_VoteConfirmTransferKey_Handler,
		},
		{
			MethodName: "CreateDeployToken",
			Handler:    _MsgService_CreateDeployToken_Handler,
		},
		{
			MethodName: "CreateBurnTokens",
			Handler:    _MsgService_CreateBurnTokens_Handler,
		},
		{
			MethodName: "SignTx",
			Handler:    _MsgService_SignTx_Handler,
		},
		{
			MethodName: "CreatePendingTransfers",
			Handler:    _MsgService_CreatePendingTransfers_Handler,
		},
		{
			MethodName: "CreateTransferOwnership",
			Handler:    _MsgService_CreateTransferOwnership_Handler,
		},
		{
			MethodName: "CreateTransferOperatorship",
			Handler:    _MsgService_CreateTransferOperatorship_Handler,
		},
		{
			MethodName: "SignCommands",
			Handler:    _MsgService_SignCommands_Handler,
		},
		{
			MethodName: "AddChain",
			Handler:    _MsgService_AddChain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "evm/v1beta1/service.proto",
}
