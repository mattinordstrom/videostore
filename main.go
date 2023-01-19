package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mattinordstrom/videostore/db"
)

func main() {
	fmt.Println("Retro Video Store - VHS & DVD")

	db.ConnectToDB()

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://127.0.0.1:8080"}
	r.Use(cors.New(config))

	r.POST("/rental", db.AddRental)
	r.PUT("/rental/:rentalid/return", db.ReturnRental)
	r.GET("/rentals", db.GetRentals)
	r.GET("/rental/receipt/:rentalid", db.GetRentalPDF)

	r.Run(":3000")
}
