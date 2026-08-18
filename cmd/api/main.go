package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/student-ankitpandit/go-server/internal/config"
	"github.com/student-ankitpandit/go-server/internal/db"
	"github.com/student-ankitpandit/go-server/internal/handlers"
)

func main() {
	cfg := config.MustLoad()
	_, err := db.Connect(cfg.DatabaseUrl)
	if err != nil {
		log.Fatalf("main.db.connect: %v", err)
	}

	fmt.Println("db connected successfully")
	fmt.Println("starting go server...")
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handlers.Health)

	srv := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	log.Printf("server is up and running on port %s", srv.Addr)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("sever failed to start: %v", err)
	}
}
