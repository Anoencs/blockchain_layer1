package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateKey(t *testing.T) {
	privKey := GeneratePrivatKey()
	pubKey := privKey.PublicKey()
	address := pubKey.Address()
	fmt.Print(address)
}

func TestSign(t *testing.T) {
	privKey := GeneratePrivatKey()

	msg := []byte("hello world")
	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)
	fmt.Println(sig)
}

func Test_KeyPair_Sign_Veriy_Valid(t *testing.T) {
	privKey := GeneratePrivatKey()
	pubKey := privKey.PublicKey()

	msg := []byte("hello world")

	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)
	assert.True(t, sig.Verify(pubKey, msg))
}

func Test_KeyPair_Sign_Verify_Failed(t *testing.T) {
	privKey := GeneratePrivatKey()
	pubKey := privKey.PublicKey()

	msg := []byte("hello world")

	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)

	otherPrivKey := GeneratePrivatKey()
	otherPublicKey := otherPrivKey.PublicKey()

	assert.False(t, sig.Verify(otherPublicKey, msg))
	assert.False(t, sig.Verify(pubKey, []byte("abcdef")))

}
