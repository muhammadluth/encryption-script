package handler

import (
	"encryption-script/src"

	"github.com/gofiber/fiber/v2"
)

type EncryptionHttpHandler struct {
	fiberRouter            fiber.Router
	iRSAEncryptionRouter   IRSAEncryptionRouter
	iHMACAESRouter         IHMACAESRouter
	iUploadSecretKeyRouter IUploadSecretKeyRouter
}

func NewEncryptionHttpHandler(fiberRouter fiber.Router, iRSAEncryptionRouter IRSAEncryptionRouter,
	iHMACAESRouter IHMACAESRouter, iUploadSecretKeyRouter IUploadSecretKeyRouter) src.IEncryptionHttpHandler {
	return &EncryptionHttpHandler{fiberRouter, iRSAEncryptionRouter, iHMACAESRouter,
		iUploadSecretKeyRouter}
}

func (h *EncryptionHttpHandler) Routers() {
	h.fiberRouter.Post("/rsa-encryption/sha-256", h.iRSAEncryptionRouter.DoRSAEncryptionSHA256)
	h.fiberRouter.Post("/rsa-encryption/md5", h.iRSAEncryptionRouter.DoRSAEncryptionMD5)

	h.fiberRouter.Post("/hmac-encryption/aes", h.iHMACAESRouter.DoHMACEncryptionAES)
	h.fiberRouter.Post("/hmac-decryption/aes", h.iHMACAESRouter.DoHMACDecryptionAES)

	h.fiberRouter.Post("/upload/secret-file", h.iUploadSecretKeyRouter.DoUploadSecretKey)
}
