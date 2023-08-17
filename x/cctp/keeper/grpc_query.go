package keeper

import (
	"github.com/circlefin/noble-cctp-router-private/x/cctp/types"
)

var _ types.QueryServer = Keeper{}
