package service

import (
	"Upload-files-to-Google-Drive-simply-using-Golang/api/repository"
	"Upload-files-to-Google-Drive-simply-using-Golang/domain"
	"Upload-files-to-Google-Drive-simply-using-Golang/entity"
	"Upload-files-to-Google-Drive-simply-using-Golang/store"

	"bytes"
	"fmt"
	"io"
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"google.golang.org/api/drive/v3"
)

type FileService struct {
	app            *fiber.App
	fileRepository *repository.FileRepository
	driveService   *drive.Service
}

func NewFileService(app *fiber.App, fileRepository *repository.FileRepository, driveService *drive.Service) *FileService {
	return &FileService{
		app:            app,
		fileRepository: fileRepository,
		driveService:   driveService,
	}
}

func (fs *FileService) GetsFile() ([]domain.FileDTO, error) {

	return []domain.FileDTO{}, nil
}

func (fs *FileService) GetFileById(id string) (domain.FileDTO, error) {
	return domain.FileDTO{}, nil
}

func (fs *FileService) SaveFile(f *multipart.FileHeader, r domain.FileDTO) (domain.FileDTO, error) {
	var data domain.FileDTO
	file, err := f.Open()
	if err != nil {
		return data, err
	}

	defer file.Close()

	filebytes, err := io.ReadAll(file)
	if err != nil {
		return data, err
	}

	readerfile := bytes.NewReader(filebytes)

	link, err := store.UploadFile(fs.driveService, r.Filename, readerfile)
	if err != nil {
		return data, err
	}

	fmt.Println(link)

	data = domain.FileDTO{
		ID:       uuid.New(),
		Filename: r.Filename,
		Link:     link,
	}

	model := entity.FileEntity{
		ID:       data.ID,
		Filename: data.Filename,
		Link:     data.Link,
	}

	if err := fs.fileRepository.SaveFile(model); err != nil {
		return data, err
	}

	return data, nil
}
