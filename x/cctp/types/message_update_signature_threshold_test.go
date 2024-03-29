package types

import (
	"testing"

	"github.com/circlefin/noble-cctp-private-builds/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgUpdateSignatureThreshold_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateSignatureThreshold
		err  error
	}{
		{
			name: "invalid from",
			msg: MsgUpdateSignatureThreshold{
				From:   "invalid_address",
				Amount: 1,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "valid from",
			msg: MsgUpdateSignatureThreshold{
				From:   sample.AccAddress(),
				Amount: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
