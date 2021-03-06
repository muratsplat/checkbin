package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/muratsplat/checkbin/service/auth/boot"
)

var (
	addr = flag.String("addr", ":8081", "Auth Service WebSocket Connection")
	max  = flag.Int("maxService", 10, "maximum number of service")
)

func main() {
	boot.Run()
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}
