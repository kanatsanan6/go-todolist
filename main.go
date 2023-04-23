package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kanatsanan6/go-todo-list/api"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load env: %v", err)
	}

	server, err := api.NewServer()
	if err != nil {
		log.Fatalf("failed to initialize a server: %v", err)
	}

	if err := server.Start("8080"); err != nil {
		log.Fatalf("failed to start a server: %v", err)
	}
}
