package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DBContext struct {
	db *sql.DB
}

type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func NewDB(config Config) (*DBContext, error) {
	db, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			config.Host, config.Port, config.User, config.Password, config.Database))
	d := &DBContext{
		db: db,
	}

	return d, err
}
