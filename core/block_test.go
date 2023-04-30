package core

import (
	"fmt"
	"testing"
	"time"

	types "github.com/Anoencs/blockchain_layer1/type"
)

func randomBlock(height uint32) *Block {
	header := Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Height:    height,
		Timestamp: time.Now().UnixNano(),
	}
	tx := Transaction{
		Data: []byte("foo"),
	}

	return NewBlock(header, []Transaction{tx})
}

func TestBlockHash(t *testing.T) {
	b := randomBlock(0)
	fmt.Print(b.Hash(BlockHasher{}))

}
