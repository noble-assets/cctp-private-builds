package types

import (
	"testing"

	"github.com/circlefin/noble-cctp-private-builds/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgReplaceDepositForBurn_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgReplaceDepositForBurn
		err  error
	}{
		{
			name: "invalid from",
			msg: MsgReplaceDepositForBurn{
				From:                 "invalid_address",
				OriginalMessage:      []byte{1, 2, 3},
				OriginalAttestation:  []byte{1, 2, 3},
				NewDestinationCaller: []byte{1, 2, 3},
				NewMintRecipient:     []byte{1, 2, 3},
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "valid from",
			msg: MsgReplaceDepositForBurn{
				From:                 sample.AccAddress(),
				OriginalMessage:      []byte{1, 2, 3},
				OriginalAttestation:  []byte{1, 2, 3},
				NewDestinationCaller: []byte{1, 2, 3},
				NewMintRecipient:     []byte{1, 2, 3},
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
