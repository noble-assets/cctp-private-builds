package keeper

import (
	"bytes"
	"encoding/binary"
	"math/big"

	"github.com/circlefin/noble-cctp-private-builds/x/router/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	channelTypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
)

type BurnMessage struct {
	Version       uint32
	BurnToken     []byte
	MintRecipient []byte
	Amount        big.Int
	MessageSender []byte
}

type Message struct {
	Version           uint32
	SourceDomain      uint32
	DestinationDomain uint32
	Nonce             uint64
	Sender            []byte
	Recipient         []byte
	DestinationCaller []byte
	MessageBody       []byte
}

const (
	// Indices of each field in message
	VersionIndex           = 0
	SourceDomainIndex      = 4
	DestinationDomainIndex = 8
	NonceIndex             = 12
	SenderIndex            = 20
	RecipientIndex         = 52
	DestinationCallerIndex = 84
	MessageBodyIndex       = 116

	// Indices of each field in BurnMessage
	BurnMsgVersionIndex = 0
	BurnTokenIndex      = 4
	BurnTokenLen        = 32
	MintRecipientIndex  = 36
	MintRecipientLen    = 32
	AmountIndex         = 68
	MsgSenderIndex      = 100
	MsgSenderLen        = 32
	// 4 byte version + 32 bytes burnToken + 32 bytes mintRecipient + 32 bytes amount + 32 bytes messageSender
	BurnMessageLen = 132

	Bytes32Len = 32
)

func DecodeIBCForward(msg []byte) (types.IBCForwardMetadata, error) {
	const (
		NonceIndex                 = 0
		NonceLength                = 8
		ChannelIndex               = NonceIndex + NonceLength
		ChannelLength              = 8
		DestinationRecipientIndex  = ChannelIndex + ChannelLength
		DestinationRecipientLength = 32
		MemoIndex                  = DestinationRecipientIndex + DestinationRecipientLength
	)

	if len(msg) < MemoIndex {
		return types.IBCForwardMetadata{}, types.ErrDecodingIBCForward
	}

	nonce := binary.BigEndian.Uint64(msg[NonceIndex:ChannelIndex])
	channel := channelTypes.FormatChannelIdentifier(
		binary.BigEndian.Uint64(msg[ChannelIndex:DestinationRecipientIndex]),
	)
	destinationRecipient := sdk.AccAddress(
		bytes.TrimLeft(msg[DestinationRecipientIndex:MemoIndex], string(byte(0))),
	)

	return types.IBCForwardMetadata{
		Nonce:                nonce,
		Port:                 "transfer",
		Channel:              channel,
		DestinationReceiver:  destinationRecipient.String(),
		Memo:                 string(msg[MemoIndex:]),
		TimeoutInNanoseconds: 0,
	}, nil
}

func bytesToBigInt(data []byte) big.Int {
	value := big.Int{}
	value.SetBytes(data)

	return value

}
