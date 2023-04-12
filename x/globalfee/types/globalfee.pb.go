// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stargaze/globalfee/v1/globalfee.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

// Params holds parameters for the globalfee module.
type Params struct {
	// Addresses which are whitelisted to modify the gas free operations
	PrivilegedAddress []string `protobuf:"bytes,1,rep,name=privileged_address,json=privilegedAddress,proto3" json:"privileged_address,omitempty"`
	// Minimum stores the minimum gas price(s) for all TX on the chain.
	MinimumGasPrices github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,2,rep,name=minimum_gas_prices,json=minimumGasPrices,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"minimum_gas_prices,omitempty" yaml:"minimum_gas_prices"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_33ee3d8d8859c0c7, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetPrivilegedAddress() []string {
	if m != nil {
		return m.PrivilegedAddress
	}
	return nil
}

func (m *Params) GetMinimumGasPrices() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.MinimumGasPrices
	}
	return nil
}

// Configuration for code Ids which can have zero gas operations
type CodeAuthorization struct {
	// authorized code ids
	CodeId uint64 `protobuf:"varint,1,opt,name=code_id,json=codeId,proto3" json:"code_id,omitempty"`
	// authorized contract operation methods
	Methods []string `protobuf:"bytes,2,rep,name=methods,proto3" json:"methods,omitempty"`
}

func (m *CodeAuthorization) Reset()         { *m = CodeAuthorization{} }
func (m *CodeAuthorization) String() string { return proto.CompactTextString(m) }
func (*CodeAuthorization) ProtoMessage()    {}
func (*CodeAuthorization) Descriptor() ([]byte, []int) {
	return fileDescriptor_33ee3d8d8859c0c7, []int{1}
}
func (m *CodeAuthorization) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CodeAuthorization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CodeAuthorization.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CodeAuthorization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeAuthorization.Merge(m, src)
}
func (m *CodeAuthorization) XXX_Size() int {
	return m.Size()
}
func (m *CodeAuthorization) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeAuthorization.DiscardUnknown(m)
}

var xxx_messageInfo_CodeAuthorization proto.InternalMessageInfo

func (m *CodeAuthorization) GetCodeId() uint64 {
	if m != nil {
		return m.CodeId
	}
	return 0
}

func (m *CodeAuthorization) GetMethods() []string {
	if m != nil {
		return m.Methods
	}
	return nil
}

// Configuration for contract addresses which can have zero gas operations
type ContractAuthorization struct {
	// authorized contract addresses
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	// authorized contract operation methods
	Methods []string `protobuf:"bytes,2,rep,name=methods,proto3" json:"methods,omitempty"`
}

func (m *ContractAuthorization) Reset()         { *m = ContractAuthorization{} }
func (m *ContractAuthorization) String() string { return proto.CompactTextString(m) }
func (*ContractAuthorization) ProtoMessage()    {}
func (*ContractAuthorization) Descriptor() ([]byte, []int) {
	return fileDescriptor_33ee3d8d8859c0c7, []int{2}
}
func (m *ContractAuthorization) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContractAuthorization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContractAuthorization.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContractAuthorization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractAuthorization.Merge(m, src)
}
func (m *ContractAuthorization) XXX_Size() int {
	return m.Size()
}
func (m *ContractAuthorization) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractAuthorization.DiscardUnknown(m)
}

var xxx_messageInfo_ContractAuthorization proto.InternalMessageInfo

func (m *ContractAuthorization) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *ContractAuthorization) GetMethods() []string {
	if m != nil {
		return m.Methods
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "publicawesome.stargaze.globalfee.v1.Params")
	proto.RegisterType((*CodeAuthorization)(nil), "publicawesome.stargaze.globalfee.v1.CodeAuthorization")
	proto.RegisterType((*ContractAuthorization)(nil), "publicawesome.stargaze.globalfee.v1.ContractAuthorization")
}

func init() {
	proto.RegisterFile("stargaze/globalfee/v1/globalfee.proto", fileDescriptor_33ee3d8d8859c0c7)
}

