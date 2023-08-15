package keeper

import (
	"github.com/strangelove-ventures/noble-cctp-router-private/x/router/types"
)

var _ types.QueryServer = Keeper{}
