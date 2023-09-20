package keeper

import (
	"context"

	"github.com/circlefin/noble-cctp-private-builds/x/cctp/types"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateTokenController(goCtx context.Context, msg *types.MsgUpdateTokenController) (*types.MsgUpdateTokenControllerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	currentOwner := k.GetOwner(ctx)
	if currentOwner != msg.From {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "this message sender cannot update the authority")
	}

	currentTokenController := k.GetTokenController(ctx)
	k.SetTokenController(ctx, msg.NewTokenController)

	event := types.TokenControllerUpdated{
		PreviousTokenController: currentTokenController,
		NewTokenController:      msg.NewTokenController,
	}

	err := ctx.EventManager().EmitTypedEvent(&event)

	return &types.MsgUpdateTokenControllerResponse{}, err
}
