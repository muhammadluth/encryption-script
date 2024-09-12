package src

import (
	"encryption-script/model"
	"mime/multipart"
)

type IRSAEncryptionUsecase interface {
	DoRSAEncryptionSHA256(traceId string, message model.Message) model.Response
	DoRSAEncryptionMD5(traceId string, message model.Message) model.Response
}

type IHMACAESUsecase interface {
	DoHMACEncryptionAES(traceId string, message model.Message) model.Response
	DoHMACDecryptionAES(traceId string, message model.Message) model.Response
}

type IUploadSecretKeyUsecase interface {
	DoUploadSecretKey(traceId string, secretKeyFile *multipart.FileHeader, message model.Message) model.Response
}
