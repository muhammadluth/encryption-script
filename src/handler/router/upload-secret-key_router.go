package router

import (
	"encryption-script/app/utils"
	"encryption-script/model"
	"encryption-script/model/constant"
	"encryption-script/src"
	"encryption-script/src/handler"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UploadSecretKeyRouter struct {
	iUploadSecretKeyUsecase src.IUploadSecretKeyUsecase
}

func NewUploadSecretKeyRouter(iUploadSecretKeyUsecase src.IUploadSecretKeyUsecase) handler.IUploadSecretKeyRouter {
	return &UploadSecretKeyRouter{iUploadSecretKeyUsecase}
}

func (r *UploadSecretKeyRouter) DoUploadSecretKey(ctx *fiber.Ctx) error {
	traceId, _ := ctx.Locals(constant.LOCALS_TRACE_ID).(string)

	// Body Parser
	secretKeyFile, _ := ctx.FormFile("secret_key_file")

	message, err := utils.GenerateMessage(traceId, ctx, nil, nil, nil, nil)
	if err != nil {
		response := model.FResponseDefault(http.StatusBadRequest, "Invalid Request")
		return ctx.Status(response.Status).JSON(response.Body)
	}
	response := r.iUploadSecretKeyUsecase.DoUploadSecretKey(traceId, secretKeyFile, message)
	return ctx.Status(response.Status).JSON(response.Body)
}
