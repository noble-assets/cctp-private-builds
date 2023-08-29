// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: router/mint.proto

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

// mints are stored so incoming IBC forward messages can check if a mint has
// been processed
// @param source_domain - source domain id
// @param source_domain_sender - address of MessageTransmitter caller on source
// domain
// @param nonce - unique message nonce
// @param amount - amount of minted tokens
// @param destination_domain - destination domain id
// @param mint_recipient - address to receive minted tokens on destination
// domain
type Mint struct {
	SourceDomain       uint32      `protobuf:"varint,1,opt,name=source_domain,json=sourceDomain,proto3" json:"source_domain,omitempty"`
	SourceDomainSender string      `protobuf:"bytes,2,opt,name=source_domain_sender,json=sourceDomainSender,proto3" json:"source_domain_sender,omitempty"`
	Nonce              uint64      `protobuf:"varint,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Amount             *types.Coin `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
	DestinationDomain  uint32      `protobuf:"varint,5,opt,name=destination_domain,json=destinationDomain,proto3" json:"destination_domain,omitempty"`
	MintRecipient      string      `protobuf:"bytes,6,opt,name=mint_recipient,json=mintRecipient,proto3" json:"mint_recipient,omitempty"`
}

func (m *Mint) Reset()         { *m = Mint{} }
func (m *Mint) String() string { return proto.CompactTextString(m) }
func (*Mint) ProtoMessage()    {}
func (*Mint) Descriptor() ([]byte, []int) {
	return fileDescriptor_650d9754adf300ce, []int{0}
}
func (m *Mint) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Mint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Mint.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Mint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mint.Merge(m, src)
}
func (m *Mint) XXX_Size() int {
	return m.Size()
}
func (m *Mint) XXX_DiscardUnknown() {
	xxx_messageInfo_Mint.DiscardUnknown(m)
}

var xxx_messageInfo_Mint proto.InternalMessageInfo

func (m *Mint) GetSourceDomain() uint32 {
	if m != nil {
		return m.SourceDomain
	}
	return 0
}

func (m *Mint) GetSourceDomainSender() string {
	if m != nil {
		return m.SourceDomainSender
	}
	return ""
}

func (m *Mint) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Mint) GetAmount() *types.Coin {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *Mint) GetDestinationDomain() uint32 {
	if m != nil {
		return m.DestinationDomain
	}
	return 0
}

func (m *Mint) GetMintRecipient() string {
	if m != nil {
		return m.MintRecipient
	}
	return ""
}

func init() {
	proto.RegisterType((*Mint)(nil), "noble.router.Mint")
}

func init() { proto.RegisterFile("router/mint.proto", fileDescriptor_650d9754adf300ce) }

