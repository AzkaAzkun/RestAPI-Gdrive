package service

import (
	"Upload-files-to-Google-Drive-simply-using-Golang/api/repository"
	"Upload-files-to-Google-Drive-simply-using-Golang/domain"
	"Upload-files-to-Google-Drive-simply-using-Golang/entity"
	"Upload-files-to-Google-Drive-simply-using-Golang/store"
	"fmt"

	"bytes"
	"io"
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
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

func (fs *FileService) GetsFile() ([]entity.FileEntity, error) {
	data, err := fs.fileRepository.GetsFile()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (fs *FileService) GetFileById(id string) (entity.FileEntity, error) {
	data, err := fs.fileRepository.GetsFileById(id)
	if err != nil {
		return data, err
	}
	return data, nil
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

	fileId, err := store.UploadFile(fs.driveService, r.Filename, readerfile)
	if err != nil {
		return data, err
	}

	link := fmt.Sprintf("https://drive.google.com/file/d/%s/view?usp=sharing", fileId)

	data = domain.FileDTO{
		ID:       fileId,
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

func (fs *FileService) UpdateFile(f *multipart.FileHeader, r domain.FileDTO, id string) (domain.FileDTO, error) {
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

	if err := store.UpdateFile(fs.driveService, r.Filename, readerfile, id); err != nil {
		return data, err
	}

	link := fmt.Sprintf("https://drive.google.com/file/d/%s/view?usp=sharing", id)

	data = domain.FileDTO{
		ID:       id,
		Filename: r.Filename,
		Link:     link,
	}

	model := entity.FileEntity{
		ID:       data.ID,
		Filename: data.Filename,
		Link:     data.Link,
	}

	if err := fs.fileRepository.UpdateFile(model, id); err != nil {
		return data, err
	}

	return data, nil
}

func (fs *FileService) DeleteFile(id string) error {
	if err := store.DeleteFile(fs.driveService, id); err != nil {
		return err
	}
	if err := fs.fileRepository.DeleteFile(id); err != nil {
		return err
	}
	return nil
}
