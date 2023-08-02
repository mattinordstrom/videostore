package main

import (
	"fmt"
	"log"

	dbhandler "github.com/mattinordstrom/videostore/db"
)

func main() {
	var gormDB = dbhandler.ConnectToDB()

	if err := gormDB.AutoMigrate(&dbhandler.Rental{}); err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
	}

	fmt.Println("Migrate DB done")
}
