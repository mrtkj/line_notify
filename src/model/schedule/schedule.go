package schedule

type Schedule struct {
	ID          int    `json:"id" db:"id"`
	UserID      int    `json:"userId" db:"user_id"`
	ExecuteDate string `json:"executeTime" db:"execute_date"`
	ExecuteTime string `json:"executeDate" db:"execute_time"`
	Task        string `json:"task" db:"task"`
}
