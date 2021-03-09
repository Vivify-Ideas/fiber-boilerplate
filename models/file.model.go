package models

import "gorm.io/gorm"

// File model
type File struct {
	gorm.Model
	Path   string `json:"path"`
	UserId int
	User   User `gorm:"foreignKey:UserId"`
}
