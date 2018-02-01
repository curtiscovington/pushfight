package server

import "github.com/gorilla/websocket"

type Queue []*websocket.Conn

func (q Queue) Len() int { return len(q) }
func (q *Queue) Push(conn *websocket.Conn) {
	*q = append(*q, conn)
}

func (q *Queue) Pop() *websocket.Conn {
	n := len(*q)
	if n == 0 {
		return nil
	}

	conn := (*q)[0]
	*q = (*q)[1:]
	return conn
}
