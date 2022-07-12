package main

import (
	"fmt"

	"golang.org/x/crypto/ripemd160"
)

func main() {
	hasher := ripemd160.New()
	hasher.Write([]byte("https://github.com/easonnong"))
	hashBytes := hasher.Sum(nil)
	hashStr := fmt.Sprintf("%x", hashBytes)
	fmt.Println(hashStr)
}
