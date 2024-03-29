package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/circlefin/noble-cctp-private-builds/testutil/keeper"
	"github.com/circlefin/noble-cctp-private-builds/testutil/nullify"
	routerkeeper "github.com/circlefin/noble-cctp-private-builds/x/router/keeper"
	"github.com/circlefin/noble-cctp-private-builds/x/router/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestAllowedSourceDomainSenderQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.RouterKeeper(t)
	queryServer := routerkeeper.NewQueryServer(keeper)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNAllowedSourceDomainSender(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryAllowedSourceDomainSenderRequest
		response *types.QueryAllowedSourceDomainSenderResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryAllowedSourceDomainSenderRequest{
				DomainId: msgs[0].DomainId,
				Address:  msgs[0].Address,
			},
			response: &types.QueryAllowedSourceDomainSenderResponse{AllowedSourceDomainSender: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryAllowedSourceDomainSenderRequest{
				DomainId: msgs[1].DomainId,
				Address:  msgs[1].Address,
			},
			response: &types.QueryAllowedSourceDomainSenderResponse{AllowedSourceDomainSender: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryAllowedSourceDomainSenderRequest{
				DomainId: uint32(32),
				Address:  []byte("32"),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := queryServer.AllowedSourceDomainSender(wctx, tc.request)
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

func TestAllowedSourceDomainSenderQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.RouterKeeper(t)
	queryServer := routerkeeper.NewQueryServer(keeper)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNAllowedSourceDomainSender(keeper, ctx, 5)
	AllowedSourceDomainSender := make([]types.AllowedSourceDomainSender, len(msgs))
	copy(AllowedSourceDomainSender, msgs)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllowedSourceDomainSendersRequest {
		return &types.QueryAllowedSourceDomainSendersRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(AllowedSourceDomainSender); i += step {
			resp, err := queryServer.AllowedSourceDomainSenders(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.AllowedSourceDomainSenders), step)
			require.Subset(t,
				nullify.Fill(AllowedSourceDomainSender),
				nullify.Fill(resp.AllowedSourceDomainSenders),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(AllowedSourceDomainSender); i += step {
			resp, err := queryServer.AllowedSourceDomainSenders(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.AllowedSourceDomainSenders), step)
			require.Subset(t,
				nullify.Fill(AllowedSourceDomainSender),
				nullify.Fill(resp.AllowedSourceDomainSenders),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := queryServer.AllowedSourceDomainSenders(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(AllowedSourceDomainSender), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(AllowedSourceDomainSender),
			nullify.Fill(resp.AllowedSourceDomainSenders),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := queryServer.AllowedSourceDomainSenders(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
