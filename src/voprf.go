package src

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

func deriveKey(unblindedPoint string, token bytes.Buffer) error {
	c := *secp256k1.S256()

	data, err := hex.DecodeString(base64.StdEncoding.EncodeToString([]byte(unblindedPoint)))
	if err != nil {
		return err
	}

	x, y := c.Unmarshal(data[2:])

	return nil
}
