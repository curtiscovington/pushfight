package main

import (
	"fmt"
	"github.com/curtiscovington/pushfight/engine"
)

func main() {

	b := engine.NewGame()
	b.DrawBoard()
	p := b.GetPiece(2,3)
	fmt.Println(p.CanMove(&b, 1,2))
}
