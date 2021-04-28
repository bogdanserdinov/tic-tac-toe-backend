package tictactoe_web

type User struct {
	ID       int    `json:"ID"`
	Name     string `json:"Name" validate:"required"`
	Password string `json:"Password" validate:"required"`
	Role     string `json:"Role"`
}

type UserStats struct {
	ID         int
	TotalGames int
	Wins       int
	Draws      int
	Losses     int
}
