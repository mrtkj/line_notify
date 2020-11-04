package repository

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/line_notify/src/db"
	"github.com/line_notify/src/model/schedule"
)

type SchedulesRepository interface {
	GetSchedules(ctx context.Context, date string) ([]schedule.Schedule, error)
}

func NewSchedulesRepository(db *db.DBContext) SchedulesRepository {
	return schedulesRepositoryImpl{
		db: db.DB,
	}
}

type schedulesRepositoryImpl struct {
	db *sql.DB
}

func (u schedulesRepositoryImpl) GetSchedules(ctx context.Context, date string) ([]schedule.Schedule, error) {
	rows, err := u.db.Query(`SELECT id, user_id, execute_date, execute_time, task
							FROM t_schedules
							WHERE
								execute_date = $1
							ORDER BY id;`, date)
	if err != nil {
		return nil, err
	}
	var list []schedule.Schedule
	for rows.Next() {
		var schedule schedule.Schedule
		err = rows.Scan(&schedule.ID, &schedule.UserID, &schedule.ExecuteDate, &schedule.ExecuteTime, &schedule.Task)
		if err != nil {
			return nil, err
		}
		list = append(list, schedule)
	}
	return list, nil
}
