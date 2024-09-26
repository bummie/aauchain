package main

import (
	"github.com/bummie/aauchain/blockchain"
)

func main() {

	chain := blockchain.CreateBlockchain()
	chain.AppendBlock(blockchain.NewBlock([]byte("cool block"), chain.Blocks[chain.EndBlockHash]))
	chain.AppendBlock(blockchain.NewBlock([]byte("block to the people"), chain.Blocks[chain.EndBlockHash]))
	chain.AppendBlock(blockchain.NewBlock([]byte("block harder this chain is sick"), chain.Blocks[chain.EndBlockHash]))
	chain.AppendBlock(blockchain.NewBlock([]byte("eyoo"), chain.Blocks[chain.EndBlockHash]))
	chain.AppendBlock(blockchain.NewBlock([]byte("yeet"), chain.Blocks[chain.EndBlockHash]))

	chain.Print()
}
