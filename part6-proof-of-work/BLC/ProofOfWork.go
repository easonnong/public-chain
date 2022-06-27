package BLC

import "math/big"

// there must be at least 16 zeros before the 256-bit hash
const targetBit = 16

type ProofOfWork struct {
	// the current block to validate
	Block *Block
	// storage of large numbers
	target *big.Int
}

func (proofOfWork *ProofOfWork) Run() ([]byte, int64) {
	return nil, 0
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
