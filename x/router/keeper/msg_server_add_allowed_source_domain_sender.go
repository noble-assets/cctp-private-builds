package keeper

import (
	"context"

	"github.com/circlefin/noble-cctp-private-builds/x/router/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (m msgServer) AddAllowedSourceDomainSender(goCtx context.Context, msg *types.MsgAddAllowedSourceDomainSender) (*types.MsgAddAllowedSourceDomainSenderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	owner := m.keeper.GetOwner(ctx)
	if owner != msg.From {
		return nil, sdkerrors.Wrapf(types.ErrUnauthorized, "this message sender cannot add allowed source domain senders")
	}

	found := m.keeper.IsAllowedSourceDomainSender(ctx, msg.DomainId, msg.Address)
	if found {
		return nil, types.ErrAllowedSourceDomainSenderAlreadyFound
	}

	if len(msg.Address) != 32 {
		return nil, sdkerrors.ErrInvalidAddress
	}

	m.keeper.AddAllowedSourceDomainSender(ctx, msg.DomainId, msg.Address)

	event := types.AllowedSourceDomainSenderAdded{
		Domain:  msg.DomainId,
		Address: msg.Address,
	}
	err := ctx.EventManager().EmitTypedEvent(&event)

	return &types.MsgAddAllowedSourceDomainSenderResponse{}, err
}
