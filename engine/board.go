package engine

import (
	"fmt"
)

type Board [][]Square

func DefaultBoard() Board {
	rows := make([][]Square, 10)
	rows[0] = []Square{{Tile: DeadlyTile},{Tile: DeadlyTile},{Tile: DeadlyTile},{Tile: DeadlyTile},{Tile: DeadlyTile},{Tile: DeadlyTile}}
	rows[1] = []Square{{Tile: DeadlyTile},{Tile: DeadlyTile},{Tile: EmptyTile},{Tile: EmptyTile},{Tile: DeadlyTile},{Tile: DeadlyTile}}
	rows[2] = []Square{{Tile: WallTile},{Tile: EmptyTile},{Tile: EmptyTile},{Tile: EmptyTile},{Tile: DeadlyTile},{Tile: DeadlyTile}}
	rows[3] = []Square{{Tile: WallTile},{Tile: EmptyTile},{Tile: EmptyTile},{Tile: EmptyTile},{Tile: EmptyTile},{Tile: WallTile}}
	rows[4] = []Square{{Tile: WallTile},{Tile: EmptyTile},{Tile: EmptyTile},{Tile: EmptyTile},{Tile: EmptyTile},{Tile: WallTile}}
	rows[5] = []Square{{Tile: WallTile},{Tile: EmptyTile},{Tile: EmptyTile},{Tile: EmptyTile},{Tile: EmptyTile},{Tile: WallTile}}
	rows[6] = []Square{{Tile: WallTile},{Tile: EmptyTile},{Tile: EmptyTile},{Tile: EmptyTile},{Tile: EmptyTile},{Tile: WallTile}}
	rows[7] = []Square{{Tile: DeadlyTile},{Tile: DeadlyTile},{Tile: EmptyTile},{Tile: EmptyTile},{Tile: EmptyTile},{Tile: WallTile}}
	rows[8] = []Square{{Tile: DeadlyTile},{Tile: DeadlyTile},{Tile: EmptyTile},{Tile: EmptyTile},{Tile: DeadlyTile},{Tile: DeadlyTile}}
	rows[9] = []Square{{Tile: DeadlyTile},{Tile: DeadlyTile},{Tile: DeadlyTile},{Tile: DeadlyTile},{Tile: DeadlyTile},{Tile: DeadlyTile}}

	
	return rows
}

func (g Game) DrawBoard() {
	b := g.Board
	for i := range b {
		fmt.Printf("    -   -   -   -   -   -\n")
		for j := range b[i] {
			if i > 0 && i < 9 && j == 0 {
				fmt.Printf("%d", i)
			} else if j == 0 {
				fmt.Printf(" ")
			}
			fmt.Printf(" | ")
			
			if b[i][j].Piece != nil {
				if b[i][j].Piece.Square {
					fmt.Printf("x")
				} else {
					fmt.Printf("o")
				}
			} else {
				switch b[i][j].Tile {
				case EmptyTile:
					fmt.Printf(" ")
					break
				case WallTile:
					fmt.Printf("W")
					break
				case DeadlyTile:
					fmt.Printf("*")
				}
			}
			if j == len(b[i])-1 {
				fmt.Printf(" | ")
			}
		}
		fmt.Printf("\n")
		if i == len(b)-1 {
			fmt.Printf("    -   -   -   -   -   -\n")
			fmt.Printf("\n        A   B   C   D   \n")
		}
	}
}
