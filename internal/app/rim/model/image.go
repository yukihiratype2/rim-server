package model

import "gorm.io/gorm"

// Image Struct
type Image struct {
	gorm.Model
	Name   string `json:"name" form:"name"`
	FileID string `json:"fileId"`
	// Folder   Folder `json:"folder"`
	Favorite        bool   `json:"favorite" form:"favorite"`
	URL             string `gorm:"-" json:"url"`
	Tags            []*Tag `gorm:"many2many:image_tags;" json:"tag"`
	ProcessComplete bool   `json:"processComplete"`
}

func (i *Image) Create() (err error) {
	return db.Create(i).Error
}

func (i *Image) First() (err error) {
	return db.Preload("Tags").Where(i).First(i).Error
}

func (i *Image) Find(images *[]Image) (err error) {
	// i.ProcessComplete = true
	return db.Where(i).Find(images).Error
}

func (i *Image) Update() (err error) {
	return db.Save(i).Error
}
