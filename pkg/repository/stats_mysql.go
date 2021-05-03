package repository

import (
	"database/sql"
	"github.com/bogdanserdinov/tic-tac-toe-web"
)

type StatsRepository struct {
	db *sql.DB
}

func NewStatsRepository(db *sql.DB) *StatsRepository{
	return &StatsRepository{
		db : db,
	}
}

func (s *StatsRepository) GetStats(user tictactoe_web.User) (tictactoe_web.UserStats,error){
	return tictactoe_web.UserStats{},nil
}