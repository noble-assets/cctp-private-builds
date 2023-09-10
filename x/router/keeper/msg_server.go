package keeper

import (
	"github.com/circlefin/noble-cctp-private-builds/x/router/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// msgServerRouterKeeper defines the router keeper methods required by the message server.
type msgServerRouterKeeper interface {
	GetOwner(ctx sdk.Context) (owner string)
	SetOwner(ctx sdk.Context, owner string)
	GetPendingOwner(ctx sdk.Context) (pendingOwner string, found bool)
	SetPendingOwner(ctx sdk.Context, pendingOwner string)
	DeletePendingOwner(ctx sdk.Context)
	IsAllowedSourceDomainSender(ctx sdk.Context, domainID uint32, sourceDomainSender []byte) (allowed bool)
	AddAllowedSourceDomainSender(ctx sdk.Context, domainID uint32, sourceDomainSender []byte)
	DeleteAllowedSourceDomainSender(ctx sdk.Context, domainID uint32, sourceDomainSender []byte)
}

type msgServer struct {
	keeper msgServerRouterKeeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper msgServerRouterKeeper) *msgServer {
	return &msgServer{keeper: keeper}
}

var _ types.MsgServer = msgServer{}
