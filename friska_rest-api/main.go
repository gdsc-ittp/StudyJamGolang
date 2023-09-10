package main

import (
	"log"
	"simple-restapi/config"
	"simple-restapi/models"
	"simple-restapi/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	// Connect to database > handle by config package
	var db *gorm.DB
	var err error
	db, err = config.ConnectDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	// Migrate models > by Gorm
	db.AutoMigrate(&models.People{})

	// Set up Gin router > on port 3000
	r := gin.Default()
	routes.PeopleRoutes(r, db)
	r.Run(":3000")
}