package main

import (
	controllers "aims_schedule_ms/controller" // new
	models "aims_schedule_ms/models"          // new
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	log.Println("Start Project")
	db := models.SetupModels() // new

	// Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	//log.Println(DB.RowsAffected)
	r.GET("/schedules", controllers.FindSchedules)
	r.GET("/schedule/:semester", controllers.FindSchedule) // find by semeter
	r.Run("localhost:4000")
}
