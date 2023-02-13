package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"

	"golang.org/x/crypto/ssh"
)

func GenerateKey() ([]byte, []byte, error) {
	// generate RSA key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return []byte{}, []byte{}, errors.New("failed to generate key")
	}

	// generate private key block
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	// generate public key
	publicKey, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		return []byte{}, []byte{}, errors.New("failed to generate public key")
	}

	// get private key bytes
	privateKeyBytes := pem.EncodeToMemory(privateKeyPEM)
	if err != nil {
		return []byte{}, []byte{}, errors.New("failed to encode private key")
	}

	// get public key bytes
	publicKeyBytes := ssh.MarshalAuthorizedKey(publicKey)

	return privateKeyBytes, publicKeyBytes, nil
}
