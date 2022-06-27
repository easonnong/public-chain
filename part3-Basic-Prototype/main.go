package main

import (
	"fmt"

	"github.com/easonnong/public-chain/part3-Basic-Prototype/BLC"
)

func main() {
	genesisBlockchain := BLC.CreateBlockchainWithGenesisBlock()
	fmt.Println(genesisBlockchain)
	fmt.Println(genesisBlockchain.Blocks)
	fmt.Println(genesisBlockchain.Blocks[0])
}
