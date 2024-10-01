package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
)

var coinbaseAmount = 50

type Transaction struct {
	ID      [32]byte
	Inputs  []TransInput
	Outputs []TransOuput
}

func (transaction *Transaction) Serialize() ([]byte, error) {
	var encoded bytes.Buffer

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(transaction)
	if err != nil {
		return []byte{}, fmt.Errorf("serializing transaction body %w\n", err)
	}

	return encoded.Bytes(), nil
}

func (transaction *Transaction) Hash() ([32]byte, error) {
	var hash [32]byte

	transactionCopy := *transaction
	transactionCopy.ID = [32]byte{}

	transactionBytes, err := transactionCopy.Serialize()

	if err != nil {
		return [32]byte{}, err
	}

	hash = sha256.Sum256(transactionBytes)

	return hash, nil
}

type TransInput struct {
	TransactionId [32]byte
	OutputIndex   int
	Signature     []byte
	PublicKey     []byte
}

type TransOuput struct {
	Value         int
	PublicKeyHash [32]byte
}

// reward to miner
func NewCoinBase(address [32]byte) (*Transaction, error) {

	data := fmt.Sprintf("Coinbase: %d awarded to %x for mining!", coinbaseAmount, address)

	transaction := Transaction{
		Inputs: []TransInput{TransInput{
			TransactionId: [32]byte{},
			OutputIndex:   -1,
			Signature:     nil,
			PublicKey:     []byte(data),
		}},
		Outputs: []TransOuput{TransOuput{
			Value:         coinbaseAmount,
			PublicKeyHash: address,
		}},
	}

	if hash, err := transaction.Hash(); err != nil {
		transaction.ID = hash
	} else {
		return &Transaction{}, fmt.Errorf("creating new coinbase '%s' %w\n", data, err)
	}

	return &transaction, nil
}
