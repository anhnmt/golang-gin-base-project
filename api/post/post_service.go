package post

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PostService struct {
	db *gorm.DB
}

func New(db *gorm.DB) PostService {
	return PostService{
		db: db,
	}
}

func (s PostService) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
