package main

import (
	"fmt"

	dbhandler "github.com/mattinordstrom/videostore/db"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = dbhandler.ConnectToDB()
}

func main() {
	db.AutoMigrate(&dbhandler.Rental{})

	fmt.Println("Migrate DB done")
}
