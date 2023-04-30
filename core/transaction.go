package core

import (
	"github.com/Anoencs/blockchain_layer1/crypto"
)

type Transaction struct {
	Data      []byte
	PublicKey crypto.PublicKey
	Signature *crypto.Signature
}
