package cryptoutil

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"io"
)

type DESKey [des.BlockSize]byte
type TripleDESKey [des.BlockSize * 3]byte

// NewDESReader creates and returns a new BlockReader
func NewDESReader(r io.Reader, key []byte, iv []byte) (*BlockReader, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return NewBlockReader(cipher.NewCBCDecrypter(block, iv), r), nil
}

// NewTripleDESReader creates and returns a new BlockReader
func NewTripleDESReader(r io.Reader, key []byte, iv []byte) (*BlockReader, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	return NewBlockReader(cipher.NewCBCDecrypter(block, iv), r), nil
}

// NewDESWriter creates and returns a new BlockWriter
func NewDESWriter(dst io.Writer, k []byte, iv []byte) (*BlockWriter, error) {
	block, err := des.NewCipher(k)
	if err != nil {
		return nil, err
	}
	return NewBlockWriter(cipher.NewCBCEncrypter(block, iv), dst), nil
}

// NewTripleDESWriter creates and returns a new BlockWriter
func NewTripleDESWriter(dst io.Writer, k []byte, iv []byte) (*BlockWriter, error) {
	block, err := des.NewTripleDESCipher(k)
	if err != nil {
		return nil, err
	}
	return NewBlockWriter(cipher.NewCBCEncrypter(block, iv), dst), nil
}

// DESEncrypt
func DESEncrypt(s []byte, key []byte, iv []byte) ([]byte, error) {
	buf := bytes.NewBuffer(s)
	out := bytes.NewBuffer([]byte{})
	w, err := NewDESWriter(out, key, iv)
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

// DESEncryptString
func DESEncryptString(s string, key []byte, iv []byte) (string, error) {
	bs, err := DESEncrypt([]byte(s), key, iv)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

// DESDecrypt
func DESDecrypt(s []byte, key []byte, iv []byte) ([]byte, error) {
	buf := bytes.NewBuffer(s)
	out := bytes.NewBuffer([]byte{})
	r, err := NewDESReader(buf, key, iv)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(out, r)
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

// DESDecryptString
func DESDecryptString(s string, key []byte, iv []byte) (string, error) {
	bs, err := DESDecrypt([]byte(s), key, iv)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

// TripleDESEncrypt
func TripleDESEncrypt(s []byte, key []byte, iv []byte) ([]byte, error) {
	buf := bytes.NewBuffer(s)
	out := bytes.NewBuffer([]byte{})
	w, err := NewTripleDESWriter(out, key, iv)
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

// TripleDESEncryptString
func TripleDESEncryptString(s string, key []byte, iv []byte) (string, error) {
	bs, err := TripleDESEncrypt([]byte(s), key, iv)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

// TripleDESDecrypt
func TripleDESDecrypt(s []byte, key []byte, iv []byte) ([]byte, error) {
	buf := bytes.NewBuffer(s)
	out := bytes.NewBuffer([]byte{})
	r, err := NewTripleDESReader(buf, key, iv)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(out, r)
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

// TripleDESDecryptString
func TripleDESDecryptString(s string, key []byte, iv []byte) (string, error) {
	bs, err := TripleDESDecrypt([]byte(s), key, iv)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}
