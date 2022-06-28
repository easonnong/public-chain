package BLC

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

type Block struct {
	// 1. bock height
	Height int64
	// 2. hash of the previous block
	PrevBlockHash []byte
	// 3. transaction data
	Data []byte
	// 4. timestamp
	Timestamp int64
	// 5. hash
	Hash []byte
	// 6. Nonce
	Nonce int64
}

// serialize
func (block *Block) SerializeBlock() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	if err := encoder.Encode(block); err != nil {
		log.Panic(err)
	}
	return result.Bytes()
}

// deserializate
func DeserializateBlock(blockBytes []byte) *Block {
	var deBlock Block
	decoder := gob.NewDecoder(bytes.NewReader(blockBytes))
	if err := decoder.Decode(&deBlock); err != nil {
		log.Panic(err)
	}
	return &deBlock
}

// create new block
func NewBlock(data string, height int64, prevBlockHash []byte) *Block {
	block := &Block{
		Height:        height,
		PrevBlockHash: prevBlockHash,
		Data:          []byte(data),
		Timestamp:     time.Now().Unix(),
		Hash:          nil,
	}

	// use pow and return hash and nonce
	pow := NewProofOfWork(block)
	hash, nonce := pow.Run()
	block.Hash = hash
	block.Nonce = nonce

	return block

}

func CreateGenesisBlock(data string) *Block {
	return NewBlock(
		data,
		1,
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	)
}
