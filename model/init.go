package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBInit(connString string) {
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	_ = db.AutoMigrate(&User{})
	_ = db.AutoMigrate(&Ad{})
	_ = db.AutoMigrate(&Image{})

	DB = db
}
