package cryptoutil

import (
	"bytes"
	"crypto/cipher"
	"io"
)

// BlockWriter block ciphertext streaming writer
type BlockWriter struct {
	io.WriteCloser
	mode      cipher.BlockMode
	blockSize int
	buffer    *bytes.Buffer
	dst       io.Writer
	count     int64
}

// wrap BlockReader block ciphertext streaming writer
func NewBlockWriter(blockMode cipher.BlockMode, dst io.Writer) *BlockWriter {
	return &BlockWriter{
		mode:      blockMode,
		blockSize: blockMode.BlockSize(),
		dst:       dst,
		buffer:    bytes.NewBuffer([]byte{}),
	}
}

// Note: Because it is block encryption, you need to manually call Close() when closing to write the block content of the cache
func (c *BlockWriter) Close() error {
	buf := PKCS7Padding(c.buffer.Next(c.buffer.Len()), c.blockSize)
	c.mode.CryptBlocks(buf, buf)
	written, err := io.Copy(c.dst, bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	c.count += written
	return nil
}

func (c *BlockWriter) Write(data []byte) (int, error) {
	if _, err := c.buffer.Write(data); err != nil {
		return 0, err
	}
	bufferLen := c.buffer.Len()
	remaining := bufferLen - bufferLen%c.blockSize
	buf := c.buffer.Next(remaining)
	c.mode.CryptBlocks(buf, buf)
	written, err := io.Copy(c.dst, bytes.NewBuffer(buf))
	if err != nil {
		return 0, err
	}
	c.count += written
	return len(data), nil
}
