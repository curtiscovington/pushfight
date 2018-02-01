package main

import (
	"net/http"

	"github.com/curtiscovington/pushfight/server"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	var q server.Queue
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil)
		go func(conn *websocket.Conn) {
			q.Push(conn)
			if q.Len() > 1 {
				c1 := q.Pop()
				c2 := q.Pop()
				game := server.NewGame(c1, c2)
				go game.Start()
			}

		}(conn)
	})

	http.ListenAndServe(":3000", nil)

}
