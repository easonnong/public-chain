package main

import (
	"github.com/easonnong/public-chain/part15-persistence-createGenesisBlock/BLC"
)

func main() {

	blockchain := BLC.CreateBlockchainWithGenesisBlock()
	defer blockchain.DB.Close()
}
