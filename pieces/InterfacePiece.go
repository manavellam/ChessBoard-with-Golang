package pieces

//Piece interface implements implicitly for each struct with the following methods.
type Piece interface {
	Move(pos string, newPos string)
	IsMoveValid(pos string, newPos string) (bool, string)
	IsPiece() string
	IsUnicode() string
}
