package main

import (
	"log"
	todo "todo-app"
	"todo-app/pkg/handler"
	"todo-app/pkg/repository"
	"todo-app/pkg/service"
)

func main() {
	repo := repository.NewRepository()
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)

	if err := srv.Run(("8080"), handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
