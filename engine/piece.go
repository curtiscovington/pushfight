package engine

type Piece struct {
	White  bool
	Square bool
	x      int
	y      int
}

func NewWhiteSquarePiece() *Piece{
	return &Piece{White: true, Square: true}
}
func NewWhiteRoundPiece() *Piece{
	return &Piece{White: true, Square: false}
}
func NewBlackSquarePiece() *Piece{
	return &Piece{White: false, Square: true}
}
func NewBlackRoundPiece() *Piece{
	return &Piece{White: false, Square: false}
}

func (p *Piece) PlacePiece(x, y int) {
	p.x = x
	p.y = y
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
		if (cl.Positions[i].X == x && cl.Positions[i].Y == y) {
			return true;
		}
	}
	return false
}

func checkSquare(g *Game, p *Piece, sq *Square, cl *CheckList,cX, cY, dX, dY int) bool {
	if cl.HasBeenChecked(cX, cY) {
		return false
	}

	cl.add(cX, cY)

	if (cX == dX) && cY == dY {
		return true
	}

	if sq.Piece != nil {
		if p.x != cX && p.y != cY {
			return false
		}
	}

	if sq.Tile != EmptyTile {
		return false
	}
	
	var newX int;
	var newY int;

	

	newY = cY - 1;
	if newY >= 0 {
		
		if checkSquare(g, p, &g.Board[newY][cX], cl, cX, newY, dX, dY) {
			return true
		}
	}

	newY = cY + 1;
	if newY < len(g.Board) {
		
		if checkSquare(g, p, &g.Board[newY][cX], cl, cX, newY, dX, dY) {
			return true
		}
	}

	newX = cX + 1;
	if newX < len(g.Board[cY]) {
		
		if checkSquare(g, p, &g.Board[cY][newX], cl, newX, cY, dX, dY) {
			return true
		}
	}

	newX = cX - 1;
	if newX > 0 {
		
		if checkSquare(g, p, &g.Board[cY][newX], cl, newX, cY, dX, dY) {
			return true
		}
	}

	return false
}