package types

import (
	"encoding/binary"
	"fmt"
)

const (
	// ModuleName defines the module name
	ModuleName = "router"

	// StoreKey defines the primary module store key
	StoreKey = "router"

	// RouterKey defines the module's message routing key
	RouterKey = StoreKey
)

var (
	IBCForwardPrefix     = []byte("forward/")
	InFlightPacketPrefix = []byte("inflight/")
	MintPrefix           = []byte("mint/")
)

func LookupKey(sourceDomain uint32, nonce uint64) []byte {
	sourceDomainBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(sourceDomainBytes, sourceDomain)
	nonceBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(nonceBytes, nonce)
	return append(nonceBytes, []byte(sourceDomainBytes)...)
}

func InFlightPacketKey(channelID, portID string, sequence uint64) []byte {
	return []byte(fmt.Sprintf("%s/%s/%d", channelID, portID, sequence))
}
