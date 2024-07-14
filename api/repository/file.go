package repository

import (
	"Upload-files-to-Google-Drive-simply-using-Golang/entity"

	"gorm.io/gorm"
)

type FileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) *FileRepository {
	return &FileRepository{db}
}

func (f *FileRepository) SaveFile(data entity.FileEntity) error {
	if err := f.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
