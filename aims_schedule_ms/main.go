package main

import (
	"log"

	"github.com/gin-gonic/gin"

	repository "aims_schedule_ms/repository" // new
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

	r.Run("localhost:4000")
}
