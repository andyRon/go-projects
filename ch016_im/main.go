package main

import "github.com/andyron/go-im/router"

func main() {
	engine := router.Router()
	engine.Run(":8080")
}
