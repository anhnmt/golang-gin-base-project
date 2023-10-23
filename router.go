package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/anhnmt/golang-gin-base-project/api/auth"
	"github.com/anhnmt/golang-gin-base-project/api/post"
)

type Controller func(api *gin.RouterGroup, db *gorm.DB)

func Router(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api")

	ctrls := []Controller{
		auth.AuthController,
		post.PostController,
	}

	for _, ctrl := range ctrls {
		ctrl(api, db)
	}
}
