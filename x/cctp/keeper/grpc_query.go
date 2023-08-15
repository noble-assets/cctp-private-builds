package keeper

import (
	"github.com/strangelove-ventures/noble-cctp-router-private/x/cctp/types"
)

var _ types.QueryServer = Keeper{}
