package usecase

import (
	"crypto/aes"
	"encoding/json"
	"encryption-script/app/utils"
	"encryption-script/model"
	"encryption-script/src"
	"net/http"
	"strings"

	aes_ecb "encryption-script/src/lib/aes-ecb"
)

type HMACAESUsecase struct {
}

func NewHMACAESUsecase() src.IHMACAESUsecase {
	return &HMACAESUsecase{}
}

func (u *HMACAESUsecase) DoHMACEncryptionAES(traceId string, message model.Message) model.Response {
	request := model.RequestHMACEncryptionAES{}
	err := json.Unmarshal(message.Body, &request)
	if err != nil {
		return model.FResponseDefault(http.StatusBadRequest, "Invalid Request")
	}

	switch strings.ToUpper(request.Type) {
	case "AES-ECB":
		return u.doEncryptionAESECB(traceId, request.SecretKey, request.Message, request.OutputEncoding)
	default:
		return model.FResponseDefault(http.StatusInternalServerError, "Invalid Type Request")
	}
}

func (u *HMACAESUsecase) DoHMACDecryptionAES(traceId string, message model.Message) model.Response {
	request := model.RequestHMACDecryptionAES{}
	err := json.Unmarshal(message.Body, &request)
	if err != nil {
		return model.FResponseDefault(http.StatusBadRequest, "Invalid Request")
	}

	switch strings.ToUpper(request.Type) {
	case "AES-ECB":
		return u.doDecryptionAESECB(traceId, request.SecretKey, request.Message, request.InputEncoding)
	default:
		return model.FResponseDefault(http.StatusInternalServerError, "Invalid Type Request")
	}
}

func (u *HMACAESUsecase) doEncryptionAESECB(traceId, key, message, outputEncoding string) model.Response {
	keyAsByte := []byte(key)
	msgAsByte := []byte(message)

	cipher, err := aes.NewCipher(keyAsByte)
	if err != nil {
		return model.FResponseDefault(http.StatusInternalServerError, "Server Error")
	}
	mode := aes_ecb.NewECBEncrypter(cipher)
	padder := aes_ecb.NewPkcs7Padding(mode.BlockSize())
	msgAsByte, err = padder.Pad(msgAsByte)
	if err != nil {
		return model.FResponseDefault(http.StatusInternalServerError, "Server Error")
	}
	encrypted := make([]byte, len(msgAsByte))
	mode.CryptBlocks(encrypted, msgAsByte)
	data := utils.SetEncoding(outputEncoding, encrypted)
	return model.FResponseData(http.StatusOK, "Success", data)
}

func (u *HMACAESUsecase) doDecryptionAESECB(traceId, key, message, inputEncoding string) model.Response {
	keyAsByte := []byte(key)
	msgAsByte, err := utils.SetDecoding(inputEncoding, message)
	if err != nil {
		return model.FResponseDefault(http.StatusInternalServerError, "Server Error")
	}

	cipher, err := aes.NewCipher(keyAsByte)
	if err != nil {
		return model.FResponseDefault(http.StatusInternalServerError, "Server Error")
	}
	mode := aes_ecb.NewECBDecrypter(cipher)

	decrypted := make([]byte, len(msgAsByte))
	mode.CryptBlocks(decrypted, msgAsByte)
	padder := aes_ecb.NewPkcs7Padding(mode.BlockSize())
	decrypted, err = padder.Unpad(decrypted) // unpad plaintext after decryption
	if err != nil {
		return model.FResponseDefault(http.StatusInternalServerError, "Server Error")
	}
	data := string(decrypted)
	return model.FResponseData(http.StatusOK, "Success", data)
}
