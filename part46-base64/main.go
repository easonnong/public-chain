package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	msg := "hello, world"

	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	fmt.Println("encoded:", encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("decode err:", err)
		return
	}
	fmt.Println("decoded:", string(decoded))
}
