package BLC

import (
	"log"

	"github.com/boltdb/bolt"
)

type BlockChainIterator struct {
	CurrentHash []byte
	DB          *bolt.DB
}

// create a blockchain iterator
func (blockchain *BlockChain) CreateIterator() *BlockChainIterator {
	return &BlockChainIterator{
		CurrentHash: blockchain.Tip,
		DB:          blockchain.DB,
	}
}

// get current iterator
func (blockChainIterator *BlockChainIterator) NextIterator() *Block {

	var block *Block

	err := blockChainIterator.DB.View(func(tx *bolt.Tx) error {

		bucket := tx.Bucket([]byte(blockTableName))

		if bucket != nil {
			// get current hash bytes
			currentBlockBytes := bucket.Get(blockChainIterator.CurrentHash)
			block = DeserializateBlock(currentBlockBytes)

			// change currentHash
			blockChainIterator.CurrentHash = block.PrevBlockHash
		}

		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	return block
}
