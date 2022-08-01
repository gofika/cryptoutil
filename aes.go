package cryptoutil

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"io"
)

type AES128Key [aes.BlockSize]byte
type AES192Key [24]byte
type AES256Key [32]byte

// NewAESReader creates and returns a new BlockReader. The key argument should be the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
func NewAESReader(r io.Reader, key []byte, iv []byte) (*BlockReader, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return NewBlockReader(cipher.NewCBCDecrypter(block, iv), r), nil
}

// NewAESWriter creates and returns a new BlockWriter. The key argument should be the AES key, either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256.
func NewAESWriter(dst io.Writer, k []byte, iv []byte) (*BlockWriter, error) {
	block, err := aes.NewCipher(k)
	if err != nil {
		return nil, err
	}
	return NewBlockWriter(cipher.NewCBCEncrypter(block, iv), dst), nil
}

// AESEncrypt
func AESEncrypt(s []byte, key []byte, iv []byte) ([]byte, error) {
	buf := bytes.NewBuffer(s)
	out := bytes.NewBuffer([]byte{})
	w, err := NewAESWriter(out, key, iv)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(w, buf)
	if err != nil {
		return nil, err
	}
	w.Close()
	return out.Bytes(), nil
}

// AESEncryptString
func AESEncryptString(s string, key []byte, iv []byte) (string, error) {
	bs, err := AESEncrypt([]byte(s), key, iv)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

// AESDecrypt
func AESDecrypt(s []byte, key []byte, iv []byte) ([]byte, error) {
	buf := bytes.NewBuffer(s)
	out := bytes.NewBuffer([]byte{})
	r, err := NewAESReader(buf, key, iv)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(out, r)
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

// AESDecryptString
func AESDecryptString(s string, key []byte, iv []byte) (string, error) {
	bs, err := AESDecrypt([]byte(s), key, iv)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}
