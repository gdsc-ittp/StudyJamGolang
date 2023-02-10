package controllers

import (
	"net/http"
	"studyjam4/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PeopleController struct {
	DB *gorm.DB
}

func (ctrl PeopleController) GetAll(c *gin.Context)  {
	var peoples []models.People
	ctrl.DB.Find(&peoples)
	c.JSON(http.StatusOK, gin.H{
		"data": peoples,
	})
}

func (ctrl PeopleController) GetByID(c *gin.Context)  {
	var people models.People
	if err := ctrl.DB.Where("id = ?", c.Param("id")).First(&people).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": people,
	})
}

func (ctrl PeopleController) Create(c *gin.Context)  {
	var  people models.People
	if err := c.ShouldBindJSON(&people); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctrl.DB.Create(&people)
	c.JSON(http.StatusOK, gin.H{
		"data": people,
	})
}

func (ctrl PeopleController) Update(c *gin.Context)  {
	var people models.People
	if err := ctrl.DB.Where("id = ?", c.Param("id")).First(&people).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}
	if err := c.ShouldBindJSON(&people); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctrl.DB.Save(&people)
	c.JSON(http.StatusOK, gin.H{
		"data": people,
	})
}

func (ctrl PeopleController) Delete(c *gin.Context)  {
	var people models.People
	if err := ctrl.DB.Where("id = ?", c.Param("id")).First(&people).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}
	ctrl.DB.Delete(&people)
	c.JSON(http.StatusOK, gin.H{
		"data": "People deleted succesfully",
	})
}