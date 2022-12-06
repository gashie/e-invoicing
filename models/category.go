package models

import "gorm.io/gorm"

type Category struct {
	Title        string `json:"title" gorm:"unique"`
	Description string `json:"description"`
	gorm.Model  
}
type CategoryRequest struct {
	Title        string `json:"title" gorm:"unique"`
}