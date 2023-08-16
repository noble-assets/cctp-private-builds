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

// mints are stored so incoming IBC forward messages can check if a mint has been processed
// @param source_domain - source domain id
// @param source_domain_sender - address of MessageTransmitter caller on source domain
// @param nonce - unique message nonce
// @param amount - amount of minted tokens
// @param destination_domain - destination domain id
// @param mint_recipient - address to receive minted tokens on destination domain
type Mint struct {
	SourceDomain       uint32      `protobuf:"varint,1,opt,name=source_domain,json=sourceDomain,proto3" json:"source_domain,omitempty"`
	SourceDomainSender string      `protobuf:"bytes,2,opt,name=source_domain_sender,json=sourceDomainSender,proto3" json:"source_domain_sender,omitempty"`
	Nonce              uint64      `protobuf:"varint,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Amount             *types.Coin `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount,omitempty"`
	DestinationDomain  string      `protobuf:"bytes,5,opt,name=destination_domain,json=destinationDomain,proto3" json:"destination_domain,omitempty"`
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

func (m *Mint) GetDestinationDomain() string {
	if m != nil {
		return m.DestinationDomain
	}
	return ""
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
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xbd, 0x4e, 0xf3, 0x30,
	0x14, 0x86, 0xeb, 0xef, 0x6b, 0x2b, 0x61, 0x5a, 0xa4, 0x5a, 0x1d, 0x42, 0x87, 0x28, 0x02, 0x21,
	0x65, 0x49, 0x4c, 0xe1, 0x0e, 0x80, 0x81, 0x85, 0x25, 0x6c, 0x2c, 0x55, 0xe2, 0x1e, 0x05, 0x4b,
	0x8d, 0x4f, 0x64, 0x9f, 0x44, 0x70, 0x17, 0x5c, 0x16, 0x63, 0x47, 0x46, 0xd4, 0xde, 0x06, 0x03,
	0xca, 0x4f, 0xa5, 0xb2, 0x1d, 0xbf, 0xe7, 0x91, 0xfc, 0xe8, 0x3d, 0x7c, 0x66, 0xb1, 0x22, 0xb0,
	0xb2, 0xd0, 0x86, 0xe2, 0xd2, 0x22, 0xa1, 0x98, 0x18, 0xcc, 0x36, 0x10, 0x77, 0x8b, 0x85, 0xaf,
	0xd0, 0x15, 0xe8, 0x64, 0x96, 0x3a, 0x90, 0xf5, 0x32, 0x03, 0x4a, 0x97, 0x52, 0xa1, 0x36, 0x1d,
	0xbd, 0x98, 0xe7, 0x98, 0x63, 0x3b, 0xca, 0x66, 0xea, 0xd2, 0x8b, 0x1f, 0xc6, 0x87, 0x4f, 0xda,
	0x90, 0xb8, 0xe4, 0x53, 0x87, 0x95, 0x55, 0xb0, 0x5a, 0x63, 0x91, 0x6a, 0xe3, 0xb1, 0x80, 0x85,
	0xd3, 0x64, 0xd2, 0x85, 0x0f, 0x6d, 0x26, 0xae, 0xf9, 0xfc, 0x0f, 0xb4, 0x72, 0x60, 0xd6, 0x60,
	0xbd, 0x7f, 0x01, 0x0b, 0x4f, 0x12, 0x71, 0xcc, 0x3e, 0xb7, 0x1b, 0x31, 0xe7, 0x23, 0x83, 0x46,
	0x81, 0xf7, 0x3f, 0x60, 0xe1, 0x30, 0xe9, 0x1e, 0x62, 0xc9, 0xc7, 0x69, 0x81, 0x95, 0x21, 0x6f,
	0x18, 0xb0, 0xf0, 0xf4, 0xe6, 0x3c, 0xee, 0xe4, 0xe3, 0x46, 0x3e, 0xee, 0xe5, 0xe3, 0x7b, 0xd4,
	0x26, 0xe9, 0x41, 0x11, 0x71, 0xb1, 0x06, 0x47, 0xda, 0xa4, 0xa4, 0xd1, 0x1c, 0x24, 0x47, 0xed,
	0xc7, 0xb3, 0xa3, 0x4d, 0x6f, 0x7a, 0xc5, 0xcf, 0x9a, 0xa6, 0x56, 0x16, 0x94, 0x2e, 0x35, 0x18,
	0xf2, 0xc6, 0x2d, 0x3a, 0x6d, 0xd2, 0xe4, 0x10, 0xde, 0x65, 0x9f, 0x3b, 0x9f, 0x6d, 0x77, 0x3e,
	0xfb, 0xde, 0xf9, 0xec, 0x63, 0xef, 0x0f, 0xb6, 0x7b, 0x7f, 0xf0, 0xb5, 0xf7, 0x07, 0x2f, 0x8f,
	0xb9, 0xa6, 0xd7, 0x2a, 0x8b, 0x15, 0x16, 0xd2, 0x91, 0x4d, 0x4d, 0x0e, 0x1b, 0xac, 0x21, 0xaa,
	0xc1, 0x50, 0x65, 0xc1, 0xc9, 0xb6, 0xfc, 0x48, 0x29, 0x2a, 0xa3, 0xee, 0x02, 0x51, 0x69, 0x75,
	0x9d, 0x12, 0xc8, 0x37, 0xd9, 0xdf, 0x8a, 0xde, 0x4b, 0x70, 0xd9, 0xb8, 0x6d, 0xfa, 0xf6, 0x37,
	0x00, 0x00, 0xff, 0xff, 0x16, 0x84, 0xce, 0xc2, 0xc2, 0x01, 0x00, 0x00,
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
	if len(m.DestinationDomain) > 0 {
		i -= len(m.DestinationDomain)
		copy(dAtA[i:], m.DestinationDomain)
		i = encodeVarintMint(dAtA, i, uint64(len(m.DestinationDomain)))
		i--
		dAtA[i] = 0x2a
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
	l = len(m.DestinationDomain)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationDomain", wireType)
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
			m.DestinationDomain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
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