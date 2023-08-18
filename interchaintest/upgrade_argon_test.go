package interchaintest_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/strangelove-ventures/interchaintest/v3/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v3/ibc"
	"github.com/stretchr/testify/require"

	cctptypes "github.com/circlefin/noble-cctp-private-builds/x/cctp/types"
)

func testPostArgonUpgrade(
	t *testing.T,
	ctx context.Context,
	noble *cosmos.CosmosChain,
	paramAuthority ibc.Wallet,
	owner ibc.Wallet,
) {

	// queryMaxMsgBodySize, _, err := noble.Validators[0].ExecQuery(ctx, "cctp", "show-max-message-body-size")
	// require.NoError(t, err, "error querying cctp max message body size")

	// t.Log("Original body size BYTE:", string(queryMaxMsgBodySize))

	var maxMsgBodySize cctptypes.MaxMessageBodySize
	// err = json.Unmarshal(queryMaxMsgBodySize, &maxMsgBodySize)
	// require.NoError(t, err, "failed to unmarshall max message body size")

	// t.Log("Original body size:", maxMsgBodySize.Amount)

	_, err := noble.Validators[0].ExecTx(ctx, owner.KeyName(), "cctp", "update-max-message-body-size", "30")
	require.NoError(t, err, "error updating max message body size")

	queryMaxMsgBodySize, _, err := noble.Validators[0].ExecQuery(ctx, "cctp", "show-max-message-body-size")
	require.NoError(t, err, "error querying cctp max message body size")

	err = json.Unmarshal(queryMaxMsgBodySize, &maxMsgBodySize)
	require.NoError(t, err, "failed to unmarshall max message body size")

	expectedSize := cctptypes.MaxMessageBodySize{
		Amount: 30,
	}
	err = json.Unmarshal(queryMaxMsgBodySize, &maxMsgBodySize)
	require.NoError(t, err, "failed to unmarshall max message body size")
	require.Equal(t, expectedSize.Amount, maxMsgBodySize.Amount)

}
