package src

import "encryption-script/model"

type IRSAEncryptionUsecase interface {
	DoRSAEncryptionSHA256(traceId string, message model.Message) model.Response
}
