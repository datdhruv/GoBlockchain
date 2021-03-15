package blockchain

// TxOutput: Contains the Value and Pubkey field
// Value is the denomination/amount of the transaction
// PubKey is the the Public key that is required to access the tokens in the Txn
// Outputs are indivisible
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

	// Here we are just saying that we are giving 100 tokens for free
	// If the data field is null
	txin := TxInput{[]byte{}, -1, data}
	txout := TxOutput{100, to}

	tx := Transaction{nil, []TxInput{txin}, []TxOutput{txout}}
	tx.SetID()

	return &tx
}

// SetID: Creates and sets hash based on the bytes that are in the tx (Transaction)
func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]hash

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
