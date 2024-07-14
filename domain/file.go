package domain

import "github.com/google/uuid"

type FileDTO struct {
	ID       uuid.UUID `json:"id"`
	Filename string    `json:"filename"`
	Link     string    `json:"Link"`
}
