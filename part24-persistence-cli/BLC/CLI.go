package BLC

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct {
	BlockChain *BlockChain
}

// print how to use cli
func (cli *CLI) printUsage() {
	fmt.Println("\nHere is a usage...")
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

// add block into blockchain
func (cli *CLI) addBlock(data string) {
	cli.BlockChain.AddBlockToBlockchain(data)
}

// print all block data
func (cli *CLI) printchain() {
	cli.BlockChain.PrintChain()
}

// cli run
func (cli *CLI) Run() {

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	flagAddBlockData := addBlockCmd.String("data", "", "block data")

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

	if printChainCmd.Parsed() {
		cli.printchain()
	}
}
