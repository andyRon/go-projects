package main

import (
	"github.com/andyron/meeting/internal/models"
	"github.com/andyron/meeting/internal/server/router"
	"log"
)

func main() {
	models.NewDB()
	engine := router.Router()
	err := engine.Run()
	if err != nil {
		log.Fatalln("run engine err:", err)
		return
	}
}
