package keeper

import (
	"context"

	"github.com/strangelove-ventures/noble-cctp-router-private/x/cctp/types"
)

func (k Keeper) BurnMessageVersion(_ context.Context, _ *types.QueryBurnMessageVersionRequest) (*types.QueryBurnMessageVersionResponse, error) {
	return &types.QueryBurnMessageVersionResponse{Version: types.MessageBodyVersion}, nil
}
