// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sifnode/oracle/v1/genesis.proto

package types

import (
	fmt "fmt"
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
	AddressWhitelist []string      `protobuf:"bytes,1,rep,name=address_whitelist,json=addressWhitelist,proto3" json:"address_whitelist,omitempty"`
	AdminAddress     string        `protobuf:"bytes,2,opt,name=admin_address,json=adminAddress,proto3" json:"admin_address,omitempty"`
	Prophecies       []*DBProphecy `protobuf:"bytes,3,rep,name=prophecies,proto3" json:"prophecies,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_1eb098560757ad78, []int{0}
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

func (m *GenesisState) GetAddressWhitelist() []string {
	if m != nil {
		return m.AddressWhitelist
	}
	return nil
}

func (m *GenesisState) GetAdminAddress() string {
	if m != nil {
		return m.AdminAddress
	}
	return ""
}

func (m *GenesisState) GetProphecies() []*DBProphecy {
	if m != nil {
		return m.Prophecies
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "sifnode.oracle.v1.GenesisState")
}

func init() { proto.RegisterFile("sifnode/oracle/v1/genesis.proto", fileDescriptor_1eb098560757ad78) }

var fileDescriptor_1eb098560757ad78 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0xce, 0x4c, 0xcb,
	0xcb, 0x4f, 0x49, 0xd5, 0xcf, 0x2f, 0x4a, 0x4c, 0xce, 0x49, 0xd5, 0x2f, 0x33, 0xd4, 0x4f, 0x4f,
	0xcd, 0x4b, 0x2d, 0xce, 0x2c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x84, 0x2a, 0xd0,
	0x83, 0x28, 0xd0, 0x2b, 0x33, 0x94, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xcb, 0xea, 0x83, 0x58,
	0x10, 0x85, 0x52, 0xb2, 0x98, 0x26, 0x95, 0x54, 0x16, 0xa4, 0x42, 0xcd, 0x51, 0x9a, 0xcf, 0xc8,
	0xc5, 0xe3, 0x0e, 0x31, 0x39, 0xb8, 0x24, 0xb1, 0x24, 0x55, 0x48, 0x9b, 0x4b, 0x30, 0x31, 0x25,
	0xa5, 0x28, 0xb5, 0xb8, 0x38, 0xbe, 0x3c, 0x23, 0xb3, 0x24, 0x35, 0x27, 0xb3, 0xb8, 0x44, 0x82,
	0x51, 0x81, 0x59, 0x83, 0x33, 0x48, 0x00, 0x2a, 0x11, 0x0e, 0x13, 0x17, 0x52, 0xe6, 0xe2, 0x4d,
	0x4c, 0xc9, 0xcd, 0xcc, 0x8b, 0x87, 0xca, 0x48, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0xf1, 0x80,
	0x05, 0x1d, 0x21, 0x62, 0x42, 0xb6, 0x5c, 0x5c, 0x05, 0x45, 0xf9, 0x05, 0x19, 0xa9, 0xc9, 0x99,
	0xa9, 0xc5, 0x12, 0xcc, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0xb2, 0x7a, 0x18, 0xee, 0xd7, 0x73, 0x71,
	0x0a, 0x80, 0x28, 0xab, 0x0c, 0x42, 0xd2, 0xe0, 0xe4, 0x72, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47,
	0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d,
	0xc7, 0x72, 0x0c, 0x51, 0x5a, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa,
	0xc1, 0x99, 0x69, 0xc9, 0x19, 0x89, 0x99, 0x79, 0xfa, 0x30, 0xef, 0x56, 0xc0, 0x3c, 0x0c, 0xf6,
	0x6d, 0x12, 0x1b, 0xd8, 0xbb, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x17, 0x1b, 0x0a,
	0x59, 0x01, 0x00, 0x00,
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
	if len(m.Prophecies) > 0 {
		for iNdEx := len(m.Prophecies) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Prophecies[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.AdminAddress) > 0 {
		i -= len(m.AdminAddress)
		copy(dAtA[i:], m.AdminAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.AdminAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.AddressWhitelist) > 0 {
		for iNdEx := len(m.AddressWhitelist) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AddressWhitelist[iNdEx])
			copy(dAtA[i:], m.AddressWhitelist[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.AddressWhitelist[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
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
	if len(m.AddressWhitelist) > 0 {
		for _, s := range m.AddressWhitelist {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = len(m.AdminAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.Prophecies) > 0 {
		for _, e := range m.Prophecies {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddressWhitelist", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AddressWhitelist = append(m.AddressWhitelist, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdminAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AdminAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prophecies", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Prophecies = append(m.Prophecies, &DBProphecy{})
			if err := m.Prophecies[len(m.Prophecies)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
