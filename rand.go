package cryptoutil

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// GenerateAES128Key generate random AES128Key
func GenerateAES128Key() []byte {
	var ret AES128Key
	_, _ = rand.Read(ret[:])
	return ret[:]
}

// GenerateAES192Key generate random AES192Key
func GenerateAES192Key() []byte {
	var ret AES192Key
	_, _ = rand.Read(ret[:])
	return ret[:]
}

// GenerateAES256Key generate random AES256Key
func GenerateAES256Key() []byte {
	var ret AES256Key
	_, _ = rand.Read(ret[:])
	return ret[:]
}

// GenerateDESKey generate random DESKey
func GenerateDESKey() []byte {
	var ret DESKey
	_, _ = rand.Read(ret[:])
	return ret[:]
}

// GenerateTripleDESKey generate random TripleDESKey
func GenerateTripleDESKey() []byte {
	var ret TripleDESKey
	_, _ = rand.Read(ret[:])
	return ret[:]
}
