package routes

import (
	"restful-api/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PeopleRoutes(r *gin.Engine, db *gorm.DB) {
	peopleCtrl := controllers.PeopleController{DB: db}

	// grouping routes
	peoples := r.Group("/peoples")
	{
		peoples.GET("/", peopleCtrl.GetAll) // memanggil fungsi getAll di controller
		peoples.GET("/:id", peopleCtrl.GetByID)
		peoples.POST("/", peopleCtrl.Create)
		peoples.PUT("/:id", peopleCtrl.Update)
		peoples.DELETE("/:id", peopleCtrl.Delete)
	}
}
