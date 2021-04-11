package model

import (
	"time"

	"gorm.io/gorm"
)

type Card struct {
	ColumnID    uint64         `gorm:"primaryKey;autoIncrement:false" json:"columnId"`
	ID          uint64         `gorm:"primaryKey;autoIncrement:false" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Description string         `gorm:"not null" json:"description"`
	Order       uint           `json:"order"`
	Status      Status         `gorm:"deafult:0" json:"status"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type Status int

const (
	NotSelected = Status(iota)
	InProgress
	Done
	Archived
)
