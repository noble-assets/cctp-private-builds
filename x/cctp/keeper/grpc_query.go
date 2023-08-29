package keeper

import (
	"github.com/circlefin/noble-cctp-private-builds/x/cctp/types"
)

var _ types.QueryServer = Keeper{}
