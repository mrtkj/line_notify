package service

import (
	"github.com/gin-gonic/gin"
	"github.com/line_notify/src/repository"
)

type ScheduleService interface {
	ExecSchedule(ctx *gin.Context) error
}

func NewScheduleService(u repository.UsersRepository) ScheduleService {
	return scheduleServiceImpl{
		userRepository: u,
	}
}

type scheduleServiceImpl struct {
	userRepository repository.UsersRepository
}

func (s scheduleServiceImpl) ExecSchedule(ctx *gin.Context) error {
	return nil
}
