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
	fmt.Println("\tcreateblockchain -data DATA")
	fmt.Println("\taddBlock -data DATA")
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
func (cli *CLI) createGenesisBlockchain(data string) {
	CreateBlockchainWithGenesisBlock(data)
}

// add block into blockchain
func (cli *CLI) addBlock(data string) {
	if !isDBExist() {
		fmt.Println("database is not exist...")
		os.Exit(1)
	}

	blockchain := GetBlockchainObject()
	blockchain.AddBlockToBlockchain(data)
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

// cli run
func (cli *CLI) Run() {

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	flagAddBlockData := addBlockCmd.String("data", "block data...", "block data")
	flagCreateBlockchainData := createBlockchainCmd.String("data", "genesis block data...", "genesis block data")

	cli.isValidArgs()

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
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

	if addBlockCmd.Parsed() {
		if *flagAddBlockData == "" {
			cli.printUsage()
			os.Exit(1)
		}
		cli.addBlock(*flagAddBlockData)
	}

	if createBlockchainCmd.Parsed() {
		if *flagCreateBlockchainData == "" {
			cli.printUsage()
			os.Exit(1)
		}
		cli.createGenesisBlockchain(*flagCreateBlockchainData)
	}

	if printChainCmd.Parsed() {
		cli.printchain()
	}
}
