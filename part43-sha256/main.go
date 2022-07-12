package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	hasher := sha256.New()
	hasher.Write([]byte("https://github.com/easonnong"))
	hashBytes := hasher.Sum(nil)
	hashStr := fmt.Sprintf("%x", hashBytes)
	fmt.Println(hashStr)
}
