package schedule

import "time"

type Schedule struct {
	ID          int       `json:"id" db:"id"`
	UserID      int       `json:"userId" db:"user_id"`
	ExecuteDate time.Time `json:"executeTime" db:"execute_date"`
	ExecuteTime time.Time `json:"executeDate" db:"execute_time"`
	Task        string    `json:"task" db:"task"`
}
