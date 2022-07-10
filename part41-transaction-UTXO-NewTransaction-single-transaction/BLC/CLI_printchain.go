package BLC

import (
	"fmt"
	"os"
)

// print all block data
func (cli *CLI) printchain() {
	if !isDBExist() {
		fmt.Println("database is not exist...")
		os.Exit(1)
	}

	blockchain := GetBlockchainObject()
	blockchain.PrintChain()
	defer blockchain.DB.Close()
}
