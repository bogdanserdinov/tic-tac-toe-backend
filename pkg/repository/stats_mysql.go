package repository

import (
	"github.com/bogdanserdinov/tic-tac-toe-web"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type StatsRepository struct {
	db *sqlx.DB
}

func NewStatsRepository(db *sqlx.DB) *StatsRepository {
	return &StatsRepository{
		db: db,
	}
}


func (s *StatsRepository) GetStats(id int) (tictactoe_web.UserStats, error) {
	//select
	var stats tictactoe_web.UserStats

	err := s.db.Get(&stats, "select * from userstats where id=$1", id)
	if err != nil {
		logrus.Errorf("could not get row from db: %s", err.Error())
		return stats, err
	}

	return stats, nil
}

func (s *StatsRepository) UpdateStats(newStats tictactoe_web.UserStats) (tictactoe_web.UserStats, error) {
	_, err := s.db.Exec("Update userstats  Set total=$1,wins=$2,draws=$3,losses=$4 where id=$5",
		newStats.TotalGames, newStats.Wins, newStats.Draws, newStats.Losses, newStats.ID)
	if err != nil {
		return tictactoe_web.UserStats{}, err
	}
	return newStats, nil
}
