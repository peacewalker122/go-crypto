package key

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"log"
	"math/big"
)

type PrivateKey struct {
	Key *ecdsa.PrivateKey
}

func (p *PrivateKey) PublicKey() *PublicKey {
	f := make([]byte, PUBLICKEYLEN)
	copy(f, p.Key.PublicKey.X.Bytes())
	copy(f[32:], p.Key.PublicKey.Y.Bytes())
	return &PublicKey{&ecdsa.PublicKey{Curve: p.Key.Curve, X: p.Key.PublicKey.X, Y: p.Key.PublicKey.Y}}
}

func NewPrivateKey() *PrivateKey {
	var (
		priv *ecdsa.PrivateKey
		err  error
	)

	if priv, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader); err != nil {
		panic(err)
	}

	log.Printf("NewPrivateKey: %s \n", hex.EncodeToString(priv.D.Bytes()))

	return &PrivateKey{
		Key: priv,
	}
}

func (p *PrivateKey) Sign(msg []byte) (*Signature, error) {
	var (
		r, s *big.Int
		err  error
	)

	if r, s, err = ecdsa.Sign(rand.Reader, p.Key, msg); err != nil {
		return nil, err
	}

	return &Signature{
		R: r.Bytes(),
		S: s.Bytes(),
	}, nil
}

func (p *PrivateKey) Bytes() []byte {
	return p.Key.D.Bytes()
}
