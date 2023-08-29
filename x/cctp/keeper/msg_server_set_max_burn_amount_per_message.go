package keeper

import (
	"context"
	"strings"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/circlefin/noble-cctp-private-builds/x/cctp/types"
)

func (k msgServer) SetMaxBurnAmountPerMessage(goCtx context.Context, msg *types.MsgSetMaxBurnAmountPerMessage) (*types.MsgSetMaxBurnAmountPerMessageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	tokenController := k.GetTokenController(ctx)
	if tokenController != msg.From {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "this message sender cannot set the max burn amount per message")
	}

	newPerMessageBurnLimit := types.PerMessageBurnLimit{
		Denom:  strings.ToLower(msg.LocalToken),
		Amount: msg.Amount,
	}
	k.SetPerMessageBurnLimit(ctx, newPerMessageBurnLimit)

	err := ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgSetMaxBurnAmountPerMessageResponse{}, err
}
