package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

const TargetBitDifficulty = 21

func TargetBits(targetDifficulty int64) *big.Int {
	//0x0000000000000000000000000000000000000000000000000000000000000001 = 1
	targetNumber := big.NewInt(1)

	targetNumber.Lsh(targetNumber, uint(256-targetDifficulty))
	// Left shift by 256-targetDifficulty
	//0x0000010000000000000000000000000000000000000000000000000000000000 = 6901746346790563787434755862277025452451108972170386555162524223799296
	// Target would be any number below the one above
	return targetNumber
}

func Mine(block *Block, target *big.Int) (int64, [32]byte) {

	publicBytes := block.PublicBytes()
	var nonce int64 = 0
	var calculatedHash big.Int

	for nonce < math.MaxInt64 {
		sha256Sum := sha256.Sum256(bytes.Join([][]byte{publicBytes, int64ToLenBytes(nonce, 8)}, []byte{}))
		calculatedHash.SetBytes(sha256Sum[:])

		// if calculated hash is smaller than our target
		if calculatedHash.Cmp(target) == -1 {
			fmt.Printf("mined -> 0x%x, nonce: %d\n", sha256Sum, nonce)
			return nonce, sha256Sum
		}

		nonce++
	}

	return 0, [32]byte{}
}
