package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/anhnmt/golang-gin-base-project/api/auth"
	"github.com/anhnmt/golang-gin-base-project/api/post"
)

func New() *gorm.DB {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(
		&auth.Auth{},
		&post.Post{},
	)

	return db
}
