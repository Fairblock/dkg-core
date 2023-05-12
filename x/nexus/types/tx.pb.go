// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nexus/v1beta1/tx.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type RegisterChainMaintainerRequest struct {
	Sender github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=sender,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"sender,omitempty"`
	Chains []string                                      `protobuf:"bytes,2,rep,name=chains,proto3" json:"chains,omitempty"`
}

func (m *RegisterChainMaintainerRequest) Reset()         { *m = RegisterChainMaintainerRequest{} }
func (m *RegisterChainMaintainerRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterChainMaintainerRequest) ProtoMessage()    {}
func (*RegisterChainMaintainerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7af3d47209cda0b3, []int{0}
}
func (m *RegisterChainMaintainerRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterChainMaintainerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterChainMaintainerRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterChainMaintainerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterChainMaintainerRequest.Merge(m, src)
}
func (m *RegisterChainMaintainerRequest) XXX_Size() int {
	return m.Size()
}
func (m *RegisterChainMaintainerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterChainMaintainerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterChainMaintainerRequest proto.InternalMessageInfo

type RegisterChainMaintainerResponse struct {
}

func (m *RegisterChainMaintainerResponse) Reset()         { *m = RegisterChainMaintainerResponse{} }
func (m *RegisterChainMaintainerResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterChainMaintainerResponse) ProtoMessage()    {}
func (*RegisterChainMaintainerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7af3d47209cda0b3, []int{1}
}
func (m *RegisterChainMaintainerResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisterChainMaintainerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisterChainMaintainerResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisterChainMaintainerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterChainMaintainerResponse.Merge(m, src)
}
func (m *RegisterChainMaintainerResponse) XXX_Size() int {
	return m.Size()
}
func (m *RegisterChainMaintainerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterChainMaintainerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterChainMaintainerResponse proto.InternalMessageInfo

type DeregisterChainMaintainerRequest struct {
	Sender github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=sender,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"sender,omitempty"`
	Chains []string                                      `protobuf:"bytes,2,rep,name=chains,proto3" json:"chains,omitempty"`
}

func (m *DeregisterChainMaintainerRequest) Reset()         { *m = DeregisterChainMaintainerRequest{} }
func (m *DeregisterChainMaintainerRequest) String() string { return proto.CompactTextString(m) }
func (*DeregisterChainMaintainerRequest) ProtoMessage()    {}
func (*DeregisterChainMaintainerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7af3d47209cda0b3, []int{2}
}
func (m *DeregisterChainMaintainerRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeregisterChainMaintainerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeregisterChainMaintainerRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeregisterChainMaintainerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeregisterChainMaintainerRequest.Merge(m, src)
}
func (m *DeregisterChainMaintainerRequest) XXX_Size() int {
	return m.Size()
}
func (m *DeregisterChainMaintainerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeregisterChainMaintainerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeregisterChainMaintainerRequest proto.InternalMessageInfo

type DeregisterChainMaintainerResponse struct {
}

func (m *DeregisterChainMaintainerResponse) Reset()         { *m = DeregisterChainMaintainerResponse{} }
func (m *DeregisterChainMaintainerResponse) String() string { return proto.CompactTextString(m) }
func (*DeregisterChainMaintainerResponse) ProtoMessage()    {}
func (*DeregisterChainMaintainerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7af3d47209cda0b3, []int{3}
}
func (m *DeregisterChainMaintainerResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeregisterChainMaintainerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeregisterChainMaintainerResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeregisterChainMaintainerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeregisterChainMaintainerResponse.Merge(m, src)
}
func (m *DeregisterChainMaintainerResponse) XXX_Size() int {
	return m.Size()
}
func (m *DeregisterChainMaintainerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeregisterChainMaintainerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeregisterChainMaintainerResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RegisterChainMaintainerRequest)(nil), "nexus.v1beta1.RegisterChainMaintainerRequest")
	proto.RegisterType((*RegisterChainMaintainerResponse)(nil), "nexus.v1beta1.RegisterChainMaintainerResponse")
	proto.RegisterType((*DeregisterChainMaintainerRequest)(nil), "nexus.v1beta1.DeregisterChainMaintainerRequest")
	proto.RegisterType((*DeregisterChainMaintainerResponse)(nil), "nexus.v1beta1.DeregisterChainMaintainerResponse")
}

func init() { proto.RegisterFile("nexus/v1beta1/tx.proto", fileDescriptor_7af3d47209cda0b3) }

