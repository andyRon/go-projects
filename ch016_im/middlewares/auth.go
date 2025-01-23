package middlewares

import (
	"github.com/andyron/go-im/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		uc, err := helper.AnalyzeToken(token)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "用户认证失败",
			})
			return
		}
		c.Set("user_claims", uc)
		c.Next()
	}
}
