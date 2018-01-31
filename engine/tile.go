package engine

type Square struct {
	Tile Tile
	Piece *Piece
}
type Tile uint8

const (
	EmptyTile Tile = iota
	WallTile
	DeadlyTile
)
