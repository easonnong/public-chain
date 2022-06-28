package main

import (
	"github.com/easonnong/public-chain/part16-persistence-addBlock/BLC"
)

func main() {

	blockchain := BLC.CreateBlockchainWithGenesisBlock()
	defer blockchain.DB.Close()
}
