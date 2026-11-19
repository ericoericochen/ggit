package ggit

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"strconv"
)

type Blob struct {
	Data []byte
}

func (b *Blob) ObjectID() string {
	encoded := encodeBlobBytes(b.Data)
	sha := computeSha1Hash(encoded)
	return sha
}

func (b *Blob) Encode(w io.Writer) error {
	_, err := w.Write(encodeBlobBytes(b.Data))
	return err
}

// Get the uncompressed representation of blob
// blob <length> \0 <content>
func encodeBlobBytes(b []byte) []byte {
	var blobBuf bytes.Buffer
	blobBuf.WriteString("blob ")
	length := len(b)
	blobBuf.WriteString(strconv.Itoa(length))
	blobBuf.WriteByte(0)
	blobBuf.Write(b)
	return blobBuf.Bytes()
}

// Compute the sha1 hash. Returns the hex string.
func computeSha1Hash(b []byte) string {
	sha1Hasher := sha1.New()
	sha1Hasher.Write(b)
	sha := hex.EncodeToString(sha1Hasher.Sum(nil))
	return sha
}
