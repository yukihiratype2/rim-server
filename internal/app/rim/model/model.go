package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() (err error) {
	db, err = gorm.Open(sqlite.Open("temp.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db.AutoMigrate(&Image{}, &Tag{}, &Folder{})
}
