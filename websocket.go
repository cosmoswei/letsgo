package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var updrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WebSocket(port int) {

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, err := updrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
		}

		for {
			messageType, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}

			fmt.Println("%s sent :%s \n", string(msg))

			if err := conn.WriteMessage(messageType, msg); err != nil {
				log.Println(err)
				return
			}

		}
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/websocket.html")
	})
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
