package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
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
func NewSimpleTransaction(from, to string, amount int, blockchain *BlockChain) *Transaction {

	//unSpentTx := blockchain.unUTXOs(from)
	//fmt.Println(unSpentTx)

	money, spendableUTXODic := blockchain.FindSpendableUTXOs(from, amount)

	var txInputs []*TXInput
	var txOutputs []*TXOutput

	//1. consumption
	for txHash, indexArray := range spendableUTXODic {
		txHashBytes, _ := hex.DecodeString(txHash)
		for _, index := range indexArray {
			txInput := &TXInput{
				TxHash:    txHashBytes,
				Vout:      index,
				ScriptSig: from,
			}
			txInputs = append(txInputs, txInput)
		}
	}

	//2. transfer accounts
	txOutput := &TXOutput{
		Value:        int64(amount),
		ScriptPubKey: to,
	}
	txOutputs = append(txOutputs, txOutput)

	//3. change note
	txOutput = &TXOutput{
		Value:        money - int64(amount),
		ScriptPubKey: from,
	}
	txOutputs = append(txOutputs, txOutput)

	tx := &Transaction{
		TxHash: []byte{},
		Vins:   txInputs,
		Vouts:  txOutputs,
	}

	// setting hash
	tx.HashTranscation()

	return tx
}

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

// is coinbase transaction
func (tx *Transaction) IsCoinbaseTransaction() bool {
	return len(tx.Vins[0].TxHash) == 0 && tx.Vins[0].Vout == -1
}
