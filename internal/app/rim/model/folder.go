package model

import (
	"gorm.io/gorm"
)

type Folder struct {
	gorm.Model
	Label  string `json:"label"`
	Images []Image
}

func (f *Folder) Create() (err error) {
	return db.Create(f).Error
}

func (f *Folder) Find(folders *[]Folder) (err error) {
	return db.Find(folders).Error
}

func (f *Folder) First() (err error) {
	return db.First(f).Error
}

func (f *Folder) Delete() (err error) {
	return db.Delete(f).Error
}
