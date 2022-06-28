package BLC

import (
	"log"

	"github.com/boltdb/bolt"
)

// database name
const dbName = "blockchain.db"

//table name
const blockTableName = "blocks"

type BlockChain struct {
	// the  latest block hash
	Tip []byte
	// blockchain database
	DB *bolt.DB
}

// 1. create genesis blockchain
func CreateBlockchainWithGenesisBlock() *BlockChain {
	var blockHash []byte

	// create or open database
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	// updata blockchain
	err = db.Update(func(tx *bolt.Tx) error {
		// create table
		bucket, err := tx.CreateBucket([]byte(blockTableName))
		if err != nil {
			log.Panic(err)
		}

		if bucket == nil {
			// create genesis block
			genesisBlock := CreateGenesisBlock("genesis block")

			//store genesis block to table
			err := bucket.Put(genesisBlock.Hash, genesisBlock.SerializeBlock())
			if err != nil {
				log.Panic(err)
			}

			//store latest block hash
			err = bucket.Put([]byte("l"), genesisBlock.Hash)
			if err != nil {
				log.Panic(err)
			}

			blockHash = genesisBlock.Hash
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	// return genesis blockchain
	return &BlockChain{
		Tip: blockHash,
		DB:  db,
	}
}

// add new block to blockchain
/* func (blockchain *BlockChain) AddBlockToBlockchain(data string, height int64, prebBlockHash []byte) {
	newBlock := NewBlock(data, height, prebBlockHash)
	blockchain.Blocks = append(blockchain.Blocks, newBlock)
} */
