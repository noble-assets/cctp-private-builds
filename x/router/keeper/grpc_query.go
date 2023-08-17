package keeper

import (
	"github.com/circlefin/noble-cctp-router-private/x/router/types"
)

var _ types.QueryServer = Keeper{}
