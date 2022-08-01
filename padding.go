package cryptoutil

import "bytes"

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	pad := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, pad...)
}

func PKCS5Padding(ciphertext []byte) []byte {
	return PKCS7Padding(ciphertext, 8)
}

func PKCSTrimming(ciphertext []byte) []byte {
	padding := int(ciphertext[len(ciphertext)-1])
	return ciphertext[:len(ciphertext)-padding]
}
