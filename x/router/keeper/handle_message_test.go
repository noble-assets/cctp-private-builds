package keeper_test

import (
	"math/big"
	"strconv"
	"testing"

	keepertest "github.com/circlefin/noble-cctp-private-builds/testutil/keeper"
	"github.com/circlefin/noble-cctp-private-builds/testutil/sample"
	cctptypes "github.com/circlefin/noble-cctp-private-builds/x/cctp/types"
	"github.com/circlefin/noble-cctp-private-builds/x/router/keeper"
	"github.com/circlefin/noble-cctp-private-builds/x/router/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestInvalidOuterMessage(t *testing.T) {
	routerKeeper, ctx := keepertest.RouterKeeper(t)

	msg := []byte("not a valid outer message")
	err := routerKeeper.HandleMessage(ctx, msg)

	require.ErrorContains(t, err, cctptypes.ErrParsingMessage.Error())
}

func TestInvalidMessageBodyNoOp(t *testing.T) {
	routerKeeper, ctx := keepertest.RouterKeeper(t)

	msg := bytesFromMessage(keeper.Message{
		Version:           1,
		SourceDomain:      2,
		DestinationDomain: 3,
		Nonce:             4,
		Sender:            fillByteArray(0, 32),
		Recipient:         fillByteArray(32, 32),
		DestinationCaller: fillByteArray(64, 32),
		MessageBody:       []byte("not a valid message body"),
	})

	err := routerKeeper.HandleMessage(ctx, msg)
	require.Nil(t, err)
}

// valid forward, found forward, ack error, existing mint -> forward packet
func TestForwardOnAckErrWithExistingMint(t *testing.T) {
	routerKeeper, ctx := keepertest.RouterKeeper(t)

	sourceDomainSender, nonce := string(fillByteArray(0, 32)), uint64(4)
	port, channel, sequence := "transfer", "channel-10", uint64(0)

	routerKeeper.SetIBCForward(ctx, types.StoreIBCForwardMetadata{
		SourceDomainSender: sourceDomainSender,
		Metadata: &types.IBCForwardMetadata{
			Nonce:                nonce,
			Port:                 port,
			Channel:              channel,
			DestinationReceiver:  "12345",
			Memo:                 "12345",
			TimeoutInNanoseconds: 0,
		},
		AckError: true,
	})

	routerKeeper.SetMint(ctx, types.Mint{
		SourceDomainSender: sourceDomainSender,
		Nonce:              nonce,
		Amount: &sdk.Coin{
			Denom:  "uusdc",
			Amount: sdk.NewInt(10000),
		},
		DestinationDomain: "",
		MintRecipient:     "12345",
	})

	_, found := routerKeeper.GetInFlightPacket(ctx, port, channel, sequence)
	require.False(t, found)

	msg := bytesFromMessage(keeper.Message{
		Version:           1,
		SourceDomain:      2,
		DestinationDomain: 3,
		Nonce:             nonce,
		Sender:            []byte(sourceDomainSender),
		Recipient:         fillByteArray(32, 32),
		DestinationCaller: fillByteArray(64, 32),
		MessageBody:       createMockMetadata(nonce, channel, sdk.Bech32PrefixAccAddr, sample.AccAddress(), "12345"),
	})

	err := routerKeeper.HandleMessage(ctx, msg)
	require.Nil(t, err)

	_, found = routerKeeper.GetInFlightPacket(ctx, channel, port, sequence)
	require.True(t, found)
}

// valid forward, found forward, ack error, no mint -> panic
func TestForwardOnAckErrWithNoMint(t *testing.T) {
	routerKeeper, ctx := keepertest.RouterKeeper(t)

	sourceDomainSender, nonce := string(fillByteArray(0, 32)), uint64(4)
	port, channel, sequence := "transfer", "channel-10", uint64(0)

	routerKeeper.SetIBCForward(ctx, types.StoreIBCForwardMetadata{
		SourceDomainSender: sourceDomainSender,
		Metadata: &types.IBCForwardMetadata{
			Nonce:                nonce,
			Port:                 port,
			Channel:              channel,
			DestinationReceiver:  "12345",
			Memo:                 "12345",
			TimeoutInNanoseconds: 0,
		},
		AckError: true,
	})

	_, found := routerKeeper.GetInFlightPacket(ctx, port, channel, sequence)
	require.False(t, found)

	msg := bytesFromMessage(keeper.Message{
		Version:           1,
		SourceDomain:      2,
		DestinationDomain: 3,
		Nonce:             nonce,
		Sender:            []byte(sourceDomainSender),
		Recipient:         fillByteArray(32, 32),
		DestinationCaller: fillByteArray(64, 32),
		MessageBody:       createMockMetadata(nonce, channel, sdk.Bech32PrefixAccAddr, sample.AccAddress(), "12345"),
	})

	require.Panics(t, func() {
		_ = routerKeeper.HandleMessage(ctx, msg)
	})

}

