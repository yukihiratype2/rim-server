package model

import "gorm.io/gorm"

// Image Struct
type Image struct {
	gorm.Model
	Name   string `json:"name"`
	FileID string `json:"fileId"`
	URL    string `gorm:"-" json:"url"`
	Tags   []*Tag `gorm:"many2many:image_tags;" json:"tag"`
}

func (i *Image) Create() (err error) {
	return db.Create(i).Error
}

func (i *Image) First() (err error) {
	return db.Preload("Tags").First(i).Error
}

func Find(images *[]Image) (err error) {
	return db.Preload("Tags").Find(images).Error
}
