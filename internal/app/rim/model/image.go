package model

import "gorm.io/gorm"

// Image Struct
type Image struct {
	// gorm.Model
	Name   string `json:"name"`
	FileID string
	// Tags   []*Tag `gorm:"many2many:image_tags;"`
}

// Tag Struct
type Tag struct {
	gorm.Model
	Label string `json:"label"`
	Color string `json:"color"`
}
