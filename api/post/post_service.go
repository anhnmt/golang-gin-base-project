package post

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/anhnmt/golang-gin-base-project/db"
	"github.com/anhnmt/golang-gin-base-project/entities"
)

type PostService struct {
	postRepository *gorm.DB
}

func NewPostService() PostService {
	s := PostService{
		postRepository: db.Repository(&entities.Post{}),
		// authRepository: db.Repository(&entities.Auth{}),
	}

	return s
}

func (t PostService) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
