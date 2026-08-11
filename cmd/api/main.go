package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"all ok"}`))
	})

	srv := http.Server{
		Addr: ":3000",
		Handler: mux,
		ReadTimeout: time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout: time.Second * 60,
	}
	
	log.Println("server is up and running on port 3000")
	
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("sever failed to start: %v", err)	
	}
}