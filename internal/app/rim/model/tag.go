package model

// Tag Struct
type Tag struct {
	ID     uint     `gorm:"primarykey" json:"id"`
	Label  string   `json:"label"`
	Color  string   `json:"color"`
	Images []*Image `gorm:"many2many:image_tags;" json:"images"`
	// gorm.Model
}
