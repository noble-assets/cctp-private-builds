// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fiattokenfactory/minters.proto

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
	return fileDescriptor_eaf83371ab173141, []int{0}
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
	proto.RegisterType((*Minters)(nil), "noble.fiattokenfactory.Minters")
}

func init() { proto.RegisterFile("fiattokenfactory/minters.proto", fileDescriptor_eaf83371ab173141) }

var fileDescriptor_eaf83371ab173141 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x63, 0x84, 0xa8, 0x1a, 0xb6, 0x08, 0xa1, 0xd0, 0xc1, 0x54, 0x4c, 0x5d, 0x62, 0xab,
	0x30, 0xb3, 0x94, 0x99, 0x25, 0x23, 0x9b, 0xed, 0x5e, 0x83, 0x45, 0xe2, 0x8b, 0xec, 0x4b, 0xa0,
	0x6f, 0xc1, 0x63, 0x75, 0xec, 0xc8, 0x84, 0x50, 0xf2, 0x22, 0xa8, 0x09, 0x08, 0x89, 0x6e, 0x77,
	0xba, 0xef, 0x3e, 0xe9, 0xff, 0x63, 0xbe, 0xb1, 0x8a, 0x08, 0x5f, 0xc0, 0x6d, 0x94, 0x21, 0xf4,
	0x5b, 0x59, 0x59, 0x47, 0xe0, 0x83, 0xa8, 0x3d, 0x12, 0x26, 0x97, 0x0e, 0x75, 0x09, 0xe2, 0x3f,
	0x35, 0xbb, 0x28, 0xb0, 0xc0, 0x01, 0x91, 0x87, 0x69, 0xa4, 0x67, 0xdc, 0x60, 0xa8, 0x30, 0x48,
	0xad, 0x02, 0xc8, 0x76, 0xa9, 0x81, 0xd4, 0x52, 0x1a, 0xb4, 0x6e, 0xbc, 0xdf, 0xe8, 0x78, 0xf2,
	0x38, 0xea, 0x93, 0x34, 0x9e, 0xa8, 0xf5, 0xda, 0x43, 0x08, 0x29, 0x9b, 0xb3, 0xc5, 0x34, 0xff,
	0x5d, 0x93, 0xfb, 0x78, 0xaa, 0xca, 0x12, 0x5f, 0x95, 0x33, 0x90, 0x9e, 0xcc, 0xd9, 0xe2, 0xfc,
	0xf6, 0x4a, 0x8c, 0x62, 0x71, 0x10, 0x8b, 0x1f, 0xb1, 0x78, 0x40, 0xeb, 0x56, 0xa7, 0xbb, 0xcf,
	0xeb, 0x28, 0xff, 0xfb, 0x58, 0x95, 0xbb, 0x8e, 0xb3, 0x7d, 0xc7, 0xd9, 0x57, 0xc7, 0xd9, 0x7b,
	0xcf, 0xa3, 0x7d, 0xcf, 0xa3, 0x8f, 0x9e, 0x47, 0x4f, 0x79, 0x61, 0xe9, 0xb9, 0xd1, 0xc2, 0x60,
	0x25, 0x03, 0x79, 0xe5, 0x0a, 0x28, 0xb1, 0x85, 0xac, 0x05, 0x47, 0x8d, 0x87, 0x20, 0x87, 0xac,
	0x99, 0x31, 0x54, 0x67, 0x1e, 0x1b, 0x02, 0x9f, 0xd5, 0xde, 0xb6, 0x8a, 0x40, 0xbe, 0xc9, 0xa3,
	0x9e, 0x68, 0x5b, 0x43, 0xd0, 0x67, 0x43, 0xb0, 0xbb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf3,
	0x89, 0xe2, 0x61, 0x48, 0x01, 0x00, 0x00,
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
