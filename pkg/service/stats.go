package service

import (
	"github.com/bogdanserdinov/tic-tac-toe-web"
	"github.com/bogdanserdinov/tic-tac-toe-web/pkg/repository"
)

type StatsService struct {
	repo *repository.Repository
}

func NewStatsService(repo *repository.Repository) *StatsService {
	return &StatsService{
		repo: repo,
	}
}

func (s *StatsService) CreateStats(id int) error{
	return nil
}

func (s *StatsService) GetStats(int int) (tictactoe_web.UserStats,error){
	return tictactoe_web.UserStats{}, nil
}

func (s *StatsService) UpdateStats(id int) (tictactoe_web.UserStats,error){
	return tictactoe_web.UserStats{},nil
}



