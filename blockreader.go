package cryptoutil

import (
	"bytes"
	"crypto/cipher"
	"io"
)

// BlockReader block ciphertext streaming reader
type BlockReader struct {
	io.ReadCloser
	mode      cipher.BlockMode
	blockSize int
	buffer    *bytes.Buffer
	peek      *bytes.Buffer
	src       io.Reader
	eof       bool
}

// wrap BlockReader block ciphertext streaming reader
func NewBlockReader(blockMode cipher.BlockMode, src io.Reader) *BlockReader {
	return &BlockReader{
		mode:      blockMode,
		blockSize: blockMode.BlockSize(),
		src:       src,
		buffer:    bytes.NewBuffer([]byte{}),
		peek:      bytes.NewBuffer([]byte{}),
		eof:       false,
	}
}

func (c *BlockReader) Read(data []byte) (int, error) {
	missing := len(data) - c.buffer.Len()
	for !c.eof && missing > 0 {
		r := missing + c.blockSize + 1
		if off := r % c.blockSize; off > 0 {
			r += c.blockSize - off
		}
		c.peek.Grow(r)
		n, err := io.CopyN(c.peek, c.src, int64(r))
		if err == io.EOF {
			c.eof = true
		} else if err != nil {
			return 0, err
		}
		rn := int(n)
		if rn%c.blockSize > 0 && c.eof {
			return 0, ShortCiphertextSizeError(rn)
		}
		buf := c.peek.Next(rn - rn%c.blockSize)
		c.mode.CryptBlocks(buf, buf)
		if _, err := c.buffer.Write(buf); err != nil {
			return 0, err
		}
		if c.eof { // padding
			pad := c.buffer.Bytes()[c.buffer.Len()-1]
			c.buffer.Truncate(c.buffer.Len() - int(pad))
		}
		missing = len(data) - c.buffer.Len()
	}
	return c.buffer.Read(data)
}

func (c *BlockReader) Close() error {
	c.eof = true
	return nil
}
