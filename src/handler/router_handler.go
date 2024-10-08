package handler

import "github.com/gofiber/fiber/v2"

type IRSAEncryptionRouter interface {
	DoRSAEncryptionSHA256(ctx *fiber.Ctx) error
	DoRSAEncryptionMD5(ctx *fiber.Ctx) error
}

type IHMACAESRouter interface {
	DoHMACEncryptionAES(ctx *fiber.Ctx) error
	DoHMACDecryptionAES(ctx *fiber.Ctx) error
}

type IUploadSecretKeyRouter interface {
	DoUploadSecretKey(ctx *fiber.Ctx) error
}
