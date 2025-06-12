package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/captcha", func(c *gin.Context) {
		// TODO
	})
}
