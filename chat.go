package main

import (
	"log"
	"time"
	"slices"
	"net/http"
)

type Message struct {
	MessageId string `json:"message_id"`
	Content string `json:"content"`
	SenderId string `json:"sender_id"`
	ChatId string `json:"chat_id"`
	CreatedAt time.Time `json:"created_at"`
	ServerReceivedAt time.Time `json:"server_received_at"`
}

func Chat(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		return slices.Contains(VALID_ORIGINS, origin)
	}
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	for {
		var msg Message
		err := c.ReadJSON(&msg)
		if err != nil {
			log.Println("error while reading incoming message:", err)
		}
	}
}


func JoinRoom(w http.ResponseWriter, r *http.Request) {
	
}