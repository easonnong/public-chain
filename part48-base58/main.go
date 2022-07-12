package main

import (
	"crypto/sha256"
	"fmt"

	"github.com/easonnong/public-chain/part48-base58/BLC"
)

func main() {
	bytes := []byte("https://github.com/easonnong")

	hasher := sha256.New()
	hasher.Write(bytes)
	hash := hasher.Sum(nil)

	fmt.Printf("hash:%x\n", hash)

	bytes58 := BLC.Base58Encode(bytes)

	fmt.Printf("%x\n", bytes58)
	fmt.Printf("%s\n", bytes58)

	bytes58Str := BLC.Base58Decode(bytes58)

	fmt.Printf("%s\n", bytes58Str)
	fmt.Printf("%s\n", bytes58Str[1:])

}
