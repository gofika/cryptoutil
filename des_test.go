package cryptoutil

import (
	"crypto/des"
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDES(t *testing.T) {
	randIV := func() []byte {
		t := [des.BlockSize]byte{}
		_, _ = rand.Read(t[:])
		return t[:]
	}
	iv := randIV()
	testFn := func(key, iv []byte) {
		// bytes
		data := []byte("foo")
		encoded, err := DESEncrypt(data, key, iv)
		assert.Nil(t, err)
		decoded, err := DESDecrypt(encoded, key, iv)
		assert.Nil(t, err)
		assert.Equal(t, data, decoded)
		// str
		dataStr := "foo"
		encodedStr, err := DESEncryptString(dataStr, key, iv)
		assert.Nil(t, err)
		decodedStr, err := DESDecryptString(encodedStr, key, iv)
		assert.Nil(t, err)
		assert.Equal(t, dataStr, decodedStr)
	}
	testFn(GenerateDESKey(), iv)
	testTripleFn := func(key, iv []byte) {
		// bytes
		data := []byte("foo")
		encoded, err := TripleDESEncrypt(data, key, iv)
		assert.Nil(t, err)
		decoded, err := TripleDESDecrypt(encoded, key, iv)
		assert.Nil(t, err)
		assert.Equal(t, data, decoded)
		// str
		dataStr := "foo"
		encodedStr, err := TripleDESEncryptString(dataStr, key, iv)
		assert.Nil(t, err)
		decodedStr, err := TripleDESDecryptString(encodedStr, key, iv)
		assert.Nil(t, err)
		assert.Equal(t, dataStr, decodedStr)
	}
	testTripleFn(GenerateTripleDESKey(), iv)
}
