package main

import (
	"github.com/0ne290/go-tasks/task1/internal"
	//"crypto/rand"
	"bufio"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
)

const privateKeyAsPemString string = `-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----`

const aesKey string = "ER9ghtUm724aCT0Eulu0AZkJw99d2hKF"

func main() {
	fmt.Print("RSA:\n\n\n")
	executeRsa()
	fmt.Print("\n\n\nAES:\n\n\n")
	executeAes()
}

func executeRsa() {
	privateKey := parseRsaPrivateKeyFromPemStr(privateKeyAsPemString)
	decryptor := internal.NewRsaDecryptor(sha256.New(), privateKey)
	publicKey := decryptor.GetPublicKey()
	encryptor := internal.NewRsaEncryptor(sha256.New(), &publicKey)
	stdinReader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter label: ")
	tmp, err := stdinReader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	label := []byte(tmp)

	fmt.Print("Enter source text: ")
	tmp, err = stdinReader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	source := []byte(tmp)

	encrypted := encryptor.Encrypt(source, label)
	decrypted := decryptor.Decrypt(encrypted, label)

	fmt.Printf("\nSource text: %s", string(source))
	fmt.Printf("Label: %s", string(label))
	fmt.Printf("\nEncryption result as Base64: %s", base64.StdEncoding.EncodeToString(encrypted))
	fmt.Printf("\nDecryption result: %s", string(decrypted))
	fmt.Printf("\n\nPrivate key as PEM string:\n\n%s", privateKeyAsPemString)
	fmt.Printf("\n\n\nPublic key of private key as PEM string:\n\n%s", exportRsaPublicKeyAsPemStr(&publicKey))
}

func executeAes() {
	aes := internal.NewAes([]byte(aesKey))
	stdinReader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter source text: ")
	tmp, err := stdinReader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	source := []byte(tmp)

	encrypted := aes.Encrypt(source)
	decrypted := aes.Decrypt(encrypted)

	fmt.Printf("\nSource text: %s", string(source))
	fmt.Printf("\nEncryption result as Base64: %s", base64.StdEncoding.EncodeToString(encrypted))
	fmt.Printf("\nDecryption result: %s", string(decrypted))
	fmt.Printf("\n\nKey: %s", aesKey)
}

func parseRsaPrivateKeyFromPemStr(privPEM string) *rsa.PrivateKey {
	block, _ := pem.Decode([]byte(privPEM))
	if block == nil {
		panic("Failed to parse PEM block containing the key!")
	}

	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	return priv
}

func exportRsaPublicKeyAsPemStr(pubkey *rsa.PublicKey) string {
	pubkey_bytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		panic(err)
	}

	pubkey_pem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubkey_bytes,
		},
	)

	return string(pubkey_pem)
}
