package main

import (
	"fmt"

	"github.com/josemarjobs/blockchain/block"
)

func main() {

	bc := block.NewBlockChain()

	bc.AddBlock("Send 1 BTC to Peter Griffin")
	bc.AddBlock("Send more 2 BTC to Peter Griffin")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n\n", block.Hash)
	}
}

// https://jeiwan.cc/posts/building-blockchain-in-go-part-1/
