// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: evm/v1beta1/query.proto

package types

import (
	fmt "fmt"
	github_com_axelarnetwork_axelar_core_x_tss_exported "github.com/axelarnetwork/axelar-core/x/tss/exported"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// DepositQueryParams describe the parameters used to query for an EVM
// deposit address
type DepositQueryParams struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Asset   string `protobuf:"bytes,2,opt,name=asset,proto3" json:"asset,omitempty"`
	Chain   string `protobuf:"bytes,3,opt,name=chain,proto3" json:"chain,omitempty"`
}

func (m *DepositQueryParams) Reset()         { *m = DepositQueryParams{} }
func (m *DepositQueryParams) String() string { return proto.CompactTextString(m) }
func (*DepositQueryParams) ProtoMessage()    {}
func (*DepositQueryParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_78a1f61c7ae3396c, []int{0}
}
func (m *DepositQueryParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DepositQueryParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DepositQueryParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DepositQueryParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DepositQueryParams.Merge(m, src)
}
func (m *DepositQueryParams) XXX_Size() int {
	return m.Size()
}
func (m *DepositQueryParams) XXX_DiscardUnknown() {
	xxx_messageInfo_DepositQueryParams.DiscardUnknown(m)
}

var xxx_messageInfo_DepositQueryParams proto.InternalMessageInfo

type QueryBatchedCommandsResponse struct {
	ID                    string                                                    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Data                  string                                                    `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Status                BatchedCommandsStatus                                     `protobuf:"varint,3,opt,name=status,proto3,enum=evm.v1beta1.BatchedCommandsStatus" json:"status,omitempty"`
	KeyID                 github_com_axelarnetwork_axelar_core_x_tss_exported.KeyID `protobuf:"bytes,4,opt,name=key_id,json=keyId,proto3,casttype=github.com/axelarnetwork/axelar-core/x/tss/exported.KeyID" json:"key_id,omitempty"`
	Signature             string                                                    `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	ExecuteData           string                                                    `protobuf:"bytes,6,opt,name=execute_data,json=executeData,proto3" json:"execute_data,omitempty"`
	PrevBatchedCommandsID string                                                    `protobuf:"bytes,7,opt,name=prev_batched_commands_id,json=prevBatchedCommandsId,proto3" json:"prev_batched_commands_id,omitempty"`
}

