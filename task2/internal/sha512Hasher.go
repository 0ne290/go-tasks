package internal

import (
	"crypto/sha512"
	"encoding/hex"
	"hash"
)

type Sha512Hasher struct {
	hasher hash.Hash
}

func NewSha512Hasher() *Sha512Hasher {
	return &Sha512Hasher{sha512.New()}
}

func (sha512Hasher *Sha512Hasher) Hash(source string) string {
	sha512Hasher.hasher.Write([]byte(source))

	ret := hex.EncodeToString(sha512Hasher.hasher.Sum(nil))

	sha512Hasher.hasher.Reset()

	return ret
}
