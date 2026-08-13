package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/student-ankitpandit/go-server/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println("starting go server...")
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"all ok"}`))
	})

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
