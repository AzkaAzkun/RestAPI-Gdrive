package store

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

type Drive struct {
	DriveService *drive.Service
}

func NewDrive() (*Drive, error) {
	driveCredentialsFile := map[string]interface{}{
		"type":                        os.Getenv("DRIVE_TYPE"),
		"project_id":                  os.Getenv("DRIVE_PROJECT_ID"),
		"private_key_id":              os.Getenv("DRIVE_PRIVATE_KEY_ID"),
		"private_key":                 strings.Replace(os.Getenv("DRIVE_PRIVATE_KEY"), "/\\n/gm", "\n", -1),
		"client_email":                os.Getenv("DRIVE_CLIENT_EMAIL"),
		"client_id":                   os.Getenv("DRIVE_CLIENT_ID"),
		"auth_uri":                    os.Getenv("DRIVE_AUTH_URI"),
		"token_uri":                   os.Getenv("DRIVE_TOKEN_URI"),
		"auth_provider_x509_cert_url": os.Getenv("DRIVE_AUTH_PROVIDER_x509_CERT_URL"),
		"client_x509_cert_url":        os.Getenv("DRIVE_CLIENT_x509_CERT_URL"),
		"universe_domain":             os.Getenv("DRIVE_UNIVERSE_DOMAIN"),
	}

	credentials, err := json.Marshal(driveCredentialsFile)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	opt := option.WithCredentialsJSON(credentials)
	// opt := option.WithCredentialsFile("key.json")

	DriveService, err := drive.NewService(ctx, opt)
	if err != nil {
		return nil, err
	}
	return &Drive{DriveService: DriveService}, nil
}

func UploadFile(service *drive.Service, filename string, content io.Reader) (string, error) {
	folderId := os.Getenv("FOLDER_ID")

	f := &drive.File{
		MimeType: "application/octet-stream",
		Name:     filename,
		Parents:  []string{folderId},
	}

	file, err := service.Files.Create(f).Media(content).Do()

	if err != nil {
		return "", err
	}
	fmt.Printf("Haloooo\n")

	link := fmt.Sprintf("https://drive.google.com/file/d/%s/view?usp=sharing", file.Id)
	fmt.Println(filename + file.Id)
	return link, nil
}
