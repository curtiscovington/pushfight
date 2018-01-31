package main

import (
	"bufio"
	"fmt"
	"github.com/curtiscovington/pushfight/engine"
	"os"
)

func main() {

	b := engine.NewGame()
	b.DrawBoard()

	fmt.Println("Enter coordinates of the piece you wish to move and then the destination.")
	fmt.Println("ex: `a4 a3`")
	reader := bufio.NewReader(os.Stdin)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			panic(err)
		}
		coords := parseLine(line)
		p := b.GetPiece(coords[0], coords[1])
		if p == nil {
			fmt.Println("No piece available")
		} else if len(coords) < 4 {
			fmt.Println(coords, p)
		} else {
			if (p.CanMove(&b, coords[2], coords[3])) {
				b.Board[coords[1]][coords[0]].Piece = nil
				b.PlacePiece(p, coords[2], coords[3])
				b.DrawBoard()
			} else {
				fmt.Println("Invalid Move")
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
