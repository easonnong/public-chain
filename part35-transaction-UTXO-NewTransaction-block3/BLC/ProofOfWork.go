package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// there must be at least 12 zeros before the 256-bit hash
const targetBit = 12

type ProofOfWork struct {
	// the current block to validate
	Block *Block
	// storage of large numbers
	target *big.Int
}

// splicing data
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevBlockHash,
			pow.Block.HashTransactions(),
			Int64ToByte(pow.Block.Timestamp),
			Int64ToByte(int64(targetBit)),
			Int64ToByte(int64(nonce)),
			Int64ToByte(int64(pow.Block.Height)),
		},
		[]byte{},
	)
	return data
}

// mining
func (pow *ProofOfWork) Run() ([]byte, int64) {
	nonce := 0
	// storage the new hash
	var hash [32]byte
	var hashInt big.Int

	for {
		// get the dataBytes
		dataBytes := pow.prepareData(nonce)
		// get hash
		hash = sha256.Sum256(dataBytes)
		fmt.Printf("\r%x", hash)
		// storage hash into hashInt
		hashInt.SetBytes(hash[:])
		// judge hashInt and target
		if pow.target.Cmp(&hashInt) == 1 {
			break
		}
		nonce += 1
	}

	return hash[:], int64(nonce)
}

// is valid of hash
func (pow *ProofOfWork) IsValid() bool {
	var hashInt big.Int
	hashInt.SetBytes(pow.Block.Hash)
	return pow.target.Cmp(&hashInt) == 1
}

// create new pow object
func NewProofOfWork(block *Block) *ProofOfWork {

	//1. create a target
	target := big.NewInt(1)
	//2. left shift 256-targetBit
	target = target.Lsh(target, 256-targetBit)

	return &ProofOfWork{
		Block:  block,
		target: target,
	}
}
