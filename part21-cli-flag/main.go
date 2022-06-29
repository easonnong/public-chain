package main

import (
	"flag"
	"fmt"
)

func main() {

	flagString := flag.String("printchain", "", "print all block data")
	flagInt := flag.Int("number", 6, "output a number")
	flagBool := flag.Bool("open", false, "true or false")

	flag.Parse()
	fmt.Printf("%s\n", *flagString)
	fmt.Printf("%d\n", *flagInt)
	fmt.Printf("%v\n", *flagBool)
}
