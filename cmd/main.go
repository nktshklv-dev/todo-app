package main

import (
	"github.com/spf13/viper"
	"log"
	todo "todo-app"
	"todo-app/pkg/handler"
	"todo-app/pkg/repository"
	"todo-app/pkg/service"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config: %v", err)
	}

	repo := repository.NewRepository()
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)
	//handler -> service -> repository

	srv := new(todo.Server)

	if err := srv.Run((viper.GetString("port")), handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
