package db

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mattinordstrom/videostore/pdf"
)

func GetRentals(context *gin.Context) {
	var rentals []Rental
	if context.Query("customer") != "" {
		gormDB.Raw("SELECT * FROM rentals WHERE customer = ? ORDER BY created_at DESC", context.Query("customer")).Scan(&rentals)
	} else {
		gormDB.Raw("SELECT * FROM rentals ORDER BY created_at DESC").Scan(&rentals)
	}

	context.JSON(200, gin.H{"rentals": rentals})
}

func AddRental(context *gin.Context) {
	var body struct {
		VideoName string `json:"videoName"`
		Customer  string `json:"customer"`
	}

	if err := context.BindJSON(&body); err != nil {
		fmt.Println("Error binding json!")

		return
	}

	rentalID := uuid.New()

	finishedPDF := make(chan int)
	go pdf.CreatePDF(finishedPDF, rentalID, body.VideoName, body.Customer)

	rental := Rental{
		VideoName: body.VideoName,
		Customer:  body.Customer,
		Status:    RentalStatusLoanedOut,
		RentalID:  rentalID,
	}

	result := gormDB.Create(&rental)

	// ERROR
	if result.Error != nil {
		fmt.Println("Error adding to db!")
		context.Status(400)

		return
	}

	// SUCCESS
	fmt.Println("Success adding to db!")

	pdfRes := <-finishedPDF

	context.JSON(200, gin.H{
		"savedtodb":  "success",
		"createdpdf": pdfRes,
	})
}

func ReturnRental(context *gin.Context) {
	result := gormDB.Model(Rental{}).Where("rental_id = ?", context.Param("rentalid")).Update("status", RentalStatusAvailable)

	// ERROR
	if result.Error != nil {
		fmt.Println("Error adding to db!")
		context.Status(400)

		return
	}

	// SUCCESS
	fmt.Println("Success updating db!")
	context.JSON(200, gin.H{"response": "success"})
}
