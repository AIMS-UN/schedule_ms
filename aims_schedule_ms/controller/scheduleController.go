package controller

import (
	service "aims_schedule_ms/service"
	"github.com/gin-gonic/gin"
)

func FindSchedule() {
	r := gin.Default()
	r.GET("/schedules/:enrrolmentId", service.FindSchedules)
}
func FindScheduleById() {
	r := gin.Default()
	r.GET("/schedule/:enrrolmentId/:semester", service.FindSchedule) // find by semeter
}
