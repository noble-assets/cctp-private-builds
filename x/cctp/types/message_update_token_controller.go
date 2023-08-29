package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateTokenController = "update_token_controller"

var _ sdk.Msg = &MsgUpdateTokenController{}

func NewMsgUpdateTokenController(from string, newTokenController string) *MsgUpdateTokenController {
	return &MsgUpdateTokenController{
		From:               from,
		NewTokenController: newTokenController,
	}
}

func (msg *MsgUpdateTokenController) Route() string {
	return RouterKey
}

func (msg *MsgUpdateTokenController) Type() string {
	return TypeMsgUpdateTokenController
}

func (msg *MsgUpdateTokenController) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (msg *MsgUpdateTokenController) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateTokenController) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid from address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.NewTokenController)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid token controller address (%s)", err)
	}
	return nil
}
