package types

import (
	"testing"

	"github.com/circlefin/noble-cctp-router-private/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgUnlinkTokenPair_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUnlinkTokenPair
		err  error
	}{
		{
			name: "invalid from",
			msg: MsgUnlinkTokenPair{
				From:         "invalid_address",
				RemoteDomain: 1,
				RemoteToken:  "0x12345",
				LocalToken:   "uusdc",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "valid from",
			msg: MsgUnlinkTokenPair{
				From:         sample.AccAddress(),
				RemoteDomain: 1,
				RemoteToken:  "12345",
				LocalToken:   "uusdc",
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
