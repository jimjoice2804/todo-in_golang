package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"todo-app/internal/config"
	"todo-app/internal/handlers"
)

func main() {

	cfg := config.MustLoad()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/home", handlers.HomeHandler)
	mux.HandleFunc("/api/slow", handlers.SlowHandler)
	//CRUD todo routes
	//create todo
	//get todo
	//update todo
	//delete todo

	fmt.Println(cfg)

	server := &http.Server{
		Addr:    cfg.HTTP.Address,
		Handler: mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting a server %v:", err)
		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	fmt.Println("\n⚡ Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.HTTP.ShutdownTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("❌ Forced shutdown:", err)
	}

	fmt.Println("✅ Server stopped cleanly")
}
