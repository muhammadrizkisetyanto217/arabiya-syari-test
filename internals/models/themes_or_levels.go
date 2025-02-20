package models

import (
	"time"
	"gorm.io/gorm"
	 "encoding/json"
)

// ThemeOrLevel Model
type ThemesOrLevel struct {
	ID                 uint           `gorm:"primaryKey" json:"id"`
	Name               string         `gorm:"type:varchar(255);not null" json:"name"`
	Status             string         `gorm:"type:varchar(10);default:'pending'" json:"status"`
	DescriptionShort   string         `gorm:"type:varchar(100)" json:"description_short"`
	DescriptionLong    string         `gorm:"type:varchar(2000)" json:"description_long"`
	GradeTotalUnit     int            `json:"grade_total_unit"`
	TotalUnit          int            `json:"total_unit"`
	CompletedTotalUnit int            `json:"completed_total_unit"`
	UpdateNews         json.RawMessage       `gorm:"type:jsonb" json:"update_news"`
	ImageURL           string         `gorm:"type:varchar(100)" json:"image_url"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	SubcategoryID uint `gorm:"column:subategories_id;not null" json:"subategories_id"`
	// Subcategory   Subcategory `gorm:"foreignKey:SubcategoryID;references:ID"`

}


func (d *ThemesOrLevel) BeforeCreate(tx *gorm.DB) error {
	d.CreatedAt = time.Now()
	return nil
}