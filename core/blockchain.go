package core

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type Blockchain struct {
	headers   []*Header
	store     Storage
	validator Validator
}

func NewBlockchain(genesis *Block) *Blockchain {
	bc := &Blockchain{
		headers: []*Header{},
		store:   NewMemoryStorage(),
	}
	bc.validator = NewBlockValidator(bc)
	bc.addBlockWithoutValidation(genesis)

	return bc
}

func (bc *Blockchain) SetValidator(v Validator) error {
	bc.validator = v
	return nil
}

func (bc *Blockchain) AddBlock(b *Block) error {
	if err := bc.validator.ValidateBlock(b); err != nil {
		return err
	}

	return bc.addBlockWithoutValidation(b)
}

func (bc *Blockchain) HasBlock(height uint32) bool {
	return height <= bc.Height()
}

func (bc *Blockchain) Height() uint32 {
	return uint32(len(bc.headers) - 1)
}

func (bc *Blockchain) addBlockWithoutValidation(b *Block) error {
	bc.headers = append(bc.headers, &b.Header)

	logrus.WithFields(logrus.Fields{
		"height": b.Height,
		"hash":   b.Hash(BlockHasher{}),
	}).Info("adding new block")

	return bc.store.Put(b)
}

func (bc *Blockchain) GetHeader(height uint32) (*Header, error) {
	if height > bc.Height() {
		return nil, fmt.Errorf("given height is too high")
	}
	return bc.headers[height], nil
}
