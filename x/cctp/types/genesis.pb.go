// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: noble/cctp/v1/genesis.proto

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

// GenesisState defines the cctp module's genesis state.
type GenesisState struct {
	Owner                             string                             `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	AttesterManager                   string                             `protobuf:"bytes,3,opt,name=attester_manager,json=attesterManager,proto3" json:"attester_manager,omitempty"`
	Pauser                            string                             `protobuf:"bytes,4,opt,name=pauser,proto3" json:"pauser,omitempty"`
	TokenController                   string                             `protobuf:"bytes,5,opt,name=token_controller,json=tokenController,proto3" json:"token_controller,omitempty"`
	AttesterList                      []Attester                         `protobuf:"bytes,6,rep,name=attester_list,json=attesterList,proto3" json:"attester_list"`
	PerMessageBurnLimitList           []PerMessageBurnLimit              `protobuf:"bytes,7,rep,name=per_message_burn_limit_list,json=perMessageBurnLimitList,proto3" json:"per_message_burn_limit_list"`
	BurningAndMintingPaused           *BurningAndMintingPaused           `protobuf:"bytes,8,opt,name=burning_and_minting_paused,json=burningAndMintingPaused,proto3" json:"burning_and_minting_paused,omitempty"`
	SendingAndReceivingMessagesPaused *SendingAndReceivingMessagesPaused `protobuf:"bytes,9,opt,name=sending_and_receiving_messages_paused,json=sendingAndReceivingMessagesPaused,proto3" json:"sending_and_receiving_messages_paused,omitempty"`
	MaxMessageBodySize                *MaxMessageBodySize                `protobuf:"bytes,10,opt,name=max_message_body_size,json=maxMessageBodySize,proto3" json:"max_message_body_size,omitempty"`
	NextAvailableNonce                *Nonce                             `protobuf:"bytes,11,opt,name=next_available_nonce,json=nextAvailableNonce,proto3" json:"next_available_nonce,omitempty"`
	SignatureThreshold                *SignatureThreshold                `protobuf:"bytes,12,opt,name=signature_threshold,json=signatureThreshold,proto3" json:"signature_threshold,omitempty"`
	TokenPairList                     []TokenPair                        `protobuf:"bytes,13,rep,name=token_pair_list,json=tokenPairList,proto3" json:"token_pair_list"`
	UsedNoncesList                    []Nonce                            `protobuf:"bytes,14,rep,name=used_nonces_list,json=usedNoncesList,proto3" json:"used_nonces_list"`
	TokenMessengerList                []RemoteTokenMessenger             `protobuf:"bytes,15,rep,name=token_messenger_list,json=tokenMessengerList,proto3" json:"token_messenger_list"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_45b63e7866d42c8d, []int{0}
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

func (m *GenesisState) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *GenesisState) GetAttesterManager() string {
	if m != nil {
		return m.AttesterManager
	}
	return ""
}

func (m *GenesisState) GetPauser() string {
	if m != nil {
		return m.Pauser
	}
	return ""
}

func (m *GenesisState) GetTokenController() string {
	if m != nil {
		return m.TokenController
	}
	return ""
}

func (m *GenesisState) GetAttesterList() []Attester {
	if m != nil {
		return m.AttesterList
	}
	return nil
}

func (m *GenesisState) GetPerMessageBurnLimitList() []PerMessageBurnLimit {
	if m != nil {
		return m.PerMessageBurnLimitList
	}
	return nil
}

func (m *GenesisState) GetBurningAndMintingPaused() *BurningAndMintingPaused {
	if m != nil {
		return m.BurningAndMintingPaused
	}
	return nil
}

func (m *GenesisState) GetSendingAndReceivingMessagesPaused() *SendingAndReceivingMessagesPaused {
	if m != nil {
		return m.SendingAndReceivingMessagesPaused
	}
	return nil
}

func (m *GenesisState) GetMaxMessageBodySize() *MaxMessageBodySize {
	if m != nil {
		return m.MaxMessageBodySize
	}
	return nil
}

func (m *GenesisState) GetNextAvailableNonce() *Nonce {
	if m != nil {
		return m.NextAvailableNonce
	}
	return nil
}

func (m *GenesisState) GetSignatureThreshold() *SignatureThreshold {
	if m != nil {
		return m.SignatureThreshold
	}
	return nil
}

func (m *GenesisState) GetTokenPairList() []TokenPair {
	if m != nil {
		return m.TokenPairList
	}
	return nil
}

func (m *GenesisState) GetUsedNoncesList() []Nonce {
	if m != nil {
		return m.UsedNoncesList
	}
	return nil
}

func (m *GenesisState) GetTokenMessengerList() []RemoteTokenMessenger {
	if m != nil {
		return m.TokenMessengerList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "noble.cctp.v1.GenesisState")
}

