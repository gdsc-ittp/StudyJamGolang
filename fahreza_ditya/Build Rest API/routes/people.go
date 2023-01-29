package routes

import (
	"studyjam4/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PeopleRoutes(r *gin.Engine, db *gorm.DB) {
	peopleCtrl := controllers.PeopleController{DB: db}

	peoples := r.Group("/people")
	{
		peoples.GET("/", peopleCtrl.GetAll)
		peoples.GET("/:id", peopleCtrl.GetByID)
		peoples.POST("/", peopleCtrl.Create)
		peoples.PUT("/:id", peopleCtrl.Update)
		peoples.DELETE("/:id", peopleCtrl.Delete)
	}
}