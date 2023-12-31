package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/anhnmt/golang-gin-base-project/entities"
)

var db *gorm.DB

func New() {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(
		&entities.Auth{},
		&entities.Post{},
	)
}

func Repository(value interface{}) *gorm.DB {
	return db.Model(value)
}
