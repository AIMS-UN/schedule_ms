package controller

import (
	service "aims_schedule_ms/service"

	"github.com/gin-gonic/gin"
)

func ScheduleRoutes(r *gin.Engine) {
	// Controllers
	r.GET("/schedule/:user_id", service.FindSchedules)          // find schedules by user id
	r.GET("/schedule/:user_id/:semester", service.FindSchedule) // find by semeter
}
