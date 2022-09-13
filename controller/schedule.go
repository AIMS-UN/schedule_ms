package controllers

import (
	models "aims_schedule_ms/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GET /schedules
// Get all schedules
func FindSchedules(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var schedules []models.Enrrollment
	db.Find(&schedules)
	c.JSON(http.StatusOK, gin.H{"Schedules": schedules})
}

// GET /schedule/:semester“
// Find a schedule
func FindSchedule(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var schedule models.Enrrollment
	if err := db.Where("semester = ?", c.Param("semester")).First(&schedule).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Schedule": schedule})
}
