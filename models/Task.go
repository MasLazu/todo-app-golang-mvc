package models

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	ID        uint           `gorm:"primaryKey" json:"id" query:"id"`
	Data      string         `json:"data" query:"data"`
	Done      bool           `json:"done" query:"done"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
