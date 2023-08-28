package keeper

import (
	"bytes"
	"context"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/circlefin/noble-cctp-private-builds/x/cctp/types"
)

func (k msgServer) SendMessageWithCaller(goCtx context.Context, msg *types.MsgSendMessageWithCaller) (*types.MsgSendMessageWithCallerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	emptyByteArr := make([]byte, types.DestinationCallerLen)
	if len(msg.DestinationCaller) != types.DestinationCallerLen || bytes.Equal(msg.DestinationCaller, emptyByteArr) {
		return nil, sdkerrors.Wrap(types.ErrSendMessage, "destination caller must be nonzero")
	}

	nonce := k.ReserveAndIncrementNonce(ctx)

	sender := make([]byte, 32)
	copy(sender[12:], sdk.MustAccAddressFromBech32(msg.From))

	err := k.sendMessage(
		ctx,
		msg.DestinationDomain,
		msg.Recipient,
		msg.DestinationCaller,
		sender,
		nonce.Nonce,
		msg.MessageBody)

	return &types.MsgSendMessageWithCallerResponse{Nonce: nonce.Nonce}, err
}
