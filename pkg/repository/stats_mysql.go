package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/bogdanserdinov/tic-tac-toe-web"
	"github.com/sirupsen/logrus"
)

type StatsRepository struct {
	db *sql.DB
}

func NewStatsRepository(db *sql.DB) *StatsRepository {
	return &StatsRepository{
		db: db,
	}
}

func (s *StatsRepository) CreateStats(stats tictactoe_web.UserStats) error {
	//insert
	if s.db == nil {
		logrus.Fatalf("db = nil")
		return errors.New("db = nil; could not open db")
	}
	query := fmt.Sprintf("Insert into %s (id,total,wins,draws,losses) values (?,?,?,?,?)", UserStatsTable)
	_, err := s.db.Exec(query, stats.ID, stats.TotalGames, stats.Wins, stats.Draws, stats.Losses)
	if err != nil {
		return err
	}
	return nil
}

func (s *StatsRepository) GetStats(id int) (tictactoe_web.UserStats, error) {
	//select
	var stats tictactoe_web.UserStats

	if s.db == nil {
		logrus.Fatalf("db = nil")
		return stats, errors.New("db = nil; could not open db")
	}

	query := fmt.Sprintf("select * from %s where id=?", UserStatsTable)
	rows := s.db.QueryRow(query, id)

	err := rows.Scan(&stats.ID, &stats.TotalGames, &stats.Wins, &stats.Draws, &stats.Losses)
	if err != nil {
		logrus.Errorf("could not get row from db: %s", err.Error())
		return stats, err
	}

	return stats, nil
}

func (s *StatsRepository) UpdateStats(newStats tictactoe_web.UserStats) (tictactoe_web.UserStats, error) {
	if s.db == nil {
		logrus.Fatalf("db = nil")
		return tictactoe_web.UserStats{}, errors.New("db = nil; could not open db")
	}
	query := fmt.Sprintf("Update %s  Set total=?,wins=?,draws=?,losses=? where id=?", UserStatsTable)
	_, err := s.db.Exec(query, newStats.TotalGames, newStats.Wins, newStats.Draws, newStats.Losses, newStats.ID)
	if err != nil {
		return tictactoe_web.UserStats{}, err
	}
	return newStats, nil
}
