package models

import (
	"time"
	"gorm.io/gorm"
)

type Category struct {
	ID               uint   `json:"id" gorm:"primaryKey"`
	Name             string `json:"name"`
	DescriptionShort string `json:"description_short"`
	DescriptionLong  string `json:"description_long"`
	TotalCategories  int    `json:"total_categories"`
	Status          string `json:"status" gorm:"default:pending"`
	CreatedAt        time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt        time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	DifficultyID     uint   `json:"difficulty_id"`
}


func (d *Category) BeforeCreate(tx *gorm.DB) error {
	d.CreatedAt = time.Now()
	return nil
}