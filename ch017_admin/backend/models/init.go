package models

import (
	"github.com/andyron/mini-admin/define"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RDB *redis.Client

func NewGormDB() {
	db, err := gorm.Open(mysql.Open(define.MiniAdminDSN), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{}, &Role{}, &Menu{}, &RoleMenu{}, &RoleFunction{}, &Function{})
	if err != nil {
		panic(err)
	}
	DB = db
}

func NewRedisDB() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     define.RedisAddr,
		Password: define.RedisPassword,
		Username: define.RedisUsername,
	})
}
