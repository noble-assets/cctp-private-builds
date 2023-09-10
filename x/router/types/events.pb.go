// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: router/events.proto

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

//
// Emitted when owner address is updated
// @param previous_owner representing the address of the previous owner
// @param new_owner representing the address of the new owner
type OwnerUpdated struct {
	PreviousOwner string `protobuf:"bytes,1,opt,name=previous_owner,json=previousOwner,proto3" json:"previous_owner,omitempty"`
	NewOwner      string `protobuf:"bytes,2,opt,name=new_owner,json=newOwner,proto3" json:"new_owner,omitempty"`
}

func (m *OwnerUpdated) Reset()         { *m = OwnerUpdated{} }
func (m *OwnerUpdated) String() string { return proto.CompactTextString(m) }
func (*OwnerUpdated) ProtoMessage()    {}
func (*OwnerUpdated) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba99373efa5638de, []int{0}
}
func (m *OwnerUpdated) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OwnerUpdated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OwnerUpdated.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OwnerUpdated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OwnerUpdated.Merge(m, src)
}
func (m *OwnerUpdated) XXX_Size() int {
	return m.Size()
}
func (m *OwnerUpdated) XXX_DiscardUnknown() {
	xxx_messageInfo_OwnerUpdated.DiscardUnknown(m)
}

var xxx_messageInfo_OwnerUpdated proto.InternalMessageInfo

func (m *OwnerUpdated) GetPreviousOwner() string {
	if m != nil {
		return m.PreviousOwner
	}
	return ""
}

func (m *OwnerUpdated) GetNewOwner() string {
	if m != nil {
		return m.NewOwner
	}
	return ""
}

//
// Emitted when an allowed source domain sender is added
// @param domain remote domain
// @param address source domain sender address on domain
type AllowedSourceDomainSenderAdded struct {
	Domain  uint32 `protobuf:"varint,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Address []byte `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *AllowedSourceDomainSenderAdded) Reset()         { *m = AllowedSourceDomainSenderAdded{} }
func (m *AllowedSourceDomainSenderAdded) String() string { return proto.CompactTextString(m) }
func (*AllowedSourceDomainSenderAdded) ProtoMessage()    {}
func (*AllowedSourceDomainSenderAdded) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba99373efa5638de, []int{1}
}
func (m *AllowedSourceDomainSenderAdded) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AllowedSourceDomainSenderAdded) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AllowedSourceDomainSenderAdded.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AllowedSourceDomainSenderAdded) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllowedSourceDomainSenderAdded.Merge(m, src)
}
func (m *AllowedSourceDomainSenderAdded) XXX_Size() int {
	return m.Size()
}
func (m *AllowedSourceDomainSenderAdded) XXX_DiscardUnknown() {
	xxx_messageInfo_AllowedSourceDomainSenderAdded.DiscardUnknown(m)
}

var xxx_messageInfo_AllowedSourceDomainSenderAdded proto.InternalMessageInfo

func (m *AllowedSourceDomainSenderAdded) GetDomain() uint32 {
	if m != nil {
		return m.Domain
	}
	return 0
}

func (m *AllowedSourceDomainSenderAdded) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

//
// Emitted when a allowed source domain sender is removed
// @param domain remote domain
// @param address source domain sender address on domain
type AllowedSourceDomainSenderRemoved struct {
	Domain  uint32 `protobuf:"varint,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Address []byte `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *AllowedSourceDomainSenderRemoved) Reset()         { *m = AllowedSourceDomainSenderRemoved{} }
func (m *AllowedSourceDomainSenderRemoved) String() string { return proto.CompactTextString(m) }
func (*AllowedSourceDomainSenderRemoved) ProtoMessage()    {}
func (*AllowedSourceDomainSenderRemoved) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba99373efa5638de, []int{2}
}
func (m *AllowedSourceDomainSenderRemoved) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AllowedSourceDomainSenderRemoved) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AllowedSourceDomainSenderRemoved.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AllowedSourceDomainSenderRemoved) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllowedSourceDomainSenderRemoved.Merge(m, src)
}
func (m *AllowedSourceDomainSenderRemoved) XXX_Size() int {
	return m.Size()
}
func (m *AllowedSourceDomainSenderRemoved) XXX_DiscardUnknown() {
	xxx_messageInfo_AllowedSourceDomainSenderRemoved.DiscardUnknown(m)
}

var xxx_messageInfo_AllowedSourceDomainSenderRemoved proto.InternalMessageInfo

func (m *AllowedSourceDomainSenderRemoved) GetDomain() uint32 {
	if m != nil {
		return m.Domain
	}
	return 0
}

func (m *AllowedSourceDomainSenderRemoved) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func init() {
	proto.RegisterType((*OwnerUpdated)(nil), "noble.router.OwnerUpdated")
	proto.RegisterType((*AllowedSourceDomainSenderAdded)(nil), "noble.router.AllowedSourceDomainSenderAdded")
	proto.RegisterType((*AllowedSourceDomainSenderRemoved)(nil), "noble.router.AllowedSourceDomainSenderRemoved")
}

func init() { proto.RegisterFile("router/events.proto", fileDescriptor_ba99373efa5638de) }

