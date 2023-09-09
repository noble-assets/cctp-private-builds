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
* Remote token messenger already found
 */

func TestAddRemoteTokenMessengerHappyPath(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	owner := sample.AccAddress()
	testkeeper.SetOwner(ctx, owner)

	message := types.MsgAddRemoteTokenMessenger{
		From:     owner,
		DomainId: 16,
		Address:  []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xAB, 0xCD},
	}

	_, err := server.AddRemoteTokenMessenger(sdk.WrapSDKContext(ctx), &message)
	require.Nil(t, err)

	actual, found := testkeeper.GetRemoteTokenMessenger(ctx, message.DomainId)
	require.True(t, found)

	require.Equal(t, message.DomainId, actual.DomainId)
	require.Equal(t, message.Address, actual.Address)
}

func TestAddRemoteTokenMessengerAuthorityNotSet(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	message := types.MsgAddRemoteTokenMessenger{
		From:     sample.AccAddress(),
		DomainId: 16,
		Address:  []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xAB, 0xCD},
	}

	require.Panicsf(t, func() {
		_, _ = server.AddRemoteTokenMessenger(sdk.WrapSDKContext(ctx), &message)
	}, "cctp owner not found in state")
}

func TestAddRemoteTokenMessengerInvalidAuthority(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	owner := sample.AccAddress()
	testkeeper.SetOwner(ctx, owner)

	message := types.MsgAddRemoteTokenMessenger{
		From:     "not the authority address",
		DomainId: 16,
		Address:  []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xAB, 0xCD},
	}

	_, err := server.AddRemoteTokenMessenger(sdk.WrapSDKContext(ctx), &message)
	require.ErrorIs(t, types.ErrUnauthorized, err)
	require.Contains(t, err.Error(), "this message sender cannot add remote token messengers")
}

func TestAddRemoteTokenMessengerTokenMessengerAlreadyFound(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	owner := sample.AccAddress()
	testkeeper.SetOwner(ctx, owner)

	existingRemoteTokenMessenger := types.RemoteTokenMessenger{
		DomainId: 3,
		Address:  []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xAB, 0xCD},
	}
	testkeeper.SetRemoteTokenMessenger(ctx, existingRemoteTokenMessenger)

	message := types.MsgAddRemoteTokenMessenger{
		From:     owner,
		DomainId: existingRemoteTokenMessenger.DomainId,
		Address:  []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xAB, 0xCD},
	}

	_, err := server.AddRemoteTokenMessenger(sdk.WrapSDKContext(ctx), &message)
	require.ErrorIs(t, types.ErrRemoteTokenMessengerAlreadyFound, err)
	require.Contains(t, err.Error(), "a remote token messenger for this domain already exists")
}
