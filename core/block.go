package core

import (
	"io"

	crypto "github.com/Anoencs/blockchain_layer1/crypto"
	types "github.com/Anoencs/blockchain_layer1/type"
)

type Header struct {
	Version   uint32
	DataHash  types.Hash
	PrevBlock types.Hash
	Timestamp int64
	Height    uint32
}

type Block struct {
	Header
	Transaction []Transaction
	Validator   crypto.PublicKey
	Signature   *crypto.Signature

	// cached version header hash
	hash types.Hash
}

func NewBlock(header Header, txx []Transaction) *Block {
	return &Block{
		Header:      header,
		Transaction: txx,
	}
}

func (b *Block) Decode(r io.Reader, dec Decoder[*Block]) error {
	return dec.Decode(r, b)
}

func (b *Block) Encode(w io.Writer, enc Encoder[*Block]) error {
	return enc.Encode(w, b)
}

func (b *Block) Hash(hasher Hasher[*Block]) types.Hash {
	if b.hash.IsZero() {
		b.hash = hasher.Hash(b)
	}

	return b.hash
}
