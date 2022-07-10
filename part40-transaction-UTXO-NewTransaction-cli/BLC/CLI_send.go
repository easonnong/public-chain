package BLC

import (
	"fmt"
	"os"
)

// transfer accounts
func (cli *CLI) send(from, to, amount []string) {
	if !isDBExist() {
		fmt.Println("database is not exist...")
		os.Exit(1)
	}

	blockchain := GetBlockchainObject()
	defer blockchain.DB.Close()

	blockchain.MineNewBlock(from, to, amount)
}
