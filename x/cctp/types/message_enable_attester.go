package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
)

const TypeMsgEnableAttester = "enable_attester"

var _ sdk.Msg = &MsgEnableAttester{}

func NewMsgEnableAttester(from string, attester string) *MsgEnableAttester {
	return &MsgEnableAttester{
		From:     from,
		Attester: attester,
	}
}

func (msg *MsgEnableAttester) Route() string {
	return RouterKey
}

func (msg *MsgEnableAttester) Type() string {
	return TypeMsgEnableAttester
}

func (msg *MsgEnableAttester) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (msg *MsgEnableAttester) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEnableAttester) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid from address (%s)", err)
	}

	attester := common.FromHex(msg.Attester)
	if len(attester) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid attester")
	}

	return nil
}
