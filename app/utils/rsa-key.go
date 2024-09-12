package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/pkcs12"
)

func GetRSAKey(bytesPublicKey, bytesPrivateKey []byte) (publicKey *rsa.PublicKey, privateKey *rsa.PrivateKey, err error) {
	blockPublicKey, _ := pem.Decode(bytesPublicKey)
	if blockPublicKey == nil {
		return nil, nil, fmt.Errorf("failed to parse PEM block public key")
	}

	publicKeyInterface, err := x509.ParsePKIXPublicKey(blockPublicKey.Bytes)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse public key")
	}

	publicKey, ok := publicKeyInterface.(*rsa.PublicKey)
	if !ok {
		return nil, nil, fmt.Errorf("invalid data public key")
	}

	blockPrivateKey, _ := pem.Decode(bytesPrivateKey)
	if blockPrivateKey == nil {
		return nil, nil, fmt.Errorf("failed to parse PEM block private key")
	}

	privateKeyInterface, err := x509.ParsePKCS8PrivateKey(blockPrivateKey.Bytes)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to parse private key")
	}

	privateKey, ok = privateKeyInterface.(*rsa.PrivateKey)
	if !ok {
		return nil, nil, fmt.Errorf("invalid data private key")
	}
	return publicKey, privateKey, err
}

func GetRSAKeyFromFile(secretKeyFilename, secretKeyPassword string) (publicKey *rsa.PublicKey, privateKey *rsa.PrivateKey, err error) {
	readFile, err := os.ReadFile(secretKeyFilename)
	if err != nil {
		panic(err)
	}

	block, err := pkcs12.ToPEM(readFile, secretKeyPassword)
	if err != nil {
		return nil, nil, err
	}

	strPrivateKey := ""
	for _, v := range block {
		if v.Type == "PRIVATE KEY" {
			data := pem.EncodeToMemory(v)
			strPrivateKey = string(data)
			break
		}
		continue
	}

	decoded, _ := pem.Decode([]byte(strings.TrimSpace(strPrivateKey)))
	if decoded == nil {
		return nil, nil, fmt.Errorf("failed to decode private key")
	}

	privateKey, err = x509.ParsePKCS1PrivateKey(decoded.Bytes)
	if err != nil {
		return nil, nil, err
	}
	publicKey = &privateKey.PublicKey
	return publicKey, privateKey, err
}
