package model

import (
	"time"

	"gorm.io/gorm"
)

type Column struct {
	ID        uint64         `gorm:"primaryKey;autoIncrement:false"`
	Name      string         `gorm:"not null" json:"name"`
	Order     uint           `json:"order"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
