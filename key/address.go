package key

import "encoding/hex"

type Address struct {
	Hash []byte
}

func (s *Address) Bytes() []byte {
	return s.Hash
}

func (s *Address) String() string {
	return hex.EncodeToString(s.Hash)
}
