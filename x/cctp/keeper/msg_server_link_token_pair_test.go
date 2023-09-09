package keeper_test

import (
	"testing"

	keepertest "github.com/circlefin/noble-cctp-private-builds/testutil/keeper"
	"github.com/circlefin/noble-cctp-private-builds/testutil/sample"
	"github.com/circlefin/noble-cctp-private-builds/x/cctp/keeper"
	"github.com/circlefin/noble-cctp-private-builds/x/cctp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

/*
 * Happy path
 * Authority not set
 * Invalid authority
 * Existing token pair found
 */
func TestLinkTokenPairHappyPath(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	tokenController := sample.AccAddress()
	testkeeper.SetTokenController(ctx, tokenController)

	message := types.MsgLinkTokenPair{
		From:         tokenController,
		RemoteDomain: 1,
		RemoteToken:  []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xAB, 0xCD},
		LocalToken:   "uusdc",
	}

	_, err := server.LinkTokenPair(sdk.WrapSDKContext(ctx), &message)
	require.NoError(t, err)

	actual, found := testkeeper.GetTokenPair(ctx, message.RemoteDomain, message.RemoteToken)
	require.True(t, found)
	require.Equal(t, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xAB, 0xCD}, actual.RemoteToken)
	require.Equal(t, message.RemoteDomain, actual.RemoteDomain)
	require.Equal(t, message.LocalToken, actual.LocalToken)
}

func TestLinkTokenPairAuthorityNotSet(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	message := types.MsgLinkTokenPair{
		From:         sample.AccAddress(),
		RemoteDomain: 1,
		RemoteToken:  []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xAB, 0xCD},
		LocalToken:   "uusdc",
	}

	require.Panicsf(t, func() {
		_, _ = server.LinkTokenPair(sdk.WrapSDKContext(ctx), &message)
	}, "cctp owner not found in state")
}

func TestLinkTokenPairInvalidAuthority(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	tokenController := sample.AccAddress()
	testkeeper.SetTokenController(ctx, tokenController)

	message := types.MsgLinkTokenPair{
		From:         "not authority",
		RemoteDomain: 1,
		RemoteToken:  []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xAB, 0xCD},
		LocalToken:   "uusdc",
	}

	_, err := server.LinkTokenPair(sdk.WrapSDKContext(ctx), &message)
	require.ErrorIs(t, types.ErrUnauthorized, err)
}

func TestLinkTokenPairExistingTokenPairFound(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	tokenController := sample.AccAddress()
	testkeeper.SetTokenController(ctx, tokenController)

	existingTokenPair := types.TokenPair{
		RemoteDomain: 1,
		RemoteToken:  []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xAB, 0xCD},
		LocalToken:   "uusdc",
	}

	testkeeper.SetTokenPair(ctx, existingTokenPair)

	message := types.MsgLinkTokenPair{
		From:         tokenController,
		RemoteDomain: existingTokenPair.RemoteDomain,
		RemoteToken:  existingTokenPair.RemoteToken,
		LocalToken:   existingTokenPair.LocalToken,
	}

	_, err := server.LinkTokenPair(sdk.WrapSDKContext(ctx), &message)
	require.ErrorIs(t, types.ErrTokenPairAlreadyFound, err)
}
