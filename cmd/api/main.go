package main

import (
	"log"
	"net/http"
	"time"

	"github.com/DigvijayWagh22/Olx-Api/internal/config"
	"github.com/DigvijayWagh22/Olx-Api/internal/handlers"
)

func main() {

	cfg := config.MustLoad()
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handlers.Health)

	srv := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed : %v", err)
	}

}