func (m *QueryBatchedCommandsResponse) Reset()         { *m = QueryBatchedCommandsResponse{} }
func (m *QueryBatchedCommandsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryBatchedCommandsResponse) ProtoMessage()    {}
func (*QueryBatchedCommandsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_78a1f61c7ae3396c, []int{1}
}
func (m *QueryBatchedCommandsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryBatchedCommandsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryBatchedCommandsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryBatchedCommandsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryBatchedCommandsResponse.Merge(m, src)
}
func (m *QueryBatchedCommandsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryBatchedCommandsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryBatchedCommandsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryBatchedCommandsResponse proto.InternalMessageInfo

type QueryAddressResponse struct {
	Address string                                                    `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	KeyID   github_com_axelarnetwork_axelar_core_x_tss_exported.KeyID `protobuf:"bytes,2,opt,name=key_id,json=keyId,proto3,casttype=github.com/axelarnetwork/axelar-core/x/tss/exported.KeyID" json:"key_id,omitempty"`
}

func (m *QueryAddressResponse) Reset()         { *m = QueryAddressResponse{} }
func (m *QueryAddressResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAddressResponse) ProtoMessage()    {}
func (*QueryAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_78a1f61c7ae3396c, []int{2}
}
func (m *QueryAddressResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAddressResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAddressResponse.Merge(m, src)
}
func (m *QueryAddressResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAddressResponse proto.InternalMessageInfo

type QueryDepositStateParams struct {
	TxID          Hash    `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3,customtype=Hash" json:"tx_id"`
	BurnerAddress Address `protobuf:"bytes,2,opt,name=burner_address,json=burnerAddress,proto3,customtype=Address" json:"burner_address"`
	Amount        uint64  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *QueryDepositStateParams) Reset()         { *m = QueryDepositStateParams{} }
func (m *QueryDepositStateParams) String() string { return proto.CompactTextString(m) }
func (*QueryDepositStateParams) ProtoMessage()    {}
func (*QueryDepositStateParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_78a1f61c7ae3396c, []int{3}
}
func (m *QueryDepositStateParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDepositStateParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDepositStateParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDepositStateParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDepositStateParams.Merge(m, src)
}
func (m *QueryDepositStateParams) XXX_Size() int {
	return m.Size()
}
func (m *QueryDepositStateParams) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDepositStateParams.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDepositStateParams proto.InternalMessageInfo

type QueryDepositStateResponse struct {
	Log    string        `protobuf:"bytes,1,opt,name=log,proto3" json:"log,omitempty"`
	Status DepositStatus `protobuf:"varint,2,opt,name=status,proto3,enum=evm.v1beta1.DepositStatus" json:"status,omitempty"`
}

func (m *QueryDepositStateResponse) Reset()         { *m = QueryDepositStateResponse{} }
func (m *QueryDepositStateResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDepositStateResponse) ProtoMessage()    {}
func (*QueryDepositStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_78a1f61c7ae3396c, []int{4}
}
func (m *QueryDepositStateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDepositStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDepositStateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDepositStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDepositStateResponse.Merge(m, src)
}
func (m *QueryDepositStateResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDepositStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDepositStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDepositStateResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*DepositQueryParams)(nil), "evm.v1beta1.DepositQueryParams")
	proto.RegisterType((*QueryBatchedCommandsResponse)(nil), "evm.v1beta1.QueryBatchedCommandsResponse")
	proto.RegisterType((*QueryAddressResponse)(nil), "evm.v1beta1.QueryAddressResponse")
	proto.RegisterType((*QueryDepositStateParams)(nil), "evm.v1beta1.QueryDepositStateParams")
	proto.RegisterType((*QueryDepositStateResponse)(nil), "evm.v1beta1.QueryDepositStateResponse")
}

func init() { proto.RegisterFile("evm/v1beta1/query.proto", fileDescriptor_78a1f61c7ae3396c) }

