package keeper

import (
	"context"

	"github.com/strangelove-ventures/noble-cctp-router-private/x/cctp/types"
)

func (k Keeper) LocalDomain(_ context.Context, _ *types.QueryLocalDomainRequest) (*types.QueryLocalDomainResponse, error) {
	return &types.QueryLocalDomainResponse{DomainId: types.NobleDomainId}, nil
}
