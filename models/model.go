package models

import (
	"time"

	"gorm.io/gorm"
)

type People struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Age       int            `json:"age"`
	CreatedAt time.Time      `json:"-"` // Omitir na resposta JSON
	UpdatedAt time.Time      `json:"-"` // Omitir na resposta JSON
	DeletedAt gorm.DeletedAt `json:"-"` // Omitir na resposta JSON
}
