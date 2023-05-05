package models

import (
	"time"
)

type Todo struct {
	ID          uint      `json:"id" gorm:"primarykey"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description" gorm:"not null;default:''"`
	Completed   bool      `json:"completed" gorm:"not null;default:false"`
	CreatedAt   time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
