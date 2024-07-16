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

func (f *FileRepository) GetsFile() ([]entity.FileEntity, error) {
	var data []entity.FileEntity
	if err := f.db.Model(entity.FileEntity{}).Find(&data); err.Error != nil {
		return data, err.Error
	}
	return data, nil
}

func (f *FileRepository) GetsFileById(id string) (entity.FileEntity, error) {
	var data entity.FileEntity
	if err := f.db.Model(entity.FileEntity{}).Where("id = ?", id).Find(&data); err.Error != nil {
		return data, err.Error
	}
	return data, nil
}

func (f *FileRepository) SaveFile(data entity.FileEntity) error {
	if err := f.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}

func (f *FileRepository) UpdateFile(data entity.FileEntity, id string) error {
	if err := f.db.Model(entity.FileEntity{}).Where("id = ?", id).Updates(&data).Error; err != nil {
		return err
	}
	return nil
}

func (f *FileRepository) DeleteFile(id string) error {
	var model entity.FileEntity
	if err := f.db.Where("id = ?", id).Delete(&model).Error; err != nil {
		return err
	}
	return nil
}
