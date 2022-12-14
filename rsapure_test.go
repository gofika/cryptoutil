package cryptoutil

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRSAPure(t *testing.T) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	assert.Nil(t, err)
	pubKey := &privateKey.PublicKey
	buf := x509.MarshalPKCS1PublicKey(pubKey)

	key := RSAPublicKeyFromBytes(pubKey.N.Bytes(), 0)
	assert.True(t, pubKey.Equal(key))

	key, err = x509.ParsePKCS1PublicKey(buf)
	assert.Nil(t, err)
	assert.True(t, pubKey.Equal(key))
}
