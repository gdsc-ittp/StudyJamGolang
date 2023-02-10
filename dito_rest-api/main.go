package main

import (
	"log"
	"restful-api/config"
	"restful-api/models"
	"restful-api/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	// connect to db
	var db *gorm.DB
	var err error
	db, err = config.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// migrate models
	db.AutoMigrate(&models.People{})

	// setup gin router
	r := gin.Default()
	routes.PeopleRoutes(r, db)

	r.Run(":3000")

}
