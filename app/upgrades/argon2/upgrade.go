package argon2

import (
	"cosmossdk.io/math"
	cctpkeeper "github.com/circlefin/noble-cctp-private-builds/x/cctp/keeper"
	cctptypes "github.com/circlefin/noble-cctp-private-builds/x/cctp/types"
	fiattokenfactorykeeper "github.com/circlefin/noble-cctp-private-builds/x/fiattokenfactory/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradeTypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	fiatTFKeeper *fiattokenfactorykeeper.Keeper,
	cctpKeeper *cctpkeeper.Keeper,
) upgradeTypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradeTypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		denom := fiatTFKeeper.GetMintingDenom(ctx)

		cctpKeeper.SetOwner(ctx, "noble1h8nzdzjgp6gn746s42lsp37pkt8wpl82j4r50k")
		cctpKeeper.SetAttesterManager(ctx, "noble1rx6vk9ll2vglwkrrlf8a7cl86lfz55uj8deunm")
		cctpKeeper.SetPauser(ctx, "noble1hvm5pxssempk3jg0tgzugtsk85js42rze7cnxd")
		cctpKeeper.SetTokenController(ctx, "noble1hl7nlkt3vyjzk0c4ytfveemmykw8ectspapcd3")
		cctpKeeper.SetPerMessageBurnLimit(ctx, cctptypes.PerMessageBurnLimit{Denom: denom.Denom, Amount: math.NewInt(99999999)})
		cctpKeeper.SetBurningAndMintingPaused(ctx, cctptypes.BurningAndMintingPaused{Paused: false})
		cctpKeeper.SetSendingAndReceivingMessagesPaused(ctx, cctptypes.SendingAndReceivingMessagesPaused{Paused: false})
		cctpKeeper.SetMaxMessageBodySize(ctx, cctptypes.MaxMessageBodySize{Amount: 8000})
		cctpKeeper.SetSignatureThreshold(ctx, cctptypes.SignatureThreshold{Amount: 2})

		return mm.RunMigrations(ctx, configurator, vm)
	}
}
