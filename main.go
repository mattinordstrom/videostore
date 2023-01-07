package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	dbhandler "github.com/mattinordstrom/videostore/db"
	"gorm.io/gorm"
)

var db *gorm.DB

func getRentals(c *gin.Context) {
	var rentals []dbhandler.Rental
	db.Find(&rentals)

	c.JSON(200, gin.H{
		"rentals": rentals,
	})
}

func addRental(c *gin.Context) {
	var body struct {
		VideoName string
		Customer  string
	}
	c.Bind(&body)

	//TODO need to specify everything in body?
	rental := dbhandler.Rental{VideoName: body.VideoName, Customer: body.Customer}
	result := db.Create(&rental)

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
	r.GET("/rentals", getRentals)
	r.Run(":3000")
}
