package db

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mattinordstrom/videostore/pdf"
)

func GetRentalPDF(c *gin.Context) {
	c.File(c.Param("rentalid") + ".pdf")
}

func GetRentals(c *gin.Context) {
	var rentals []Rental
	if c.Query("customer") != "" {
		gormDB.Raw("SELECT * FROM rentals WHERE customer = ? ORDER BY created_at DESC", c.Query("customer")).Scan(&rentals)
		//gormDB.Where("customer = ?", c.Query("customer")).Order("created_at DESC").Find(&rentals)
	} else {
		gormDB.Order("created_at DESC").Find(&rentals)
	}

	c.JSON(200, gin.H{"rentals": rentals})
}

func AddRental(c *gin.Context) {
	var body struct {
		VideoName string `json:"VideoName"`
		Customer  string `json:"Customer"`
	}
	c.BindJSON(&body)

	rentalId := uuid.New()

	finishedPDF := make(chan int)
	go pdf.CreatePDF(finishedPDF, rentalId, body.VideoName, body.Customer)

	rental := Rental{
		VideoName: body.VideoName,
		Customer:  body.Customer,
		Status:    RentalStatusLoanedOut,
		RentalID:  rentalId,
	}
	result := gormDB.Create(&rental)

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

func ReturnRental(c *gin.Context) {
	result := gormDB.Model(Rental{}).Where("rental_id = ?", c.Param("rentalid")).Update("status", RentalStatusAvailable)

	//ERROR
	if result.Error != nil {
		fmt.Println("Error adding to db!")
		c.Status(400)
		return
	}

	//SUCCESS
	fmt.Println("Success updating db!")
	c.JSON(200, gin.H{"response": "success"})
}
