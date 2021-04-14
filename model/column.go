package model

import (
	"time"

	"gorm.io/gorm"
)

type Column struct {
	ID        uint64         `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	Order     uint           `gorm:"index" json:"order"`
	Cards     CardList       `json:"cards"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type ColumnList []*Column

func (c *Column) Update(other *Column) {
	c.Name = other.Name
}
