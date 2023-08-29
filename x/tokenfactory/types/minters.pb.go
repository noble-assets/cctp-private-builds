// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tokenfactory/minters.proto

package types

import (
	fmt "fmt"
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

type Minters struct {
	Address   string     `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Allowance types.Coin `protobuf:"bytes,2,opt,name=allowance,proto3" json:"allowance"`
}

func (m *Minters) Reset()         { *m = Minters{} }
func (m *Minters) String() string { return proto.CompactTextString(m) }
func (*Minters) ProtoMessage()    {}
func (*Minters) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac9d7080b5299f2f, []int{0}
}
func (m *Minters) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Minters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Minters.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Minters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Minters.Merge(m, src)
}
func (m *Minters) XXX_Size() int {
	return m.Size()
}
func (m *Minters) XXX_DiscardUnknown() {
	xxx_messageInfo_Minters.DiscardUnknown(m)
}

var xxx_messageInfo_Minters proto.InternalMessageInfo

func (m *Minters) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Minters) GetAllowance() types.Coin {
	if m != nil {
		return m.Allowance
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*Minters)(nil), "noble.tokenfactory.Minters")
}

func init() { proto.RegisterFile("tokenfactory/minters.proto", fileDescriptor_ac9d7080b5299f2f) }

var fileDescriptor_ac9d7080b5299f2f = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x63, 0x84, 0xa8, 0x1a, 0xb6, 0x88, 0x21, 0x64, 0x30, 0x15, 0x53, 0x97, 0xd8, 0x2a,
	0xcc, 0x2c, 0xed, 0xcc, 0xd2, 0x11, 0x89, 0xc1, 0x76, 0xaf, 0xc1, 0x22, 0xf1, 0x45, 0xf6, 0xb5,
	0xd0, 0xb7, 0xe0, 0xb1, 0x3a, 0x76, 0x64, 0x42, 0x28, 0x79, 0x11, 0xd4, 0x18, 0x04, 0x6c, 0x77,
	0xfa, 0xff, 0xfb, 0xa4, 0xef, 0xd2, 0x82, 0xf0, 0x19, 0xdc, 0x5a, 0x19, 0x42, 0xbf, 0x93, 0x8d,
	0x75, 0x04, 0x3e, 0x88, 0xd6, 0x23, 0x61, 0x96, 0x39, 0xd4, 0x35, 0x88, 0xbf, 0x8d, 0xe2, 0xa2,
	0xc2, 0x0a, 0x87, 0x58, 0x1e, 0xa7, 0xd8, 0x2c, 0xb8, 0xc1, 0xd0, 0x60, 0x90, 0x5a, 0x05, 0x90,
	0xdb, 0x99, 0x06, 0x52, 0x33, 0x69, 0xd0, 0xba, 0x98, 0x5f, 0xeb, 0x74, 0x74, 0x1f, 0xd1, 0x59,
	0x9e, 0x8e, 0xd4, 0x6a, 0xe5, 0x21, 0x84, 0x9c, 0x4d, 0xd8, 0x74, 0xbc, 0xfc, 0x59, 0xb3, 0xbb,
	0x74, 0xac, 0xea, 0x1a, 0x5f, 0x94, 0x33, 0x90, 0x9f, 0x4c, 0xd8, 0xf4, 0xfc, 0xe6, 0x52, 0x44,
	0xb0, 0x38, 0x82, 0xc5, 0x37, 0x58, 0x2c, 0xd0, 0xba, 0xf9, 0xe9, 0xfe, 0xe3, 0x2a, 0x59, 0xfe,
	0x5e, 0xcc, 0x1f, 0xf7, 0x1d, 0x67, 0x87, 0x8e, 0xb3, 0xcf, 0x8e, 0xb3, 0xb7, 0x9e, 0x27, 0x87,
	0x9e, 0x27, 0xef, 0x3d, 0x4f, 0x1e, 0x16, 0x95, 0xa5, 0xa7, 0x8d, 0x16, 0x06, 0x1b, 0x69, 0xac,
	0x37, 0x35, 0xac, 0xad, 0x93, 0x83, 0x5c, 0x69, 0x0c, 0xb5, 0xa5, 0xc7, 0x0d, 0x81, 0x2f, 0x5b,
	0x6f, 0xb7, 0x8a, 0x40, 0xbe, 0xca, 0x7f, 0x4f, 0xa1, 0x5d, 0x0b, 0x41, 0x9f, 0x0d, 0x26, 0xb7,
	0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1a, 0xe4, 0xad, 0xd9, 0x31, 0x01, 0x00, 0x00,
}

func (m *Minters) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Minters) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Minters) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Allowance.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMinters(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintMinters(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMinters(dAtA []byte, offset int, v uint64) int {
	offset -= sovMinters(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Minters) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovMinters(uint64(l))
	}
	l = m.Allowance.Size()
	n += 1 + l + sovMinters(uint64(l))
	return n
}

func sovMinters(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMinters(x uint64) (n int) {
	return sovMinters(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Minters) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMinters
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
			return fmt.Errorf("proto: Minters: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Minters: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinters
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
				return ErrInvalidLengthMinters
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMinters
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allowance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMinters
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
				return ErrInvalidLengthMinters
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMinters
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Allowance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMinters(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMinters
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
func skipMinters(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMinters
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
					return 0, ErrIntOverflowMinters
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
					return 0, ErrIntOverflowMinters
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
				return 0, ErrInvalidLengthMinters
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMinters
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMinters
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMinters        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMinters          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMinters = fmt.Errorf("proto: unexpected end of group")
)
