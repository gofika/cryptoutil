package cryptoutil

import (
	"crypto/aes"
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAES(t *testing.T) {
	randIV := func() []byte {
		t := [aes.BlockSize]byte{}
		_, _ = rand.Read(t[:])
		return t[:]
	}
	iv := randIV()
	testFn := func(key, iv []byte) {
		// bytes
		data := []byte("foo")
		encoded, err := AESEncrypt(data, key, iv)
		assert.Nil(t, err)
		decoded, err := AESDecrypt(encoded, key, iv)
		assert.Nil(t, err)
		assert.Equal(t, data, decoded)
		// str
		dataStr := "foo"
		encodedStr, err := AESEncryptString(dataStr, key, iv)
		assert.Nil(t, err)
		decodedStr, err := AESDecryptString(encodedStr, key, iv)
		assert.Nil(t, err)
		assert.Equal(t, dataStr, decodedStr)
	}
	testFn(GenerateAES128Key(), iv)
	testFn(GenerateAES192Key(), iv)
	testFn(GenerateAES256Key(), iv)
}
