package board

import "Test/pieces"

//Tile represents a tile of the 64 tiles in the game
type Tile struct {
	IsWhite  bool
	Occupied bool
	Piece    pieces.Piece
}
