package handlers

import (
	"Upload-files-to-Google-Drive-simply-using-Golang/api/service"
	"Upload-files-to-Google-Drive-simply-using-Golang/domain"
	"Upload-files-to-Google-Drive-simply-using-Golang/helper"

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
	data, err := fh.fileservice.GetsFile()
	if err != nil {
		return helper.SendResponse(c, domain.ServiceResponse{
			Code:    fiber.ErrBadRequest.Code,
			Success: false,
			Error:   "Failed to Gets file: " + err.Error(),
			Data:    nil,
		})
	}
	return helper.SendResponse(c, domain.ServiceResponse{
		Code:    fiber.StatusAccepted,
		Success: true,
		Error:   "",
		Data:    data,
	})
}

func (fh *FileHandlers) GetFileById(c *fiber.Ctx) error {
	idParams := c.Params("id")
	if idParams == "" {
		return helper.SendResponse(c, domain.ServiceResponse{
			Code:    fiber.ErrBadRequest.Code,
			Success: false,
			Error:   "Bad request, Id is requred",
			Data:    nil,
		})
	}
	data, err := fh.fileservice.GetFileById(idParams)
	if err != nil {
		return helper.SendResponse(c, domain.ServiceResponse{
			Code:    fiber.ErrBadRequest.Code,
			Success: false,
			Error:   "Failed to Get file: " + err.Error(),
			Data:    nil,
		})
	}
	return helper.SendResponse(c, domain.ServiceResponse{
		Code:    fiber.StatusAccepted,
		Success: true,
		Error:   "",
		Data:    data,
	})
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

	idParams := c.Params("id")
	if idParams == "" {
		return helper.SendResponse(c, domain.ServiceResponse{
			Code:    fiber.ErrBadRequest.Code,
			Success: false,
			Error:   "Bad request, Id is requred",
			Data:    nil,
		})
	}
	payload, err := fh.fileservice.UpdateFile(file, req, idParams)

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

func (fh *FileHandlers) DeleteFile(c *fiber.Ctx) error {
	idParams := c.Params("id")
	if idParams == "" {
		return helper.SendResponse(c, domain.ServiceResponse{
			Code:    fiber.ErrBadRequest.Code,
			Success: false,
			Error:   "Bad request, Id is requred",
			Data:    nil,
		})
	}
	if err := fh.fileservice.DeleteFile(idParams); err != nil {
		return helper.SendResponse(c, domain.ServiceResponse{
			Code:    fiber.ErrBadRequest.Code,
			Success: false,
			Error:   "Failed to delete file: " + err.Error(),
			Data:    nil,
		})
	}
	return helper.SendResponse(c, domain.ServiceResponse{
		Code:    fiber.StatusAccepted,
		Success: true,
		Error:   "",
		Data:    "",
	})
}

func (fh *FileHandlers) Route() {
	fh.app.Route("/api/v1", func(router fiber.Router) {
		router.Get("/file", fh.GetsFile)
		router.Get("/file/:id", fh.GetFileById)

		router.Post("/file", fh.SaveFile)
		router.Put("/file/:id", fh.UpdateFile)
		router.Delete("/file/:id", fh.DeleteFile)
	})
}
