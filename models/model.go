package models

import (
	"time"

	"gorm.io/gorm"
)

// Model base model
type Model struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// Modeler model interface
type Modeler interface {
	GetID() uint
}

// GetID get id
func (m Model) GetID() uint {
	return m.ID
}
