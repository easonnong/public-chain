package main

import (
	"fmt"

	"github.com/easonnong/public-chain/part1-Basic-Prototype/BLC"
)

func main() {
	block := BLC.NewBlock(
		"genesis block",
		1,
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	)
	fmt.Println(block)
}
