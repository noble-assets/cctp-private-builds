// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: noble/cctp/v1/sending_and_receiving_messages_paused.proto

package types

import (
	fmt "fmt"
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

// *
// Message format for SendingAndReceivingMessagesPaused
// @param paused true if paused, false if not paused
type SendingAndReceivingMessagesPaused struct {
	Paused bool `protobuf:"varint,1,opt,name=paused,proto3" json:"paused,omitempty"`
}

func (m *SendingAndReceivingMessagesPaused) Reset()         { *m = SendingAndReceivingMessagesPaused{} }
func (m *SendingAndReceivingMessagesPaused) String() string { return proto.CompactTextString(m) }
func (*SendingAndReceivingMessagesPaused) ProtoMessage()    {}
func (*SendingAndReceivingMessagesPaused) Descriptor() ([]byte, []int) {
	return fileDescriptor_484d35aa55c816b9, []int{0}
}
func (m *SendingAndReceivingMessagesPaused) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SendingAndReceivingMessagesPaused) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SendingAndReceivingMessagesPaused.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SendingAndReceivingMessagesPaused) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendingAndReceivingMessagesPaused.Merge(m, src)
}
func (m *SendingAndReceivingMessagesPaused) XXX_Size() int {
	return m.Size()
}
func (m *SendingAndReceivingMessagesPaused) XXX_DiscardUnknown() {
	xxx_messageInfo_SendingAndReceivingMessagesPaused.DiscardUnknown(m)
}

var xxx_messageInfo_SendingAndReceivingMessagesPaused proto.InternalMessageInfo

func (m *SendingAndReceivingMessagesPaused) GetPaused() bool {
	if m != nil {
		return m.Paused
	}
	return false
}

func init() {
	proto.RegisterType((*SendingAndReceivingMessagesPaused)(nil), "noble.cctp.v1.SendingAndReceivingMessagesPaused")
}

func init() {
	proto.RegisterFile("noble/cctp/v1/sending_and_receiving_messages_paused.proto", fileDescriptor_484d35aa55c816b9)
}

var fileDescriptor_484d35aa55c816b9 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0xcc, 0xcb, 0x4f, 0xca,
	0x49, 0xd5, 0x4f, 0x4e, 0x2e, 0x29, 0xd0, 0x2f, 0x33, 0xd4, 0x2f, 0x4e, 0xcd, 0x4b, 0xc9, 0xcc,
	0x4b, 0x8f, 0x4f, 0xcc, 0x4b, 0x89, 0x2f, 0x4a, 0x4d, 0x4e, 0xcd, 0x2c, 0x03, 0xf1, 0x72, 0x53,
	0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x8b, 0xe3, 0x0b, 0x12, 0x4b, 0x8b, 0x53, 0x53, 0xf4, 0x0a, 0x8a,
	0xf2, 0x4b, 0xf2, 0x85, 0x78, 0xc1, 0x5a, 0xf5, 0x40, 0x5a, 0xf5, 0xca, 0x0c, 0x95, 0xac, 0xb9,
	0x14, 0x83, 0x21, 0xba, 0x1d, 0xf3, 0x52, 0x82, 0x60, 0x7a, 0x7d, 0xa1, 0x5a, 0x03, 0xc0, 0x3a,
	0x85, 0xc4, 0xb8, 0xd8, 0x20, 0x66, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x04, 0x41, 0x79, 0x4e,
	0x6e, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7,
	0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x93, 0x9e, 0x59, 0x92,
	0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f, 0x9c, 0x59, 0x94, 0x9c, 0x93, 0x9a, 0x96, 0x99,
	0xa7, 0x0f, 0xb6, 0x5a, 0x17, 0xec, 0xea, 0x0a, 0x88, 0xe3, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0x93,
	0xd8, 0xc0, 0x4e, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x42, 0xe6, 0x58, 0x30, 0xd7, 0x00,
	0x00, 0x00,
}

func (m *SendingAndReceivingMessagesPaused) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SendingAndReceivingMessagesPaused) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SendingAndReceivingMessagesPaused) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Paused {
		i--
		if m.Paused {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSendingAndReceivingMessagesPaused(dAtA []byte, offset int, v uint64) int {
	offset -= sovSendingAndReceivingMessagesPaused(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SendingAndReceivingMessagesPaused) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Paused {
		n += 2
	}
	return n
}

func sovSendingAndReceivingMessagesPaused(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSendingAndReceivingMessagesPaused(x uint64) (n int) {
	return sovSendingAndReceivingMessagesPaused(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SendingAndReceivingMessagesPaused) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSendingAndReceivingMessagesPaused
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
			return fmt.Errorf("proto: SendingAndReceivingMessagesPaused: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SendingAndReceivingMessagesPaused: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Paused", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSendingAndReceivingMessagesPaused
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Paused = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipSendingAndReceivingMessagesPaused(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSendingAndReceivingMessagesPaused
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
func skipSendingAndReceivingMessagesPaused(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSendingAndReceivingMessagesPaused
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
					return 0, ErrIntOverflowSendingAndReceivingMessagesPaused
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
					return 0, ErrIntOverflowSendingAndReceivingMessagesPaused
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
				return 0, ErrInvalidLengthSendingAndReceivingMessagesPaused
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSendingAndReceivingMessagesPaused
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSendingAndReceivingMessagesPaused
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSendingAndReceivingMessagesPaused        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSendingAndReceivingMessagesPaused          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSendingAndReceivingMessagesPaused = fmt.Errorf("proto: unexpected end of group")
)
