package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/strangelove-ventures/noble-cctp-router-private/testutil/keeper"
	"github.com/strangelove-ventures/noble-cctp-router-private/testutil/nullify"
	"github.com/strangelove-ventures/noble-cctp-router-private/x/tokenfactory/types"
)

func TestBlacklisterQuery(t *testing.T) {
	keeper, ctx := keepertest.TokenfactoryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestBlacklister(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetBlacklisterRequest
		response *types.QueryGetBlacklisterResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetBlacklisterRequest{},
			response: &types.QueryGetBlacklisterResponse{Blacklister: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Blacklister(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}