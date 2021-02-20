package model

import "gorm.io/gorm"

// Tag Struct
type Tag struct {
	gorm.Model
	Label  string   `json:"label"`
	Color  string   `json:"color"`
	Images []*Image `gorm:"many2many:image_tags;" json:"images"`
}

func (t *Tag) Find(tags *[]Tag) (err error) {
	return db.Find(tags).Error
}

func (t *Tag) Create() (err error) {
	return db.Create(t).Error
}

func (t *Tag) Update() (err error) {
	return db.Save(t).Error
}

func (t *Tag) Delete() (err error) {
	return db.Delete(t).Error
}
