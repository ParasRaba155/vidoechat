package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	flag.Parse()

	r := mux.NewRouter()
	hub := newHub()

	go hub.run()

	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWebSocket(hub, w, r)
	})

	server := http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	log.Printf("starting the server on address %s", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Printf("error in serving the server: %s", err)
	}
}
