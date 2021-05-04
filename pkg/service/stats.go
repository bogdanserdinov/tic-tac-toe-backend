package service

import (
	"github.com/bogdanserdinov/tic-tac-toe-web"
	"github.com/bogdanserdinov/tic-tac-toe-web/pkg/repository"
	"github.com/sirupsen/logrus"
)

type StatsService struct {
	repo *repository.Repository
}

func NewStatsService(repo *repository.Repository) *StatsService {
	return &StatsService{
		repo: repo,
	}
}

func (s *StatsService) GetStats(id int) (tictactoe_web.UserStats, error) {
	stats, err := s.repo.Stats.GetStats(id)
	if err != nil {
		logrus.Errorf("statsService: could not get stats: %s", err.Error())
		return tictactoe_web.UserStats{}, err
	}

	return stats, err
}

func (s *StatsService) UpdateStats(id int, result string) (tictactoe_web.UserStats, error) {
	stats, err := s.GetStats(id)
	if err != nil {
		logrus.Errorf("statsService: could not get stats: %s", err.Error())
		return tictactoe_web.UserStats{}, err
	}
	stats.TotalGames++
	if result == "win" {
		stats.Wins++
	}
	if result == "draw" {
		stats.Draws++
	}
	if result == "lose" {
		stats.Losses++
	}

	newStats, err := s.repo.Stats.UpdateStats(stats)
	if err != nil {
		return tictactoe_web.UserStats{}, err
	}
	return newStats, err
}
