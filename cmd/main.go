package main

import (
	"github.com/bogdanserdinov/tic-tac-toe-web"
	"github.com/bogdanserdinov/tic-tac-toe-web/pkg/handler"
	"github.com/bogdanserdinov/tic-tac-toe-web/pkg/repository"
	"github.com/bogdanserdinov/tic-tac-toe-web/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := InitConfigs(); err != nil{
		log.Fatalf("error initialazing configs %s",err.Error())
	}
	repos := repository.NewRepository()
	ser := service.NewService(repos)
	handle := handler.NewHandler(ser) //struct, but method InitRoutes == http.Handler

	server := new(tictactoe_web.Server)
	if err := server.Run(viper.GetString("port"), handle.InitRoutes()); err != nil {
		log.Fatalln("error occurred while running http server", err.Error())
	}
}

func InitConfigs() error{
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
