package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jasonbirchall/tools-api-poc/controllers"
	"github.com/jasonbirchall/tools-api-poc/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/tools", controllers.FindTools)
	r.POST("/tools", controllers.CreateTool)
	r.GET("/tools/:name", controllers.FindTool)
	r.PATCH("/tools/:name", controllers.UpdateTool)
	r.DELETE("/tools/:name", controllers.DeleteTool)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
