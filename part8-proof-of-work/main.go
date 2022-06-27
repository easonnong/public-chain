package main

import (
	"fmt"

	"github.com/easonnong/public-chain/part8-proof-of-work/BLC"
)

func main() {
	// genesis block
	/* blockchain := BLC.CreateBlockchainWithGenesisBlock()
	fmt.Printf("\n%v\n", blockchain.Blocks)
	// new block
	blockchain.AddBlockToBlockchain(
		"send 1 bitcoin to satoshi",
		blockchain.Blocks[len(blockchain.Blocks)-1].Height+1,
		blockchain.Blocks[len(blockchain.Blocks)-1].Hash,
	)
	fmt.Printf("\n%v\n", blockchain.Blocks)
	blockchain.AddBlockToBlockchain(
		"send 2 bitcoin to satoshi",
		blockchain.Blocks[len(blockchain.Blocks)-1].Height+1,
		blockchain.Blocks[len(blockchain.Blocks)-1].Hash,
	)
	fmt.Printf("\n%v\n", blockchain.Blocks) */
	block := BLC.NewBlock(
		"test",
		1,
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	)
	fmt.Printf("%d\n", block.Nonce)
	fmt.Printf("%x\n", block.Hash)

	pow := BLC.NewProofOfWork(block)
	fmt.Printf("%v\n", pow.IsValid())
}
