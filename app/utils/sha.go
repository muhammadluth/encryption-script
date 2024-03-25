package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
)

func EncryptSHA256(privateKey *rsa.PrivateKey, data []byte) (string, error) {
	dataSumSHA256 := sha256.Sum256(data)
	dataRSASignPSS, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, dataSumSHA256[:], nil)
	if err != nil {
		return "", err
	}
	dataEncodeToString := hex.EncodeToString(dataRSASignPSS)
	return dataEncodeToString, err
}

func VerifySHA256(publicKey *rsa.PublicKey, data, dataEncrypted []byte) error {
	dataSumSHA256 := sha256.Sum256(data)
	decodeDataEncrypted, err := hex.DecodeString(string(dataEncrypted))
	if err != nil {
		return err
	}
	err = rsa.VerifyPSS(publicKey, crypto.SHA256, dataSumSHA256[:], decodeDataEncrypted, nil)
	if err != nil {
		return err
	}
	return err
}
