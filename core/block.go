package core

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"

	crypto "github.com/Anoencs/blockchain_layer1/crypto"
	types "github.com/Anoencs/blockchain_layer1/type"
)

type Header struct {
	Version       uint32
	DataHash      types.Hash
	PrevBlockHash types.Hash
	Timestamp     int64
	Height        uint32
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

func (b *Block) Hash(hasher Hasher[*Header]) types.Hash {
	if b.hash.IsZero() {
		b.hash = hasher.Hash(&b.Header)
	}

	return b.hash
}

func (b *Block) Sign(privKey crypto.PrivateKey) error {
	sig, err := privKey.Sign(b.Header.Bytes())
	if err != nil {
		return err
	}

	b.Validator = privKey.PublicKey()
	b.Signature = sig

	return nil
}

func (b *Block) Verify() error {
	if b.Signature == nil {
		return fmt.Errorf("block have no signature")
	}
	if !b.Signature.Verify(b.Validator, b.Header.Bytes()) {
		return fmt.Errorf("invalid signature")
	}

	for _, tx := range b.Transaction {
		if err := tx.Verify(); err != nil {
			return err
		}
	}

	return nil

}

func (h *Header) Bytes() []byte {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	enc.Encode(h)

	return buf.Bytes()
}

func (b *Block) AddTx(tx *Transaction) {
	b.Transaction = append(b.Transaction, *tx)
}
