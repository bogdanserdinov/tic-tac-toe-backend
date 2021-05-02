package tictactoe_web

type User struct {
	ID       int    `json:"ID" db:"id"`
	Name     string `json:"Name" validate:"required" db:"name"`
	Password string `json:"Password" validate:"required" db:"password"`
	Role     string `json:"Role" db:"role"`
}

type UserStats struct {
	ID         int
	TotalGames int
	Wins       int
	Draws      int
	Losses     int
}