func init() { proto.RegisterFile("noble/cctp/v1/genesis.proto", fileDescriptor_45b63e7866d42c8d) }

var fileDescriptor_45b63e7866d42c8d = []byte{
	// 684 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xcf, 0x4e, 0xdc, 0x3a,
	0x14, 0xc6, 0x67, 0xf8, 0x77, 0xc1, 0x30, 0x80, 0x7c, 0xe7, 0x5e, 0x52, 0xa8, 0xa6, 0x40, 0xd5,
	0x16, 0xaa, 0x36, 0x53, 0xe8, 0xaa, 0xcb, 0x99, 0x56, 0x54, 0xaa, 0x98, 0x0a, 0x65, 0x58, 0xb5,
	0x8b, 0xc8, 0x49, 0x0e, 0xc1, 0x6a, 0x62, 0x47, 0xb6, 0x67, 0x3a, 0xb0, 0xec, 0x13, 0x54, 0xea,
	0x4b, 0xb1, 0x64, 0xd9, 0x55, 0x55, 0xc1, 0x8b, 0x54, 0xb1, 0x63, 0x50, 0x42, 0xaa, 0x76, 0x17,
	0x9f, 0xf3, 0x9d, 0xdf, 0xe7, 0x13, 0x1f, 0x1b, 0x6d, 0x30, 0x1e, 0x24, 0xd0, 0x0d, 0x43, 0x95,
	0x75, 0xc7, 0x7b, 0xdd, 0x18, 0x18, 0x48, 0x2a, 0xdd, 0x4c, 0x70, 0xc5, 0x71, 0x4b, 0x27, 0xdd,
	0x3c, 0xe9, 0x8e, 0xf7, 0xd6, 0xdb, 0x31, 0x8f, 0xb9, 0xce, 0x74, 0xf3, 0x2f, 0x23, 0x5a, 0xbf,
	0x5f, 0x26, 0x10, 0xa5, 0x40, 0x2a, 0x10, 0x45, 0xd6, 0x2d, 0x67, 0x83, 0x91, 0x60, 0x94, 0xc5,
	0x3e, 0x61, 0x91, 0x9f, 0x52, 0xa6, 0xf2, 0xef, 0x8c, 0x8c, 0x24, 0x44, 0x85, 0x7e, 0xb7, 0xac,
	0x4f, 0xc9, 0xc4, 0x4f, 0x41, 0x4a, 0x12, 0x83, 0x1f, 0xf0, 0xe8, 0xcc, 0x97, 0xf4, 0x1c, 0x0a,
	0xe9, 0xbd, 0xb2, 0x94, 0x71, 0x16, 0xda, 0xd4, 0xd3, 0x72, 0x2a, 0x03, 0x71, 0x4b, 0x19, 0x09,
	0xe6, 0x27, 0x34, 0xa5, 0xaa, 0x5e, 0x2b, 0x20, 0xe5, 0x0a, 0x7c, 0xc5, 0x3f, 0x01, 0xd3, 0x45,
	0xc0, 0xe2, 0x9b, 0x6e, 0x5e, 0x95, 0xb5, 0x12, 0x58, 0x64, 0xbb, 0x11, 0x10, 0x02, 0x1d, 0xe7,
	0xab, 0xc2, 0x49, 0x96, 0x1b, 0x7b, 0x52, 0x29, 0xa5, 0x31, 0x23, 0x6a, 0x24, 0xc0, 0x57, 0xa7,
	0x02, 0xe4, 0x29, 0x4f, 0xac, 0xb0, 0x53, 0x16, 0x9a, 0x8d, 0x64, 0x84, 0x16, 0x7b, 0xd8, 0xfe,
	0x36, 0x8f, 0x96, 0xde, 0x9a, 0x63, 0x1a, 0x2a, 0xa2, 0x00, 0xb7, 0xd1, 0x2c, 0xff, 0xcc, 0x40,
	0x38, 0x53, 0x9b, 0xcd, 0x9d, 0x05, 0xcf, 0x2c, 0xf0, 0x2e, 0x5a, 0xb5, 0x47, 0xe1, 0xa7, 0x84,
	0x91, 0x18, 0x84, 0x33, 0xad, 0x05, 0x2b, 0x36, 0x3e, 0x30, 0x61, 0xfc, 0x3f, 0x9a, 0xd3, 0x5b,
	0x15, 0xce, 0x8c, 0x16, 0x14, 0xab, 0x1c, 0x61, 0xdc, 0x43, 0xce, 0x94, 0xe0, 0x49, 0x02, 0xc2,
	0x99, 0x35, 0x08, 0x1d, 0x7f, 0x7d, 0x13, 0xc6, 0x7d, 0xd4, 0xba, 0x71, 0x4b, 0xa8, 0x54, 0xce,
	0xdc, 0xe6, 0xf4, 0xce, 0xe2, 0xfe, 0x9a, 0x5b, 0x9a, 0x20, 0xb7, 0x57, 0x68, 0xfa, 0x33, 0x17,
	0x3f, 0x1e, 0x34, 0xbc, 0x25, 0x5b, 0x73, 0x48, 0xa5, 0xc2, 0x27, 0x68, 0xa3, 0xfe, 0xa0, 0x0c,
	0xf1, 0x1f, 0x4d, 0xdc, 0xae, 0x10, 0x8f, 0x40, 0x0c, 0x4c, 0x41, 0x7f, 0x24, 0xd8, 0x61, 0x2e,
	0x2f, 0xe0, 0x6b, 0xd9, 0xdd, 0x94, 0xf6, 0x09, 0xd1, 0xfa, 0xef, 0xc7, 0xd0, 0x99, 0xdf, 0x6c,
	0xee, 0x2c, 0xee, 0x3f, 0xae, 0xd8, 0xf4, 0x4d, 0x41, 0x8f, 0x45, 0x03, 0x23, 0x3f, 0xd2, 0x6a,
	0x6f, 0x2d, 0xa8, 0x4f, 0xe0, 0x2f, 0x4d, 0xf4, 0xe8, 0xaf, 0xc6, 0xc3, 0x59, 0xd0, 0x86, 0x2f,
	0x2a, 0x86, 0x43, 0x53, 0xdb, 0x63, 0x91, 0x67, 0x2b, 0x8b, 0x66, 0x64, 0x61, 0xbd, 0x25, 0xff,
	0x24, 0xc1, 0xc7, 0xe8, 0xbf, 0xda, 0x0b, 0xe4, 0x20, 0xed, 0xb9, 0x55, 0xf1, 0x1c, 0x90, 0x89,
	0xfd, 0x61, 0x3c, 0x3a, 0x1b, 0xd2, 0x73, 0xf0, 0x70, 0x7a, 0x27, 0x86, 0x0f, 0x50, 0x9b, 0xc1,
	0x44, 0xf9, 0x64, 0x4c, 0x68, 0x42, 0x82, 0x04, 0x7c, 0x7d, 0xf5, 0x9c, 0x45, 0x0d, 0x6d, 0x57,
	0xa0, 0xef, 0xf3, 0x9c, 0x87, 0xf3, 0x8a, 0x9e, 0x2d, 0xd0, 0x31, 0xec, 0xa1, 0x7f, 0x6b, 0x6e,
	0x81, 0xb3, 0x54, 0xbb, 0xb7, 0xa1, 0x55, 0x1e, 0x5b, 0xa1, 0x87, 0xe5, 0x9d, 0x18, 0x3e, 0x40,
	0x2b, 0xb7, 0x17, 0xc6, 0xcc, 0x4d, 0x4b, 0xcf, 0x8d, 0x53, 0xe1, 0x1d, 0xe7, 0xaa, 0x23, 0x42,
	0xed, 0x28, 0xb6, 0x94, 0x0d, 0xe8, 0x19, 0x79, 0x83, 0x56, 0xf3, 0x3f, 0x68, 0x3a, 0x93, 0x06,
	0xb4, 0xac, 0x41, 0xb5, 0xfd, 0x15, 0x90, 0xe5, 0xbc, 0x46, 0x07, 0xa4, 0xa6, 0x7c, 0x44, 0xed,
	0xca, 0x3b, 0x62, 0x48, 0x2b, 0x9a, 0xf4, 0xb0, 0x42, 0xf2, 0xf4, 0xcb, 0xa3, 0x37, 0x36, 0xb0,
	0xfa, 0x02, 0x8c, 0x55, 0x29, 0x9a, 0xc3, 0xdf, 0xcd, 0xcc, 0x37, 0x57, 0xa7, 0xf2, 0xbb, 0x2a,
	0x48, 0x2a, 0xfb, 0x07, 0x17, 0x57, 0x9d, 0xe6, 0xe5, 0x55, 0xa7, 0xf9, 0xf3, 0xaa, 0xd3, 0xfc,
	0x7a, 0xdd, 0x69, 0x5c, 0x5e, 0x77, 0x1a, 0xdf, 0xaf, 0x3b, 0x8d, 0x0f, 0xcf, 0x62, 0xaa, 0x4e,
	0x47, 0x81, 0x1b, 0xf2, 0xb4, 0x1b, 0x52, 0x11, 0x26, 0x70, 0x42, 0x59, 0x57, 0x5b, 0x3f, 0xd7,
	0x8f, 0xcc, 0xc4, 0xbc, 0x35, 0xea, 0x2c, 0x03, 0x19, 0xcc, 0xe9, 0x47, 0xe6, 0xe5, 0xaf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x68, 0x59, 0x52, 0xeb, 0x18, 0x06, 0x00, 0x00,
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
	if len(m.TokenMessengerList) > 0 {
		for iNdEx := len(m.TokenMessengerList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TokenMessengerList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x7a
		}
	}
	if len(m.UsedNoncesList) > 0 {
		for iNdEx := len(m.UsedNoncesList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UsedNoncesList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x72
		}
	}
	if len(m.TokenPairList) > 0 {
		for iNdEx := len(m.TokenPairList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TokenPairList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x6a
		}
	}
	if m.SignatureThreshold != nil {
		{
			size, err := m.SignatureThreshold.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x62
	}
	if m.NextAvailableNonce != nil {
		{
			size, err := m.NextAvailableNonce.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x5a
	}
	if m.MaxMessageBodySize != nil {
		{
			size, err := m.MaxMessageBodySize.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	if m.SendingAndReceivingMessagesPaused != nil {
		{
			size, err := m.SendingAndReceivingMessagesPaused.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	if m.BurningAndMintingPaused != nil {
		{
			size, err := m.BurningAndMintingPaused.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if len(m.PerMessageBurnLimitList) > 0 {
		for iNdEx := len(m.PerMessageBurnLimitList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PerMessageBurnLimitList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.AttesterList) > 0 {
		for iNdEx := len(m.AttesterList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AttesterList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.TokenController) > 0 {
		i -= len(m.TokenController)
		copy(dAtA[i:], m.TokenController)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.TokenController)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Pauser) > 0 {
		i -= len(m.Pauser)
		copy(dAtA[i:], m.Pauser)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Pauser)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.AttesterManager) > 0 {
		i -= len(m.AttesterManager)
		copy(dAtA[i:], m.AttesterManager)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.AttesterManager)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
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
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.AttesterManager)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.Pauser)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.TokenController)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.AttesterList) > 0 {
		for _, e := range m.AttesterList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PerMessageBurnLimitList) > 0 {
		for _, e := range m.PerMessageBurnLimitList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.BurningAndMintingPaused != nil {
		l = m.BurningAndMintingPaused.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.SendingAndReceivingMessagesPaused != nil {
		l = m.SendingAndReceivingMessagesPaused.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.MaxMessageBodySize != nil {
		l = m.MaxMessageBodySize.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.NextAvailableNonce != nil {
		l = m.NextAvailableNonce.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.SignatureThreshold != nil {
		l = m.SignatureThreshold.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.TokenPairList) > 0 {
		for _, e := range m.TokenPairList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.UsedNoncesList) > 0 {
		for _, e := range m.UsedNoncesList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.TokenMessengerList) > 0 {
		for _, e := range m.TokenMessengerList {
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
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
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttesterManager", wireType)
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
			m.AttesterManager = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pauser", wireType)
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
			m.Pauser = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenController", wireType)
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
			m.TokenController = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AttesterList", wireType)
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
			m.AttesterList = append(m.AttesterList, Attester{})
			if err := m.AttesterList[len(m.AttesterList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PerMessageBurnLimitList", wireType)
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
			m.PerMessageBurnLimitList = append(m.PerMessageBurnLimitList, PerMessageBurnLimit{})
			if err := m.PerMessageBurnLimitList[len(m.PerMessageBurnLimitList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurningAndMintingPaused", wireType)
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
			if m.BurningAndMintingPaused == nil {
				m.BurningAndMintingPaused = &BurningAndMintingPaused{}
			}
			if err := m.BurningAndMintingPaused.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SendingAndReceivingMessagesPaused", wireType)
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
			if m.SendingAndReceivingMessagesPaused == nil {
				m.SendingAndReceivingMessagesPaused = &SendingAndReceivingMessagesPaused{}
			}
			if err := m.SendingAndReceivingMessagesPaused.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxMessageBodySize", wireType)
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
			if m.MaxMessageBodySize == nil {
				m.MaxMessageBodySize = &MaxMessageBodySize{}
			}
			if err := m.MaxMessageBodySize.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextAvailableNonce", wireType)
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
			if m.NextAvailableNonce == nil {
				m.NextAvailableNonce = &Nonce{}
			}
			if err := m.NextAvailableNonce.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignatureThreshold", wireType)
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
			if m.SignatureThreshold == nil {
				m.SignatureThreshold = &SignatureThreshold{}
			}
			if err := m.SignatureThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenPairList", wireType)
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
			m.TokenPairList = append(m.TokenPairList, TokenPair{})
			if err := m.TokenPairList[len(m.TokenPairList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UsedNoncesList", wireType)
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
			m.UsedNoncesList = append(m.UsedNoncesList, Nonce{})
			if err := m.UsedNoncesList[len(m.UsedNoncesList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenMessengerList", wireType)
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
			m.TokenMessengerList = append(m.TokenMessengerList, RemoteTokenMessenger{})
			if err := m.TokenMessengerList[len(m.TokenMessengerList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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