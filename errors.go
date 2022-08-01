package cryptoutil

import (
	"strconv"
)

type KeySizeError int

func (k KeySizeError) Error() string {
	return "cryptoutil: invalid key size " + strconv.Itoa(int(k))
}

type ShortCiphertextSizeError int

func (s ShortCiphertextSizeError) Error() string {
	return "cryptoutil: short ciphertext size " + strconv.Itoa(int(s))
}
