package model

import "time"

// Image Struct
type Image struct {
	Name    string    `json:"name"`
	FileID  string    `json:"fileId"`
	Created time.Time `json:"createdAt"`
}

// Tag Struct
type Tag struct {
	Label string `json:"label"`
	Color string `json:"color"`
}
