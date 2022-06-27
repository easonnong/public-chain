package main

import (
	"fmt"

	"github.com/easonnong/public-chain/part2-Basic-Prototype/BLC"
)

func main() {
	block := BLC.CreateGenesisBlock("genesis block")
	fmt.Println(block)
}
