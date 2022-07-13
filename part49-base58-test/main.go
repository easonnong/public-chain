package main

import (
	"fmt"

	"github.com/easonnong/public-chain/part49-base58-test/BLC"
)

func main() {
	hash := []byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	fmt.Println("len(hash):", len(hash))
	fmt.Printf("hash:%x\n", hash)

	bytes58 := BLC.Base58Encode(hash)

	//fmt.Printf("%x\n", bytes58)
	fmt.Printf("bytes58:%s\n", bytes58)

	bytes58Str := BLC.Base58Decode(bytes58)

	fmt.Printf("bytes58Str:%x\n", bytes58Str)
	fmt.Printf("bytes58Str[1:]:%x\n", bytes58Str[1:])

}
