package entities

import (
	"time"

	"gorm.io/gorm"
)

type ApiResponse struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

type Base struct {
	ID        string         `json:"id" gorm:"primarykey;type:uuid;"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}

type DeleteRequest struct {
	ID string `json:"id" binding:"required"`
}
