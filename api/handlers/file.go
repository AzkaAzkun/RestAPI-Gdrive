package handlers

import (
	"Upload-files-to-Google-Drive-simply-using-Golang/api/service"
	"Upload-files-to-Google-Drive-simply-using-Golang/domain"
	"Upload-files-to-Google-Drive-simply-using-Golang/helper"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type FileHandlers struct {
	app         *fiber.App
	fileservice *service.FileService
}

func NewFileHandlers(app *fiber.App, fileservice *service.FileService) {
	FileHandlers := &FileHandlers{
		app:         app,
		fileservice: fileservice,
	}

	FileHandlers.Route()
}

func (fh *FileHandlers) GetsFile(c *fiber.Ctx) error {
	return nil
}

func (fh *FileHandlers) GetFileById(c *fiber.Ctx) error {
	return nil
}

func (fh *FileHandlers) SaveFile(c *fiber.Ctx) error {
	var req domain.FileDTO
	if err := c.BodyParser(&req); err != nil {
		return helper.SendResponse(c, domain.ServiceResponse{
			Code:    fiber.ErrBadRequest.Code,
			Success: false,
			Error:   "Failed to parser body request: " + err.Error(),
			Data:    nil,
		})
	}

	file, err := c.FormFile("file")
	if err != nil {
		return helper.SendResponse(c, domain.ServiceResponse{
			Code:    fiber.ErrBadRequest.Code,
			Success: false,
			Error:   "Failed to get file form: " + err.Error(),
			Data:    nil,
		})
	}
	fmt.Println(file.Filename)
	payload, err := fh.fileservice.SaveFile(file, req)

	if err != nil {
		return helper.SendResponse(c, domain.ServiceResponse{
			Code:    fiber.ErrBadRequest.Code,
			Success: false,
			Error:   "Failed to save file: " + err.Error(),
			Data:    nil,
		})
	}

	return helper.SendResponse(c, domain.ServiceResponse{
		Code:    fiber.StatusAccepted,
		Success: true,
		Error:   "",
		Data:    payload,
	})
}

func (fh *FileHandlers) UpdateFile(c *fiber.Ctx) error {
	return nil
}

func (fh *FileHandlers) RemoveFile(c *fiber.Ctx) error {
	return nil
}

func (fh *FileHandlers) Route() {
	fh.app.Route("/api", func(router fiber.Router) {
		router.Post("/v1/file", fh.SaveFile)
	})
}
