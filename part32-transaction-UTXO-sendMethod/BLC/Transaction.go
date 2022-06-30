package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

type Transaction struct {

	//1. transcation hash
	TxHash []byte
	//2. input
	Vins []*TXInput
	//3. output
	Vouts []*TXOutput
}

// 1. genesis transaction
func NewCoinbaseTransaction(address string) *Transaction {
	// consumption
	txInput := &TXInput{
		TxHash:    []byte{},
		Vout:      -1,
		ScriptSig: "genesis data",
	}

	txOutput := &TXOutput{
		Value:        10,
		ScriptPubKey: address,
	}

	txCoinbase := &Transaction{
		TxHash: []byte{},
		Vins:   []*TXInput{txInput},
		Vouts:  []*TXOutput{txOutput},
	}

	// setting hash
	txCoinbase.HashTranscation()

	return txCoinbase
}

//2. transfer accounts transaction

// transfer hash
func (tx *Transaction) HashTranscation() {
	var result bytes.Buffer

	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	hash := sha256.Sum256(result.Bytes())

	tx.TxHash = hash[:]
}
