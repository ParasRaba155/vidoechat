package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var env = flag.String("env", "dev", "the environment of the application")

func main() {
	flag.Parse()

	r := mux.NewRouter()
	r.HandleFunc("/", echo)
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
