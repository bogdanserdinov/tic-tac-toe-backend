package main

import (
	"github.com/bogdanserdinov/tic-tac-toe-web"
	"github.com/bogdanserdinov/tic-tac-toe-web/pkg/handler"
	"github.com/bogdanserdinov/tic-tac-toe-web/pkg/repository"
	"github.com/bogdanserdinov/tic-tac-toe-web/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	ser := service.NewService(repos)
	hndl := handler.NewHandler(ser) //struct, but method InitRoutes == http.Handler

	server := new(tictactoe_web.Server)
	if err := server.Run("8080", hndl.InitRoutes()); err != nil {
		log.Fatalln("error occurred while running http server", err.Error())
	}
}
