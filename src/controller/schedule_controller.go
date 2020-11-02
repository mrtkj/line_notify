package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line_notify/src/service"
)

func NewScheduleController(ss service.ScheduleService) *ScheduleController {
	return &ScheduleController{
		ss: ss,
	}
}

func (s *ScheduleController) Endpoint(api *gin.RouterGroup) {
	api.GET("/schedules/exec", s.ExecSchedule)
}

type ScheduleController struct {
	ss service.ScheduleService
}

func (s *ScheduleController) ExecSchedule(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
	})
}
