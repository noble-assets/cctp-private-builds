package argon

import (
	cctptypes "github.com/circlefin/noble-cctp-router-private/x/cctp/types"
	routertypes "github.com/circlefin/noble-cctp-router-private/x/router/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	storeTypes "github.com/cosmos/cosmos-sdk/store/types"
	upgradeTypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

func CreateStoreLoader(upgradeHeight int64) baseapp.StoreLoader {
	storeUpgrades := storeTypes.StoreUpgrades{
		Added: []string{
			cctptypes.ModuleName, routertypes.ModuleName,
		},
	}

	return upgradeTypes.UpgradeStoreLoader(upgradeHeight, &storeUpgrades)
}
