package main

import (
	"blockchain/blockchain"
	"fmt"
	"strconv"
)

func main() {
	chain := blockchain.InitBlockChain()

	// Create a wallet for Alice
	aliceWallet, err := blockchain.NewWallet()
	if err != nil {
		fmt.Println("Error creating wallet for Alice:", err)
		return
	}

	fmt.Println("Alice's wallet created successfully")

	// Create a wallet for Bob
	bobWallet, err := blockchain.NewWallet()
	if err != nil {
		fmt.Println("Error creating wallet for Bob:", err)
		return
	}

	fmt.Println("Bob's wallet created successfully")

	// Create a transaction from Alice to Bob
	tx := &blockchain.Transaction{
		Sender:   aliceWallet.PublicKey.N.String(),
		Receiver: bobWallet.PublicKey.N.String(),
		Amount:   5.0,
	}
	fmt.Println("Alice to Bob transaction created successfully")

	// Sign the transaction with Alice's wallet
	signature, err := aliceWallet.SignTransaction(tx)
	if err != nil {
		fmt.Println("Error signing transaction:", err)
		return
	}

	// Verify the transaction with Alice's wallet
	err = blockchain.VerifiyTransaction(tx, aliceWallet.PublicKey, signature)
	if err != nil {
		fmt.Println("Error verifying transaction:", err)
		return
	}
	fmt.Println("Transaction verified successfully")

	// Add the transaction to the block
	chain.AddBlock("Block 1", "Alice", []*blockchain.Transaction{tx})
	fmt.Println()

	for _, block := range chain.Blocks {
		fmt.Printf("Previous hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash of block: %x\n", block.Hash)

		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("IsValidPow: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		fmt.Println("Transactions:")

		for _, tx := range block.Transactions {
			fmt.Printf("Sender: %s\n", tx.Sender)
			fmt.Printf("Receiver: %s\n", tx.Receiver)
			fmt.Printf("Amount: %f\n", tx.Amount)
			fmt.Printf("Coinbase: %t\n", tx.Coinbase)
			fmt.Println()
		}
		fmt.Println()
	}
}