var fileDescriptor_7af3d47209cda0b3 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x91, 0xbf, 0x4a, 0x33, 0x41,
	0x14, 0xc5, 0x77, 0xbe, 0x0f, 0x02, 0x0e, 0xda, 0x04, 0x09, 0x21, 0xc8, 0xe4, 0x8f, 0x4d, 0x9a,
	0xec, 0x10, 0x7d, 0x82, 0x44, 0x1b, 0x0b, 0x41, 0xb6, 0xb4, 0x9b, 0xcc, 0x5e, 0x26, 0x43, 0x92,
	0xb9, 0xeb, 0xdc, 0x89, 0xae, 0xb5, 0xd8, 0xfb, 0x58, 0x29, 0x53, 0x5a, 0x89, 0x66, 0xdf, 0xc2,
	0x4a, 0xb2, 0xbb, 0x85, 0x8d, 0xb6, 0x56, 0x33, 0xe7, 0x9c, 0xcb, 0xbd, 0x3f, 0x38, 0xbc, 0xe5,
	0x20, 0x5f, 0x93, 0xbc, 0x1f, 0xcf, 0x20, 0xa8, 0xb1, 0x0c, 0x79, 0x9c, 0x79, 0x0c, 0xd8, 0x3c,
	0x2a, 0xfd, 0xb8, 0xf6, 0x3b, 0xc7, 0x06, 0x0d, 0x96, 0x89, 0xdc, 0xff, 0xaa, 0xa1, 0xce, 0x89,
	0x41, 0x34, 0x4b, 0x90, 0x2a, 0xb3, 0x52, 0x39, 0x87, 0x41, 0x05, 0x8b, 0x8e, 0xaa, 0x74, 0xf0,
	0xc4, 0xb8, 0x48, 0xc0, 0x58, 0x0a, 0xe0, 0x2f, 0xe6, 0xca, 0xba, 0x6b, 0x65, 0x5d, 0x50, 0xd6,
	0x81, 0x4f, 0xe0, 0x6e, 0x0d, 0x14, 0x9a, 0x57, 0xbc, 0x41, 0xe0, 0x52, 0xf0, 0x6d, 0xd6, 0x63,
	0xc3, 0xc3, 0xe9, 0xf8, 0xf3, 0xad, 0x3b, 0x32, 0x36, 0xcc, 0xd7, 0xb3, 0x58, 0xe3, 0x4a, 0x6a,
	0xa4, 0x15, 0x52, 0xfd, 0x8c, 0x28, 0x5d, 0xc8, 0xf0, 0x98, 0x01, 0xc5, 0x13, 0xad, 0x27, 0x69,
	0xea, 0x81, 0x28, 0xa9, 0x17, 0x34, 0x5b, 0xbc, 0xa1, 0xf7, 0x47, 0xa8, 0xfd, 0xaf, 0xf7, 0x7f,
	0x78, 0x90, 0xd4, 0x6a, 0xd0, 0xe7, 0xdd, 0x1f, 0x21, 0x28, 0x43, 0x47, 0x30, 0x78, 0x66, 0xbc,
	0x77, 0x09, 0xfe, 0xcf, 0x51, 0x4f, 0x79, 0xff, 0x17, 0x8c, 0x0a, 0x76, 0x7a, 0xb3, 0xf9, 0x10,
	0xd1, 0x66, 0x27, 0xd8, 0x76, 0x27, 0xd8, 0xfb, 0x4e, 0xb0, 0x97, 0x42, 0x44, 0xdb, 0x42, 0x44,
	0xaf, 0x85, 0x88, 0x6e, 0xcf, 0xbe, 0x11, 0xa9, 0x1c, 0x96, 0xca, 0x3b, 0x08, 0x0f, 0xe8, 0x17,
	0xb5, 0x1a, 0x69, 0xf4, 0x20, 0x73, 0x59, 0xb5, 0x5e, 0x12, 0xce, 0x1a, 0x65, 0x5d, 0xe7, 0x5f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xbf, 0x46, 0xd6, 0x4a, 0x0b, 0x02, 0x00, 0x00,
}

func (m *RegisterChainMaintainerRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterChainMaintainerRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterChainMaintainerRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Chains) > 0 {
		for iNdEx := len(m.Chains) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Chains[iNdEx])
			copy(dAtA[i:], m.Chains[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Chains[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RegisterChainMaintainerResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisterChainMaintainerResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisterChainMaintainerResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *DeregisterChainMaintainerRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeregisterChainMaintainerRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeregisterChainMaintainerRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Chains) > 0 {
		for iNdEx := len(m.Chains) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Chains[iNdEx])
			copy(dAtA[i:], m.Chains[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Chains[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DeregisterChainMaintainerResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeregisterChainMaintainerResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeregisterChainMaintainerResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RegisterChainMaintainerRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Chains) > 0 {
		for _, s := range m.Chains {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *RegisterChainMaintainerResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *DeregisterChainMaintainerRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Chains) > 0 {
		for _, s := range m.Chains {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *DeregisterChainMaintainerResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RegisterChainMaintainerRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: RegisterChainMaintainerRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterChainMaintainerRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = append(m.Sender[:0], dAtA[iNdEx:postIndex]...)
			if m.Sender == nil {
				m.Sender = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chains", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chains = append(m.Chains, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *RegisterChainMaintainerResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: RegisterChainMaintainerResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisterChainMaintainerResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *DeregisterChainMaintainerRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: DeregisterChainMaintainerRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeregisterChainMaintainerRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = append(m.Sender[:0], dAtA[iNdEx:postIndex]...)
			if m.Sender == nil {
				m.Sender = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chains", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chains = append(m.Chains, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *DeregisterChainMaintainerResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: DeregisterChainMaintainerResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeregisterChainMaintainerResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
