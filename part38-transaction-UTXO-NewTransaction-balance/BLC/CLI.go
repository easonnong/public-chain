package BLC

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct{}

// print how to use cli
func (cli *CLI) printUsage() {
	fmt.Println("\nHere is a usage...")
	fmt.Println("\tgetbalance -address ADDRESS")
	fmt.Println("\tcreateblockchain -address ADDRESS")
	fmt.Println("\tsend -from FROM -to TO -amount AMOUNT")
	fmt.Println("\tprintchain")
}

// judge args is valid
func (cli *CLI) isValidArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

// create genesis block
func (cli *CLI) createGenesisBlockchain(address string) {
	CreateBlockchainWithGenesisBlock(address)
}

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

// get balance
func (cli *CLI) getBalance(address string) {
	fmt.Println("address:", address)

	blockchain := GetBlockchainObject()
	defer blockchain.DB.Close()

	amount := blockchain.GetBalance(address)

	fmt.Printf("%s has %d token\n", address, amount)

}

// cli run
func (cli *CLI) Run() {

	sendBlockCmd := flag.NewFlagSet("send", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)

	flagFrom := sendBlockCmd.String("from", "", "transfer from")
	flagTo := sendBlockCmd.String("to", "", "transfer to")
	flagAmount := sendBlockCmd.String("amount", "", "transfer amount")

	flagCreateBlockchainAddress := createBlockchainCmd.String("address", "genesis block address...", "genesis block address")

	flagGetBalance := getBalanceCmd.String("address", "", "get balance")

	cli.isValidArgs()

	switch os.Args[1] {
	case "send":
		err := sendBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "createblockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "getbalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if sendBlockCmd.Parsed() {
		if *flagFrom == "" || *flagTo == "" || *flagAmount == "" {
			cli.printUsage()
			os.Exit(1)
		}

		from := JsonToArray(*flagFrom)
		to := JsonToArray(*flagTo)
		amount := JsonToArray(*flagAmount)

		cli.send(from, to, amount)
	}

	if createBlockchainCmd.Parsed() {
		if *flagCreateBlockchainAddress == "" {
			cli.printUsage()
			os.Exit(1)
		}
		cli.createGenesisBlockchain(*flagCreateBlockchainAddress)
	}

	if getBalanceCmd.Parsed() {
		if *flagGetBalance == "" {
			cli.printUsage()
			os.Exit(1)
		}
		cli.getBalance(*flagGetBalance)
	}

	if printChainCmd.Parsed() {
		cli.printchain()
	}
}
