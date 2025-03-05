package encryption

import (
	"crypto/aes"
    "crypto/cipher"
    "crypto/rand"
	"io"
)

type Aes struct {
	key []byte
}

func NewAes(key []byte) *Aes {
	return &Aes{key}
}

func (myAes *Aes) Encrypt(plaintext []byte) []byte {
	block, err := aes.NewCipher(myAes.key)
    if err != nil {
        panic(err)
    }

    ciphertext := make([]byte, aes.BlockSize+len(plaintext))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        panic(err)
    }

    stream := cipher.NewCTR(block, iv)
    stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
    return ciphertext
}

func (myAes *Aes) Decrypt(ciphertext []byte) []byte {
    block, err := aes.NewCipher(myAes.key)
    if err != nil {
        panic(err)
    }

    if len(ciphertext) < aes.BlockSize {
		panic("Ciphertext too short!")
    }
    iv := ciphertext[:aes.BlockSize]
    ciphertext = ciphertext[aes.BlockSize:]

    stream := cipher.NewCTR(block, iv)
    stream.XORKeyStream(ciphertext, ciphertext)
    return ciphertext
}