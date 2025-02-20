package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// Subcategory Model
type Subcategory struct {
	ID                            uint           `gorm:"primaryKey" json:"id"`
	Name                          string         `gorm:"type:varchar(255);not null" json:"name"`
	Status                        string         `gorm:"type:varchar(10);default:'pending'" json:"status"`
	DescriptionLong               string         `gorm:"type:varchar(2000)" json:"description_long"`
	GreatTotalThemesOrLevels      int            `json:"great_total_themes_or_levels"`
	TotalThemesOrLevels           int            `json:"total_themes_or_levels"`
	CompletedTotalThemesOrLevels  int            `json:"completed_total_themes_or_levels"`
	UpdateNews                    json.RawMessage       `gorm:"type:jsonb" json:"update_news"`
	ImageURL                      string         `gorm:"type:varchar(100)" json:"image_url"`
	CreatedAt                     time.Time      `json:"created_at"`
	UpdatedAt                     time.Time      `json:"updated_at"`
	DeletedAt                     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CategoryID uint `gorm:"column:categories_id;not null" json:"categories_id"`
}


func (d *Subcategory) BeforeCreate(tx *gorm.DB) error {
	d.CreatedAt = time.Now()
	return nil
}