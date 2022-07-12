package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	msg := "https://github.com/easonnong"

	encoded := base64.URLEncoding.EncodeToString([]byte(msg))
	fmt.Println("encoded:", encoded)

	decoded, err := base64.URLEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("decode err:", err)
		return
	}
	fmt.Println("decoded:", string(decoded))
}
