package blockchain

import (
	"crypto/ed25519"
	"crypto/sha256"
	"fmt"

	"github.com/bummie/aauchain/utils"
)

type Wallet struct {
	PublicKey  []byte
	PrivateKey []byte
}

func CreateWallet() Wallet {
	pubKey, privKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		fmt.Errorf("creating new wallet %w\n", err)
		return Wallet{}
	}

	return Wallet{PublicKey: pubKey, PrivateKey: privKey}
}

func (wallet *Wallet) Address() string {
	sha256HashPubKey := sha256.Sum256(wallet.PublicKey)
	return utils.Base58Encode(sha256HashPubKey[:])
}

/*
msg := []byte("The quick brown fox jumps over the lazy dog")

sig, err := priv.Sign(nil, msg, &Options{
	Context: "Example_ed25519ctx",
})
if err != nil {
	log.Fatal(err)
}

if err := VerifyWithOptions(pub, msg, sig, &Options{
	Context: "Example_ed25519ctx",
}); err != nil {
	log.Fatal("invalid signature")
}
*/
