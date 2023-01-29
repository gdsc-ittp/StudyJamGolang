package controllers

import (
	"gdsc-go/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PeopleContrtoller struct {
	DB *gorm.DB
}

func (ctrl PeopleContrtoller) GetAll(c *gin.Context) {
	var peoples []models.People

	ctrl.DB.Find(&peoples)

	c.JSON(200, gin.H{
		"message": "Get People Succes",
		"peoples": peoples,
	})
}

func (ctrl PeopleContrtoller) GetOne(c *gin.Context) {
	var people models.People

	if err := ctrl.DB.Where("id = ?", c.Param("id")).First(&people).Error; err != nil {
		c.JSON(404, gin.H{
			"message": "Not Found",
		})
		return
	}

	c.JSON(200, gin.H{
		"people": people,
	})
}

func (ctrl PeopleContrtoller) Create(c *gin.Context) {
	var people models.People

	if err := c.ShouldBindJSON(&people); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}

	ctrl.DB.Create(&people)

	c.JSON(201, gin.H{
		"message": "Create People Succes",
		"people":  people,
	})
}

func (ctrl PeopleContrtoller) Update(c *gin.Context) {
	var people models.People

	if err := ctrl.DB.Where("id = ?", c.Param("id")).First(&people).Error; err != nil {
		c.JSON(404, gin.H{
			"message": "Not Found",
		})
		return
	}

	if err := c.ShouldBindJSON(&people); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}

	ctrl.DB.Save(&people)

	c.JSON(200, gin.H{
		"message": "Update People Succes",
		"people":  people,
	})
}

func (ctrl PeopleContrtoller) Delete(c *gin.Context) {
	var people models.People

	if err := ctrl.DB.Where("id = ?", c.Param("id")).First(&people).Error; err != nil {
		c.JSON(404, gin.H{
			"message": "Not Found",
		})
		return
	}

	ctrl.DB.Delete(&people)

	c.JSON(200, gin.H{
		"message": "Delete Success",
	})
}
