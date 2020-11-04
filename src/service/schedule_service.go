package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/line_notify/src/client/line"
	"github.com/line_notify/src/constants"
	"github.com/line_notify/src/model/user"
	"github.com/line_notify/src/repository"
)

type ScheduleService interface {
	ExecSchedule(ctx *gin.Context) error
}

func NewScheduleService(u repository.UsersRepository, s repository.SchedulesRepository, c line.Client) ScheduleService {
	return scheduleServiceImpl{
		userRepository:     u,
		scheduleRepository: s,
		client:             c,
	}
}

type scheduleServiceImpl struct {
	userRepository     repository.UsersRepository
	scheduleRepository repository.SchedulesRepository
	client             line.Client
}

func (s scheduleServiceImpl) ExecSchedule(ctx *gin.Context) error {
	users, err := s.userRepository.GetUsers(ctx)
	if err != nil {
		return err
	}

	now := time.Now()
	dt := now.Format("20060102")
	schedules, err := s.scheduleRepository.GetSchedules(ctx, dt)

	tasks := make([]string, 0, len(schedules))
	for _, schedule := range schedules {
		user := getUser(users, schedule.UserID)
		tasks = append(tasks, fmt.Sprintf("%s: %s", user.Name, schedule.Task))
	}
	msg := fmt.Sprintf(constants.MessageTemplate,
		fmt.Sprintf("%d/%d", int(now.Month()), now.Day()), strings.Join(tasks, "\n"))
	s.client.SendMessage(msg)

	return nil
}

func getUser(list []user.User, userID int) *user.User {
	var user user.User
	for _, u := range list {
		if u.ID == userID {
			user = u
		}
	}
	return &user
}
