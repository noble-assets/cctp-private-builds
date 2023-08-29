// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: noble/cctp/v1/per_message_burn_limit.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
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

// PerMessageBurnLimit is the maximum amount of a certain denom that can be burned in an single burn
// @param denom the denom
// @param amount the amount that can be burned (in microunits).  An amount of 1000000 uusdc is equivalent to 1USDC
type PerMessageBurnLimit struct {
	Denom  string                `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Amount cosmossdk_io_math.Int `protobuf:"bytes,2,opt,name=amount,proto3,customtype=cosmossdk.io/math.Int" json:"amount"`
}

func (m *PerMessageBurnLimit) Reset()         { *m = PerMessageBurnLimit{} }
func (m *PerMessageBurnLimit) String() string { return proto.CompactTextString(m) }
func (*PerMessageBurnLimit) ProtoMessage()    {}
func (*PerMessageBurnLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_957ecb70ef2a5e4f, []int{0}
}
func (m *PerMessageBurnLimit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PerMessageBurnLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PerMessageBurnLimit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PerMessageBurnLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PerMessageBurnLimit.Merge(m, src)
}
func (m *PerMessageBurnLimit) XXX_Size() int {
	return m.Size()
}
func (m *PerMessageBurnLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_PerMessageBurnLimit.DiscardUnknown(m)
}

var xxx_messageInfo_PerMessageBurnLimit proto.InternalMessageInfo

func (m *PerMessageBurnLimit) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func init() {
	proto.RegisterType((*PerMessageBurnLimit)(nil), "noble.cctp.v1.PerMessageBurnLimit")
}

func init() {
	proto.RegisterFile("noble/cctp/v1/per_message_burn_limit.proto", fileDescriptor_957ecb70ef2a5e4f)
}

var fileDescriptor_957ecb70ef2a5e4f = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0x8f, 0xbd, 0x4a, 0xf4, 0x40,
	0x14, 0x86, 0x93, 0x0f, 0xbe, 0x05, 0x03, 0x36, 0x71, 0x85, 0x45, 0x70, 0x56, 0xac, 0x44, 0xc8,
	0x0c, 0x8b, 0x58, 0xd9, 0xa5, 0x13, 0x14, 0x64, 0xc1, 0xc6, 0x26, 0x24, 0xb3, 0xc7, 0xec, 0x60,
	0x66, 0x4e, 0x98, 0x39, 0x13, 0xf4, 0x2e, 0xbc, 0xac, 0x2d, 0xb7, 0x14, 0x8b, 0x45, 0x92, 0x1b,
	0x91, 0x49, 0xec, 0xce, 0xcf, 0xc3, 0xcb, 0xfb, 0x24, 0xd7, 0x06, 0xab, 0x06, 0x84, 0x94, 0xd4,
	0x8a, 0x6e, 0x25, 0x5a, 0xb0, 0x85, 0x06, 0xe7, 0xca, 0x1a, 0x8a, 0xca, 0x5b, 0x53, 0x34, 0x4a,
	0x2b, 0xe2, 0xad, 0x45, 0xc2, 0xf4, 0x78, 0x64, 0x79, 0x60, 0x79, 0xb7, 0x3a, 0x9b, 0xd7, 0x58,
	0xe3, 0xf8, 0x11, 0x61, 0x9a, 0xa0, 0xcb, 0x2a, 0x39, 0x79, 0x02, 0xfb, 0x38, 0x65, 0xe4, 0xde,
	0x9a, 0x87, 0x90, 0x90, 0xce, 0x93, 0xff, 0x1b, 0x30, 0xa8, 0x17, 0xf1, 0x45, 0x7c, 0x75, 0xb4,
	0x9e, 0x96, 0xf4, 0x36, 0x99, 0x95, 0x1a, 0xbd, 0xa1, 0xc5, 0xbf, 0x70, 0xce, 0xcf, 0x77, 0x87,
	0x65, 0xf4, 0x7d, 0x58, 0x9e, 0x4a, 0x74, 0x1a, 0x9d, 0xdb, 0xbc, 0x71, 0x85, 0x42, 0x97, 0xb4,
	0xe5, 0xf7, 0x86, 0xd6, 0x7f, 0x70, 0xfe, 0xbc, 0xeb, 0x59, 0xbc, 0xef, 0x59, 0xfc, 0xd3, 0xb3,
	0xf8, 0x73, 0x60, 0xd1, 0x7e, 0x60, 0xd1, 0xd7, 0xc0, 0xa2, 0x97, 0xbb, 0x5a, 0xd1, 0xd6, 0x57,
	0x5c, 0xa2, 0x16, 0x52, 0x59, 0xd9, 0xc0, 0xab, 0x32, 0x62, 0xec, 0x9d, 0x85, 0xde, 0x99, 0x45,
	0x4f, 0x60, 0xb3, 0xd6, 0xaa, 0xae, 0x24, 0x10, 0xef, 0x93, 0x39, 0x7d, 0xb4, 0xe0, 0xaa, 0xd9,
	0x68, 0x70, 0xf3, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x19, 0xb9, 0x2b, 0x14, 0x01, 0x00, 0x00,
}

func (m *PerMessageBurnLimit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PerMessageBurnLimit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PerMessageBurnLimit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPerMessageBurnLimit(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintPerMessageBurnLimit(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPerMessageBurnLimit(dAtA []byte, offset int, v uint64) int {
	offset -= sovPerMessageBurnLimit(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PerMessageBurnLimit) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovPerMessageBurnLimit(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovPerMessageBurnLimit(uint64(l))
	return n
}

func sovPerMessageBurnLimit(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPerMessageBurnLimit(x uint64) (n int) {
	return sovPerMessageBurnLimit(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PerMessageBurnLimit) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPerMessageBurnLimit
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
			return fmt.Errorf("proto: PerMessageBurnLimit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PerMessageBurnLimit: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerMessageBurnLimit
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
				return ErrInvalidLengthPerMessageBurnLimit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPerMessageBurnLimit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPerMessageBurnLimit
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
				return ErrInvalidLengthPerMessageBurnLimit
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPerMessageBurnLimit
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPerMessageBurnLimit(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPerMessageBurnLimit
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
func skipPerMessageBurnLimit(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPerMessageBurnLimit
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
					return 0, ErrIntOverflowPerMessageBurnLimit
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
					return 0, ErrIntOverflowPerMessageBurnLimit
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
				return 0, ErrInvalidLengthPerMessageBurnLimit
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPerMessageBurnLimit
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPerMessageBurnLimit
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPerMessageBurnLimit        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPerMessageBurnLimit          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPerMessageBurnLimit = fmt.Errorf("proto: unexpected end of group")
)
