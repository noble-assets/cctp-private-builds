package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/strangelove-ventures/noble-cctp-router-private/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgPauseSendingAndReceivingMessages_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgPauseSendingAndReceivingMessages
		err  error
	}{
		{
			name: "invalid from",
			msg: MsgPauseSendingAndReceivingMessages{
				From: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid from",
			msg: MsgPauseSendingAndReceivingMessages{
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
