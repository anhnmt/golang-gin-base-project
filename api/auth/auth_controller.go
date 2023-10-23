package auth

import (
	"github.com/gin-gonic/gin"
)

func AuthController(r *gin.RouterGroup) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
