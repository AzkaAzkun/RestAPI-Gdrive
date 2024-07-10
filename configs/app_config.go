package configs

import (
	"github.com/gofiber/fiber"
	"google.golang.org/api/drive/v3"
	"gorm.io/gorm"
)

type AppConfig struct {
	db    *gorm.DB
	app   *fiber.App
	drive *drive.Service
}

func NewApp(db *gorm.DB, app *fiber.App, drive *drive.Service) *AppConfig {
	return nil
}

func (ap *AppConfig) Run() {

}
