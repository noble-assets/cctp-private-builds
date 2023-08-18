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
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "this message sender cannot update the authority")
	}

	k.SetAttesterManager(ctx, msg.NewAttesterManager)

	return &types.MsgUpdateAttesterManagerResponse{}, nil
}
