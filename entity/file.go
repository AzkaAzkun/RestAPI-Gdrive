package entity

import (
	"time"
)

type FileEntity struct {
	ID       string `gorm:"primary;keycolumn:id;"`
	Filename string `gorm:"filename"`
	Link     string `gorm:"Link"`

	CreatedAt time.Time `gorm:"column:created_at"`
	CreatedBy string    `gorm:"column:created_by"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	UpdatedBy string    `gorm:"column:updated_by"`
}

func (ff *FileEntity) TableName() string {
	return "file_entities"
}
