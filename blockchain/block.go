package blockchain

import (
	"bytes"
	"crypto/md5"
	"math/rand"
	"time"
)

// Task 2: Add your import here
// Task 1: Add your code here
type Block struct {
	Hash         string
	Data         string
	PrevHash     string
	Nonce        int
	Transactions []*Transaction
}

type Blockchain struct {
	Blocks []*Block
}

// Task 2: Compute the hash of the data here
func (b *Block) ComputeHash() {
	concatenatedData := bytes.Join([][]byte{[]byte(b.Data), []byte(b.PrevHash)}, []byte{})
	computedHash := md5.Sum(concatenatedData)
	b.Hash = string(computedHash[:])
}

// Task 3: Add your code here(Genesis Block)
func CreateBlock(data string, prevHash string, transactions []*Transaction) *Block {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	initialNonce := rand.Intn(10000)

	block := &Block{"", data, prevHash, initialNonce, transactions}

	newPow := NewProofOfWork(block)

	nonce, hash := newPow.MineBlock()

	block.Hash = string(hash[:])
	block.Nonce = nonce

	return block
}

func Genesis() *Block {
	coinbaseTransaction := &Transaction{
		Sender:   "Coinbase",
		Receiver: "Genesis",
		Amount:   0.0,
		Coinbase: true,
	}

	return CreateBlock("Genesis", "", []*Transaction{coinbaseTransaction})
}
