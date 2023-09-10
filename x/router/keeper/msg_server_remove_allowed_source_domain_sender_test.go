package keeper_test

import (
	"testing"

	keepertest "github.com/circlefin/noble-cctp-private-builds/testutil/keeper"
	"github.com/circlefin/noble-cctp-private-builds/testutil/sample"
	"github.com/circlefin/noble-cctp-private-builds/x/router/keeper"
	"github.com/circlefin/noble-cctp-private-builds/x/router/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

/*
* Happy path
* Authority not set
* Invalid authority
* Allowed source domain sender not found
 */

func TestRemoveAllowedSourceDomainSenderHappyPath(t *testing.T) {
	testkeeper, ctx := keepertest.RouterKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	owner := sample.AccAddress()
	testkeeper.SetOwner(ctx, owner)

	addMessage := types.MsgAddAllowedSourceDomainSender{
		From:     owner,
		DomainId: 16,
		Address:  []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xAB, 0xCD},
	}

	_, err := server.AddAllowedSourceDomainSender(sdk.WrapSDKContext(ctx), &addMessage)
	require.Nil(t, err)

	removeMessage := types.MsgRemoveAllowedSourceDomainSender{
		From:     owner,
		DomainId: addMessage.DomainId,
		Address:  addMessage.Address,
	}

	_, err = server.RemoveAllowedSourceDomainSender(sdk.WrapSDKContext(ctx), &removeMessage)
	require.Nil(t, err)

	allowed := testkeeper.IsAllowedSourceDomainSender(ctx, removeMessage.DomainId, removeMessage.Address)
	require.False(t, allowed)
}

func TestRemoveAllowedSourceDomainSenderAuthorityNotSet(t *testing.T) {
	testkeeper, ctx := keepertest.RouterKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	message := types.MsgRemoveAllowedSourceDomainSender{
		From:     sample.AccAddress(),
		DomainId: 16,
	}

	require.Panicsf(t, func() {
		_, _ = server.RemoveAllowedSourceDomainSender(sdk.WrapSDKContext(ctx), &message)
	}, "router owner not found in state")
}

func TestRemoveAllowedSourceDomainSenderInvalidAuthority(t *testing.T) {
	testkeeper, ctx := keepertest.RouterKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	owner := sample.AccAddress()
	testkeeper.SetOwner(ctx, owner)

	message := types.MsgRemoveAllowedSourceDomainSender{
		From:     "not the authority address",
		DomainId: 16,
	}

	_, err := server.RemoveAllowedSourceDomainSender(sdk.WrapSDKContext(ctx), &message)
	require.ErrorIs(t, types.ErrUnauthorized, err)
	require.Contains(t, err.Error(), "this message sender cannot remove allowed source domain senders")
}

func TestRemoveAllowedSourceDomainSenderTokenMessengerNotFound(t *testing.T) {
	testkeeper, ctx := keepertest.RouterKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	owner := sample.AccAddress()
	testkeeper.SetOwner(ctx, owner)

	message := types.MsgRemoveAllowedSourceDomainSender{
		From:     owner,
		DomainId: 1,
	}

	_, err := server.RemoveAllowedSourceDomainSender(sdk.WrapSDKContext(ctx), &message)
	require.ErrorIs(t, types.ErrAllowedSourceDomainSenderNotFound, err)
	require.Contains(t, err.Error(), "source domain sender not found")
}
