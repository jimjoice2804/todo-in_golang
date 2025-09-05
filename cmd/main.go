package main

import (
	"fmt"
	"log"
	"net/http"
	"todo-app/internal/config"
	"todo-app/internal/handlers"
)

func main() {

	cfg := config.MustLoad()
	handler := handlers.HomeHandler

	mux := http.NewServeMux()
	mux.HandleFunc("/api/home", handler)

	fmt.Println(cfg)

	if err := http.ListenAndServe(cfg.HTTP.Address, mux); err != nil {
		log.Fatalf("Error starting a server %v:", err)
	}
}
