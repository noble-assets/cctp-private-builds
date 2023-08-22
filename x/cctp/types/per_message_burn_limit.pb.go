// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: circle/cctp/v1/per_message_burn_limit.proto

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

//
// PerMessageBurnLimit is the maximum amount of a certain denom that can be
// burned in an single burn
// @param denom the denom
// @param amount the amount that can be burned (in microunits).  An amount of
// 1000000 uusdc is equivalent to 1USDC
type PerMessageBurnLimit struct {
	Denom  string                `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Amount cosmossdk_io_math.Int `protobuf:"bytes,2,opt,name=amount,proto3,customtype=cosmossdk.io/math.Int" json:"amount"`
}

func (m *PerMessageBurnLimit) Reset()         { *m = PerMessageBurnLimit{} }
func (m *PerMessageBurnLimit) String() string { return proto.CompactTextString(m) }
func (*PerMessageBurnLimit) ProtoMessage()    {}
func (*PerMessageBurnLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_c255171867807834, []int{0}
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
	proto.RegisterType((*PerMessageBurnLimit)(nil), "circle.cctp.v1.PerMessageBurnLimit")
}

func init() {
	proto.RegisterFile("circle/cctp/v1/per_message_burn_limit.proto", fileDescriptor_c255171867807834)
}

var fileDescriptor_c255171867807834 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0x8f, 0xcb, 0x4a, 0xf4, 0x30,
	0x18, 0x86, 0x9b, 0x1f, 0xfe, 0x01, 0xbb, 0x70, 0x51, 0x47, 0x18, 0x04, 0x33, 0xe2, 0x4a, 0x90,
	0x26, 0x0c, 0xe2, 0xca, 0x5d, 0x77, 0x82, 0x82, 0x0c, 0xb8, 0x71, 0x53, 0x9a, 0x34, 0x76, 0x82,
	0x4d, 0xbe, 0x90, 0x43, 0xd1, 0xbb, 0xf0, 0xb2, 0x66, 0x39, 0x4b, 0x71, 0x31, 0x48, 0x7b, 0x23,
	0x92, 0x89, 0xbb, 0xef, 0xf0, 0xf0, 0xf2, 0x3e, 0xf9, 0x35, 0x97, 0x96, 0xf7, 0x82, 0x72, 0xee,
	0x0d, 0x1d, 0x56, 0xd4, 0x08, 0x5b, 0x2b, 0xe1, 0x5c, 0xd3, 0x89, 0x9a, 0x05, 0xab, 0xeb, 0x5e,
	0x2a, 0xe9, 0x89, 0xb1, 0xe0, 0xa1, 0x38, 0x4e, 0x30, 0x89, 0x30, 0x19, 0x56, 0x67, 0xf3, 0x0e,
	0x3a, 0x38, 0xbc, 0x68, 0x9c, 0x12, 0x75, 0xc9, 0xf2, 0x93, 0x27, 0x61, 0x1f, 0x53, 0x48, 0x15,
	0xac, 0x7e, 0x88, 0x11, 0xc5, 0x3c, 0xff, 0xdf, 0x0a, 0x0d, 0x6a, 0x81, 0x2e, 0xd0, 0xd5, 0xd1,
	0x3a, 0x2d, 0xc5, 0x6d, 0x3e, 0x6b, 0x14, 0x04, 0xed, 0x17, 0xff, 0xe2, 0xb9, 0x3a, 0xdf, 0xee,
	0x97, 0xd9, 0xf7, 0x7e, 0x79, 0xca, 0xc1, 0x29, 0x70, 0xae, 0x7d, 0x23, 0x12, 0xa8, 0x6a, 0xfc,
	0x86, 0xdc, 0x6b, 0xbf, 0xfe, 0x83, 0xab, 0xe7, 0xed, 0x88, 0xd1, 0x6e, 0xc4, 0xe8, 0x67, 0xc4,
	0xe8, 0x73, 0xc2, 0xd9, 0x6e, 0xc2, 0xd9, 0xd7, 0x84, 0xb3, 0x97, 0xbb, 0x4e, 0xfa, 0x4d, 0x60,
	0x84, 0x83, 0xa2, 0xa9, 0xee, 0xab, 0xd4, 0x54, 0x03, 0xeb, 0x45, 0x19, 0x7b, 0x97, 0xc6, 0xca,
	0xa1, 0xf1, 0xa2, 0x64, 0x41, 0xf6, 0xad, 0xa3, 0xef, 0x49, 0xdd, 0x7f, 0x18, 0xe1, 0xd8, 0xec,
	0x60, 0x70, 0xf3, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x04, 0x9a, 0x74, 0x0c, 0x16, 0x01, 0x00, 0x00,
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
