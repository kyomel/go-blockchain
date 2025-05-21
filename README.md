# Blockchain Implementation in Go

A simple blockchain implementation in Go that demonstrates the core concepts of blockchain technology, including block creation, cryptographic hashing, proof-of-work consensus mechanism, transactions, wallet management, and digital signatures.

## Table of Contents

- [Overview](#overview)
- [Project Structure](#project-structure)
- [Core Components](#core-components)
  - [Block](#block)
  - [Blockchain](#blockchain)
  - [Proof of Work](#proof-of-work)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Running the Project](#running-the-project)
- [How It Works](#how-it-works)
- [API Reference](#api-reference)
- [Contributing](#contributing)
- [License](#license)

## Overview

This project implements a basic blockchain that includes:

- Block creation with cryptographic hashing using MD5
- Blockchain structure with linked blocks using cryptographic hashes
- Genesis block creation
- Transaction support with sender, receiver, and amount
- Wallet generation using RSA public-key cryptography
- Transaction signing and verification
- Coinbase transactions for block rewards
- Proof-of-Work (PoW) consensus mechanism with configurable difficulty
- Mining simulation with nonce finding

## Project Structure

```
blockchain-go/
├── blockchain/
│   ├── block.go         # Block structure and operations
│   ├── blockchain.go   # Blockchain implementation
│   ├── proof.go        # Proof of Work implementation
│   └── wallet.go      # Wallet and transaction signing implementation
├── main.go            # Example usage
├── go.mod            # Go module definition
└── README.md         # This documentation
```

## Core Components

### Wallet

The wallet implementation provides:

- RSA key pair generation (2048-bit)
- Transaction signing with private key
- Transaction verification with public key
- Secure key storage

Example usage:

```go
// Create a new wallet
aliceWallet, err := blockchain.NewWallet()
if err != nil {
    // handle error
}

// Create a transaction
tx := &blockchain.Transaction{
    Sender:   aliceWallet.PublicKey.N.String(),
    Receiver: "recipient-public-key",
    Amount:   5.0,
}

// Sign the transaction
signature, err := aliceWallet.SignTransaction(tx)
if err != nil {
    // handle error
}

// Verify the transaction
err = blockchain.VerifiyTransaction(tx, aliceWallet.PublicKey, signature)
if err != nil {
    // invalid signature
}
```

### Block

Each block in the blockchain contains:

- `Data`: The actual data stored in the block (string)
- `Hash`: The cryptographic hash of the block (MD5)
- `PrevHash`: The hash of the previous block in the chain (links blocks together)
- `Nonce`: A number used once in the mining process to find a valid hash
- `Transactions`: A list of transactions included in the block

### Transaction

Each transaction in the blockchain contains:

- `Sender`: The sender's address (string)
- `Receiver`: The recipient's address (string)
- `Amount`: The amount being transferred (float64)
- `Coinbase`: Boolean indicating if this is a coinbase transaction (mining reward)

### Blockchain

The blockchain is a series of blocks where each block is linked to its predecessor through the `PrevHash` field. The chain starts with a special block called the Genesis block.

### Proof of Work

The blockchain uses a Proof of Work (PoW) consensus mechanism to secure the network and create new blocks. The implementation includes:

- Configurable mining difficulty (set to 10 by default)
- Nonce-based mining to find a valid hash
- MD5 hashing algorithm (for demonstration purposes)
- Target difficulty adjustment through left bit shifting

The mining process involves:

1. Creating a target value by left-shifting 1 by (256 - Difficulty) bits
2. Iterating through nonce values until a valid hash is found
3. Validating that the hash meets the target difficulty
4. Including the nonce and difficulty in the block header

## Getting Started

### Prerequisites

- Go 1.15 or higher
- Git

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/blockchain-go.git
   cd blockchain-go/blockChainScratch
   ```

2. Build the project:
   ```bash
   go build
   ```

### Running the Project

Run the main program:

```bash
go run main.go
```

## Security Features

- **Digital Signatures**: All transactions are signed using RSA-PKCS1v15 with SHA-256
- **Key Management**: Secure generation and handling of RSA key pairs
- **Transaction Integrity**: Each transaction is hashed and signed to prevent tampering
- **Proof of Work**: Protects against spam and Sybil attacks

## How It Works

1. **Wallet Creation**:

   - Generates a new RSA key pair (2048-bit)
   - Provides methods for transaction signing and verification

2. **Transaction Flow**:

   - Sender creates a transaction with recipient and amount
   - Transaction is signed with sender's private key
   - Signature is verified using sender's public key
   - Valid transactions are added to the mempool for mining

3. **Block Creation**:
   - The program initializes a new blockchain with a genesis block
   - It adds three sample blocks to the chain, each requiring proof-of-work mining
   - Miners compete to find a valid nonce that produces a hash below the target difficulty
   - Once a valid nonce is found, the block is added to the chain
   - It prints out the details of each block, including its hash, previous block's hash, and mining information

## API Reference

### Block

```go
// ComputeHash calculates the hash of the block
func (b *Block) ComputeHash()

// CreateBlock creates a new block with the given data and previous hash
func CreateBlock(data string, prevHash string) *Block

// Genesis creates the first block in the blockchain
func Genesis() *Block
```

### Blockchain

```go
// InitBlockChain initializes a new blockchain with a genesis block
func InitBlockChain() *Blockchain

// AddBlock adds a new block to the blockchain with the given data, coinbase recipient, and transactions
// It automatically creates a coinbase transaction as the first transaction in the block
func (chain *Blockchain) AddBlock(data string, coinbaseRcpt string, transactions []*Transaction)
```

## Example Usage

```go
package main

import (
    "blockchain/blockchain"
    "fmt"
    "strconv"
)

func main() {
    // Initialize a new blockchain with genesis block
    chain := blockchain.InitBlockChain()

    // Add some blocks with transactions to the chain
    chain.AddBlock("Block 1", "Alice", []*blockchain.Transaction{
        {Sender: "Alice", Receiver: "Bob", Amount: 1.5},
        {Sender: "Alice", Receiver: "Charlie", Amount: 19.5},
    })

    chain.AddBlock("Block 2", "Bob", []*blockchain.Transaction{
        {Sender: "Bob", Receiver: "Charlie", Amount: 2.3},
    })


    // Print block information
    for _, block := range chain.Blocks {
        fmt.Printf("Previous hash: %x\n", block.PrevHash)
        fmt.Printf("Data in Block: %s\n", block.Data)
        fmt.Printf("Hash of block: %x\n", block.Hash)

        // Validate the proof of work
        pow := blockchain.NewProofOfWork(block)
        fmt.Printf("IsValidPow: %s\n\n", strconv.FormatBool(pow.Validate()))

        // Print transactions
        fmt.Println("Transactions:")
        for _, tx := range block.Transactions {
            fmt.Printf("  Sender: %s\n", tx.Sender)
            fmt.Printf("  Receiver: %s\n", tx.Receiver)
            fmt.Printf("  Amount: %f\n", tx.Amount)
            fmt.Printf("  Coinbase: %t\n\n", tx.Coinbase)
        }
    }
}
```

Note: To test what you have developed, sign and verify new transactions and add them to different blocks. Add those blocks to the blockchain and run the proof of work algorithm on these newly created blocks.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
