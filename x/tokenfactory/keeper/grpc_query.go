package keeper

import (
	"github.com/circlefin/noble-cctp-router-private/x/tokenfactory/types"
)

var _ types.QueryServer = Keeper{}
