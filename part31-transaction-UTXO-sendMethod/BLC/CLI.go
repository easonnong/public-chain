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

// add block into blockchain
func (cli *CLI) addBlock(txs []*Transaction) {
	if !isDBExist() {
		fmt.Println("database is not exist...")
		os.Exit(1)
	}

	blockchain := GetBlockchainObject()
	blockchain.AddBlockToBlockchain(txs)
	defer blockchain.DB.Close()
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
	MineNewBlock(from, to, amount)
}

// mine new block
func MineNewBlock(from, to, amount []string) {
	fmt.Println(from)
	fmt.Println(to)
	fmt.Println(amount)
}

// cli run
func (cli *CLI) Run() {

	sendBlockCmd := flag.NewFlagSet("send", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)

	flagFrom := sendBlockCmd.String("from", "", "transfer from")
	flagTo := sendBlockCmd.String("to", "", "transfer to")
	flagAmount := sendBlockCmd.String("amount", "", "transfer amount")

	flagCreateBlockchainAddress := createBlockchainCmd.String("address", "genesis block address...", "genesis block address")

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
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if sendBlockCmd.Parsed() {
		if *flagFrom == "" || *flagTo == "" || *flagAmount == "" {
			cli.printUsage()
			os.Exit(1)
		}

		fmt.Println(*flagFrom)
		fmt.Println(flagTo)
		fmt.Println(*flagAmount)

		from := JsonToArray(*flagFrom)
		to := JsonToArray(*flagTo)
		amount := JsonToArray(*flagAmount)

		cli.send(from, to, amount)

		cli.addBlock([]*Transaction{})
	}

	if createBlockchainCmd.Parsed() {
		if *flagCreateBlockchainAddress == "" {
			cli.printUsage()
			os.Exit(1)
		}
		cli.createGenesisBlockchain(*flagCreateBlockchainAddress)
	}

	if printChainCmd.Parsed() {
		cli.printchain()
	}
}
