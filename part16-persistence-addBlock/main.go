package main

import (
	"fmt"

	"github.com/easonnong/public-chain/part16-persistence-addBlock/BLC"
)

func main() {

	blockchain := BLC.CreateBlockchainWithGenesisBlock()
	fmt.Println("\ngenesis ok")
	defer blockchain.DB.Close()

	blockchain.AddBlockToBlockchain("send 2 bitcoin to satoshi")
	fmt.Println("\n2 ok")
	blockchain.AddBlockToBlockchain("send 3 bitcoin to satoshi")
	fmt.Println("\n3 ok")
	blockchain.AddBlockToBlockchain("send 4 bitcoin to satoshi")
	fmt.Println("\n4 ok")
}
