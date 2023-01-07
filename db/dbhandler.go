package dbhandler

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// TODO move model to separate file
type Rental struct {
	gorm.Model
	VideoName string
	Customer  string
}

func ConnectToDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=docker dbname=videostore port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to DB")
	}

	return db
}
