package BLC

import (
	"bytes"
	"crypto/sha256"
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
	Txs []*Transaction
	// 4. timestamp
	Timestamp int64
	// 5. hash
	Hash []byte
	// 6. Nonce
	Nonce int64
}

// Txs to []byte
func (block *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range block.Txs {
		txHashes = append(txHashes, tx.TxHash)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
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
func NewBlock(txs []*Transaction, height int64, prevBlockHash []byte) *Block {
	block := &Block{
		Height:        height,
		PrevBlockHash: prevBlockHash,
		Txs:           txs,
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

func CreateGenesisBlock(txs []*Transaction) *Block {
	return NewBlock(
		txs,
		1,
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	)
}
