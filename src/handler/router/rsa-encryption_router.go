package router

import (
	"encoding/json"
	"encryption-script/app/utils"
	"encryption-script/model"
	"encryption-script/model/constant"
	"encryption-script/src"
	"encryption-script/src/handler"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type RSAEncryptionRouter struct {
	iRSAEncryptionUsecase src.IRSAEncryptionUsecase
}

func NewRSAEncryptionRouter(iRSAEncryptionUsecase src.IRSAEncryptionUsecase) handler.IRSAEncryptionRouter {
	return &RSAEncryptionRouter{iRSAEncryptionUsecase}
}

func (r *RSAEncryptionRouter) DoRSAEncryptionSHA256(ctx *fiber.Ctx) error {
	traceId, _ := ctx.Locals(constant.LOCALS_TRACE_ID).(string)

	// Body Parser
	request := new(model.RequestRSAEncryption)
	if err := ctx.BodyParser(request); err != nil {
		response := model.FResponseDefault(http.StatusBadRequest, "Invalid Request")
		return ctx.Status(response.Status).JSON(response.Body)
	}
	requestAsByteJSON, _ := json.Marshal(request)

	message, err := utils.GenerateMessage(traceId, ctx, nil, nil, nil, requestAsByteJSON)
	if err != nil {
		response := model.FResponseDefault(http.StatusBadRequest, "Invalid Request")
		return ctx.Status(response.Status).JSON(response.Body)
	}
	response := r.iRSAEncryptionUsecase.DoRSAEncryptionSHA256(traceId, message)
	return ctx.Status(response.Status).JSON(response.Body)
}

func (r *RSAEncryptionRouter) DoRSAEncryptionMD5(ctx *fiber.Ctx) error {
	traceId, _ := ctx.Locals(constant.LOCALS_TRACE_ID).(string)

	// Body Parser
	request := new(model.RequestRSAEncryption)
	if err := ctx.BodyParser(request); err != nil {
		response := model.FResponseDefault(http.StatusBadRequest, "Invalid Request")
		return ctx.Status(response.Status).JSON(response.Body)
	}
	requestAsByteJSON, _ := json.Marshal(request)

	message, err := utils.GenerateMessage(traceId, ctx, nil, nil, nil, requestAsByteJSON)
	if err != nil {
		response := model.FResponseDefault(http.StatusBadRequest, "Invalid Request")
		return ctx.Status(response.Status).JSON(response.Body)
	}
	response := r.iRSAEncryptionUsecase.DoRSAEncryptionMD5(traceId, message)
	return ctx.Status(response.Status).JSON(response.Body)
}
