package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jasonbirchall/tools-api-poc/models"
)

type CreateToolInput struct {
	Name    string `json:"name" binding:"required"`
	Version int    `json:"version" binding:"required"`
}

type UpdateToolInput struct {
	Name    string `json:"name"`
	Version int    `json:"version"`
}

func FindTools(c *gin.Context) {
	var tools []models.Tools
	models.DB.Find(&tools)

	c.JSON(http.StatusOK, gin.H{"data:": tools})
}

func CreateTool(c *gin.Context) {
	var input CreateToolInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tool := models.Tools{Name: input.Name, Version: input.Version}
	models.DB.Create(&tool)

	c.JSON(http.StatusOK, gin.H{"data": tool})
}

func FindTool(c *gin.Context) {
	var tool models.Tools
	if err := models.DB.Where("name = ?", c.Param("name")).First(&tool).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tool})
}

func UpdateTool(c *gin.Context) {
	var tool models.Tools

	if err := models.DB.Where("name = ?", c.Param("name")).First(&tool).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input UpdateToolInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&tool).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": tool})
}

func DeleteTool(c *gin.Context) {
	var tool models.Tools

	if err := models.DB.Where("name = ?", c.Param("name")).First(&tool).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Delete(&tool)

	c.JSON(http.StatusOK, gin.H{"data": tool})
}
