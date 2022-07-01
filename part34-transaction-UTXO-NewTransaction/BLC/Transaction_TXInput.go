package BLC

type TXInput struct {
	//1. transaction hash
	TxHash []byte
	//2. store index in Vout
	Vout int
	//3. user name
	ScriptSig string
}
