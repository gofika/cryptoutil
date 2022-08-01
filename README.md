[![codecov](https://codecov.io/gh/gofika/cryptoutil/branch/main/graph/badge.svg)](https://codecov.io/gh/gofika/cryptoutil)
[![Build Status](https://github.com/gofika/cryptoutil/workflows/build/badge.svg)](https://github.com/gofika/cryptoutil)
[![go.dev](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/gofika/cryptoutil)
[![Go Report Card](https://goreportcard.com/badge/github.com/gofika/cryptoutil)](https://goreportcard.com/report/github.com/gofika/cryptoutil)
[![Licenses](https://img.shields.io/github/license/gofika/cryptoutil)](LICENSE)


# CryptoUtil

Crypto algorithm wrapper for easy use. ex: AES,DES and others.

## Basic Usage

### Installation

To get the package, execute:

```bash
go get -u github.com/gofika/cryptoutil
```

### BlockReader

```go
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "fmt"
    "io"
    "math/rand"

    "github.com/gofika/cryptoutil"
)

func main() {
    // generate random iv for testing
    randIV := func() []byte {
        t := [aes.BlockSize]byte{}
        _, _ = rand.Read(t[:])
        return t[:]
    }
    iv := randIV()
    // generate random aes key for testing
    key := cryptoutil.GenerateAES256Key()

    f, _ := os.Open("src.data")
    defer f.Close()
    // target
    decrypted, _ := os.Create("decrypted.data")
    defer decrypted.Close()
    block, _ := aes.NewCipher(key)
    // block read wrapper for io.Reader
    r := cryptoutil.NewBlockReader(cipher.NewCBCDecrypter(block, iv), f)
    n, _ := io.Copy(decrypted, r)
    fmt.Printf("decrypted size: %d\n", n)
}
```

### BlockWriter

```go
package main

import (
    "crypto/aes"
    "crypto/cipher"
    "fmt"
    "io"
    "math/rand"

    "github.com/gofika/cryptoutil"
)

func main() {
    // generate random iv for testing
    randIV := func() []byte {
        t := [aes.BlockSize]byte{}
        _, _ = rand.Read(t[:])
        return t[:]
    }
    iv := randIV()
    // generate random aes key for testing
    key := cryptoutil.GenerateAES256Key()

    f, _ := os.Open("src.data")
    defer f.Close()
    // target
    encrypted, _ := os.Create("encrypted.data")
    defer encrypted.Close()
    block, _ := aes.NewCipher(key)
    // block wrapper for io.Writer
    w := cryptoutil.NewBlockWriter(cipher.NewCBCEncrypter(block, iv), encrypted)
    n, _ := io.Copy(w, f)
    // Note: Because it is block encryption, you need to manually call Close() when closing to write the block content of the cache
    w.Close()
    fmt.Printf("encrypted size: %d\n", n)
}
```

### AES

```go
package main

import (
    "crypto/aes"
    "fmt"

    "github.com/gofika/cryptoutil"
)

func main() {
    randIV := func() []byte {
        t := [aes.BlockSize]byte{}
        _, _ = rand.Read(t[:])
        return t[:]
    }
    iv := randIV()
    // generate random aes key for testing
    key := cryptoutil.GenerateAES256Key()

    // encrypt/decrypt bytes
    data := []byte("foo")
    encoded, _ := cryptoutil.AESEncrypt(data, key, iv)
    fmt.Printf("bytes encoded: %v\n", encoded)
    decoded, _ := cryptoutil.AESDecrypt(encoded, key, iv)
    fmt.Printf("bytes decoded: %s\n", decoded)

    // encrypt/decrypt string
    dataStr := "foo"
    encodedStr, _ := cryptoutil.AESEncryptString(dataStr, key, iv)
    fmt.Printf("string encoded: %v\n", encodedStr)
    decodedStr, err := cryptoutil.AESDecryptString(encodedStr, key, iv)
    fmt.Printf("string decoded: %s\n", decodedStr)
}
```

### DES

```go
package main

import (
    "crypto/des"
    "fmt"

    "github.com/gofika/cryptoutil"
)

func main() {
    randIV := func() []byte {
        t := [des.BlockSize]byte{}
        _, _ = rand.Read(t[:])
        return t[:]
    }
    iv := randIV()
    // generate random des key for testing
    key := cryptoutil.GenerateDESKey()

    // encrypt/decrypt bytes
    data := []byte("foo")
    encoded, _ := cryptoutil.DESEncrypt(data, key, iv)
    fmt.Printf("bytes encoded: %v\n", encoded)
    decoded, _ := cryptoutil.DESDecrypt(encoded, key, iv)
    fmt.Printf("bytes decoded: %s\n", decoded)

    // encrypt/decrypt string
    dataStr := "foo"
    encodedStr, _ := cryptoutil.DESEncryptString(dataStr, key, iv)
    fmt.Printf("string encoded: %v\n", encodedStr)
    decodedStr, err := cryptoutil.DESDecryptString(encodedStr, key, iv)
    fmt.Printf("string decoded: %s\n", decodedStr)
}
```