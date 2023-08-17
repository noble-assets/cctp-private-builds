package types

import (
	"context"

	cctptypes "github.com/circlefin/noble-cctp-private-builds/x/cctp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
)

// TransferKeeper defines the expected transfer keeper
type TransferKeeper interface {
	Transfer(ctx context.Context, msg *types.MsgTransfer) (*types.MsgTransferResponse, error)
}

// CctpKeeper defines the expected cctp keeper
type CctpKeeper interface {
	GetTokenPair(ctx sdk.Context, remoteDomain uint32, remoteToken []byte) (val cctptypes.TokenPair, found bool)
}
