package BLC

import (
	"time"
)

type Block struct {
	// 1. bock height.
	Height int64
	// 2. hash of the previous block.
	PrevBlockHash []byte
	// 3. transaction data.
	Data []byte
	// 4. timestamp.
	Timestamp int64
	// 5. hash.
	Hash []byte
}

// create new block.
func NewBlock(data string, height int64, prevBlockHash []byte) *Block {
	block := &Block{
		Height:        height,
		PrevBlockHash: prevBlockHash,
		Data:          []byte(data),
		Timestamp:     time.Now().Unix(),
		Hash:          nil,
	}
	return block
}