// valid forward, found forward, no ack error -> ErrHandleMessage
func TestForwardWithFoundForwardAndNoAckError(t *testing.T) {
	routerKeeper, ctx := keepertest.RouterKeeper(t)

	sourceDomainSender, nonce := string(fillByteArray(0, 32)), uint64(4)
	port, channel, sequence := "transfer", "channel-10", uint64(0)

	routerKeeper.SetIBCForward(ctx, types.StoreIBCForwardMetadata{
		SourceDomainSender: sourceDomainSender,
		Metadata: &types.IBCForwardMetadata{
			Nonce:                nonce,
			Port:                 port,
			Channel:              channel,
			DestinationReceiver:  "12345",
			Memo:                 "12345",
			TimeoutInNanoseconds: 0,
		},
		AckError: false,
	})

	_, found := routerKeeper.GetInFlightPacket(ctx, port, channel, sequence)
	require.False(t, found)

	msg := bytesFromMessage(keeper.Message{
		Version:           1,
		SourceDomain:      2,
		DestinationDomain: 3,
		Nonce:             nonce,
		Sender:            []byte(sourceDomainSender),
		Recipient:         fillByteArray(32, 32),
		DestinationCaller: fillByteArray(64, 32),
		MessageBody:       createMockMetadata(nonce, channel, sdk.Bech32PrefixAccAddr, sample.AccAddress(), "12345"),
	})

	err := routerKeeper.HandleMessage(ctx, msg)
	require.ErrorIs(t, err, types.ErrHandleMessage)
}

// valid forward, no forward -> set forward, if existing mint -> forward packet
func TestForwardWithNoForwardFoundAndExistingMint(t *testing.T) {
	routerKeeper, ctx := keepertest.RouterKeeper(t)

	sourceDomain, sourceDomainSender, nonce := uint32(1), string(fillByteArray(0, 32)), uint64(4)
	port, channel, sequence := "transfer", "channel-10", uint64(0)

	routerKeeper.SetMint(ctx, types.Mint{
		SourceDomainSender: sourceDomainSender,
		Nonce:              nonce,
		Amount: &sdk.Coin{
			Denom:  "uusdc",
			Amount: sdk.NewInt(10000),
		},
		DestinationDomain: "",
		MintRecipient:     "12345",
	})

	_, found := routerKeeper.GetInFlightPacket(ctx, port, channel, sequence)
	require.False(t, found)

	msg := bytesFromMessage(keeper.Message{
		Version:           1,
		SourceDomain:      2,
		DestinationDomain: 3,
		Nonce:             nonce,
		Sender:            []byte(sourceDomainSender),
		Recipient:         fillByteArray(32, 32),
		DestinationCaller: fillByteArray(64, 32),
		MessageBody:       createMockMetadata(nonce, channel, sdk.Bech32PrefixAccAddr, sample.AccAddress(), "12345"),
	})

	err := routerKeeper.HandleMessage(ctx, msg)
	require.Nil(t, err)

	_, found = routerKeeper.GetInFlightPacket(ctx, channel, port, sequence)
	require.True(t, found)

	forward, found := routerKeeper.GetIBCForward(ctx, sourceDomain, sourceDomainSender, nonce)
	require.True(t, found)
	require.Equal(t, sourceDomainSender, forward.SourceDomainSender)
	require.Equal(t, nonce, forward.Metadata.Nonce)
}

