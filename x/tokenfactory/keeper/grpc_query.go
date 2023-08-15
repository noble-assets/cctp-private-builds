package keeper

import (
	"github.com/strangelove-ventures/noble-cctp-router-private/x/tokenfactory/types"
)

var _ types.QueryServer = Keeper{}
