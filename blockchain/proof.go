package blockchain

/* How does this work?

   The NewProof function creates a "big integer" whose value is 1 intially.
   A Big integer is an integer with alot of bits reserved to be filled later.

   Next we left-shift the target (which we had set to 1) by "256 - Difficulty" bits

   To better explain this, let us consider a 8 bit numeber

   0000 0001

   say we shift the abovenumber by 5, we get

   0010 0000

   Now, we have to find a hash, which is lesser than this number,
   ie first 3 digits are 0, only then will we have the unique hash.

   To find the hash, the data is fixed, so we vary the nonce number continously, by putting it in a "for loop"

   Finally the Validate function checks of the final hash value is actually less than the target hash value
*/

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

/* adjust this number based on the numeber of computers on the network
   and also as time passes and computational energy increases */
const Difficulty = 18

type ProofOfWork struct {
	Block  *Block
	Target *big.Int // target hash
}

// NewProof: creates the PoW and links the block to this PoW

func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)

	// Lsh stands for left shift.
	// 256 is the number of bytes in our hash (sha256.sum256)
	target.Lsh(target, uint(256-Difficulty))

	return &ProofOfWork{b, target}

}

func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)
	return data
}

// ToHex: We need this function because our nonce(int) and Difficulty(big.Int) are not of type byte
func ToHex(num int64) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buf.Bytes()
}

// Run: is the Main function for proof
func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}

	}
	fmt.Println()
	return nonce, hash[:]
}

// Validate: checks if the hash generated is actully less than the target hash or not.
func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int
	data := pow.InitData(pow.Block.Nonce)
	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}
