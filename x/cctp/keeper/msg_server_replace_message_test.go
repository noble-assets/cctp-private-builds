package keeper_test

import (
	"testing"

	"cosmossdk.io/math"
	keepertest "github.com/circlefin/noble-cctp-private-builds/testutil/keeper"
	"github.com/circlefin/noble-cctp-private-builds/x/cctp/keeper"
	"github.com/circlefin/noble-cctp-private-builds/x/cctp/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

/*
 * Happy path
 * Fails when paused
 * Signature threshold not found
 * Signature verification failed
 * Message body too short
 * Invalid sender
 * Message not originally sent from this domain
 */

func TestReplaceMessageHappyPath(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	paused := types.SendingAndReceivingMessagesPaused{Paused: false}
	testkeeper.SetSendingAndReceivingMessagesPaused(ctx, paused)

	burnMessage := types.BurnMessage{
		Version:       1,
		BurnToken:     make([]byte, 32),
		MintRecipient: make([]byte, 32),
		Amount:        math.NewInt(123456),
		MessageSender: make([]byte, 32),
	}

	burnMessageBytes, err := burnMessage.Bytes()
	require.Nil(t, err)

	// we encode the message sender when sending messages, so we must use an encoded message in the original message
	senderBech32 := "cosmos1x8rynykqla7cnc0tf2f3xn0wa822ztt788yd5a"
	senderEncoded := make([]byte, 32)
	copy(senderEncoded[12:], sdk.MustAccAddressFromBech32(senderBech32))

	originalMessage := types.Message{
		Version:           1,
		SourceDomain:      4, // Noble domain id
		DestinationDomain: 3,
		Nonce:             2,
		Sender:            senderEncoded,
		Recipient:         []byte("recipient01234567890123456789012"),
		DestinationCaller: []byte("destination caller90123456789012"),
		MessageBody:       burnMessageBytes,
	}
	originalMessageBytes, _ := originalMessage.Bytes()

	// generate attestation, set attesters, signature threshold
	signatureThreshold := uint32(2)
	privKeys := generateNPrivateKeys(int(signatureThreshold))
	attesters := getAttestersFromPrivateKeys(privKeys)
	originalAttestation := generateAttestation(originalMessageBytes, privKeys)
	for _, attester := range attesters {
		testkeeper.SetAttester(ctx, attester)
	}
	testkeeper.SetSignatureThreshold(ctx, types.SignatureThreshold{Amount: signatureThreshold})

	msg := types.MsgReplaceMessage{
		From:                 senderBech32,
		OriginalMessage:      originalMessageBytes,
		OriginalAttestation:  originalAttestation,
		NewMessageBody:       []byte("123"),
		NewDestinationCaller: []byte("new destination caller3456789012"),
	}

	_, err = server.ReplaceMessage(sdk.WrapSDKContext(ctx), &msg)
	require.Nil(t, err)
}

func TestReplaceMessageFailsWhenPaused(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	paused := types.SendingAndReceivingMessagesPaused{Paused: true}
	testkeeper.SetSendingAndReceivingMessagesPaused(ctx, paused)

	_, err := server.ReplaceMessage(sdk.WrapSDKContext(ctx), &types.MsgReplaceMessage{})
	require.ErrorIs(t, types.ErrReplaceMessage, err)
	require.Contains(t, err.Error(), "sending and receiving messages are paused")
}

func TestReplaceMessageSignatureThresholdNotFound(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	paused := types.SendingAndReceivingMessagesPaused{Paused: false}
	testkeeper.SetSendingAndReceivingMessagesPaused(ctx, paused)

	burnMessage := types.BurnMessage{
		Version:       1,
		BurnToken:     make([]byte, 32),
		MintRecipient: make([]byte, 32),
		Amount:        math.NewInt(123456),
		MessageSender: make([]byte, 32),
	}

	burnMessageBytes, err := burnMessage.Bytes()
	require.Nil(t, err)

	originalMessage := types.Message{
		Version:           1,
		SourceDomain:      4, // Noble domain id
		DestinationDomain: 3,
		Nonce:             2,
		Sender:            []byte("sender78901234567890123456789012"),
		Recipient:         []byte("recipient01234567890123456789012"),
		DestinationCaller: []byte("destination caller90123456789012"),
		MessageBody:       burnMessageBytes,
	}
	originalMessageBytes, _ := originalMessage.Bytes()

	// generate attestation, set attesters, signature threshold
	signatureThreshold := uint32(2)
	privKeys := generateNPrivateKeys(int(signatureThreshold))
	attesters := getAttestersFromPrivateKeys(privKeys)
	originalAttestation := generateAttestation(originalMessageBytes, privKeys)
	for _, attester := range attesters {
		testkeeper.SetAttester(ctx, attester)
	}

	msg := types.MsgReplaceMessage{
		From:                 string(originalMessage.Sender),
		OriginalMessage:      originalMessageBytes,
		OriginalAttestation:  originalAttestation,
		NewMessageBody:       []byte("123"),
		NewDestinationCaller: []byte("new destination caller3456789012"),
	}

	_, err = server.ReplaceMessage(sdk.WrapSDKContext(ctx), &msg)
	require.ErrorIs(t, types.ErrReplaceMessage, err)
	require.Contains(t, err.Error(), "signature threshold not found")
}

func TestReplaceMessageSignatureVerificationFailed(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	paused := types.SendingAndReceivingMessagesPaused{Paused: false}
	testkeeper.SetSendingAndReceivingMessagesPaused(ctx, paused)

	burnMessage := types.BurnMessage{
		Version:       1,
		BurnToken:     make([]byte, 32),
		MintRecipient: make([]byte, 32),
		Amount:        math.NewInt(123456),
		MessageSender: make([]byte, 32),
	}

	burnMessageBytes, err := burnMessage.Bytes()
	require.Nil(t, err)

	originalMessage := types.Message{
		Version:           1,
		SourceDomain:      4, // Noble domain id
		DestinationDomain: 3,
		Nonce:             2,
		Sender:            []byte("sender78901234567890123456789012"),
		Recipient:         []byte("recipient01234567890123456789012"),
		DestinationCaller: []byte("destination caller90123456789012"),
		MessageBody:       burnMessageBytes,
	}
	originalMessageBytes, _ := originalMessage.Bytes()

	// generate attestation, set attesters, signature threshold
	signatureThreshold := uint32(2)
	privKeys := generateNPrivateKeys(int(signatureThreshold))
	attesters := getAttestersFromPrivateKeys(privKeys)
	originalAttestation := generateAttestation(originalMessageBytes, privKeys)
	for _, attester := range attesters {
		testkeeper.SetAttester(ctx, attester)
	}
	testkeeper.SetSignatureThreshold(ctx, types.SignatureThreshold{Amount: signatureThreshold})

	// corrupt attestation
	originalAttestation[10] = 1

	msg := types.MsgReplaceMessage{
		From:                 string(originalMessage.Sender),
		OriginalMessage:      originalMessageBytes,
		OriginalAttestation:  originalAttestation,
		NewMessageBody:       []byte("123"),
		NewDestinationCaller: []byte("new destination caller3456789012"),
	}

	_, err = server.ReplaceMessage(sdk.WrapSDKContext(ctx), &msg)
	require.ErrorIs(t, types.ErrSignatureVerification, err)
	require.Contains(t, err.Error(), "unable to verify signatures")
}

func TestReplaceMessageMessageBodyTooShort(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	paused := types.SendingAndReceivingMessagesPaused{Paused: false}
	testkeeper.SetSendingAndReceivingMessagesPaused(ctx, paused)

	burnMessage := types.BurnMessage{
		Version:       1,
		BurnToken:     make([]byte, 32),
		MintRecipient: make([]byte, 32),
		Amount:        math.NewInt(123456),
		MessageSender: make([]byte, 32),
	}

	burnMessageBytes, err := burnMessage.Bytes()
	require.Nil(t, err)

	originalMessage := types.Message{
		Version:           1,
		SourceDomain:      4, // Noble domain id
		DestinationDomain: 3,
		Nonce:             2,
		Sender:            []byte("sender78901234567890123456789012"),
		Recipient:         []byte("recipient01234567890123456789012"),
		DestinationCaller: []byte("destination caller90123456789012"),
		MessageBody:       burnMessageBytes,
	}
	originalMessageBytes, _ := originalMessage.Bytes()
	// make it too small
	originalMessageBytes = originalMessageBytes[0:115]

	// generate attestation, set attesters, signature threshold
	signatureThreshold := uint32(2)
	privKeys := generateNPrivateKeys(int(signatureThreshold))
	attesters := getAttestersFromPrivateKeys(privKeys)
	originalAttestation := generateAttestation(originalMessageBytes, privKeys)
	for _, attester := range attesters {
		testkeeper.SetAttester(ctx, attester)
	}
	testkeeper.SetSignatureThreshold(ctx, types.SignatureThreshold{Amount: signatureThreshold})

	msg := types.MsgReplaceMessage{
		From:                 string(originalMessage.Sender),
		OriginalMessage:      originalMessageBytes,
		OriginalAttestation:  originalAttestation,
		NewMessageBody:       []byte("123"),
		NewDestinationCaller: []byte("new destination caller3456789012"),
	}

	_, err = server.ReplaceMessage(sdk.WrapSDKContext(ctx), &msg)
	require.ErrorIs(t, types.ErrParsingMessage, err)
	require.Contains(t, err.Error(), "cctp message must be at least 116 bytes, got 115: error while parsing message into bytes")
}

func TestReplaceMessageInvalidSender(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	paused := types.SendingAndReceivingMessagesPaused{Paused: false}
	testkeeper.SetSendingAndReceivingMessagesPaused(ctx, paused)

	burnMessage := types.BurnMessage{
		Version:       1,
		BurnToken:     make([]byte, 32),
		MintRecipient: make([]byte, 32),
		Amount:        math.NewInt(123456),
		MessageSender: make([]byte, 32),
	}

	burnMessageBytes, err := burnMessage.Bytes()
	require.Nil(t, err)

	// we encode the message sender when sending messages, so we must use an encoded message in the original message
	senderBech32 := "cosmos1x8rynykqla7cnc0tf2f3xn0wa822ztt788yd5a"
	senderEncoded := make([]byte, 32)
	copy(senderEncoded[12:], sdk.MustAccAddressFromBech32(senderBech32))

	originalMessage := types.Message{
		Version:           1,
		SourceDomain:      4, // Noble domain id
		DestinationDomain: 3,
		Nonce:             2,
		Sender:            senderEncoded,
		Recipient:         []byte("recipient01234567890123456789012"),
		DestinationCaller: []byte("destination caller90123456789012"),
		MessageBody:       burnMessageBytes,
	}
	originalMessageBytes, _ := originalMessage.Bytes()

	// generate attestation, set attesters, signature threshold
	signatureThreshold := uint32(2)
	privKeys := generateNPrivateKeys(int(signatureThreshold))
	attesters := getAttestersFromPrivateKeys(privKeys)
	originalAttestation := generateAttestation(originalMessageBytes, privKeys)
	for _, attester := range attesters {
		testkeeper.SetAttester(ctx, attester)
	}
	testkeeper.SetSignatureThreshold(ctx, types.SignatureThreshold{Amount: signatureThreshold})

	msg := types.MsgReplaceMessage{
		From:                 "cosmos169xaqmxumqa829gg73nxrenkhhd2mrs36j3vrz", // different sender than in original message
		OriginalMessage:      originalMessageBytes,
		OriginalAttestation:  originalAttestation,
		NewMessageBody:       []byte("123"),
		NewDestinationCaller: []byte("new destination caller3456789012"),
	}

	_, err = server.ReplaceMessage(sdk.WrapSDKContext(ctx), &msg)
	require.ErrorIs(t, types.ErrReplaceMessage, err)
	require.Contains(t, err.Error(), "sender not permitted to use nonce")
}

func TestReplaceMessageMessageNotOriginallySentFromThisDomain(t *testing.T) {
	testkeeper, ctx := keepertest.CctpKeeper(t)
	server := keeper.NewMsgServerImpl(testkeeper)

	paused := types.SendingAndReceivingMessagesPaused{Paused: false}
	testkeeper.SetSendingAndReceivingMessagesPaused(ctx, paused)

	burnMessage := types.BurnMessage{
		Version:       1,
		BurnToken:     make([]byte, 32),
		MintRecipient: make([]byte, 32),
		Amount:        math.NewInt(123456),
		MessageSender: make([]byte, 32),
	}

	burnMessageBytes, err := burnMessage.Bytes()
	require.Nil(t, err)

	// we encode the message sender when sending messages, so we must use an encoded message in the original message
	senderBech32 := "cosmos1x8rynykqla7cnc0tf2f3xn0wa822ztt788yd5a"
	senderEncoded := make([]byte, 32)
	copy(senderEncoded[12:], sdk.MustAccAddressFromBech32(senderBech32))

	originalMessage := types.Message{
		Version:           1,
		SourceDomain:      8, // not Noble's domain id
		DestinationDomain: 3,
		Nonce:             2,
		Sender:            senderEncoded,
		Recipient:         []byte("recipient01234567890123456789012"),
		DestinationCaller: []byte("destination caller90123456789012"),
		MessageBody:       burnMessageBytes,
	}
	originalMessageBytes, _ := originalMessage.Bytes()

	// generate attestation, set attesters, signature threshold
	signatureThreshold := uint32(2)
	privKeys := generateNPrivateKeys(int(signatureThreshold))
	attesters := getAttestersFromPrivateKeys(privKeys)
	originalAttestation := generateAttestation(originalMessageBytes, privKeys)
	for _, attester := range attesters {
		testkeeper.SetAttester(ctx, attester)
	}
	testkeeper.SetSignatureThreshold(ctx, types.SignatureThreshold{Amount: signatureThreshold})

	msg := types.MsgReplaceMessage{
		From:                 senderBech32,
		OriginalMessage:      originalMessageBytes,
		OriginalAttestation:  originalAttestation,
		NewMessageBody:       []byte("123"),
		NewDestinationCaller: []byte("new destination caller3456789012"),
	}

	_, err = server.ReplaceMessage(sdk.WrapSDKContext(ctx), &msg)
	require.ErrorIs(t, types.ErrReplaceMessage, err)
	require.Contains(t, err.Error(), "message not originally sent from this domain")
}
