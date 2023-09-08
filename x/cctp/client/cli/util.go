package cli

import (
	"errors"

	"github.com/cosmos/btcutil/base58"
	"github.com/ethereum/go-ethereum/common"
)

// parseAddress parses an encoded address into a 32 length byte array.
// Currently supported encodings: base58, hex.
func parseAddress(address string) ([]byte, error) {
	if address[:2] == "0x" {
		bz := common.FromHex(address)
		return padBytes(bz)
	}

	bz := base58.Decode(address)
	return padBytes(bz)
}

// padBytes left pads a byte array to be length 32.
func padBytes(bz []byte) ([]byte, error) {
	res := make([]byte, 32)

	if len(bz) > 32 {
		return nil, errors.New("decoded bytes too big")
	}

	copy(res[32-len(bz):], bz)
	return res, nil
}
