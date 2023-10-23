package main

import (
	"github.com/gin-gonic/gin"

	"github.com/anhnmt/golang-gin-base-project/api/auth"
	"github.com/anhnmt/golang-gin-base-project/api/post"
)

type Controller func(api *gin.RouterGroup)

func Router(r *gin.Engine) {
	api := r.Group("/api")

	ctrls := []Controller{
		auth.AuthController,
		post.PostController,
	}

	for _, ctrl := range ctrls {
		ctrl(api)
	}
}
