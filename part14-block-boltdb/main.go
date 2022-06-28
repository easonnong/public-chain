package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/easonnong/public-chain/part14-block-boltdb/BLC"
)

func main() {

	block := BLC.NewBlock(
		"test",
		1,
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	)
	fmt.Printf("%d\n", block.Nonce)
	fmt.Printf("%x\n", block.Hash)

	// create or open database
	db, err := bolt.Open("database.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// update database
	err = db.Update(func(tx *bolt.Tx) error {
		// get table object
		bucket := tx.Bucket([]byte("blocks"))
		if bucket == nil {
			bucket, err = tx.CreateBucket([]byte("blocks"))
			if err != nil {
				log.Panic("create bucket failed,err=", err)
			}
		}

		err = bucket.Put([]byte("l"), block.SerializeBlock())
		if err != nil {
			log.Panic("serialize block failed,err=", err)
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	// view database
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("blocks"))
		if bucket != nil {
			blockData := bucket.Get([]byte("l"))
			fmt.Println("blockData->", blockData)
			block := block.DeserializateBlock(blockData)
			fmt.Println("block->", block)
		}
		return nil
	})
	if err != nil {
		log.Panic("view failed,err=", err)
	}
}
