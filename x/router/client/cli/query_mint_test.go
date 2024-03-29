package cli_test

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"google.golang.org/grpc/codes"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/status"

	"github.com/circlefin/noble-cctp-private-builds/testutil/network"
	"github.com/circlefin/noble-cctp-private-builds/testutil/nullify"
	"github.com/circlefin/noble-cctp-private-builds/x/router/client/cli"
	"github.com/circlefin/noble-cctp-private-builds/x/router/types"
)

func networkWithMintObjects(t *testing.T, n uint32) (*network.Network, []types.Mint) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := uint32(1); i <= n; i++ {
		addr := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		binary.BigEndian.PutUint32(addr[28:], i)
		Mints := types.Mint{
			SourceDomain:       i,
			SourceDomainSender: addr,
			Nonce:              uint64(i),
			Amount: &sdk.Coin{
				Denom:  "uusdc",
				Amount: sdk.NewInt(1),
			},
			DestinationDomain: 4,
			MintRecipient:     "cosmos1x8rynykqla7cnc0tf2f3xn0wa822ztt788yd5a",
		}
		state.Mints = append(state.Mints, Mints)
		state.AllowedSourceDomainSenders = append(state.AllowedSourceDomainSenders, types.AllowedSourceDomainSender{
			DomainId: i,
			Address:  addr,
		})
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.Mints
}

func TestShowMint(t *testing.T) {
	net, objs := networkWithMintObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc         string
		sourceDomain uint32
		nonce        string

		args []string
		err  error
		obj  types.Mint
	}{
		{
			desc:         "found",
			sourceDomain: objs[0].SourceDomain,
			nonce:        strconv.Itoa(int(objs[0].Nonce)),
			args:         common,
			obj:          objs[0],
		},
		{
			desc:         "not found",
			sourceDomain: uint32(99),
			nonce:        "456",
			args:         common,
			err:          status.Error(codes.NotFound, "not found"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				strconv.Itoa(int(tc.sourceDomain)),
				tc.nonce,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowMint(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetMintResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.Mint)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.Mint),
				)
			}
		})
	}
}
