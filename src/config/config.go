package config

import (
	"os"

	"github.com/line_notify/src/client/line"
	"github.com/line_notify/src/db"
)

type Config struct {
	DB struct {
		User     string
		Password string
		Host     string
		Port     string
		Database string
	}
	LINE struct {
		Token string
	}
}

func (c *Config) GetDBConfig() db.Config {
	return db.Config{
		User:     c.DB.User,
		Password: c.DB.Password,
		Host:     c.DB.Host,
		Port:     c.DB.Port,
		Database: c.DB.Database,
	}
}

func (c *Config) GetLINEConfig() line.Config {
	return line.Config{
		Token: c.LINE.Token,
	}
}

func NewConfig() *Config {

	User := os.Getenv("DB_USER")
	Password := os.Getenv("DB_PASSWORD")
	Host := os.Getenv("DB_HOST")
	Port := os.Getenv("DB_PORT")
	Database := os.Getenv("DB_DATABASE")

	db := db.Config{
		User:     User,
		Password: Password,
		Host:     Host,
		Port:     Port,
		Database: Database,
	}

	token := os.Getenv("LINE_NOTIFY_TOKEN")
	line := line.Config{
		Token: token,
	}

	config := Config{
		DB:   db,
		LINE: line,
	}
	return &config
}
