package main

import (
	"github.com/bogdanserdinov/tic-tac-toe-web"
	"github.com/bogdanserdinov/tic-tac-toe-web/pkg/handler"
	"log"
)

func main() {
	hndl := new(handler.Handler) //struct, but method InitRoutes == http.Handler

	server := new(tictactoe_web.Server)
	if err := server.Run("8080", hndl.InitRoutes()); err != nil {
		log.Fatalln("error occurred while running http server", err.Error())
	}
}
