package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateSignatureThreshold = "update_signature_threshold"

var _ sdk.Msg = &MsgUpdateSignatureThreshold{}

func NewMsgUpdateSignatureThreshold(from string, amount uint32) *MsgUpdateSignatureThreshold {
	return &MsgUpdateSignatureThreshold{
		From:   from,
		Amount: amount,
	}
}

func (msg *MsgUpdateSignatureThreshold) Route() string {
	return RouterKey
}

func (msg *MsgUpdateSignatureThreshold) Type() string {
	return TypeMsgUpdateSignatureThreshold
}

func (msg *MsgUpdateSignatureThreshold) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (msg *MsgUpdateSignatureThreshold) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateSignatureThreshold) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid from address (%s)", err)
	}
	return nil
}
