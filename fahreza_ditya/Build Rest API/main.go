package main

import (
	"log"
	"studyjam4/config"
	"studyjam4/models"
	"studyjam4/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	var db *gorm.DB
	var err error

	db, err = config.ConnectDB()
	if err != nil {
		log.Fatalf("Fail to connect to database : %v", err)
	}

	db.AutoMigrate(&models.People{})

	r := gin.Default()
	routes.PeopleRoutes(r, db)
	
	r.Run(":8080")
}