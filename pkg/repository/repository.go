package repository

import (
	"database/sql"
	tictactoe_web "github.com/bogdanserdinov/tic-tac-toe-web"
)

type Authorization interface {
	CreateUser(user tictactoe_web.User) (int,error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sql.DB) *Repository{
	return &Repository{
		Authorization : NewAuthRepository(db),
	}
}
