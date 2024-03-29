package cli_test

import (
	"fmt"
	"github.com/circlefin/noble-cctp-private-builds/x/router/client/cli"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/circlefin/noble-cctp-private-builds/testutil/network"
	"github.com/circlefin/noble-cctp-private-builds/testutil/nullify"
	"github.com/circlefin/noble-cctp-private-builds/x/router/types"
)

func networkWithInFlightPacketObjects(t *testing.T, n uint32) (*network.Network, []types.InFlightPacket) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := uint32(0); i < n; i++ {
		InFlightPacket := types.InFlightPacket{
			SourceDomain: i,
			Channel:      "channel-" + strconv.Itoa(int(i)),
			Port:         "port-" + strconv.Itoa(int(i)),
			Nonce:        uint64(i),
			Sequence:     uint64(i),
		}
		nullify.Fill(&InFlightPacket)
		state.InFlightPackets = append(state.InFlightPackets, InFlightPacket)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.InFlightPackets
}

func TestShowInFlightPacket(t *testing.T) {
	net, objs := networkWithInFlightPacketObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc      string
		channelId string
		portId    string
		sequence  string

		args []string
		err  error
		obj  types.InFlightPacket
	}{
		{
			desc:      "found",
			channelId: objs[0].Channel,
			portId:    objs[0].Port,
			sequence:  strconv.FormatUint(objs[0].Sequence, 10),
			args:      common,
			obj:       objs[0],
		},
		{
			desc:      "not found",
			channelId: "123",
			portId:    "456",
			sequence:  strconv.FormatUint(789, 10),
			args:      common,
			err:       status.Error(codes.NotFound, "not found"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.channelId,
				tc.portId,
				tc.sequence,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowInFlightPacket(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetInFlightPacketResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.InFlightPacket)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.InFlightPacket),
				)
			}
		})
	}
}
