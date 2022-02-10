package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jasonbirchall/tools-api-poc/controllers"
	"github.com/jasonbirchall/tools-api-poc/models"
)

func main() {
	r := gin.Default()

	err := models.ConnectDatabase()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	r.GET("/tools", controllers.FindTools)
	r.POST("/tools", controllers.CreateTool)
	r.POST("/tools/testData", controllers.PopulateTools)
	r.GET("/tools/:name", controllers.FindTool)
	r.PATCH("/tools/:name", controllers.UpdateTool)
	r.DELETE("/tools/:name", controllers.DeleteTool)

	err = r.Run(":5000")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
