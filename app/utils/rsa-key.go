package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
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
