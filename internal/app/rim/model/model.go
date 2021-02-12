package model

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() (err error) {
	dsn := "host=localhost user=natsuki dbname=natsuki port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db.AutoMigrate(&Tag{}, &Folder{}, &Image{})
}
