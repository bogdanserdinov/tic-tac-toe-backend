package repository

import (
	"database/sql"
	"github.com/bogdanserdinov/tic-tac-toe-web"
)

type Authorization interface {
	CreateUser(user tictactoe_web.User) (int,error)
	GetUser(name string,password string) (tictactoe_web.User,error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sql.DB) *Repository{
	return &Repository{
		Authorization : NewAuthRepository(db),
	}
}
