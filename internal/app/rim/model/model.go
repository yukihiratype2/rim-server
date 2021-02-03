package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Model struct {
	DB *gorm.DB
}

func NewModel() (model *Model, err error) {
	db, err := gorm.Open(sqlite.Open("temp.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Image{}, &Tag{})
	model = &Model{db}
	return model, err
}
