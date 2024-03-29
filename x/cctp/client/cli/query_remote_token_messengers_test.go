package cli_test

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"google.golang.org/grpc/codes"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/status"

	"github.com/circlefin/noble-cctp-private-builds/testutil/network"
	"github.com/circlefin/noble-cctp-private-builds/testutil/nullify"
	"github.com/circlefin/noble-cctp-private-builds/x/cctp/client/cli"
	"github.com/circlefin/noble-cctp-private-builds/x/cctp/types"
)

func networkWithRemoteTokenMessengerObjects(t *testing.T, n uint32) (*network.Network, []types.RemoteTokenMessenger) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := uint32(1); i <= n; i++ {
		addr := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		binary.BigEndian.PutUint32(addr[28:], i)
		remoteTokenMessenger := types.RemoteTokenMessenger{
			DomainId: i,
			Address:  addr,
		}
		state.TokenMessengerList = append(state.TokenMessengerList, remoteTokenMessenger)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.TokenMessengerList
}

func TestRemoteTokenMessenger(t *testing.T) {
	net, objs := networkWithRemoteTokenMessengerObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc         string
		remoteDomain string
		remoteToken  []byte

		args []string
		err  error
		obj  types.RemoteTokenMessenger
	}{
		{
			desc:         "found",
			remoteDomain: strconv.Itoa(int(objs[0].DomainId)),
			remoteToken:  objs[0].Address,
			args:         common,
			obj:          objs[0],
		},
		{
			desc:         "not found",
			remoteDomain: "notakey",
			remoteToken:  objs[0].Address,
			args:         common,
			err:          status.Error(codes.NotFound, "not found"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{tc.remoteDomain}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdRemoteTokenMessenger(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryRemoteTokenMessengerResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.RemoteTokenMessenger.DomainId)
				require.NotNil(t, resp.RemoteTokenMessenger.Address)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.RemoteTokenMessenger),
				)
			}
		})
	}
}

func TestRemoteTokenMessengers(t *testing.T) {
	net, objs := networkWithRemoteTokenMessengerObjects(t, 5)

	ctx := net.Validators[0].ClientCtx
	request := func(next []byte, offset, limit uint64, total bool) []string {
		args := []string{
			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
		}
		if next == nil {
			args = append(args, fmt.Sprintf("--%s=%d", flags.FlagOffset, offset))
		} else {
			args = append(args, fmt.Sprintf("--%s=%s", flags.FlagPageKey, next))
		}
		args = append(args, fmt.Sprintf("--%s=%d", flags.FlagLimit, limit))
		if total {
			args = append(args, fmt.Sprintf("--%s", flags.FlagCountTotal))
		}
		return args
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(objs); i += step {
			args := request(nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdRemoteTokenMessengers(), args)
			require.NoError(t, err)
			var resp types.QueryRemoteTokenMessengersResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.RemoteTokenMessengers), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.RemoteTokenMessengers),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdRemoteTokenMessengers(), args)
			require.NoError(t, err)
			var resp types.QueryRemoteTokenMessengersResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.RemoteTokenMessengers), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.RemoteTokenMessengers),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdRemoteTokenMessengers(), args)
		require.NoError(t, err)
		var resp types.QueryRemoteTokenMessengersResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(objs),
			nullify.Fill(resp.RemoteTokenMessengers),
		)
	})
}
