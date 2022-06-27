package main

import (
	"fmt"

	"github.com/easonnong/public-chain/part5-proof-of-work/BLC"
)

func main() {
	// genesis block
	blockchain := BLC.CreateBlockchainWithGenesisBlock()

	// new block
	blockchain.AddBlockToBlockchain(
		"send 1 bitcoin to satoshi",
		blockchain.Blocks[len(blockchain.Blocks)-1].Height+1,
		blockchain.Blocks[len(blockchain.Blocks)-1].Hash,
	)
	blockchain.AddBlockToBlockchain(
		"send 2 bitcoin to satoshi",
		blockchain.Blocks[len(blockchain.Blocks)-1].Height+1,
		blockchain.Blocks[len(blockchain.Blocks)-1].Hash,
	)

	fmt.Println(blockchain)
}