var fileDescriptor_33ee3d8d8859c0c7 = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xf6, 0xd1, 0x2a, 0x55, 0x8e, 0x81, 0xd6, 0x02, 0x11, 0xaa, 0xca, 0x8e, 0x8c, 0x90, 0x82,
	0x20, 0x3e, 0x05, 0x06, 0x44, 0xb7, 0x36, 0x08, 0xc4, 0x56, 0x3c, 0x22, 0xa4, 0xe8, 0x7c, 0xf7,
	0x70, 0x4e, 0xf8, 0xfc, 0xac, 0xbb, 0x4b, 0x20, 0xfd, 0x15, 0x8c, 0x8c, 0xcc, 0xfc, 0x06, 0x7e,
	0x40, 0xc7, 0x8e, 0x4c, 0x01, 0x25, 0x1b, 0x23, 0x33, 0x03, 0x4a, 0x6c, 0x93, 0xa0, 0x8a, 0xe9,
	0xee, 0xdd, 0xf7, 0xde, 0xf7, 0xe9, 0x7d, 0xf7, 0xd1, 0x7b, 0xd6, 0x71, 0x93, 0xf1, 0x73, 0x60,
	0x59, 0x8e, 0x29, 0xcf, 0xdf, 0x02, 0xb0, 0xe9, 0x60, 0x53, 0xc4, 0xa5, 0x41, 0x87, 0xfe, 0xdd,
	0x72, 0x92, 0xe6, 0x4a, 0xf0, 0xf7, 0x60, 0x51, 0x43, 0xdc, 0x0c, 0xc5, 0x9b, 0xbe, 0xe9, 0xe0,
	0xf0, 0x66, 0x86, 0x19, 0xae, 0xfb, 0xd9, 0xea, 0x56, 0x8d, 0x1e, 0x06, 0x02, 0xad, 0x46, 0xcb,
	0x52, 0x6e, 0x57, 0xd4, 0x29, 0x38, 0x3e, 0x60, 0x02, 0x55, 0x51, 0xe1, 0xd1, 0x6f, 0x42, 0x5b,
	0x67, 0xdc, 0x70, 0x6d, 0xfd, 0x3e, 0xf5, 0x4b, 0xa3, 0xa6, 0x2a, 0x87, 0x0c, 0xe4, 0x88, 0x4b,
	0x69, 0xc0, 0xda, 0x0e, 0xe9, 0xee, 0xf4, 0xda, 0xc9, 0xc1, 0x06, 0x39, 0xa9, 0x00, 0xff, 0x2b,
	0xa1, 0xbe, 0x56, 0x85, 0xd2, 0x13, 0x3d, 0xca, 0xb8, 0x1d, 0x95, 0x46, 0x09, 0xb0, 0x9d, 0x6b,
	0xdd, 0x9d, 0xde, 0xf5, 0x47, 0x47, 0x71, 0xa5, 0x1b, 0xaf, 0x74, 0xe3, 0x5a, 0x37, 0x7e, 0x06,
	0x62, 0x88, 0xaa, 0x38, 0x2d, 0x2f, 0xe6, 0xa1, 0xf7, 0x73, 0x1e, 0x1e, 0x5d, 0x9d, 0x7f, 0x88,
	0x5a, 0x39, 0xd0, 0xa5, 0x9b, 0xfd, 0x9a, 0x87, 0x77, 0x66, 0x5c, 0xe7, 0xc7, 0xd1, 0xd5, 0xae,
	0xe8, 0xcb, 0xf7, 0xf0, 0x41, 0xa6, 0xdc, 0x78, 0x92, 0xc6, 0x02, 0x35, 0xab, 0x97, 0xac, 0x8e,
	0xbe, 0x95, 0xef, 0x98, 0x9b, 0x95, 0x60, 0x1b, 0x41, 0x9b, 0xec, 0xd7, 0x1c, 0x2f, 0xb8, 0x3d,
	0x5b, 0x33, 0x1c, 0xef, 0x7e, 0xfa, 0x1c, 0x7a, 0xd1, 0x73, 0x7a, 0x30, 0x44, 0x09, 0x27, 0x13,
	0x37, 0x46, 0xa3, 0xce, 0xb9, 0x53, 0x58, 0xf8, 0xb7, 0xe9, 0x9e, 0x40, 0x09, 0x23, 0x25, 0x3b,
	0xa4, 0x4b, 0x7a, 0xbb, 0x49, 0x6b, 0x55, 0xbe, 0x94, 0x7e, 0x87, 0xee, 0x69, 0x70, 0x63, 0x94,
	0xd5, 0x9a, 0xed, 0xa4, 0x29, 0xa3, 0x37, 0xf4, 0xd6, 0x10, 0x0b, 0x67, 0xb8, 0x70, 0xff, 0x72,
	0xdd, 0xa7, 0xfb, 0xa2, 0x06, 0xb6, 0x2c, 0x25, 0xbd, 0x76, 0x72, 0xa3, 0x79, 0x6f, 0x0c, 0xfd,
	0x2f, 0xfb, 0xe9, 0xab, 0x8b, 0x45, 0x40, 0x2e, 0x17, 0x01, 0xf9, 0xb1, 0x08, 0xc8, 0xc7, 0x65,
	0xe0, 0x5d, 0x2e, 0x03, 0xef, 0xdb, 0x32, 0xf0, 0x5e, 0x3f, 0xd9, 0x32, 0xa1, 0x0a, 0x49, 0xbf,
	0x4e, 0x09, 0xfb, 0x1b, 0xad, 0xe9, 0x53, 0xf6, 0x61, 0x2b, 0x5f, 0x6b, 0x67, 0xd2, 0xd6, 0xfa,
	0xfb, 0x1f, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x17, 0x65, 0x2b, 0x1b, 0x82, 0x02, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MinimumGasPrices) > 0 {
		for iNdEx := len(m.MinimumGasPrices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MinimumGasPrices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGlobalfee(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.PrivilegedAddress) > 0 {
		for iNdEx := len(m.PrivilegedAddress) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PrivilegedAddress[iNdEx])
			copy(dAtA[i:], m.PrivilegedAddress[iNdEx])
			i = encodeVarintGlobalfee(dAtA, i, uint64(len(m.PrivilegedAddress[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *CodeAuthorization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CodeAuthorization) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CodeAuthorization) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Methods) > 0 {
		for iNdEx := len(m.Methods) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Methods[iNdEx])
			copy(dAtA[i:], m.Methods[iNdEx])
			i = encodeVarintGlobalfee(dAtA, i, uint64(len(m.Methods[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.CodeId != 0 {
		i = encodeVarintGlobalfee(dAtA, i, uint64(m.CodeId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ContractAuthorization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractAuthorization) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContractAuthorization) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Methods) > 0 {
		for iNdEx := len(m.Methods) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Methods[iNdEx])
			copy(dAtA[i:], m.Methods[iNdEx])
			i = encodeVarintGlobalfee(dAtA, i, uint64(len(m.Methods[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintGlobalfee(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGlobalfee(dAtA []byte, offset int, v uint64) int {
	offset -= sovGlobalfee(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PrivilegedAddress) > 0 {
		for _, s := range m.PrivilegedAddress {
			l = len(s)
			n += 1 + l + sovGlobalfee(uint64(l))
		}
	}
	if len(m.MinimumGasPrices) > 0 {
		for _, e := range m.MinimumGasPrices {
			l = e.Size()
			n += 1 + l + sovGlobalfee(uint64(l))
		}
	}
	return n
}

func (m *CodeAuthorization) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CodeId != 0 {
		n += 1 + sovGlobalfee(uint64(m.CodeId))
	}
	if len(m.Methods) > 0 {
		for _, s := range m.Methods {
			l = len(s)
			n += 1 + l + sovGlobalfee(uint64(l))
		}
	}
	return n
}

func (m *ContractAuthorization) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovGlobalfee(uint64(l))
	}
	if len(m.Methods) > 0 {
		for _, s := range m.Methods {
			l = len(s)
			n += 1 + l + sovGlobalfee(uint64(l))
		}
	}
	return n
}

func sovGlobalfee(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGlobalfee(x uint64) (n int) {
	return sovGlobalfee(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGlobalfee
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrivilegedAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGlobalfee
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
				return ErrInvalidLengthGlobalfee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGlobalfee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrivilegedAddress = append(m.PrivilegedAddress, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinimumGasPrices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGlobalfee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGlobalfee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGlobalfee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinimumGasPrices = append(m.MinimumGasPrices, types.DecCoin{})
			if err := m.MinimumGasPrices[len(m.MinimumGasPrices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGlobalfee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGlobalfee
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
func (m *CodeAuthorization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGlobalfee
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
			return fmt.Errorf("proto: CodeAuthorization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CodeAuthorization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeId", wireType)
			}
			m.CodeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGlobalfee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CodeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Methods", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGlobalfee
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
				return ErrInvalidLengthGlobalfee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGlobalfee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Methods = append(m.Methods, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGlobalfee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGlobalfee
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
func (m *ContractAuthorization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGlobalfee
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
			return fmt.Errorf("proto: ContractAuthorization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractAuthorization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGlobalfee
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
				return ErrInvalidLengthGlobalfee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGlobalfee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Methods", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGlobalfee
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
				return ErrInvalidLengthGlobalfee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGlobalfee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Methods = append(m.Methods, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGlobalfee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGlobalfee
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
func skipGlobalfee(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGlobalfee
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
					return 0, ErrIntOverflowGlobalfee
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
					return 0, ErrIntOverflowGlobalfee
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
				return 0, ErrInvalidLengthGlobalfee
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGlobalfee
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGlobalfee
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGlobalfee        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGlobalfee          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGlobalfee = fmt.Errorf("proto: unexpected end of group")
)
