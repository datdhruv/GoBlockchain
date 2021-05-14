package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

// TxOutput: Contains the Value and Pubkey field
// Value is the denomination/amount of the transaction
// PubKey is the the Public key that is required to access the tokens in the Txn
// Outputs are indivisible
// TODO: Check, probably value is the amount given for a succesful mine
type TxOutput struct {
	Value  int
	PubKey string
}

// TxInput: Are references to previous outputs
type TxInput struct {
	ID  []byte
	Out int
	Sig string
}

type Transaction struct {
	ID      []byte
	Inputs  []TxInput
	Outputs []TxOutput
}

func CoinbaseTx(to string, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Coins to %s", to)
	}

	txin := TxInput{[]byte{}, -1, data}
	txout := TxOutput{100, to}

	tx := Transaction{nil, []TxInput{txin}, []TxOutput{txout}}
	tx.SetID()

	return &tx
}

// SetID: Creates and sets hash based on the bytes that are in the tx (Transaction)
func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	encode := gob.NewEncoder(&encoded)
	err := encode.Encode(tx)
	Handle(err)

	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]

}

// IsCoinbase: Checks if the transaction is a base transaction
func (tx *Transaction) IsCoinbase() bool {
	return len(tx.Inputs) == 1 && len(tx.Inputs[0].ID) == 0 && tx.Inputs[0].Out == -1
}

func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}

func NewTransaction(from, to string, amount int, chain *Blockchain) *Transaction {
	var inputs []TxInput
	var outputs []TxOutput

	acc, validOutpus := chain.FindSpendableOutputs(from, amount)
	if acc < amount {
		log.Panic("Erron: not enough funds")
	}

	for txid, outs := range validOutpus {
		txID, err := hex.DecodeString(txid)
		Handle(err)

		for _, out := range outs {
			input := TxInput{txID, out, from}
			inputs = append(inputs, input)
		}
	}

	outputs = append(outputs, TxOutput{amount, to})

	if acc > amount {
		outputs = append(outputs, TxOutput{acc - amount, from})
	}

	tx := Transaction{nil, inputs, outputs}
	tx.SetID()

	return &tx
}