// valid forward, no forward -> set forward, if no mint -> return nil (mint hasn't come in yet)
func TestForwardWithNoForwardFoundAndNoMint(t *testing.T) {
	routerKeeper, ctx := keepertest.RouterKeeper(t)

	sourceDomain, sourceDomainSender, nonce := uint32(1), string(fillByteArray(0, 32)), uint64(4)
	_, channel, _ := "transfer", "channel-10", uint64(0)

	msg := bytesFromMessage(keeper.Message{
		Version:           1,
		SourceDomain:      2,
		DestinationDomain: 3,
		Nonce:             nonce,
		Sender:            []byte(sourceDomainSender),
		Recipient:         fillByteArray(32, 32),
		DestinationCaller: fillByteArray(64, 32),
		MessageBody:       createMockMetadata(nonce, channel, sdk.Bech32PrefixAccAddr, sample.AccAddress(), "12345"),
	})

	err := routerKeeper.HandleMessage(ctx, msg)
	require.Nil(t, err)

	forward, found := routerKeeper.GetIBCForward(ctx, sourceDomain, sourceDomainSender, nonce)
	require.True(t, found)
	require.Equal(t, sourceDomainSender, forward.SourceDomainSender)
	require.Equal(t, nonce, forward.Metadata.Nonce)
}

// valid mint, set mint, no forward -> return nil (forward hasn't come in yet)
func TestMintWithNoForward(t *testing.T) {
	routerKeeper, ctx := keepertest.RouterKeeper(t)

	sourceDomain, sourceDomainSender, nonce := uint32(1), string(fillByteArray(0, 32)), uint64(4)
	port, channel, sequence := "transfer", "channel-10", uint64(0)

	_, found := routerKeeper.GetInFlightPacket(ctx, port, channel, sequence)
	require.False(t, found)

	msg := bytesFromMessage(keeper.Message{
		Version:           1,
		SourceDomain:      2,
		DestinationDomain: 3,
		Nonce:             nonce,
		Sender:            []byte(sourceDomainSender),
		Recipient:         fillByteArray(32, 32),
		DestinationCaller: fillByteArray(64, 32),
		MessageBody: bytesFromBurnMessage(keeper.BurnMessage{
			Version:       0,
			BurnToken:     fillByteArray(0, 32),
			MintRecipient: fillByteArray(0, 32),
			Amount:        *big.NewInt(10000),
			MessageSender: fillByteArray(0, 32),
		}),
	})

	err := routerKeeper.HandleMessage(ctx, msg)
	require.Nil(t, err)

	mint, found := routerKeeper.GetMint(ctx, sourceDomain, sourceDomainSender, nonce)
	require.True(t, found)
	require.Equal(t, sourceDomainSender, mint.SourceDomainSender)
	require.Equal(t, nonce, mint.Nonce)

}

// TODO add test for valid mint with no token pair found once integrated with cctp

// valid mint, set mint, existing forward -> forward packet
func TestMintWithExistingForward(t *testing.T) {
	routerKeeper, ctx := keepertest.RouterKeeper(t)

	sourceDomain, sourceDomainSender, nonce := uint32(1), string(fillByteArray(0, 32)), uint64(4)
	port, channel, sequence := "transfer", "channel-10", uint64(0)

	routerKeeper.SetIBCForward(ctx, types.StoreIBCForwardMetadata{
		SourceDomainSender: sourceDomainSender,
		Metadata: &types.IBCForwardMetadata{
			Nonce:                nonce,
			Port:                 port,
			Channel:              channel,
			DestinationReceiver:  "12345",
			Memo:                 "12345",
			TimeoutInNanoseconds: 0,
		},
		AckError: false,
	})

	_, found := routerKeeper.GetInFlightPacket(ctx, port, channel, sequence)
	require.False(t, found)

	msg := bytesFromMessage(keeper.Message{
		Version:           1,
		SourceDomain:      2,
		DestinationDomain: 3,
		Nonce:             nonce,
		Sender:            []byte(sourceDomainSender),
		Recipient:         fillByteArray(32, 32),
		DestinationCaller: fillByteArray(64, 32),
		MessageBody: bytesFromBurnMessage(keeper.BurnMessage{
			Version:       0,
			BurnToken:     fillByteArray(0, 32),
			MintRecipient: fillByteArray(0, 32),
			Amount:        *big.NewInt(10000),
			MessageSender: fillByteArray(0, 32),
		}),
	})

	err := routerKeeper.HandleMessage(ctx, msg)
	require.Nil(t, err)

	_, found = routerKeeper.GetInFlightPacket(ctx, channel, port, sequence)
	require.True(t, found)

	mint, found := routerKeeper.GetMint(ctx, sourceDomain, sourceDomainSender, nonce)
	require.True(t, found)
	require.Equal(t, sourceDomainSender, mint.SourceDomainSender)
	require.Equal(t, nonce, mint.Nonce)

}
