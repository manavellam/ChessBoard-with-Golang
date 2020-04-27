package pieces

//King represents the king
type King struct {
	IsWhite bool
}

//Move moves the piece to its new location if possible
func (p *King) Move(pos string, newPos string) {

}

//IsMoveValid validates current moove
func (p *King) IsMoveValid(pos string, newPos string) (bool, string) {
	//Piece in the way
	//Piece of the same color

	return true, ""
}

//IsPiece returns the kind and colour of the piece
func (p *King) IsPiece() string {
	if p.IsWhite {
		return "White King"
	}
	return "Black King"
}

//IsUnicode prints unicode carachter
func (p *King) IsUnicode() string {
	if p.IsWhite {
		return `♔`
	}
	return `♚`
}
