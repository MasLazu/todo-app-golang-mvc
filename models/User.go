package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey" json:"id" query:"id"`
	Name      string `json:"name" query:"name"`
	Email     string `json:"email" query:"email"`
	Password  string `json:"password" query:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
