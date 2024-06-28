package src

import "encryption-script/model"

type IRSAEncryptionUsecase interface {
	DoRSAEncryptionSHA256(traceId string, message model.Message) model.Response
}

type IHMACAESUsecase interface {
	DoHMACEncryptionAES(traceId string, message model.Message) model.Response
	DoHMACDecryptionAES(traceId string, message model.Message) model.Response
}
