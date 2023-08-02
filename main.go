package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mattinordstrom/videostore/db"
)

func main() {
	fmt.Println("Retro Video Store - VHS & DVD")

	db.ConnectToDB()

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://127.0.0.1:8080"}
	router.Use(cors.New(config))

	router.POST("/rental", db.AddRental)
	router.PUT("/rental/:rentalid/return", db.ReturnRental)
	router.GET("/rentals", db.GetRentals)
	router.GET("/rental/receipt/:rentalid", getRentalPDF)

	if err := router.Run(":3000"); err != nil {
		log.Fatalf("Server failed to run: %v", err)
	}
}

func getRentalPDF(c *gin.Context) {
	c.File("pdf_output/" + c.Param("rentalid") + ".pdf")
}
