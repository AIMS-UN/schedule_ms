package services

import (
	//new
	model "aims_schedule_ms/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GET /schedules/:user_id“
// Find a schedules by user_id
func FindSchedules(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var schedules []model.Enrrollment
	if err := db.Where("user_id = ?", c.Param("user_id")).Find(&schedules).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "User id not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": schedules})
}

// GET /schedule/:user_id/:semester“
// Find a schedule by user_id and semester
func FindSchedule(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var schedule model.Enrrollment
	if err := db.Where("user_id = ? AND semester = ?", c.Param("user_id"), c.Param("semester")).First(&schedule).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": schedule})
}
