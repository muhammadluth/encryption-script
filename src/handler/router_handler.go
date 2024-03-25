package handler

import "github.com/gofiber/fiber/v2"

type IRSAEncryptionRouter interface {
	DoRSAEncryptionSHA256(ctx *fiber.Ctx) error
}
