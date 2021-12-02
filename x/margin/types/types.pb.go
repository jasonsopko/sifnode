// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sifnode/margin/v1/types.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type GenesisState struct {
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3994728d56e8650, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

type MTP struct {
	Address          string                                  `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	CollateralAsset  string                                  `protobuf:"bytes,2,opt,name=collateral_asset,json=collateralAsset,proto3" json:"collateral_asset,omitempty"`
	CollateralAmount github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,3,opt,name=collateral_amount,json=collateralAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"collateral_amount"`
	LiabilitiesP     github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,4,opt,name=liabilities_p,json=liabilitiesP,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"liabilities_p"`
	LiabilitiesI     github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,5,opt,name=liabilities_i,json=liabilitiesI,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"liabilities_i"`
	CustodyAsset     string                                  `protobuf:"bytes,6,opt,name=custody_asset,json=custodyAsset,proto3" json:"custody_asset,omitempty"`
	CustodyAmount    github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,7,opt,name=custody_amount,json=custodyAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"custody_amount"`
	Leverage         github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,8,opt,name=leverage,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"leverage"`
	MtpHealth        github_com_cosmos_cosmos_sdk_types.Dec  `protobuf:"bytes,9,opt,name=mtp_health,json=mtpHealth,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"mtp_health"`
}

func (m *MTP) Reset()         { *m = MTP{} }
func (m *MTP) String() string { return proto.CompactTextString(m) }
func (*MTP) ProtoMessage()    {}
func (*MTP) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3994728d56e8650, []int{1}
}
func (m *MTP) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MTP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MTP.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MTP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MTP.Merge(m, src)
}
func (m *MTP) XXX_Size() int {
	return m.Size()
}
func (m *MTP) XXX_DiscardUnknown() {
	xxx_messageInfo_MTP.DiscardUnknown(m)
}

var xxx_messageInfo_MTP proto.InternalMessageInfo

func (m *MTP) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MTP) GetCollateralAsset() string {
	if m != nil {
		return m.CollateralAsset
	}
	return ""
}

func (m *MTP) GetCustodyAsset() string {
	if m != nil {
		return m.CustodyAsset
	}
	return ""
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "sifnode.margin.v1.GenesisState")
	proto.RegisterType((*MTP)(nil), "sifnode.margin.v1.MTP")
}

func init() { proto.RegisterFile("sifnode/margin/v1/types.proto", fileDescriptor_b3994728d56e8650) }

