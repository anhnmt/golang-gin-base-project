package main

import (
	"github.com/gin-gonic/gin"

	"github.com/anhnmt/golang-gin-base-project/database"
)

func main() {
	db := database.New()

	r := gin.Default()

	Router(r, db)

	r.Run() // listen and serve on 0.0.0.0:8080
}
