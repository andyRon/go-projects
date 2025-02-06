package middleware

import (
	"github.com/andyron/meeting/internal/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		userClaims, err := helper.AnalyseToken(auth)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "Unauthorized Authorization",
			})
			return
		}
		if userClaims == nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "Unauthorized Admin",
			})
			return
		}
		c.Set("user_claims", userClaims)
		c.Next()
	}
}
