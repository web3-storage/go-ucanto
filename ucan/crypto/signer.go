package crypto

import (
	"github.com/web3-storage/go-ucanto/ucan/crypto/signature"
)

// Signer is an entity that can sign a payload.
type Signer interface {
	// Sign takes a byte encoded message and produces a verifiable signature.
	Sign(msg []byte) signature.SignatureView
}
