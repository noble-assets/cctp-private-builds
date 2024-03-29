package types

import (
	"testing"

	"github.com/circlefin/noble-cctp-private-builds/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgAcceptOwner_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgAcceptOwner
		err  error
	}{
		{
			name: "invalid from",
			msg: MsgAcceptOwner{
				From: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "valid from",
			msg: MsgAcceptOwner{
				From: sample.AccAddress(),
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
