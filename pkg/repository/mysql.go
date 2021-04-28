package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	UserTable      = "user"
	UserStatsTable = "userstats"
	GameTable      = "game"
	GameStatsTable = "gamestats"
)

type Config struct {
	Username string
	Password string
	DBName   string
}

func InitMySQL(conf Config) (*sql.DB, error) {
	configuration := fmt.Sprintf("%s:%s@/%s", conf.Username, conf.Password, conf.DBName)
	fmt.Println(configuration)
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
