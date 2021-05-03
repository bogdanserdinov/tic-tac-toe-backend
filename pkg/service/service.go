package service

import (
	"github.com/bogdanserdinov/tic-tac-toe-web"
	"github.com/bogdanserdinov/tic-tac-toe-web/pkg/repository"
)

type Authorization interface {
	CreateUser(user tictactoe_web.User) (int, error)
	GenerateToken(name, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type Stats interface {
	GetStats(int int) (tictactoe_web.UserStats, error)
	UpdateStats(id int,result string) (tictactoe_web.UserStats,error)
	CreateStats(id int) error
}

type Service struct {
	Authorization
	Stats
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository),
		Stats:         NewStatsService(repository),
	}
}
