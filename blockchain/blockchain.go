package blockchain

import (
	"fmt"
	"time"
)

type Blockchain struct {
	StartBlockHash [32]byte
	EndBlockHash   [32]byte
	Blocks         map[[32]byte]*Block
}

func CreateBlockchain() *Blockchain {
	genesisBlock := NewBlock([]byte("genesis block"), &Block{})

	blockchain := &Blockchain{StartBlockHash: genesisBlock.Hash, EndBlockHash: genesisBlock.Hash}
	blockchain.Blocks = make(map[[32]byte]*Block)

	blockchain.Blocks[genesisBlock.Hash] = genesisBlock
	return blockchain
}

func (blockchain *Blockchain) AppendBlock(block *Block) {
	blockchain.Blocks[block.Hash] = block
	blockchain.EndBlockHash = block.Hash
}

func (blockchain *Blockchain) Print() {

	currentHash := blockchain.EndBlockHash

	for {
		block := blockchain.Blocks[currentHash]
		fmt.Printf("Hash: 0x%x\nPrevHash: 0x%x\nTimestamp: %s\nNonce: %d\nTransactions: %x\n\n", block.Hash, block.PrevHash, time.Unix(block.Timestamp, 0).String(), block.Nonce, block.TransactionsHash())

		if block.PrevHash == [32]byte{} {
			// genesisblock, lets get out of here
			return
		}

		currentHash = block.PrevHash
	}
}
