package dbhandler

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	RentalStatusAvailable = "available"
	RentalStatusLoanedOut = "loanedout"
)

type Rental struct {
	gorm.Model
	VideoName string
	Customer  string
	Status    string
}

func ConnectToDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=docker dbname=videostore port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to DB")
	}

	return db
}
