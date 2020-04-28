package board

//Queen represents a queen
type Queen struct {
	IsWhite bool
}

//Move moves the piece to its new location if possible
func (p *Queen) Move(pos string, newPos string) {

}

//White tells if piece is white
func (p *Queen) White() bool {
	if p.IsWhite {
		return true
	}
	return false
}

//MoveIfValid move the piece if valid
func (p *Queen) MoveIfValid(pos string, newPos string, B *Board) bool {
	//Piece in the way
	//Piece of the same color

	return true
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
	if !p.IsWhite {
		return `♕`
	}
	return `♛`
}