var fileDescriptor_ba99373efa5638de = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x31, 0x4b, 0xc4, 0x30,
	0x14, 0x80, 0xaf, 0x0e, 0xa7, 0x17, 0xee, 0x1c, 0xaa, 0xc8, 0xa1, 0x10, 0x8e, 0x03, 0xc1, 0xa5,
	0xcd, 0xe0, 0xec, 0x70, 0xe2, 0x2e, 0xe4, 0x14, 0xc1, 0x45, 0xda, 0xe4, 0x59, 0x03, 0x6d, 0x5e,
	0x48, 0xd2, 0x56, 0xff, 0x85, 0x3f, 0xcb, 0xf1, 0x46, 0x47, 0x69, 0xff, 0x88, 0x5c, 0xae, 0x37,
	0xba, 0xb8, 0xe5, 0xbd, 0xef, 0xe3, 0x1b, 0x5e, 0xc8, 0x89, 0xc5, 0xda, 0x83, 0x65, 0xd0, 0x80,
	0xf6, 0x2e, 0x35, 0x16, 0x3d, 0xc6, 0x53, 0x8d, 0x79, 0x09, 0xe9, 0x0e, 0x9d, 0x9f, 0x16, 0x58,
	0x60, 0x00, 0x6c, 0xfb, 0xda, 0x39, 0x4b, 0x4e, 0xa6, 0xf7, 0xad, 0x06, 0xfb, 0x68, 0x64, 0xe6,
	0x41, 0xc6, 0x97, 0xe4, 0xd8, 0x58, 0x68, 0x14, 0xd6, 0xee, 0x05, 0xb7, 0x60, 0x1e, 0x2d, 0xa2,
	0xab, 0x09, 0x9f, 0xed, 0xb7, 0xc1, 0x8e, 0x2f, 0xc8, 0x44, 0x43, 0x3b, 0x18, 0x07, 0xc1, 0x38,
	0xd2, 0xd0, 0x06, 0xb8, 0xe4, 0x84, 0xae, 0xca, 0x12, 0x5b, 0x90, 0x6b, 0xac, 0xad, 0x80, 0x3b,
	0xac, 0x32, 0xa5, 0xd7, 0xa0, 0x25, 0xd8, 0x95, 0x94, 0x20, 0xe3, 0x33, 0x32, 0x96, 0x61, 0x19,
	0xea, 0x33, 0x3e, 0x4c, 0xf1, 0x9c, 0x1c, 0x66, 0x52, 0x5a, 0x70, 0x2e, 0x44, 0xa7, 0x7c, 0x3f,
	0x2e, 0x1f, 0xc8, 0xe2, 0xcf, 0x26, 0x87, 0x0a, 0x9b, 0xff, 0x54, 0x6f, 0x9f, 0xbe, 0x3a, 0x1a,
	0x6d, 0x3a, 0x1a, 0xfd, 0x74, 0x34, 0xfa, 0xec, 0xe9, 0x68, 0xd3, 0xd3, 0xd1, 0x77, 0x4f, 0x47,
	0xcf, 0x37, 0x85, 0xf2, 0x6f, 0x75, 0x9e, 0x0a, 0xac, 0x98, 0x50, 0x56, 0x94, 0xf0, 0xaa, 0x34,
	0x0b, 0x07, 0x4d, 0x84, 0xf0, 0x26, 0x31, 0x56, 0x35, 0x99, 0x87, 0x24, 0xaf, 0x55, 0x29, 0x1d,
	0x7b, 0x67, 0xc3, 0x0f, 0xf8, 0x0f, 0x03, 0x2e, 0x1f, 0x87, 0xeb, 0x5e, 0xff, 0x06, 0x00, 0x00,
	0xff, 0xff, 0x0f, 0x79, 0x18, 0x6e, 0x98, 0x01, 0x00, 0x00,
}

func (m *OwnerUpdated) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OwnerUpdated) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OwnerUpdated) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NewOwner) > 0 {
		i -= len(m.NewOwner)
		copy(dAtA[i:], m.NewOwner)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.NewOwner)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PreviousOwner) > 0 {
		i -= len(m.PreviousOwner)
		copy(dAtA[i:], m.PreviousOwner)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.PreviousOwner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AllowedSourceDomainSenderAdded) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllowedSourceDomainSenderAdded) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AllowedSourceDomainSenderAdded) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.Domain != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Domain))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *AllowedSourceDomainSenderRemoved) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllowedSourceDomainSenderRemoved) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AllowedSourceDomainSenderRemoved) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.Domain != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.Domain))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OwnerUpdated) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PreviousOwner)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.NewOwner)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *AllowedSourceDomainSenderAdded) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Domain != 0 {
		n += 1 + sovEvents(uint64(m.Domain))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func (m *AllowedSourceDomainSenderRemoved) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Domain != 0 {
		n += 1 + sovEvents(uint64(m.Domain))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OwnerUpdated) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: OwnerUpdated: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OwnerUpdated: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousOwner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PreviousOwner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewOwner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewOwner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *AllowedSourceDomainSenderAdded) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: AllowedSourceDomainSenderAdded: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllowedSourceDomainSenderAdded: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Domain", wireType)
			}
			m.Domain = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Domain |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *AllowedSourceDomainSenderRemoved) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: AllowedSourceDomainSenderRemoved: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllowedSourceDomainSenderRemoved: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Domain", wireType)
			}
			m.Domain = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Domain |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
