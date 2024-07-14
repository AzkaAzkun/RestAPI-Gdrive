package configs

import (
	"Upload-files-to-Google-Drive-simply-using-Golang/api/handlers"
	"Upload-files-to-Google-Drive-simply-using-Golang/api/repository"
	"Upload-files-to-Google-Drive-simply-using-Golang/api/service"

	"github.com/gofiber/fiber/v2"

	"google.golang.org/api/drive/v3"
	"gorm.io/gorm"
)

type AppConfig struct {
	db    *gorm.DB
	app   *fiber.App
	drive *drive.Service
}

func NewApp(db *gorm.DB, app *fiber.App, drive *drive.Service) *AppConfig {
	ap := &AppConfig{db: db, app: app, drive: drive}

	fileRepository := repository.NewFileRepository(db)
	fileService := service.NewFileService(app, fileRepository, drive)
	handlers.NewFileHandlers(app, fileService)

	return ap
}

func (ap *AppConfig) Run() {
	ap.app.Listen(":8080")
}
