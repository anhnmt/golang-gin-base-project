package post

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PostController(r *gin.RouterGroup, db *gorm.DB) {
	post := New(db)

	r.GET("/pong", post.Ping)
}
