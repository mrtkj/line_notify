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

func (s *ScheduleController) ExecSchedule(ctx *gin.Context) {
	err := s.ss.ExecSchedule(ctx)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"isSuccess": true,
	})
}
