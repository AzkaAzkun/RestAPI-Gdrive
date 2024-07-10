package main

import (
	"Upload-files-to-Google-Drive-simply-using-Golang/configs"
	"Upload-files-to-Google-Drive-simply-using-Golang/database"
	"Upload-files-to-Google-Drive-simply-using-Golang/drive"
	"os"

	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	env := os.Getenv("ENV")
	if err != nil || env == "" {
		panic("Failed to load env file")
	}

	db := database.NewDatabase()

	drive, err := drive.NewDriveService()
	if err != nil {
		panic(err)
	}

	fiber := fiber.New()
	app := configs.NewApp(db, fiber, drive)

	app.Run()
}
