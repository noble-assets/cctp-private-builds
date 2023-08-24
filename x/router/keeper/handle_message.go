package keeper

import (
	"strconv"
	"time"

	cctptypes "github.com/circlefin/noble-cctp-router-private/x/cctp/types"
	"github.com/circlefin/noble-cctp-router-private/x/router/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
)

// MinimumRelativePacketTimeoutTimestamp is the default minimum packet timeout timestamp (in nanoseconds)
// relative to the current block timestamp of the counterparty chain provided by the client state. Incoming packets
// with IBCForwardMetadata cannot have a timeout smaller than this value in order to prevent DoS styled attacks
// on relayer resources.
var MinimumRelativePacketTimeoutTimestamp = uint64((time.Duration(30) * time.Second).Nanoseconds())

func (k Keeper) HandleMessage(ctx sdk.Context, msg []byte) error {

	// parse outer message
	outerMessage, err := new(cctptypes.Message).Parse(msg)
	if err != nil {
		return err
	}

	// parse internal message into IBCForward
	if ibcForward, err := DecodeIBCForward(outerMessage.MessageBody); err == nil {
		if storedForward, ok := k.GetIBCForward(ctx, outerMessage.SourceDomain, string(outerMessage.Sender), ibcForward.Nonce); ok {
			if storedForward.AckError {
				if existingMint, ok := k.GetMint(ctx, outerMessage.SourceDomain, string(outerMessage.Sender), ibcForward.Nonce); ok {
					return k.ForwardPacket(ctx, ibcForward, existingMint)
				}
				panic("unexpected state")
			}

			return sdkerrors.Wrapf(types.ErrHandleMessage, "previous operation still in progress")
		}
		// this is the first time we are seeing this forward info -> store it.
		k.SetIBCForward(ctx, types.StoreIBCForwardMetadata{
			SourceDomain:       outerMessage.SourceDomain,
			SourceDomainSender: string(outerMessage.Sender),
			Metadata:           &ibcForward,
		})
		if existingMint, ok := k.GetMint(ctx, outerMessage.SourceDomain, string(outerMessage.Sender), ibcForward.Nonce); ok {
			return k.ForwardPacket(ctx, ibcForward, existingMint)
		}
		return nil
	}

	// try to parse internal message into burn (representing a remote burn -> local mint)
	if burnMessage, err := new(cctptypes.BurnMessage).Parse(outerMessage.MessageBody); err == nil {
		tokenPair, found := k.cctpKeeper.GetTokenPair(ctx, outerMessage.SourceDomain, burnMessage.BurnToken)
		if !found {
			return sdkerrors.Wrapf(types.ErrHandleMessage, "unable to find local token denom for this burn")
		}

		addr, err := sdk.Bech32ifyAddressBytes("noble", burnMessage.MintRecipient[12:])
		if err != nil {
			return sdkerrors.Wrapf(types.ErrHandleMessage, "error bech32 encoding mint recipient address")
		}

		// message is a Mint
		mint := types.Mint{
			SourceDomain:       outerMessage.SourceDomain,
			SourceDomainSender: string(outerMessage.Sender),
			Nonce:              outerMessage.Nonce,
			Amount: &sdk.Coin{
				Denom:  tokenPair.LocalToken,
				Amount: sdk.NewIntFromBigInt(burnMessage.Amount.BigInt()),
			},
			DestinationDomain: strconv.Itoa(int(outerMessage.DestinationDomain)),
			MintRecipient:     addr,
		}
		k.SetMint(ctx, mint)
		if existingIBCForward, found := k.GetIBCForward(ctx, outerMessage.SourceDomain, string(burnMessage.MessageSender), outerMessage.Nonce); found {
			return k.ForwardPacket(ctx, *existingIBCForward.Metadata, mint)
		}
	}

	return nil
}

func (k Keeper) ForwardPacket(ctx sdk.Context, ibcForward types.IBCForwardMetadata, mint types.Mint) error {
	timeout := ibcForward.TimeoutInNanoseconds
	if timeout < MinimumRelativePacketTimeoutTimestamp {
		timeout = transfertypes.DefaultRelativePacketTimeoutTimestamp
	}

	transfer := &transfertypes.MsgTransfer{
		SourcePort:    ibcForward.Port,
		SourceChannel: ibcForward.Channel,
		Token:         *mint.Amount,
		Sender:        mint.MintRecipient,
		Receiver:      ibcForward.DestinationReceiver,
		TimeoutHeight: clienttypes.Height{
			RevisionNumber: 0,
			RevisionHeight: 0,
		},
		TimeoutTimestamp: uint64(ctx.BlockTime().UnixNano()) + timeout,
		Memo:             ibcForward.Memo,
	}

	res, err := k.transferKeeper.Transfer(sdk.WrapSDKContext(ctx), transfer)
	if err != nil {
		return err
	}

	inFlightPacket := types.InFlightPacket{
		SourceDomain:       mint.SourceDomain,
		SourceDomainSender: mint.SourceDomainSender,
		Nonce:              mint.Nonce,

		Channel:  ibcForward.Channel,
		Port:     ibcForward.Port,
		Sequence: res.Sequence,
	}

	k.SetInFlightPacket(ctx, inFlightPacket)

	return nil
}
