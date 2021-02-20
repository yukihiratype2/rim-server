package model

import (
	"gorm.io/gorm"
)

// Image Struct
type Image struct {
	gorm.Model
	Name            string `json:"name,omitempty" form:"name"`
	FileID          string `json:"fileId,omitempty"`
	FolderID        uint   `json:"folderId,omitempty" form:"folderId"`
	Favorite        *bool  `json:"favorite,omitempty" form:"favorite"`
	URL             string `gorm:"-" json:"url,omitempty"`
	Tags            []*Tag `gorm:"many2many:image_tags;" json:"tags,omitempty" form:"tags"`
	ProcessComplete bool   `json:"-"`
}

// Create image
func (i *Image) Create() (err error) {
	return db.Create(i).Error
}

// First find image
func (i *Image) First() (err error) {
	return db.Preload("Tags").Where(i).First(i).Error
}

// Find all image by filter
func (i *Image) Find(images *[]Image) (err error) {
	i.ProcessComplete = true
	return db.Preload("Tags").Where(i).Find(images).Error
}

// Update image
func (i *Image) Update() (err error) {
	err = db.Updates(i).Error
	if err != nil {
		return err
	}
	// Check
	return i.First()
}
