package types

import (
	"testing"

	"github.com/circlefin/noble-cctp-private-builds/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestAddRemoteTokenMessenger_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgAddRemoteTokenMessenger
		err  error
	}{
		{
			name: "invalid from",
			msg: MsgAddRemoteTokenMessenger{
				From:     "invalid_address",
				DomainId: uint32(123),
				Address:  []byte{0x1, 0x23},
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "valid from",
			msg: MsgAddRemoteTokenMessenger{
				From:     sample.AccAddress(),
				DomainId: uint32(123),
				Address:  []byte{0x1, 0x23},
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
