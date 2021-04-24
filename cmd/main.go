package main

import (
	"github.com/bogdanserdinov/tic-tac-toe-web"
	"github.com/bogdanserdinov/tic-tac-toe-web/pkg/handler"
	"github.com/bogdanserdinov/tic-tac-toe-web/pkg/repository"
	"github.com/bogdanserdinov/tic-tac-toe-web/pkg/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	if err := InitConfigs(); err != nil{
		log.Fatalf("error initialazing configs %s",err.Error())
	}

	if err := godotenv.Load(); err != nil{
		log.Fatalf("error initialazing environment file %s",err.Error())
	}

	db,err := repository.InitMySQL(repository.Config{
		Username: viper.GetString("db.name"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName: viper.GetString("db.dbname"),
	})
	if err != nil{
		log.Printf("error in initialization MySQL %s",err.Error())
	}

	repos := repository.NewRepository(db)
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
