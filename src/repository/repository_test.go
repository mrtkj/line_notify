package repository

import (
	"os"

	"github.com/line_notify/src/config"
	"github.com/line_notify/src/db"
)

func InitDB() *db.DBContext {
	setEnv()
	config := config.NewConfig()
	d, _ := db.NewDB(config.DB)
	return d
}

func CloseDB(d *db.DBContext) {
	d.DB.Close()
}

func setEnv() {
	os.Setenv("DB_USER", "user")
	os.Setenv("DB_PASSWORD", "password")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_DATABASE", "db")
}