var fileDescriptor_78a1f61c7ae3396c = []byte{
	// 563 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0xdd, 0xa4, 0xbb, 0x5b, 0xd5, 0x2d, 0x05, 0x59, 0xdb, 0x36, 0xad, 0xaa, 0x04, 0xf6, 0x04,
	0x07, 0x12, 0x5a, 0x24, 0x24, 0xb8, 0xb1, 0x44, 0x88, 0x08, 0x09, 0x95, 0xc0, 0xa9, 0x12, 0x8a,
	0xbc, 0xf1, 0x68, 0x37, 0x6a, 0x13, 0x07, 0xdb, 0x59, 0xb2, 0x5f, 0x01, 0x07, 0x3e, 0xaa, 0xc7,
	0x1e, 0x11, 0x87, 0x08, 0xd2, 0xbf, 0xe8, 0x09, 0xc5, 0x71, 0xbb, 0x6d, 0x01, 0x89, 0x0b, 0x37,
	0xcf, 0xf3, 0x8c, 0x67, 0xde, 0xf3, 0x1b, 0xb4, 0x05, 0xb3, 0xd4, 0x9b, 0xed, 0x8d, 0x41, 0x92,
	0x3d, 0xef, 0x63, 0x01, 0x7c, 0xee, 0xe6, 0x9c, 0x49, 0x86, 0x57, 0x61, 0x96, 0xba, 0xfa, 0x62,
	0x67, 0x30, 0x61, 0x13, 0xa6, 0x70, 0xaf, 0x39, 0xb5, 0x29, 0x3b, 0xd7, 0x6a, 0xe5, 0x3c, 0x07,
	0xd1, 0x5e, 0x0c, 0x0f, 0x11, 0xf6, 0x21, 0x67, 0x22, 0x91, 0x6f, 0x9b, 0x17, 0x0f, 0x08, 0x27,
	0xa9, 0xc0, 0x16, 0x5a, 0x26, 0x94, 0x72, 0x10, 0xc2, 0x32, 0xee, 0x1a, 0xf7, 0x57, 0xc2, 0x8b,
	0x10, 0x0f, 0x50, 0x8f, 0x08, 0x01, 0xd2, 0x32, 0x15, 0xde, 0x06, 0x0d, 0x1a, 0x4f, 0x49, 0x92,
	0x59, 0x4b, 0x2d, 0xaa, 0x82, 0xe1, 0xb9, 0x89, 0x76, 0xd5, 0xab, 0x23, 0x22, 0xe3, 0x29, 0xd0,
	0x17, 0x2c, 0x4d, 0x49, 0x46, 0x45, 0x08, 0x22, 0x67, 0x99, 0x00, 0xbc, 0x89, 0xcc, 0x84, 0xb6,
	0x1d, 0x46, 0xfd, 0xba, 0x72, 0xcc, 0xc0, 0x0f, 0xcd, 0x84, 0x62, 0x8c, 0xba, 0x94, 0x48, 0xa2,
	0x7b, 0xa8, 0x33, 0x7e, 0x86, 0xfa, 0x42, 0x12, 0x59, 0x08, 0xd5, 0x63, 0x7d, 0x7f, 0xe8, 0x5e,
	0x61, 0xed, 0xde, 0xe8, 0xf0, 0x4e, 0x65, 0x86, 0xba, 0x02, 0x7f, 0x40, 0xfd, 0x23, 0x98, 0x47,
	0x09, 0xb5, 0xba, 0xaa, 0xd7, 0xcb, 0xba, 0x72, 0x7a, 0xaf, 0x61, 0x1e, 0xf8, 0xe7, 0x95, 0xf3,
	0x74, 0x92, 0xc8, 0x69, 0x31, 0x76, 0x63, 0x96, 0x7a, 0xa4, 0x84, 0x63, 0xc2, 0x33, 0x90, 0x9f,
	0x18, 0x3f, 0xd2, 0xd1, 0xc3, 0x98, 0x71, 0xf0, 0x4a, 0x4f, 0x0a, 0xe1, 0x41, 0x99, 0x33, 0x2e,
	0x81, 0xba, 0xaa, 0x38, 0xec, 0x1d, 0xc1, 0x3c, 0xa0, 0x78, 0x17, 0xad, 0x88, 0x64, 0x92, 0x11,
	0x59, 0x70, 0xb0, 0x7a, 0x6a, 0xe6, 0x05, 0x80, 0xef, 0xa1, 0x35, 0x28, 0x21, 0x2e, 0x24, 0x44,
	0x8a, 0x54, 0x5f, 0x25, 0xac, 0x6a, 0xcc, 0x6f, 0xb8, 0x85, 0xc8, 0xca, 0x39, 0xcc, 0xa2, 0x71,
	0xcb, 0x22, 0x8a, 0x35, 0x8d, 0x66, 0xe2, 0x65, 0x35, 0xf1, 0x76, 0x5d, 0x39, 0x1b, 0x07, 0x1c,
	0x66, 0x37, 0x88, 0x06, 0x7e, 0xb8, 0x91, 0xff, 0x01, 0xa6, 0xc3, 0xcf, 0x06, 0x1a, 0x28, 0xf1,
	0x9f, 0xb7, 0x3f, 0x77, 0x29, 0xfa, 0xdf, 0xff, 0x76, 0x21, 0x93, 0xf9, 0x1f, 0x64, 0x1a, 0x7e,
	0x35, 0xd0, 0x96, 0x9a, 0x48, 0x1b, 0xae, 0xf9, 0x24, 0xd0, 0x86, 0x7b, 0x80, 0x7a, 0xb2, 0x8c,
	0xb4, 0x19, 0xd6, 0x46, 0x83, 0x93, 0xca, 0xe9, 0x7c, 0xaf, 0x9c, 0xee, 0x2b, 0x22, 0xa6, 0x75,
	0xe5, 0x74, 0xdf, 0x97, 0x81, 0x1f, 0x76, 0x65, 0x19, 0x50, 0xfc, 0x04, 0xad, 0x8f, 0x0b, 0x9e,
	0x01, 0x8f, 0x2e, 0x68, 0x98, 0xaa, 0xe6, 0xb6, 0xae, 0x59, 0xbe, 0x20, 0x7c, 0xab, 0x4d, 0xd3,
	0x21, 0xde, 0x44, 0x7d, 0x92, 0xb2, 0x22, 0x93, 0xca, 0x40, 0xdd, 0x50, 0x47, 0x43, 0x82, 0xb6,
	0x7f, 0x9b, 0xea, 0x52, 0xac, 0x3b, 0x68, 0xe9, 0x98, 0x4d, 0xb4, 0x50, 0xcd, 0x11, 0xef, 0x5f,
	0xfa, 0xd0, 0x54, 0x3e, 0xdc, 0xb9, 0xe6, 0xc3, 0x2b, 0x8f, 0x2c, 0xfc, 0x37, 0x7a, 0x73, 0xf2,
	0xd3, 0xee, 0x9c, 0xd4, 0xb6, 0x71, 0x5a, 0xdb, 0xc6, 0x8f, 0xda, 0x36, 0xbe, 0x9c, 0xd9, 0x9d,
	0xd3, 0x33, 0xbb, 0xf3, 0xed, 0xcc, 0xee, 0x1c, 0x3e, 0xfa, 0x47, 0x65, 0x9b, 0x15, 0x56, 0xab,
	0x3b, 0xee, 0xab, 0xdd, 0x7d, 0xfc, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x07, 0xa0, 0x9f, 0x40, 0x12,
	0x04, 0x00, 0x00,
}

