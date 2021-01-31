package models

import (
	"time"

	"gorm.io/gorm"
)

// BaseModel - Holds basic fields
type BaseModel struct {
	// ID        uint      `gorm:"primarykey", json: "id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	// gorm.Model
	IsDelete bool `json:"is_delete"`
	IsActive bool `json:"is_active"`
}