var fileDescriptor_b3994728d56e8650 = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xcd, 0x6a, 0xea, 0x40,
	0x1c, 0xc5, 0x93, 0xab, 0xd7, 0x8f, 0x41, 0xbd, 0xd7, 0xd0, 0xc5, 0x50, 0x68, 0x2c, 0x16, 0xfa,
	0x05, 0x4d, 0x90, 0x3e, 0x81, 0x22, 0xb4, 0xa5, 0x08, 0xa2, 0xb6, 0x8b, 0x52, 0x90, 0x31, 0x19,
	0x93, 0xa1, 0x49, 0x26, 0x64, 0x46, 0xa9, 0x6f, 0xd1, 0xc7, 0x72, 0xe9, 0xb2, 0x74, 0x21, 0x45,
	0xd7, 0x7d, 0x87, 0xe2, 0x64, 0x62, 0x43, 0x57, 0xc5, 0xae, 0x66, 0xe6, 0xff, 0x3f, 0xfc, 0x0e,
	0x73, 0x38, 0xe0, 0x80, 0x91, 0x71, 0x40, 0x6d, 0x6c, 0xfa, 0x28, 0x72, 0x48, 0x60, 0x4e, 0x1b,
	0x26, 0x9f, 0x85, 0x98, 0x19, 0x61, 0x44, 0x39, 0xd5, 0xaa, 0x72, 0x6d, 0xc4, 0x6b, 0x63, 0xda,
	0xd8, 0xdf, 0x73, 0xa8, 0x43, 0xc5, 0xd6, 0xdc, 0xdc, 0x62, 0x61, 0xbd, 0x02, 0x4a, 0x57, 0x38,
	0xc0, 0x8c, 0xb0, 0x3e, 0x47, 0x1c, 0xd7, 0x3f, 0xb2, 0x20, 0xd3, 0x19, 0x74, 0x35, 0x08, 0xf2,
	0xc8, 0xb6, 0x23, 0xcc, 0x18, 0x54, 0x0f, 0xd5, 0xd3, 0x62, 0x2f, 0x79, 0x6a, 0x67, 0xe0, 0xbf,
	0x45, 0x3d, 0x0f, 0x71, 0x1c, 0x21, 0x6f, 0x88, 0x18, 0xc3, 0x1c, 0xfe, 0x11, 0x92, 0x7f, 0x5f,
	0xf3, 0xe6, 0x66, 0xac, 0x3d, 0x82, 0x6a, 0x5a, 0xea, 0xd3, 0x49, 0xc0, 0x61, 0x66, 0xa3, 0x6d,
	0x99, 0xf3, 0x65, 0x4d, 0x79, 0x5b, 0xd6, 0x4e, 0x1c, 0xc2, 0xdd, 0xc9, 0xc8, 0xb0, 0xa8, 0x6f,
	0x5a, 0x94, 0xf9, 0x94, 0xc9, 0xe3, 0x82, 0xd9, 0x4f, 0xf2, 0x4b, 0x77, 0x24, 0xe0, 0xbd, 0x94,
	0x69, 0x53, 0x80, 0xb4, 0x01, 0x28, 0x7b, 0x04, 0x8d, 0x88, 0x47, 0x38, 0xc1, 0x6c, 0x18, 0xc2,
	0xec, 0x6e, 0xe4, 0x52, 0x8a, 0xd2, 0xfd, 0x4e, 0x25, 0xf0, 0xef, 0xef, 0xa9, 0x37, 0xda, 0x11,
	0x28, 0x5b, 0x13, 0xc6, 0xa9, 0x3d, 0x93, 0x89, 0xe5, 0x44, 0x62, 0x25, 0x39, 0x8c, 0xe3, 0xba,
	0x07, 0x95, 0xad, 0x28, 0xce, 0x2a, 0xbf, 0x9b, 0x77, 0xe2, 0x25, 0x83, 0xba, 0x05, 0x05, 0x0f,
	0x4f, 0x71, 0x84, 0x1c, 0x0c, 0x0b, 0xbb, 0x11, 0xb7, 0x00, 0xad, 0x03, 0x80, 0xcf, 0xc3, 0xa1,
	0x8b, 0x91, 0xc7, 0x5d, 0x58, 0x14, 0x38, 0x43, 0xe2, 0x8e, 0x7f, 0x80, 0x6b, 0x63, 0xab, 0x57,
	0xf4, 0x79, 0x78, 0x2d, 0x00, 0xad, 0xf6, 0x7c, 0xa5, 0xab, 0x8b, 0x95, 0xae, 0xbe, 0xaf, 0x74,
	0xf5, 0x65, 0xad, 0x2b, 0x8b, 0xb5, 0xae, 0xbc, 0xae, 0x75, 0xe5, 0xe1, 0x3c, 0x05, 0xeb, 0x93,
	0xb1, 0xe5, 0x22, 0x12, 0x98, 0x49, 0xeb, 0x9f, 0x93, 0xde, 0x0b, 0xe8, 0x28, 0x27, 0xca, 0x7c,
	0xf9, 0x19, 0x00, 0x00, 0xff, 0xff, 0xc1, 0x4e, 0x29, 0x1c, 0x16, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MTP) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MTP) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MTP) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MtpHealth.Size()
		i -= size
		if _, err := m.MtpHealth.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size := m.Leverage.Size()
		i -= size
		if _, err := m.Leverage.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.CustodyAmount.Size()
		i -= size
		if _, err := m.CustodyAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if len(m.CustodyAsset) > 0 {
		i -= len(m.CustodyAsset)
		copy(dAtA[i:], m.CustodyAsset)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.CustodyAsset)))
		i--
		dAtA[i] = 0x32
	}
	{
		size := m.LiabilitiesI.Size()
		i -= size
		if _, err := m.LiabilitiesI.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.LiabilitiesP.Size()
		i -= size
		if _, err := m.LiabilitiesP.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.CollateralAmount.Size()
		i -= size
		if _, err := m.CollateralAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.CollateralAsset) > 0 {
		i -= len(m.CollateralAsset)
		copy(dAtA[i:], m.CollateralAsset)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.CollateralAsset)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MTP) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.CollateralAsset)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.CollateralAmount.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.LiabilitiesP.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.LiabilitiesI.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = len(m.CustodyAsset)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.CustodyAmount.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.Leverage.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.MtpHealth.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *MTP) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: MTP: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MTP: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralAsset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollateralAsset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CollateralAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiabilitiesP", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LiabilitiesP.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiabilitiesI", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LiabilitiesI.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CustodyAsset", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CustodyAsset = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CustodyAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CustodyAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Leverage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Leverage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MtpHealth", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MtpHealth.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
