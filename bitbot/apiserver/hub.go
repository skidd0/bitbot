// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Copyright 2018 Bren Briggs. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file for this package namespace

package apiserver

import (
	"fmt"

	//"github.com/gorilla/websocket"
	"github.com/whyrusleeping/hellabot"
)

type Hub struct {
	// Connection to IRC server
	IRC *hbot.Bot

	// Registered clients.
	Clients map[*Client]bool

	// Outbound messages from the server.
	Outbound chan []byte

	// Inbound messages from clients
	Inbound chan []byte

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client
}

func NewHub(irc *hbot.Bot) *Hub {
	return &Hub{
		IRC:        irc,
		Outbound:   make(chan []byte),
		Inbound:    make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.send)
			}
		case message := <-h.Outbound:
			for client := range h.Clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.Clients, client)
				}
			}
		case message := <-h.Outbound:
			fmt.Println(message)
		}
	}
}
