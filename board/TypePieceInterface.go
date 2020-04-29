package board

//Piece interface implements implicitly for each struct with the following methods.
type Piece interface {
	Move(pos string, newPos string, b *Board)
	MoveIfValid(pos string, newPos string, B *Board) bool
	IsPiece() string
	IsUnicode() string
	White() bool
}
