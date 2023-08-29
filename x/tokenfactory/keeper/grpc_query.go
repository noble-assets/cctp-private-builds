package keeper

import (
	"github.com/circlefin/noble-cctp-private-builds/x/tokenfactory/types"
)

var _ types.QueryServer = Keeper{}
