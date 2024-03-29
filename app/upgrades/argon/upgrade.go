package argon

import (
	"fmt"

	"cosmossdk.io/math"
	cctpkeeper "github.com/circlefin/noble-cctp-private-builds/x/cctp/keeper"
	cctptypes "github.com/circlefin/noble-cctp-private-builds/x/cctp/types"
	fiattokenfactorykeeper "github.com/circlefin/noble-cctp-private-builds/x/fiattokenfactory/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradeTypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	paramauthoritykeeper "github.com/strangelove-ventures/paramauthority/x/params/keeper"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	fiatTFKeeper *fiattokenfactorykeeper.Keeper,
	paramauthoritykeeper paramauthoritykeeper.Keeper,
	cctpKeeper *cctpkeeper.Keeper,
) upgradeTypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradeTypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		var authority string
		if ctx.ChainID() == TestnetChainID {
			authority = paramauthoritykeeper.GetAuthority(ctx)
		} else {
			owner, ok := fiatTFKeeper.GetOwner(ctx)
			if !ok {
				return nil, fmt.Errorf("fiat token factory owner not found")
			}

			authority = owner.Address
		}

		denom := fiatTFKeeper.GetMintingDenom(ctx)

		cctpKeeper.SetOwner(ctx, authority)
		cctpKeeper.SetAttesterManager(ctx, authority)
		cctpKeeper.SetPauser(ctx, authority)
		cctpKeeper.SetTokenController(ctx, authority)
		cctpKeeper.SetPerMessageBurnLimit(ctx, cctptypes.PerMessageBurnLimit{Denom: denom.Denom, Amount: math.NewInt(99999999)})
		cctpKeeper.SetBurningAndMintingPaused(ctx, cctptypes.BurningAndMintingPaused{Paused: false})
		cctpKeeper.SetSendingAndReceivingMessagesPaused(ctx, cctptypes.SendingAndReceivingMessagesPaused{Paused: false})
		cctpKeeper.SetMaxMessageBodySize(ctx, cctptypes.MaxMessageBodySize{Amount: 8000})
		cctpKeeper.SetSignatureThreshold(ctx, cctptypes.SignatureThreshold{Amount: 2})

		return mm.RunMigrations(ctx, configurator, vm)
	}
}
