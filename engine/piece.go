package engine

import (
	"math"
)

type Piece struct {
	White  bool
	Square bool
	x      int
	y      int
}

type direction int

const (
	up direction = iota
	down
	left
	right
)

func NewWhiteSquarePiece() *Piece {
	return &Piece{White: true, Square: true}
}
func NewWhiteRoundPiece() *Piece {
	return &Piece{White: true, Square: false}
}
func NewBlackSquarePiece() *Piece {
	return &Piece{White: false, Square: true}
}
func NewBlackRoundPiece() *Piece {
	return &Piece{White: false, Square: false}
}

func (p *Piece) PlacePiece(x, y int) {
	p.x = x
	p.y = y
}

func (p Piece) CanPush(g *Game, x, y int) bool {
	if !p.Square || g.Board[y][x].Piece == nil {
		return false
	}

	if p.x == x && p.y == y {
		return false
	}
	dX := float64(x - p.x)
	dXSq := math.Pow(dX, 2)
	dY := float64(y - p.y)
	dYSq := math.Pow(dY, 2)
	D := math.Sqrt(dXSq + dYSq)
	if D == 0 || D > 1 {
		return false
	}

	if dX < 0 {
		return checkDir(g, left, p.x, p.y)
	} else if dX > 0 {
		return checkDir(g, right, p.x, p.y)
	} else if dY < 0 {
		return checkDir(g, up, p.x, p.y)
	} else if dY > 0 {
		return checkDir(g, down, p.x, p.y)
	}

	return false
}

func checkDir(g *Game, dir direction, x, y int) bool {
	switch dir {
	case down:
		if y+1 > len(g.Board)-1 {
			return false
		}
		if g.Board[y+1][x].Tile == WallTile || (y+1 == g.AnchorY && x == g.AnchorX) {
			return false
		}
		if g.Board[y+1][x].Piece != nil {
			if checkDir(g, down, x, y+1) {
				g.Board[y+1][x].Piece = g.Board[y][x].Piece
				g.Board[y+1][x].Piece.PlacePiece(x, y+1)
				g.Board[y][x].Piece = nil
				return true
			}
		} else if g.Board[y+1][x].Tile == EmptyTile || g.Board[y+1][x].Tile == DeadlyTile {
			if g.Board[y+1][x].Tile == DeadlyTile {
				g.DeadPiece = g.Board[y][x].Piece
			}
			g.Board[y+1][x].Piece = g.Board[y][x].Piece
			g.Board[y+1][x].Piece.PlacePiece(x, y+1)
			g.Board[y][x].Piece = nil
			return true
		}
	case up:
		if y-1 < 0 {
			return false
		}
		if g.Board[y-1][x].Tile == WallTile || (y-1 == g.AnchorY && x == g.AnchorX) {
			return false
		}
		if g.Board[y-1][x].Piece != nil {
			if checkDir(g, up, x, y-1) {
				g.Board[y-1][x].Piece = g.Board[y][x].Piece
				g.Board[y-1][x].Piece.PlacePiece(x, y-1)
				g.Board[y][x].Piece = nil
				return true
			}
		} else if g.Board[y-1][x].Tile == EmptyTile || g.Board[y-1][x].Tile == DeadlyTile {

			if g.Board[y-1][x].Tile == DeadlyTile {
				g.DeadPiece = g.Board[y][x].Piece
			}
			g.Board[y-1][x].Piece = g.Board[y][x].Piece
			g.Board[y-1][x].Piece.PlacePiece(x, y-1)
			g.Board[y][x].Piece = nil
			return true
		}
	case left:
		if x-1 < 1 {
			return false
		}
		if g.Board[y][x-1].Tile == WallTile || (y == g.AnchorY && x-1 == g.AnchorX) {
			return false
		}
		if g.Board[y][x-1].Piece != nil {
			if checkDir(g, left, x-1, y) {
				g.Board[y][x-1].Piece = g.Board[y][x].Piece
				g.Board[y][x-1].Piece.PlacePiece(x-1, y)
				g.Board[y][x].Piece = nil
				return true
			}
		} else if g.Board[y][x-1].Tile == EmptyTile || g.Board[y][x-1].Tile == DeadlyTile {
			if g.Board[y][x-1].Tile == DeadlyTile {
				g.DeadPiece = g.Board[y][x].Piece
			}
			g.Board[y][x-1].Piece = g.Board[y][x].Piece
			g.Board[y][x-1].Piece.PlacePiece(x-1, y)
			g.Board[y][x].Piece = nil
			return true
		}
	case right:
		if x+1 > len(g.Board)-2 {
			return false
		}

		if g.Board[y][x+1].Tile == WallTile || (y == g.AnchorY && x+1 == g.AnchorX) {
			return false
		}
		if g.Board[y][x+1].Piece != nil {
			if checkDir(g, right, x+1, y) {
				g.Board[y][x+1].Piece = g.Board[y][x].Piece
				g.Board[y][x+1].Piece.PlacePiece(x+1, y)
				g.Board[y][x].Piece = nil
				return true
			}
		} else if g.Board[y][x+1].Tile == EmptyTile || g.Board[y][x+1].Tile == DeadlyTile {
			if g.Board[y][x+1].Tile == DeadlyTile {
				g.DeadPiece = g.Board[y][x].Piece
			}

			g.Board[y][x+1].Piece = g.Board[y][x].Piece
			g.Board[y][x+1].Piece.PlacePiece(x+1, y)
			g.Board[y][x].Piece = nil
			return true
		}
	}
	return false
}

func (p Piece) CanMove(g *Game, x, y int) bool {
	start := g.Board[p.y][p.x]

	dest := g.Board[y][x]
	if dest.Piece != nil {
		return false
	}
	if dest.Tile != EmptyTile {
		return false
	}

	stX := p.x
	stY := p.y

	cl := CheckList{}
	return checkSquare(g, &p, &start, &cl, stX, stY, x, y)

}

type position struct {
	X, Y int
}
type CheckList struct {
	Positions []position
}

func (cl *CheckList) add(x, y int) {
	cl.Positions = append(cl.Positions, position{x, y})
}

func (cl CheckList) HasBeenChecked(x, y int) bool {
	for i := range cl.Positions {
		if cl.Positions[i].X == x && cl.Positions[i].Y == y {
			return true
		}
	}
	return false
}

func checkSquare(g *Game, p *Piece, sq *Square, cl *CheckList, cX, cY, dX, dY int) bool {
	if cl.HasBeenChecked(cX, cY) {
		return false
	}

	cl.add(cX, cY)

	if (cX == dX) && cY == dY {
		return true
	}

	if sq.Piece != nil && len(cl.Positions) > 1 {
		return false
	}

	if sq.Tile != EmptyTile {
		return false
	}

	var newX int
	var newY int

	newY = cY - 1
	if newY >= 0 {

		if checkSquare(g, p, &g.Board[newY][cX], cl, cX, newY, dX, dY) {
			return true
		}
	}

	newY = cY + 1
	if newY < len(g.Board) {

		if checkSquare(g, p, &g.Board[newY][cX], cl, cX, newY, dX, dY) {
			return true
		}
	}

	newX = cX + 1
	if newX < len(g.Board[cY]) {

		if checkSquare(g, p, &g.Board[cY][newX], cl, newX, cY, dX, dY) {
			return true
		}
	}

	newX = cX - 1
	if newX > 0 {

		if checkSquare(g, p, &g.Board[cY][newX], cl, newX, cY, dX, dY) {
			return true
		}
	}

	return false
}
