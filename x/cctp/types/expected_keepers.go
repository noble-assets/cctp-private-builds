package types

import (
	fiattokenfactorytypes "github.com/circlefin/noble-cctp-private-builds/x/fiattokenfactory/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
	// Methods imported from account should be defined here
}

type BankKeeper interface {
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
}

type FiatTokenfactoryKeeper interface {
	Burn(ctx sdk.Context, msg *fiattokenfactorytypes.MsgBurn) (*fiattokenfactorytypes.MsgBurnResponse, error)
	Mint(ctx sdk.Context, msg *fiattokenfactorytypes.MsgMint) (*fiattokenfactorytypes.MsgMintResponse, error)
	GetMintingDenom(ctx sdk.Context) (val fiattokenfactorytypes.MintingDenom)
}

type RouterKeeper interface {
	HandleMessage(ctx sdk.Context, msg []byte) error
}
