package blockchain

import (
	"fmt"
	"testing"
)

func TestBlockchain(t *testing.T) {
	chain := InitBlockchain()
	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Hash: %x\n", block.Hash)
	}

}