func (m *DepositQueryParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DepositQueryParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DepositQueryParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Asset) > 0 {
		i -= len(m.Asset)
		copy(dAtA[i:], m.Asset)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Asset)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryBatchedCommandsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryBatchedCommandsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryBatchedCommandsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PrevBatchedCommandsID) > 0 {
		i -= len(m.PrevBatchedCommandsID)
		copy(dAtA[i:], m.PrevBatchedCommandsID)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.PrevBatchedCommandsID)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.ExecuteData) > 0 {
		i -= len(m.ExecuteData)
		copy(dAtA[i:], m.ExecuteData)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ExecuteData)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.KeyID) > 0 {
		i -= len(m.KeyID)
		copy(dAtA[i:], m.KeyID)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.KeyID)))
		i--
		dAtA[i] = 0x22
	}
	if m.Status != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAddressResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAddressResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAddressResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.KeyID) > 0 {
		i -= len(m.KeyID)
		copy(dAtA[i:], m.KeyID)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.KeyID)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDepositStateParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDepositStateParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDepositStateParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Amount != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.BurnerAddress.Size()
		i -= size
		if _, err := m.BurnerAddress.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.TxID.Size()
		i -= size
		if _, err := m.TxID.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryDepositStateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDepositStateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDepositStateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Log) > 0 {
		i -= len(m.Log)
		copy(dAtA[i:], m.Log)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Log)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DepositQueryParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Asset)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryBatchedCommandsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovQuery(uint64(m.Status))
	}
	l = len(m.KeyID)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.ExecuteData)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.PrevBatchedCommandsID)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAddressResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.KeyID)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDepositStateParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TxID.Size()
	n += 1 + l + sovQuery(uint64(l))
	l = m.BurnerAddress.Size()
	n += 1 + l + sovQuery(uint64(l))
	if m.Amount != 0 {
		n += 1 + sovQuery(uint64(m.Amount))
	}
	return n
}

func (m *QueryDepositStateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Log)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovQuery(uint64(m.Status))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DepositQueryParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DepositQueryParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DepositQueryParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Asset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryBatchedCommandsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryBatchedCommandsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryBatchedCommandsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= BatchedCommandsStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyID = github_com_axelarnetwork_axelar_core_x_tss_exported.KeyID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExecuteData", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExecuteData = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevBatchedCommandsID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrevBatchedCommandsID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAddressResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAddressResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAddressResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyID = github_com_axelarnetwork_axelar_core_x_tss_exported.KeyID(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryDepositStateParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryDepositStateParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDepositStateParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxID", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TxID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnerAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.BurnerAddress.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryDepositStateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryDepositStateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDepositStateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Log", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Log = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= DepositStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
