package blockchain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp int64
	Data      []byte
	Hash      [32]byte
	PrevHash  [32]byte
	Nonce     int64
}

func (block *Block) CalculateHash() {
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))
	headers := bytes.Join([][]byte{block.PrevHash[:], block.Data, timestamp}, []byte{})

	block.Hash = sha256.Sum256(headers)
}

func (block *Block) PublicBytes() []byte {
	return bytes.Join([][]byte{block.Hash[:], block.PrevHash[:], block.Data, int64ToLenBytes(block.Timestamp, 8), TargetBits(TargetBitDifficulty).Bytes()}, []byte{})
}

func NewBlock(data []byte, prevBlock *Block) *Block {
	block := &Block{Timestamp: time.Now().Unix(), PrevHash: prevBlock.Hash, Data: data}
	nonce, hash := Mine(block, TargetBits(TargetBitDifficulty))
	block.Nonce = nonce
	block.Hash = hash

	return block
}
