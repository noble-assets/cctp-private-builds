package types

import (
	"testing"

	"github.com/circlefin/noble-cctp-private-builds/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgUnpauseSendingAndReceivingMessages_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUnpauseSendingAndReceivingMessages
		err  error
	}{
		{
			name: "invalid from",
			msg: MsgUnpauseSendingAndReceivingMessages{
				From: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid from",
			msg: MsgUnpauseSendingAndReceivingMessages{
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
