package server

import (
	"encryption-script/model"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type SetupServer struct {
	fiberApp      *fiber.App
	svcProperties model.ServiceProperties
}

func NewSetupServer(svcProperties model.ServiceProperties) SetupServer {
	fiberApp := fiber.New(
		fiber.Config{
			CaseSensitive: true,
			Concurrency:   svcProperties.ServicePoolSizeConnection,
		},
	)
	return SetupServer{fiberApp, svcProperties}
}

func (s *SetupServer) InitServerConfiguration() (*fiber.App, fiber.Router) {
	s.fiberApp.Use(etag.New())
	s.fiberApp.Use(compress.New())
	s.fiberApp.Use(requestid.New())
	s.fiberApp.Use(recover.New())
	s.fiberApp.Use(cors.New(cors.Config{
		Next:         cors.ConfigDefault.Next,
		AllowOrigins: strings.Join([]string{"http://localhost:3000"}, ","),
		AllowMethods: fmt.Sprintf("%s, %s, %s, %s, %s",
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodPut,
			fiber.MethodPatch,
			fiber.MethodDelete),
		AllowHeaders:     strings.Join([]string{"Content-Type", "Accept", "Origin"}, ","),
		AllowCredentials: true,
		ExposeHeaders:    strings.Join([]string{"Content-Type", "Accept", "Content-Length"}, ","),
	}))
	// init api version 1
	routerAPIv1 := s.fiberApp.Group("api/v1", s.InitServerMiddleware())
	return s.fiberApp, routerAPIv1
}

func (s *SetupServer) InitServer() {
	// HEALTH CHECK
	s.fiberApp.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"message": "Hello, Welcome to My API!"})
	})
	svcPort := s.svcProperties.ServicePort
	s.fiberApp.Listen(fmt.Sprintf(":%d", svcPort))
}
