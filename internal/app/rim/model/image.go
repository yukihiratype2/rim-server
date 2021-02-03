package model

import "gorm.io/gorm"

// Image Struct
type Image struct {
	gorm.Model
	Name   string `json:"name"`
	FileID string
	Tags   []*Tag `gorm:"many2many:image_tags;" json:"tag"`
}

// Tag Struct
type Tag struct {
	ID     uint     `gorm:"primarykey" json:"id"`
	Label  string   `json:"label"`
	Color  string   `json:"color"`
	Images []*Image `gorm:"many2many:image_tags;" json:"images"`
	// gorm.Model
}
