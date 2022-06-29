package main

import (
	"github.com/easonnong/public-chain/part24-persistence-cli/BLC"
)

func main() {

	blockchain := BLC.CreateBlockchainWithGenesisBlock()
	defer blockchain.DB.Close()

	cli := &BLC.CLI{
		BlockChain: blockchain,
	}

	cli.Run()
}
