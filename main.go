package main

import (
	"github.com/gin-gonic/gin"

	"github.com/anhnmt/golang-gin-base-project/db"
)

func main() {
	db.New()

	r := gin.Default()

	Router(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}
