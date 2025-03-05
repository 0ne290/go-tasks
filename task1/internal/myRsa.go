package encryption

import (
	"hash"
	"crypto/rsa"
	"crypto/rand"
)

type RsaEncryptor struct {
	hasher hash.Hash
	publicKey *rsa.PublicKey
}

type RsaDecryptor struct {
	hasher hash.Hash
	privateKey *rsa.PrivateKey
	publicKey *rsa.PublicKey
}

func NewRsaEncryptor(hasher hash.Hash, publicKey *rsa.PublicKey) *RsaEncryptor {
	return &RsaEncryptor{hasher, publicKey}
}

func NewRsaDecryptor(hasher hash.Hash, privateKey *rsa.PrivateKey) *RsaDecryptor {
	return &RsaDecryptor{hasher, privateKey, &privateKey.PublicKey}
}

func (encryptor *RsaEncryptor) Encrypt(plaintext []byte, label []byte) []byte {
    ciphertext, err := rsa.EncryptOAEP(encryptor.hasher, rand.Reader, encryptor.publicKey, plaintext, label)
    if err != nil {
        panic(err)
    }
    return ciphertext
}

func (decryptor *RsaDecryptor) Decrypt(plaintext []byte, label []byte) []byte {
    ciphertext, err := rsa.DecryptOAEP(decryptor.hasher, rand.Reader, decryptor.privateKey, plaintext, label)
    if err != nil {
        panic(err)
    }
    return ciphertext
}

func (decryptor *RsaDecryptor) GetPublicKey() rsa.PublicKey {
    return *decryptor.publicKey
}