package radon

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	storeTypes "github.com/cosmos/cosmos-sdk/store/types"

	// GlobalFee
	globalFeeTypes "github.com/circlefin/noble-cctp-private-builds/x/globalfee/types"
	// Tariff
	tariffTypes "github.com/circlefin/noble-cctp-private-builds/x/tariff/types"
	// Upgrade
	upgradeTypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

func CreateStoreLoader(upgradeHeight int64) baseapp.StoreLoader {
	storeUpgrades := storeTypes.StoreUpgrades{
		Added: []string{
			globalFeeTypes.ModuleName, tariffTypes.ModuleName,
		},
	}

	return upgradeTypes.UpgradeStoreLoader(upgradeHeight, &storeUpgrades)
}
