package main

import (
	"fmt"

	"github.com/easonnong/snowcash/part1-Basic-Prototype/BLC"
)

func main() {
	block := BLC.NewBlock(
		"the first block",
		1,
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	)
	fmt.Println(block)
}
