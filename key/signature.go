package key

import (
	"crypto/ecdsa"
	"math/big"
)

type Signature struct {
	R, S []byte
}

func (s *Signature) Verify(pub *PublicKey, msg []byte) bool {
	return ecdsa.Verify(pub.Key, msg, new(big.Int).SetBytes(s.R), new(big.Int).SetBytes(s.S))
}
