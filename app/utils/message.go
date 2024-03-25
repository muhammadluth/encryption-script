package utils

import (
	"encryption-script/model"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GenerateMessage(traceId string, ctx *fiber.Ctx, header http.Header, param, query, rawReqBody []byte) (message model.Message, err error) {
	message = model.Message{
		Header: model.Header{
			URL:    string(ctx.BaseURL() + ctx.OriginalURL()),
			Query:  string(ctx.Context().QueryArgs().QueryString()),
			Header: header,
		},
		Body:  rawReqBody,
		Param: param,
		Query: query,
	}
	return message, err
}
