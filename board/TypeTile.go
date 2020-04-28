package board

//Tile represents a tile of the 64 tiles in the game
type Tile struct {
	IsWhite  bool
	Occupied bool
	Piece    Piece
}
