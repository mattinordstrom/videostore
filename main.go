package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Rental struct {
	VideoName string
	Customer  string
}

func main() {
	fmt.Println("Hello world")

	// GIN HTTP
	/*
		r := gin.Default()
		r.GET("/rental", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"resp": "abc123",
			})
		})
		r.Run(":3000")
	*/
	//GORM DB
	dsn := "host=localhost user=postgres password=docker dbname=videostore port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to DB")
	}

	rental := Rental{VideoName: "Die Hard", Customer: "John Smith"}
	result := db.Create(&rental)

	if result.Error != nil {
		fmt.Println("Error adding to db!")
		return
	}

	fmt.Println("Success adding to db!")
}
