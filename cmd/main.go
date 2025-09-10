package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
	todo "todo-app"
	"todo-app/pkg/handler"
	"todo-app/pkg/repository"
	"todo-app/pkg/service"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing config: %v", err)
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("error initializing DB: %v", err)
	}

	repo := repository.NewRepository(db)
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
