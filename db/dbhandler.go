package db

import (
	"log"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	RentalStatusAvailable = "available"
	RentalStatusLoanedOut = "loanedout"
)

var gormDB *gorm.DB

type Rental struct {
	gorm.Model
	VideoName string
	Customer  string //TODO make new table Customers and change this to CustomerID
	Status    string
	RentalID  uuid.UUID
}

func ConnectToDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=docker dbname=videostore port=5432 sslmode=disable"
	newGormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to DB")
	}

	gormDB = newGormDB
	return gormDB
}
