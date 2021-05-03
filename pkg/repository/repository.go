package repository

import (
	"database/sql"
	"github.com/bogdanserdinov/tic-tac-toe-web"
)

type Authorization interface {
	CreateUser(user tictactoe_web.User) (int,error)
	GetUser(name string,password string) (tictactoe_web.User,error)
}

type Stats interface {
	GetStats(user tictactoe_web.User) (tictactoe_web.UserStats,error)
}

type Repository struct {
	Authorization
	Stats
}

func NewRepository(db *sql.DB) *Repository{
	return &Repository{
		Authorization : NewAuthRepository(db),
		Stats : NewStatsRepository(db),

	}
}
