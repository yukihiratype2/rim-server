package model

// Tag Struct
type Tag struct {
	ID     uint     `gorm:"primarykey" json:"id"`
	Label  string   `json:"label"`
	Color  string   `json:"color"`
	Images []*Image `gorm:"many2many:image_tags;" json:"images"`
	// gorm.Model
}

func (t Tag) Create() (err error) {
	return db.Create(t).Error
}

func (t Tag) Update() (err error) {
	return db.Save(t).Error
}

func (t Tag) Delete() (err error) {
	return db.Delete(t).Error
}
