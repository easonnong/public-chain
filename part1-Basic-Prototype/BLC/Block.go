package BLC

import (
	"bytes"
	"crypto/sha256"
	"strconv"
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
}

// set hash
func (block *Block) SetHash() {
	// 1. height to byte
	heightBytes := Int64ToByte(block.Height)
	// 2. timestamp to byte
	timestampString := strconv.FormatInt(block.Timestamp, 2)
	timestampBytes := []byte(timestampString)
	// 3. sum all member
	blockBytes := bytes.Join(
		[][]byte{heightBytes, block.PrevBlockHash, block.Data, timestampBytes},
		[]byte{},
	)
	// 4. get the hash
	hash := sha256.Sum256(blockBytes)
	block.Hash = hash[:]
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
	block.SetHash()
	return block
}