var fileDescriptor_650d9754adf300ce = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xbf, 0x4e, 0xf3, 0x30,
	0x14, 0xc5, 0xeb, 0xef, 0x6b, 0x2b, 0x61, 0x5a, 0xa4, 0x5a, 0x1d, 0x42, 0x87, 0x28, 0x02, 0x21,
	0x65, 0x49, 0x4c, 0x61, 0x66, 0x01, 0x56, 0x96, 0x30, 0x20, 0xb1, 0x54, 0x89, 0x63, 0xca, 0x95,
	0x9a, 0x7b, 0x23, 0xdb, 0xa9, 0xe0, 0x2d, 0x78, 0x2c, 0xc6, 0x8e, 0x8c, 0xa8, 0x7d, 0x0d, 0x06,
	0x94, 0x3f, 0x95, 0xca, 0x76, 0xfd, 0x3b, 0x67, 0xf8, 0xe9, 0x98, 0x4f, 0x0c, 0x55, 0x4e, 0x1b,
	0x59, 0x00, 0xba, 0xb8, 0x34, 0xe4, 0x48, 0x8c, 0x90, 0xb2, 0x95, 0x8e, 0xdb, 0x60, 0xe6, 0x2b,
	0xb2, 0x05, 0x59, 0x99, 0xa5, 0x56, 0xcb, 0xf5, 0x3c, 0xd3, 0x2e, 0x9d, 0x4b, 0x45, 0x80, 0x6d,
	0x7b, 0x36, 0x5d, 0xd2, 0x92, 0x9a, 0x53, 0xd6, 0x57, 0x4b, 0xcf, 0x7e, 0x18, 0xef, 0x3f, 0x00,
	0x3a, 0x71, 0xce, 0xc7, 0x96, 0x2a, 0xa3, 0xf4, 0x22, 0xa7, 0x22, 0x05, 0xf4, 0x58, 0xc0, 0xc2,
	0x71, 0x32, 0x6a, 0xe1, 0x7d, 0xc3, 0xc4, 0x25, 0x9f, 0xfe, 0x29, 0x2d, 0xac, 0xc6, 0x5c, 0x1b,
	0xef, 0x5f, 0xc0, 0xc2, 0xa3, 0x44, 0x1c, 0x76, 0x1f, 0x9b, 0x44, 0x4c, 0xf9, 0x00, 0x09, 0x95,
	0xf6, 0xfe, 0x07, 0x2c, 0xec, 0x27, 0xed, 0x43, 0xcc, 0xf9, 0x30, 0x2d, 0xa8, 0x42, 0xe7, 0xf5,
	0x03, 0x16, 0x1e, 0x5f, 0x9d, 0xc6, 0xad, 0x7c, 0x5c, 0xcb, 0xc7, 0x9d, 0x7c, 0x7c, 0x47, 0x80,
	0x49, 0x57, 0x14, 0x11, 0x17, 0xb9, 0xb6, 0x0e, 0x30, 0x75, 0x40, 0xb8, 0x97, 0x1c, 0x34, 0x92,
	0x93, 0x83, 0xa4, 0x33, 0xbd, 0xe0, 0x27, 0xf5, 0x52, 0x0b, 0xa3, 0x15, 0x94, 0xa0, 0xd1, 0x79,
	0xc3, 0xc6, 0x71, 0x5c, 0xd3, 0x64, 0x0f, 0x6f, 0x9f, 0x3e, 0xb7, 0x3e, 0xdb, 0x6c, 0x7d, 0xf6,
	0xbd, 0xf5, 0xd9, 0xc7, 0xce, 0xef, 0x6d, 0x76, 0x7e, 0xef, 0x6b, 0xe7, 0xf7, 0x9e, 0x6f, 0x96,
	0xe0, 0x5e, 0xab, 0x2c, 0x56, 0x54, 0x48, 0x05, 0x46, 0xad, 0xf4, 0x0b, 0xa0, 0x6c, 0x16, 0x8f,
	0x94, 0x72, 0x65, 0x54, 0x1a, 0x58, 0xa7, 0x4e, 0x47, 0x59, 0x05, 0xab, 0xdc, 0xca, 0x37, 0xd9,
	0x7d, 0x90, 0x7b, 0x2f, 0xb5, 0xcd, 0x86, 0xcd, 0xbc, 0xd7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x6c, 0xaa, 0x1c, 0x3e, 0xb7, 0x01, 0x00, 0x00,
}

func (m *Mint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Mint) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Mint) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MintRecipient) > 0 {
		i -= len(m.MintRecipient)
		copy(dAtA[i:], m.MintRecipient)
		i = encodeVarintMint(dAtA, i, uint64(len(m.MintRecipient)))
		i--
		dAtA[i] = 0x32
	}
	if m.DestinationDomain != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.DestinationDomain))
		i--
		dAtA[i] = 0x28
	}
	if m.Amount != nil {
		{
			size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMint(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Nonce != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x18
	}
	if len(m.SourceDomainSender) > 0 {
		i -= len(m.SourceDomainSender)
		copy(dAtA[i:], m.SourceDomainSender)
		i = encodeVarintMint(dAtA, i, uint64(len(m.SourceDomainSender)))
		i--
		dAtA[i] = 0x12
	}
	if m.SourceDomain != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.SourceDomain))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMint(dAtA []byte, offset int, v uint64) int {
	offset -= sovMint(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Mint) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SourceDomain != 0 {
		n += 1 + sovMint(uint64(m.SourceDomain))
	}
	l = len(m.SourceDomainSender)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	if m.Nonce != 0 {
		n += 1 + sovMint(uint64(m.Nonce))
	}
	if m.Amount != nil {
		l = m.Amount.Size()
		n += 1 + l + sovMint(uint64(l))
	}
	if m.DestinationDomain != 0 {
		n += 1 + sovMint(uint64(m.DestinationDomain))
	}
	l = len(m.MintRecipient)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	return n
}

func sovMint(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMint(x uint64) (n int) {
	return sovMint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Mint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: Mint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Mint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceDomain", wireType)
			}
			m.SourceDomain = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SourceDomain |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceDomainSender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceDomainSender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Amount == nil {
				m.Amount = &types.Coin{}
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationDomain", wireType)
			}
			m.DestinationDomain = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DestinationDomain |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintRecipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintRecipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func skipMint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
				return 0, ErrInvalidLengthMint
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMint
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMint
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMint        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMint          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMint = fmt.Errorf("proto: unexpected end of group")
)
