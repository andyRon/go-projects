package main

import (
	"github.com/andyron/mini-admin/define"
	"github.com/andyron/mini-admin/models"
	"github.com/andyron/mini-admin/router"
)

func main() {
	define.InitEnv()
	models.NewGormDB()
	models.NewRedisDB()

	r := router.App()
	r.Run(define.Port)
}
