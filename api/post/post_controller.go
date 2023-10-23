package post

import (
	"github.com/gin-gonic/gin"
)

func PostController(r *gin.RouterGroup) {
	service := NewPostService()

	r.GET("/pong", service.Ping)
}
