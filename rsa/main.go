package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
)

func main() {
	text := []byte("This is a secret!")
	fmt.Printf("Text is: %x(%s)\n", text, string(text))
	encrypted, e := Encrypt(text)
	if e != nil {
		panic(e)
	}
	fmt.Printf("Encrypted text is: %x(%s)\n", encrypted, string(encrypted))
	orgData, e := Decrypt(encrypted)
	if e != nil {
		panic(e)
	}
	fmt.Printf("Decrypted text is: %x(%s)\n", orgData, string(orgData))
}

func Encrypt(orgData []byte) ([]byte, error) {
	pubKey, _ := ioutil.ReadFile("rsa/public.pem")
	block, _ := pem.Decode(pubKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, orgData)
}

func Decrypt(cipherText []byte) ([]byte, error) {
	prvKey, _ := ioutil.ReadFile("rsa/private.pem")
	block, _ := pem.Decode(prvKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
}
