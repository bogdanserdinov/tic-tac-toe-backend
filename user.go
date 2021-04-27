package tictactoe_web

type User struct {
	ID       int
	Name     string
	Password []byte
	Role     string
}

type UserStats struct {
	ID         int
	TotalGames int
	Wins       int
	Draws      int
	Losses     int
}
