package pieces

//Tower represents a tower
type Tower struct {
	IsWhite bool
}

//Move try to move the pawn to its new position
func (p *Tower) Move(pos string, newPos string) {

}

//IsMoveValid validates current moove
func (p *Tower) IsMoveValid(pos string, newPos string) (bool, string) {
	//Piece in the way
	//Piece of the same color

	return true, ""
}

//IsPiece returns the kind and colour of the piece
func (p *Tower) IsPiece() string {
	if p.IsWhite {
		return "White Tower"
	}
	return "Black Tower"
}

//IsUnicode prints unicode carachter
func (p *Tower) IsUnicode() string {
	if p.IsWhite {
		return `♖`
	}
	return `♜`
}
