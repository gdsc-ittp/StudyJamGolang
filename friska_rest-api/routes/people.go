package routes

import (
	"simple-restapi/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PeopleRoutes(r *gin.Engine, db *gorm.DB) {
	peopleCtrl := controllers.PeopleController{DB: db}

	//Grouping routes > to organize routing from endpoint user
	peoples := r.Group("/peoples") 
	{
		peoples.GET("/", peopleCtrl.GetAll)
		peoples.GET("/:id", peopleCtrl.GetByID)
		peoples.POST("/", peopleCtrl.Create)
		peoples.PUT("/:id", peopleCtrl.Update)
		peoples.DELETE("/:id", peopleCtrl.Delete)
	}
}