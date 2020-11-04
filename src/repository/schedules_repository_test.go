package repository

import (
	"context"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/line_notify/src/model/schedule"
)

func TestSchedulesRepository_GetSchedules(t *testing.T) {
	db := InitDB()
	sr := NewSchedulesRepository(db)
	ctx := context.TODO()

	tests := []struct {
		name string
		err  error
		dt   string
		want []schedule.Schedule
	}{
		{
			name: "getUsers_20201231",
			err:  nil,
			dt:   "20201231",
			want: []schedule.Schedule{
				{
					ID:          1,
					UserID:      1,
					ExecuteDate: "20201231",
					ExecuteTime: "080000",
					Task:        "test task",
				},
			},
		},
		{
			name: "getUsers_20201230",
			err:  nil,
			dt:   "20201230",
			want: []schedule.Schedule{
				{
					ID:          2,
					UserID:      2,
					ExecuteDate: "20201230",
					ExecuteTime: "090000",
					Task:        "test task2",
				},
				{
					ID:          3,
					UserID:      3,
					ExecuteDate: "20201230",
					ExecuteTime: "100000",
					Task:        "test task3",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, _ := sr.GetSchedules(ctx, test.dt)

			if test.err != nil {
				return
			}

			assert.Equal(t, test.want, result)
		})
	}

	CloseDB(db)
}
