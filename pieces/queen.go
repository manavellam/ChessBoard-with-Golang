package pieces

//Queen represents a queen
type Queen struct {
	IsWhite bool
}

//Move moves the piece to its new location if possible
func (p *Queen) Move(pos string, newPos string) {

}

//IsMoveValid validates current moove
func (p *Queen) IsMoveValid(pos string, newPos string) (bool, string) {
	//Piece in the way
	//Piece of the same color

	return true, ""
}

//IsPiece returns the kind and colour of the piece
func (p *Queen) IsPiece() string {
	if p.IsWhite {
		return "White Queen"
	}
	return "Black Queen"
}

//IsUnicode prints unicode carachter
func (p *Queen) IsUnicode() string {
	if p.IsWhite {
		return `♕`
	}
	return `♛`
}
