package engine

type Game struct {
	Board     Board
	Pieces    []Piece
	AnchorX   int
	AnchorY   int
	DeadPiece *Piece
}

func (g Game) GetPiece(x, y int) *Piece {
	return g.Board[y][x].Piece
}

func (g *Game) PlacePiece(p *Piece, x, y int) {
	p.PlacePiece(x, y)
	g.Board[y][x].Piece = p
}

func NewGame() Game {
	g := Game{
		Board:   DefaultBoard(),
		Pieces:  make([]Piece, 10),
		AnchorX: -1,
		AnchorY: -1,
	}
	// White
	g.Pieces[0] = *NewWhiteSquarePiece()
	g.PlacePiece(&g.Pieces[0], 1, 4)
	g.Pieces[1] = *NewWhiteSquarePiece()
	g.PlacePiece(&g.Pieces[1], 2, 4)
	g.Pieces[2] = *NewWhiteSquarePiece()
	g.PlacePiece(&g.Pieces[2], 3, 4)
	g.Pieces[3] = *NewWhiteRoundPiece()
	g.PlacePiece(&g.Pieces[3], 4, 4)
	g.Pieces[4] = *NewWhiteRoundPiece()
	g.PlacePiece(&g.Pieces[4], 2, 3)
	// Black
	g.Pieces[5] = *NewBlackSquarePiece()
	g.PlacePiece(&g.Pieces[5], 1, 5)
	g.Pieces[6] = *NewBlackSquarePiece()
	g.PlacePiece(&g.Pieces[6], 2, 5)
	g.Pieces[7] = *NewBlackSquarePiece()
	g.PlacePiece(&g.Pieces[7], 3, 5)
	g.Pieces[8] = *NewBlackRoundPiece()
	g.PlacePiece(&g.Pieces[8], 4, 5)
	g.Pieces[9] = *NewBlackRoundPiece()
	g.PlacePiece(&g.Pieces[9], 3, 6)
	return g
}
