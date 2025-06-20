package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB() {
	dsn := "user:password@tcp(192.168.1.8:3306)/meeting?charset=utf8mb4&parseTime=True&loc=Local" // TODO
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Room{}, &RoomUser{}, &User{})
	DB = db
}
