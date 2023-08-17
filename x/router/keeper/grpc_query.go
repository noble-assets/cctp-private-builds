package keeper

import (
	"github.com/circlefin/noble-cctp-private-builds/x/router/types"
)

var _ types.QueryServer = Keeper{}
