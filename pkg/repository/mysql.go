package repository

import (
	"database/sql"
	"fmt"
)

const (
	User      = "user"
	UserStats = "userstats"
	Game      = "game"
	GameStats = "gamestats"
)

type Config struct {
	Username string
	Password string
	DBName   string
}

func InitDB(conf Config) (*sql.DB, error) {
	configuration := fmt.Sprintf("%s:%s@/%s", conf.Username, conf.Password, conf.DBName)
	db, err := sql.Open("mysql", configuration)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
