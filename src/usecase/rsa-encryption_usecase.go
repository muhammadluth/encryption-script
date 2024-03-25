package usecase

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encryption-script/app/utils"
	"encryption-script/model"
	"encryption-script/src"
	"net/http"
)

type RSAEncryptionUsecase struct {
}

func NewRSAEncryptionUsecase() src.IRSAEncryptionUsecase {
	return &RSAEncryptionUsecase{}
}

func (u *RSAEncryptionUsecase) DoRSAEncryptionSHA256(traceId string, message model.Message) model.Response {
	request := model.RequestRSAEncryption{}
	err := json.Unmarshal(message.Body, &request)
	if err != nil {
		return model.FResponseDefault(http.StatusBadRequest, "Invalid Request")
	}

	_, privateKey, err := utils.GetRSAKey([]byte(request.PublicKey), []byte(request.PrivateKey))
	if err != nil {
		return model.FResponseDefault(http.StatusBadRequest, utils.ToTitleCase(err.Error()))
	}

	hashed := sha256.Sum256([]byte(request.Message))
	dataInBytes, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return model.FResponseDefault(http.StatusInternalServerError, "Server Error")
	}

	data := ""
	switch request.OutputEncoding {
	case "HEX":
		data = hex.EncodeToString(dataInBytes)
	case "BASE64":
		data = base64.StdEncoding.EncodeToString(dataInBytes)
	}
	return model.FResponseData(http.StatusOK, "Success", data)
}
