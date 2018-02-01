package server

import (
	"github.com/curtiscovington/pushfight/engine"
	"github.com/gorilla/websocket"
)

type Game struct {
	conns     [2]*websocket.Conn
	isStarted bool
	turn      int
	engine    engine.Game
	whiteCh   chan []byte
	blackCh   chan []byte
}

func NewGame(c1, c2 *websocket.Conn) *Game {
	return &Game{
		conns:     [2]*websocket.Conn{c1, c2},
		isStarted: false,
		turn:      0,
		engine:    engine.NewGame(),
		whiteCh:   make(chan []byte),
		blackCh:   make(chan []byte),
	}
}

func (g *Game) Start() {
	go func() {
		for {
			_, msg, _ := g.conns[0].ReadMessage()
			g.whiteCh <- msg
		}
	}()
	go func() {
		for {
			_, msg, _ := g.conns[1].ReadMessage()
			g.blackCh <- msg
		}
	}()

	g.conns[0].WriteMessage(websocket.TextMessage, []byte("white"))
	g.conns[1].WriteMessage(websocket.TextMessage, []byte("black"))
	g.run()
}
func (g *Game) run() {
	for {
		select {
		case msg := <-g.whiteCh:
			coords := parseLine(msg)
			if len(coords) == 4 {
				p := g.engine.GetPiece(coords[0], coords[1])
				if p != nil && p.White {
					if p.CanMove(&g.engine, coords[2], coords[3]) {
						g.engine.Board[coords[1]][coords[0]].Piece = nil
						g.engine.PlacePiece(p, coords[2], coords[3])
						g.conns[0].WriteMessage(websocket.TextMessage, msg)
						g.conns[1].WriteMessage(websocket.TextMessage, msg)
						g.engine.DrawBoard()
					} else {
						g.conns[0].WriteMessage(websocket.TextMessage, []byte("0"))
					}
				} else {
					g.conns[0].WriteMessage(websocket.TextMessage, []byte("0"))
				}
			}
		case msg := <-g.blackCh:
			coords := parseLine(msg)
			if len(coords) == 4 {
				p := g.engine.GetPiece(coords[0], coords[1])
				if p != nil && !p.White {
					if p.CanMove(&g.engine, coords[2], coords[3]) {
						g.engine.Board[coords[1]][coords[0]].Piece = nil
						g.engine.PlacePiece(p, coords[2], coords[3])
						g.engine.DrawBoard()
						g.conns[1].WriteMessage(websocket.TextMessage, msg)
						g.conns[0].WriteMessage(websocket.TextMessage, msg)
					} else {
						g.conns[1].WriteMessage(websocket.TextMessage, []byte("0"))
					}
				} else {
					g.conns[1].WriteMessage(websocket.TextMessage, []byte("0"))
				}
			}
		}
	}
}

func parseLine(line []byte) []int {
	var coords []int
	coords = make([]int, 0)
	for _, r := range line {
		switch r {
		case '1':
			fallthrough
		case 'a':
			fallthrough
		case 'A':
			coords = append(coords, 1)
			break
		case '2':
			fallthrough
		case 'b':
			fallthrough
		case 'B':
			coords = append(coords, 2)
			break
		case '3':
			fallthrough
		case 'c':
			fallthrough
		case 'C':
			coords = append(coords, 3)
			break
		case '4':
			fallthrough
		case 'd':
			fallthrough
		case 'D':
			coords = append(coords, 4)
			break
		case '5':
			coords = append(coords, 5)
			break
		case '6':
			coords = append(coords, 6)
			break
		case '7':
			coords = append(coords, 7)
			break
		case '8':
			coords = append(coords, 8)
			break
		}
	}

	return coords
}
