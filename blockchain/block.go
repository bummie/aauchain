package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"log"
	"strconv"
	"time"
)

type Block struct {
	Timestamp    int64
	Transactions []*Transaction
	Hash         [32]byte
	PrevHash     [32]byte
	Nonce        int64
}

func (block *Block) CalculateHash() {
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))

	headers := bytes.Join([][]byte{block.PrevHash[:], block.TransactionsHash(), timestamp}, []byte{})

	block.Hash = sha256.Sum256(headers)
}

func (block *Block) TransactionsHash() []byte {

	transactionBytes := [][]byte{}

	for _, transaction := range block.Transactions {
		transactionBytes, err := transaction.Serialize()
		if err != nil {
			log.Panic("could not serialize transaction " + fmt.Sprintf("%x", transaction.ID) + "\n" + err.Error())
		}

		transactionBytes = append(transactionBytes, transactionBytes...)
	}

	transactionHash := sha256.Sum256(bytes.Join(transactionBytes, []byte{}))
	return transactionHash[:]
}

func (block *Block) PublicBytes() []byte {
	return bytes.Join([][]byte{block.Hash[:], block.PrevHash[:], block.TransactionsHash(), int64ToLenBytes(block.Timestamp, 8), TargetBits(TargetBitDifficulty).Bytes()}, []byte{})
}

func NewBlock(data []byte, prevBlock *Block) *Block {
	block := &Block{Timestamp: time.Now().Unix(), PrevHash: prevBlock.Hash, Transactions: []*Transaction{}}
	nonce, hash := Mine(block, TargetBits(TargetBitDifficulty))
	block.Nonce = nonce
	block.Hash = hash

	return block
}
