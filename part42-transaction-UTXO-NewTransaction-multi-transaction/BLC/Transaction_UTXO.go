package BLC

type UTXO struct {
	TxHash []byte
	Index  int
	OutPut *TXOutput
}
