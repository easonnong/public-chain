package BLC

type BlockChain struct {
	// ordered storage block
	Blocks []*Block
}

// 1. create genesis blockchain
func CreateBlockchainWithGenesisBlock() *BlockChain {
	// create genesis block
	genesisBlock := CreateGenesisBlock("genesis block")
	// return genesis blockchain
	return &BlockChain{
		Blocks: []*Block{genesisBlock},
	}
}
