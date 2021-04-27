package tictactoe_web

type Game struct {
	ID           int
	UserIDFirst  int
	UserIDSecond int
}

type GameStats struct {
	ID     int
	GameID int
	UserID int
	Step   int
}
