package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthController(r *gin.RouterGroup, db *gorm.DB) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
