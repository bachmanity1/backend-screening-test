package model

import (
	"time"

	"gorm.io/gorm"
)

type Card struct {
	ID          uint64         `gorm:"primaryKey" json:"id"`
	ColumnID    uint64         `gorm:"index" json:"columnId"`
	Name        string         `gorm:"not null" json:"name"`
	Description string         `gorm:"not null" json:"description"`
	Order       uint           `json:"order"`
	Status      Status         `gorm:"default:1" json:"status"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type CardList []*Card

type Status byte

const (
	NotSelected = Status(iota + 1)
	InProgress
	Done
	Archived
)
