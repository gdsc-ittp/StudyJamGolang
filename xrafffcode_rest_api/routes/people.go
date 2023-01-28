package routes

import (
	"gdsc-go/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PeopleRoutes(r *gin.Engine, db *gorm.DB) {
	peopleCtrl := controllers.PeopleContrtoller{DB: db}

	people := r.Group("/people")

	people.GET("/", peopleCtrl.GetAll)
	people.GET("/:id", peopleCtrl.GetOne)
	people.POST("/", peopleCtrl.Create)
	people.PUT("/:id", peopleCtrl.Update)
	people.DELETE("/:id", peopleCtrl.Delete)

}
