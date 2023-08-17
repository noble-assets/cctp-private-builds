package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/circlefin/noble-cctp-router-private/testutil/keeper"
	"github.com/circlefin/noble-cctp-router-private/testutil/nullify"
	"github.com/circlefin/noble-cctp-router-private/x/cctp/types"
)

func TestBurningAndMintingPausedQuery(t *testing.T) {
	keeper, ctx := keepertest.CctpKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	BurningAndMintingPaused := types.BurningAndMintingPaused{Paused: true}
	keeper.SetBurningAndMintingPaused(ctx, BurningAndMintingPaused)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetBurningAndMintingPausedRequest
		response *types.QueryGetBurningAndMintingPausedResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetBurningAndMintingPausedRequest{},
			response: &types.QueryGetBurningAndMintingPausedResponse{Paused: BurningAndMintingPaused},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.BurningAndMintingPaused(wctx, tc.request)
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
