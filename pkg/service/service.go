package service

import (
	"github.com/bogdanserdinov/tic-tac-toe-web"
	"github.com/bogdanserdinov/tic-tac-toe-web/pkg/repository"
)

type Authorization interface {
	SignIn(user tictactoe_web.User) (int,error)
	//SignUp(user tictactoe_web.User) (int,error)
}

type Service struct {
	Authorization
}

func NewService(repository *repository.Repository) *Service{
	return &Service{
		Authorization: NewAuthService(repository),
	}
}