package types

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSetMaxBurnAmountPerMessage = "set_max_burn_amount_per_message"

var _ sdk.Msg = &MsgSetMaxBurnAmountPerMessage{}

func NewMsgSetMaxBurnAmountPerMessage(from string, localToken string, amount math.Int) *MsgSetMaxBurnAmountPerMessage {
	return &MsgSetMaxBurnAmountPerMessage{
		From:       from,
		LocalToken: localToken,
		Amount:     amount,
	}
}

func (msg *MsgSetMaxBurnAmountPerMessage) Route() string {
	return RouterKey
}

func (msg *MsgSetMaxBurnAmountPerMessage) Type() string {
	return TypeMsgSetMaxBurnAmountPerMessage
}

func (msg *MsgSetMaxBurnAmountPerMessage) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (msg *MsgSetMaxBurnAmountPerMessage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetMaxBurnAmountPerMessage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid from address (%s)", err)
	}

	if len(msg.LocalToken) == 0 {
		return sdkerrors.Wrapf(ErrInvalidToken, "local token cannot be empty")
	}

	if msg.Amount.IsNegative() {
		return sdkerrors.Wrapf(ErrInvalidAmount, "amount must be >= 0")
	}
	return nil
}
