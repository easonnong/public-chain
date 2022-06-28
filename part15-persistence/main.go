package main

import (
	"github.com/easonnong/public-chain/part15-persistence/BLC"
)

func main() {

	blockchain := BLC.CreateBlockchainWithGenesisBlock()
	defer blockchain.DB.Close()
}
