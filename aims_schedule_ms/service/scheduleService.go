package services

import (
	dto "aims_schedule_ms/dto" //new
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GET /schedules
// Get all schedules
func FindSchedules(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var schedules []dto.Scheduledto
	if err := db.Where("user_id = ?", c.Param("enrrolmentId")).Find(&schedules).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User Id not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Schedules": schedules})
}

// GET /schedule/:semester“
// Find a schedule
func FindSchedule(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var schedule []dto.Scheduledto
	if err := db.Where("semester = ? AND user_id = ?", c.Param("semester"), c.Param("enrrolmentId")).Find(&schedule).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Schedule": schedule})
}
