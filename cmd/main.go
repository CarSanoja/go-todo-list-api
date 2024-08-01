package main

import (
	"go-todo-list-api/config"
	"go-todo-list-api/handlers"
	"log"
)

func main() {
	config.LoadConfig()
	if err := handlers.LoadToDos(); err != nil {
		log.Fatalf("Failed to load todos: %v", err)
	}

	r := handlers.SetupRouter()
	log.Printf("Server running at port %s", config.GetConfig().Port)
	if err := r.Run(":" + config.GetConfig().Port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
