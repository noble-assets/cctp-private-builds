package keeper_test

import (
	"testing"

	keepertest "github.com/circlefin/noble-cctp-private-builds/testutil/keeper"
	"github.com/circlefin/noble-cctp-private-builds/x/cctp/keeper"
	"github.com/circlefin/noble-cctp-private-builds/x/cctp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

/*
 * Happy path
 * Sending and receiving messages is paused
 * Message body is too long
 * Recipient is empty
 */
func TestSendMessageHappyPath(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	paused := types.SendingAndReceivingMessagesPaused{Paused: false}
	testkeeper.SetSendingAndReceivingMessagesPaused(ctx, paused)

	nonce := types.Nonce{
		Nonce: 5,
	}
	testkeeper.SetNextAvailableNonce(ctx, nonce)

	msg := types.MsgSendMessage{
		From:              "cosmos1x8rynykqla7cnc0tf2f3xn0wa822ztt788yd5a",
		DestinationDomain: 3,
		Recipient:         []byte("12345678901234567890123456789012"),
		MessageBody:       []byte("It's not about money, it's about sending a message"),
	}

	resp, err := server.SendMessage(sdk.WrapSDKContext(ctx), &msg)
	require.Nil(t, err)
	require.Equal(t, nonce.Nonce, resp.Nonce)

	nextNonce, found := testkeeper.GetNextAvailableNonce(ctx)
	require.True(t, found)
	require.Equal(t, nonce.Nonce+1, nextNonce.Nonce)
}

func TestSendMessageSendingAndReceivingMessagesPaused(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	paused := types.SendingAndReceivingMessagesPaused{Paused: true}
	testkeeper.SetSendingAndReceivingMessagesPaused(ctx, paused)

	_, err := server.SendMessage(sdk.WrapSDKContext(ctx), &types.MsgSendMessage{From: "cosmos1x8rynykqla7cnc0tf2f3xn0wa822ztt788yd5a"})
	require.ErrorIs(t, types.ErrSendMessage, err)
	require.Contains(t, err.Error(), "sending and receiving messages is paused")
}

func TestSendMessageMessageBodyTooLong(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	paused := types.SendingAndReceivingMessagesPaused{Paused: false}
	testkeeper.SetSendingAndReceivingMessagesPaused(ctx, paused)

	max := types.MaxMessageBodySize{Amount: 5}
	testkeeper.SetMaxMessageBodySize(ctx, max)

	nonce := types.Nonce{
		Nonce: 5,
	}
	testkeeper.SetNextAvailableNonce(ctx, nonce)

	msg := types.MsgSendMessage{
		From:              "cosmos1x8rynykqla7cnc0tf2f3xn0wa822ztt788yd5a",
		DestinationDomain: 3,
		Recipient:         []byte("12345678901234567890123456789012"),
		MessageBody:       []byte("It's not about money, it's about sending a message"),
	}

	_, err := server.SendMessage(sdk.WrapSDKContext(ctx), &msg)
	require.ErrorIs(t, types.ErrSendMessage, err)
	require.Contains(t, err.Error(), "message body exceeds max size")
}

func TestSendMessageRecipientEmpty(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	paused := types.SendingAndReceivingMessagesPaused{Paused: false}
	testkeeper.SetSendingAndReceivingMessagesPaused(ctx, paused)

	nonce := types.Nonce{
		Nonce: 5,
	}
	testkeeper.SetNextAvailableNonce(ctx, nonce)

	msg := types.MsgSendMessage{
		From:              "cosmos1x8rynykqla7cnc0tf2f3xn0wa822ztt788yd5a",
		DestinationDomain: 3,
		Recipient:         make([]byte, types.MintRecipientLen),
		MessageBody:       []byte("It's not about money, it's about sending a message"),
	}

	_, err := server.SendMessage(sdk.WrapSDKContext(ctx), &msg)
	require.ErrorIs(t, types.ErrSendMessage, err)
	require.Contains(t, err.Error(), "recipient must not be nonzero")
}
