package helper

import (
	"Upload-files-to-Google-Drive-simply-using-Golang/domain"

	"github.com/gofiber/fiber/v2"
)

func SendResponse(ctx *fiber.Ctx, data domain.ServiceResponse) error {
	res := &domain.Response{
		Success: data.Success,
		Error:   data.Error,
		Data:    data.Data,
	}

	return ctx.Status(data.Code).JSON(res)
}
