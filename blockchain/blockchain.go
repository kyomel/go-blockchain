package blockchain

// Task 1: Add your code here
type Transaction struct {
	Sender   string
	Receiver string
	Amount   float64
	Coinbase bool
}

// Task 4: Add your code here
func InitBlockChain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}

func (chain *Blockchain) AddBlock(data string, coinbaseRcpt string, transactions []*Transaction) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]

	coinbaseTransaction := &Transaction{
		Sender:   "Coinbase",
		Receiver: coinbaseRcpt,
		Amount:   10.0,
		Coinbase: true,
	}

	newBlock := CreateBlock(data, prevBlock.Hash, append([]*Transaction{coinbaseTransaction}, transactions...))

	chain.Blocks = append(chain.Blocks, newBlock)
}

// Task 10: Add your code here
