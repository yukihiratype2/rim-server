package model

import (
	"gorm.io/gorm"
)

// Image Struct
type Image struct {
	gorm.Model
	Name            string `json:"name" form:"name"`
	FileID          string `json:"fileId"`
	FolderID        uint   `json:"folderId" form:"folderId"`
	Favorite        *bool  `json:"favorite,omitempty" form:"favorite"`
	URL             string `gorm:"-" json:"url"`
	Tags            []*Tag `gorm:"many2many:image_tags;" json:"tag" form:"tag"`
	ProcessComplete bool   `json:"-"`
}

func (i *Image) Create() (err error) {
	return db.Create(i).Error
}

func (i *Image) First() (err error) {
	return db.Preload("Tags").Where(i).First(i).Error
}

func (i *Image) Find(images *[]Image) (err error) {
	i.ProcessComplete = true
	return db.Preload("Tags").Where(i).Find(images).Error
}

func (i *Image) Update() (err error) {
	return db.Updates(i).Error
}
