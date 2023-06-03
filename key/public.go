package key

import (
	"crypto/ecdsa"
)

type PublicKey struct {
	Key *ecdsa.PublicKey
}

func (p *PublicKey) Bytes() []byte {
	f := make([]byte, PUBLICKEYLEN)
	copy(f, p.Key.X.Bytes())
	copy(f[32:], p.Key.Y.Bytes())
	return f
}

func (p *PublicKey) Address() *Address {
	return &Address{Hash: p.Key.X.Bytes()[len(p.Key.X.Bytes())-ADDRESSLEN:]}
}
