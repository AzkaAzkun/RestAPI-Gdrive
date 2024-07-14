package entity

import (
	"time"

	"github.com/google/uuid"
)

type FileEntity struct {
	ID       uuid.UUID `gorm:"primary;keycolumn:id;type:uuid;"`
	Filename string    `gorm:"filename"`
	Link     string    `gorm:"Link"`

	CreatedAt time.Time `gorm:"column:created_at"`
	CreatedBy string    `gorm:"column:created_by"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	UpdatedBy string    `gorm:"column:updated_by"`
}

func (ff *FileEntity) TableName() string {
	return "file_entities"
}
