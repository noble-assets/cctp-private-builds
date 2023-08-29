package keeper

import (
	"context"

	"github.com/circlefin/noble-cctp-private-builds/x/cctp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RemoveRemoteTokenMessenger(goCtx context.Context, msg *types.MsgRemoveRemoteTokenMessenger) (*types.MsgRemoveRemoteTokenMessengerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	owner := k.GetOwner(ctx)
	if owner != msg.From {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "this message sender cannot remove remote token messengers")
	}

	existingRemoteTokenMessenger, found := k.GetRemoteTokenMessenger(ctx, msg.DomainId)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrRemoteTokenMessengerNotFound, "no remote token messenger was found for this domain")
	}

	k.DeleteRemoteTokenMessenger(ctx, msg.DomainId)

	event := types.RemoteTokenMessengerRemoved{
		Domain:               msg.DomainId,
		RemoteTokenMessenger: existingRemoteTokenMessenger.Address,
	}
	err := ctx.EventManager().EmitTypedEvent(&event)

	return &types.MsgRemoveRemoteTokenMessengerResponse{}, err
}
