package main

import (
	"bytes"
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"os"
)

type DataSource interface {
	write(data []byte)
}

type FileDataSource struct {
	filename string
}

func (f *FileDataSource) write(data []byte) {
	file, _ := os.OpenFile(f.filename, os.O_WRONLY|os.O_CREATE, 0700)
	file.Write(data)
}

// decorators
type EncryptionDecorator struct {
	wrappee DataSource
	key     []byte
}

func (e *EncryptionDecorator) write(data []byte) {
	block, err := aes.NewCipher(e.key)
	if err != nil {
		panic(err.Error())
	}

	cipherText := make([]byte, aes.BlockSize+len(data))
	iv := cipherText[:aes.BlockSize]
	_, _ = io.ReadFull(rand.Reader, iv)

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], data)
	e.wrappee.write([]byte(base64.StdEncoding.EncodeToString(cipherText)))
}

type CompressionDecorator struct {
	wrappee DataSource
}

func (c *CompressionDecorator) write(data []byte) {
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	w.Write(data)
	c.wrappee.write(buf.Bytes())
}

func main() {
	fds := FileDataSource{filename: "test.txt"}
	ed := EncryptionDecorator{wrappee: &fds, key: []byte("passphrasewhichneedstobe32bytes!")}
	cd := CompressionDecorator{wrappee: &ed}

	cd.write([]byte("hello world. This is really long message.sdndjfusjlase"))
}
