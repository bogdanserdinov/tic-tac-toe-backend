package repository

import (
	"database/sql"
	tictactoe_web "github.com/bogdanserdinov/tic-tac-toe-web"
)

type AuthRepository struct{
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository{
	return &AuthRepository{
		db : db,
	}
}

func(r *AuthRepository) CreateUser(user tictactoe_web.User) (int,error){
	return 0,nil
}