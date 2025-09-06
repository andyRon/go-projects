package router

import (
	"github.com/andyron/meeting/internal/middleware"
	"github.com/andyron/meeting/internal/server/service"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Use(middleware.Cors())

	r.POST("/user/login", service.UserLogin)

	// ws
	r.GET("/ws/p2p/:room_identity/:user_identity", service.Wsp2PConnection)

	//r.POST("/meeting/create", service.MeetingCreate)
	auth := r.Group("/auth", middleware.Auth())
	{
		auth.GET("/meeting/list", service.MeetingList)
		auth.POST("/meeting/create", service.MeetingCreate)
		auth.POST("/meeting/edit", service.MeetingEdit)
		auth.DELETE("/meeting/delete", service.MeetingDelete)
	}

	return r
}
