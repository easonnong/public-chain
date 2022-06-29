package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args

	fmt.Printf("%v\n", args)
	fmt.Printf("%v\n", args[1])

}
