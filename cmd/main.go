package main

import (
	"log"
	todo "todo-app"
	"todo-app/pkg/handler"
)

func main() {
	srv := new(todo.Server)
	handlers := new(handler.Handler)

	if err := srv.Run(("8080"), handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
