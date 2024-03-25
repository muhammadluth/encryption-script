package handler

import (
	"encryption-script/src"

	"github.com/gofiber/fiber/v2"
)

type EncryptionHttpHandler struct {
	fiberRouter          fiber.Router
	iRSAEncryptionRouter IRSAEncryptionRouter
}

func NewEncryptionHttpHandler(fiberRouter fiber.Router, iRSAEncryptionRouter IRSAEncryptionRouter) src.IEncryptionHttpHandler {
	return &EncryptionHttpHandler{fiberRouter, iRSAEncryptionRouter}
}

func (h *EncryptionHttpHandler) Routers() {
	h.fiberRouter.Post("/rsa-encryption/sha-256", h.iRSAEncryptionRouter.DoRSAEncryptionSHA256)
}
