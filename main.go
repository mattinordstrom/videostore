package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	dbhandler "github.com/mattinordstrom/videostore/db"
	pdfhandler "github.com/mattinordstrom/videostore/pdf"
	"gorm.io/gorm"
)

var db *gorm.DB

func getRentalPDF(c *gin.Context) {
	c.File(c.Param("rentalid") + ".pdf")
}

func getRentals(c *gin.Context) {
	var rentals []dbhandler.Rental
	if c.Query("customer") != "" {
		db.Where("customer = ?", c.Query("customer")).Find(&rentals)
	} else {
		db.Find(&rentals)
	}

	c.JSON(200, gin.H{
		"rentals": rentals,
	})
}

func addRental(c *gin.Context) {
	rentalId := uuid.New()

	finishedPDF := make(chan int)
	go pdfhandler.CreatePDF(finishedPDF, rentalId)

	var body struct {
		VideoName string
		Customer  string
	}
	c.Bind(&body)

	rental := dbhandler.Rental{
		VideoName: body.VideoName,
		Customer:  body.Customer,
		Status:    dbhandler.RentalStatusLoanedOut,
		RentalID:  rentalId,
	}
	result := db.Create(&rental)

	//ERROR
	if result.Error != nil {
		fmt.Println("Error adding to db!")
		c.Status(400)
		return
	}

	//SUCCESS
	fmt.Println("Success adding to db!")

	pdfRes := <-finishedPDF

	c.JSON(200, gin.H{
		"savedtodb":  "success",
		"createdpdf": pdfRes,
	})
}

func returnRental(c *gin.Context) {
	result := db.Model(dbhandler.Rental{}).Where("rental_id = ?", c.Param("rentalid")).Update("status", dbhandler.RentalStatusAvailable)

	//ERROR
	if result.Error != nil {
		fmt.Println("Error adding to db!")
		c.Status(400)
		return
	}

	//SUCCESS
	fmt.Println("Success adding to db!")
	c.JSON(200, gin.H{
		"response": "success",
	})
}

func main() {
	fmt.Println("Retro Video Store - VHS & DVD")
	db = dbhandler.ConnectToDB()

	r := gin.Default()
	r.POST("/rental", addRental)
	r.PUT("/rental/:rentalid/return", returnRental)
	r.GET("/rentals", getRentals)
	r.GET("/rental/receipt/:rentalid", getRentalPDF)
	r.Run(":3000")
}
