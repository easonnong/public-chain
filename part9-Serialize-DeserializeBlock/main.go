package main

import (
	"fmt"

	"github.com/easonnong/public-chain/part9-Serialize-DeserializeBlock/BLC"
)

func main() {
	block := BLC.NewBlock(
		"test",
		1,
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	)
	fmt.Printf("%d\n", block.Nonce)
	fmt.Printf("%x\n", block.Hash)

	bytes := block.SerializeBlock()
	fmt.Println(bytes)

	block = block.DeserializateBlock(bytes)
	fmt.Printf("%d\n", block.Nonce)
	fmt.Printf("%x\n", block.Hash)
}
