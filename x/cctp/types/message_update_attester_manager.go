package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateAttesterManager = "update_attester_manager"

var _ sdk.Msg = &MsgUpdateAttesterManager{}

func NewMsgUpdateAttesterManager(from string, newAttesterManager string) *MsgUpdateAttesterManager {
	return &MsgUpdateAttesterManager{
		From:               from,
		NewAttesterManager: newAttesterManager,
	}
}

func (msg *MsgUpdateAttesterManager) Route() string {
	return RouterKey
}

func (msg *MsgUpdateAttesterManager) Type() string {
	return TypeMsgUpdateAttesterManager
}

func (msg *MsgUpdateAttesterManager) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (msg *MsgUpdateAttesterManager) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateAttesterManager) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid from address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.NewAttesterManager)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid attester manager address (%s)", err)
	}
	return nil
}
