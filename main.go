package main

import (
	"os"

	"fmt"

	"log"

	"github.com/gin-gonic/gin"

	repository "aims_schedule_ms/repository"

	controller "aims_schedule_ms/controller"
)

func main() {
	r := gin.Default()
	log.Println("Start Project")
	db := repository.SetupModels() // new

	// Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	port := os.Getenv("PORT")
	controller.ScheduleRoutes(r)
	r.Run(fmt.Sprintf(":%s", port))
}
