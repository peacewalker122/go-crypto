package block

import (
	"fmt"
	"runtime/debug"
	"strconv"
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
				bc.AddBlock("Send 1 more BTC to Ivan")
				iter := bc.Iterator()

				for {
					block := iter.Next()

					fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
					fmt.Printf("Data: %s\n", block.Data)
					fmt.Printf("Hash: %x\n", block.Hash)
					pow := NewProofOfWork(block)
					fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
					fmt.Println()

					if len(block.PrevBlockHash) == 0 {
						break
					}
				}

			},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			tc.fn(t)

			if err := recover(); err != nil {
				debug.PrintStack()
				t.Errorf("TestNewBlock() failed: %v", err)
			}
		})
	}
}
