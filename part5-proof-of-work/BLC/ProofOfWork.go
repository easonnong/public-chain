package BLC

type ProofOfWork struct {
	Block *Block
}

func (proofOfWork *ProofOfWork) Run() ([]byte, int64) {
	return nil, 0
}

// create new pow object
func NewProofOfWork(block *Block) *ProofOfWork {
	return &ProofOfWork{
		Block: block,
	}
}
