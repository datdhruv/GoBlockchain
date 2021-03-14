package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

//func (b *Block) DeriveHash() {
//	// bytes.Join() takes two arguments,
//	// A 2d array of bytes and a sep of type []byte
//	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
//	hash := sha256.Sum256(info)
//	b.Hash = hash[:]
//}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	//block.DeriveHash()
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Nonce = nonce
	block.Hash = hash
	return block
}

func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// Genesis is the First Block in the chain
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// Serialize: is used because the badger database takes only serialized byte array
// What is serialization in this context? It is a method with which you send the Type information of the variable
// This is where the gob.NewEncoder come into picture
func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)
	Handle(err)

	return res.Bytes()

}

// Deserialize: is used because the badger database takes only serialized byte array
func Deserialize(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)

	Handle(err)

	return &block
}
