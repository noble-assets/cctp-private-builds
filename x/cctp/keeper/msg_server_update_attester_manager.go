package keeper

import (
	"context"

	"github.com/circlefin/noble-cctp-private-builds/x/cctp/types"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateAttesterManager(goCtx context.Context, msg *types.MsgUpdateAttesterManager) (*types.MsgUpdateAttesterManagerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	currentOwner := k.GetOwner(ctx)
	if currentOwner != msg.From {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "this message sender cannot update the attester manager")
	}

	currentAttesterManager := k.GetAttesterManager(ctx)
	k.SetAttesterManager(ctx, msg.NewAttesterManager)

	event := types.AttesterManagerUpdated{
		PreviousAttesterManager: currentAttesterManager,
		NewAttesterManager:      msg.NewAttesterManager,
	}
	err := ctx.EventManager().EmitTypedEvent(&event)

	return &types.MsgUpdateAttesterManagerResponse{}, err
}
