package models

import (
	"time"
	"gorm.io/gorm"
)

// Difficulty Model
type Difficulty struct {
	ID                uint      `gorm:"primaryKey" json:"id"`
	Name              string    `gorm:"type:varchar(255);not null" json:"name"`
	DescriptionShort  string    `gorm:"type:varchar(100)" json:"description_short"`
	DescriptionLong   string    `gorm:"type:varchar(2000)" json:"description_long"`
	TotalCategories   int       `json:"total_categories"`
	Status            string    `gorm:"type:varchar(10);check:status IN ('active', 'pending', 'archived');default:'pending'" json:"status"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

// BeforeCreate Hook untuk menangani created_at
func (d *Difficulty) BeforeCreate(tx *gorm.DB) error {
	d.CreatedAt = time.Now()
	return nil
}

// BeforeUpdate Hook untuk menangani updated_at
func (d *Difficulty) BeforeUpdate(tx *gorm.DB) error {
	d.UpdatedAt = time.Now()
	return nil
}
