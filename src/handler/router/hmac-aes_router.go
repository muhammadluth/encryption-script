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

type HMACAESRouter struct {
	iHMACAESUsecase src.IHMACAESUsecase
}

func NewHMACAESRouter(iHMACAESUsecase src.IHMACAESUsecase) handler.IHMACAESRouter {
	return &HMACAESRouter{iHMACAESUsecase}
}

func (r *HMACAESRouter) DoHMACEncryptionAES(ctx *fiber.Ctx) error {
	traceId, _ := ctx.Locals(constant.LOCALS_TRACE_ID).(string)

	// Body Parser
	request := new(model.RequestHMACEncryptionAES)
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
	response := r.iHMACAESUsecase.DoHMACEncryptionAES(traceId, message)
	return ctx.Status(response.Status).JSON(response.Body)
}

func (r *HMACAESRouter) DoHMACDecryptionAES(ctx *fiber.Ctx) error {
	traceId, _ := ctx.Locals(constant.LOCALS_TRACE_ID).(string)

	// Body Parser
	request := new(model.RequestHMACDecryptionAES)
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
	response := r.iHMACAESUsecase.DoHMACDecryptionAES(traceId, message)
	return ctx.Status(response.Status).JSON(response.Body)
}
