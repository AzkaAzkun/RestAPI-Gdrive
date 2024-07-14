package main

import (
	"Upload-files-to-Google-Drive-simply-using-Golang/configs"
	"Upload-files-to-Google-Drive-simply-using-Golang/store"

	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	env := os.Getenv("ENV")
	if err != nil && env == "" {
		panic("Failed to load env file")
	}

	db := configs.NewDatabase()

	drive, err := store.NewDrive()
	if err != nil {
		fmt.Printf("Error Connect to drive: ")
		panic(err)
	}

	fiber := fiber.New()
	app := configs.NewApp(db, fiber, drive.DriveService)

	app.Run()
}
