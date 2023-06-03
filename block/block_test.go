package block

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBlock(t *testing.T) {
	testcases := []struct {
		name string
		fn   func(t *testing.T)
	}{
		{
			name: "NewBlock",
			fn: func(t *testing.T) {
				bc := NewBlockchain()
				bc.AddBlock("Send 1 BTC to Ivan")
				bc.AddBlock("Send 2 more BTC to Ivan")

				assert.Equal(t, 3, len(bc.blocks), "NewBlock() = %v, want %v", len(bc.blocks), 3)

				for _, block := range bc.blocks {
					pow := NewProofOfWork(block)
					assert.True(t, pow.Validate(), "NewBlock() = %v, want %v", pow.Validate(), true)
					t.Logf("Valid POW: %v\n", pow.Validate())
					t.Logf("Prev. hash: %x\n", block.PrevBlockHash)
					t.Logf("Data: %s\n", block.Data)
					t.Logf("Hash: %x\n", block.Hash)
				}
			},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, tc.fn)
	}
}
