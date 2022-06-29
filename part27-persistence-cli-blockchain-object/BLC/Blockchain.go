package BLC

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

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

// add new block to blockchain
func (blockchain *BlockChain) AddBlockToBlockchain(data string) {

	// update database
	err := blockchain.DB.Update(func(tx *bolt.Tx) error {
		// 1.get table
		bucket := tx.Bucket([]byte(blockTableName))

		// 2.create new block
		if bucket != nil {
			// 3.get latest block
			blockBytes := bucket.Get(blockchain.Tip)
			// deserialize
			block := DeserializateBlock(blockBytes)

			// 4.store new block
			newBlock := NewBlock(data, block.Height+1, block.Hash)
			err := bucket.Put(newBlock.Hash, newBlock.SerializeBlock())
			if err != nil {
				log.Panic(err)
			}

			// 4.update "l"
			err = bucket.Put([]byte("l"), newBlock.Hash)
			if err != nil {
				log.Panic(err)
			}
			// 5.update Tip
			blockchain.Tip = newBlock.Hash
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

// print blockchain database
func (blockchain *BlockChain) PrintChain() {

	// create interator
	blockChainIterator := blockchain.CreateIterator()
	for {
		// get current block
		block := blockChainIterator.NextIterator()

		// print block data
		fmt.Printf("Height: %d\n", block.Height)
		fmt.Printf("PrevBlockHash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Timestamp: %s\n", time.Unix(block.Timestamp, 0).Format("2006-01-02 03:04:05 PM"))
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Println()

		// doing cycle
		var hashInt big.Int
		hashInt.SetBytes(block.PrevBlockHash)

		if big.NewInt(0).Cmp(&hashInt) == 0 {
			break
		}

	}
}

// get blockchain object
func GetBlockchainObject() *BlockChain {

	// open database
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	var tip []byte

	err = db.View(func(tx *bolt.Tx) error {

		bucket := tx.Bucket([]byte(blockTableName))

		if bucket != nil {
			tip = bucket.Get([]byte("l"))
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	return &BlockChain{
		Tip: tip,
		DB:  db,
	}
}

// 1. create genesis blockchain
func CreateBlockchainWithGenesisBlock(data string) {

	// is database exist
	if isDBExist() {
		fmt.Println("genesis block had exist.")
		os.Exit(1)
	}

	// create or open database
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	// updata blockchain
	err = db.Update(func(tx *bolt.Tx) error {

		// 1.create table
		bucket, err := tx.CreateBucket([]byte(blockTableName))
		if err != nil {
			log.Panic(err)
		}

		if bucket != nil {
			// 2.create genesis block
			genesisBlock := CreateGenesisBlock(data)

			// 3.store genesis block to table
			err := bucket.Put(genesisBlock.Hash, genesisBlock.SerializeBlock())
			if err != nil {
				log.Panic(err)
			}

			// 4.store latest block hash
			err = bucket.Put([]byte("l"), genesisBlock.Hash)
			if err != nil {
				log.Panic(err)
			}
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

}

// judge isExist database
func isDBExist() bool {
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		return false
	}
	return true
}
