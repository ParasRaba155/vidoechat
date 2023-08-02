package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func echo(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("could not upgrade the handler: %s", err)
		w.Write([]byte(`{"err":"could not upgrade"}`))
		w.WriteHeader(400)
		return
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		log.Printf("messageType: %#+v", messageType)
		log.Printf("message from client: %s", p)
		if err != nil {
			log.Printf("error in reading the message: %s", err)
			w.Write([]byte(`{"err":"could not read message"}`))
			w.WriteHeader(400)
			return
		}
		time.Sleep(500 * time.Millisecond)
		err = conn.WriteMessage(messageType, p)
		log.Printf("message to client: %s", p)
		if err != nil {
			log.Printf("error in writing the message: %s", err)
			w.Write([]byte(`{"err":"could not write message"}`))
			w.WriteHeader(400)
			return
		}
	}
}
