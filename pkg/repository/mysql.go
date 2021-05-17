package repository

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
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

func InitMySQL(conf Config) (*sqlx.DB, error) {
	configuration := fmt.Sprintf("%s:%s@(localhost:3306)/%s", conf.Username, conf.Password, conf.DBName)
	fmt.Println(configuration)
	db, err := sqlx.Connect("mysql", configuration)
	if err != nil {
		return nil, err
	}

	return db, nil
}
