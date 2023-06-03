package block

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-DIFFICULTY))
	return &ProofOfWork{b, target}
}

func (p *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join([][]byte{
		p.Block.PrevBlockHash,
		p.Block.Data,
		IntToHex(p.Block.Timestamp),
		IntToHex(int64(DIFFICULTY)),
		IntToHex(int64(nonce)),
	}, []byte{})
	return data
}

func (p *ProofOfWork) Run() (int, []byte) {
	var (
		hashInt big.Int
		hash    [32]byte
		nonce   int = 0
	)

	for nonce < maxNonce {
		data := p.prepareData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(p.Target) == -1 {
			fmt.Printf("found hash: %x\n", hash)
			break
		} else {
			nonce++
		}
	}

	return nonce, hash[:]
}

func (p *ProofOfWork) Validate() bool {
	var hashInt big.Int
	data := p.prepareData(p.Block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	return hashInt.Cmp(p.Target) == -1
}

func IntToHex(num int64) []byte {
	return []byte(fmt.Sprintf("%x", num))
}
