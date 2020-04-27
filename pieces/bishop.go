package pieces

//Bishop represents a bishop
type Bishop struct {
	IsWhite bool
}

//Move moves the piece to its new location if possible
func (p *Bishop) Move(pos string, newPos string) {

}

//IsMoveValid validates current moove
func (p *Bishop) IsMoveValid(pos string, newPos string) (bool, string) {
	//Piece in the way
	//Piece of the same color

	return true, ""
}

//IsPiece returns the kind and colour of the piece
func (p *Bishop) IsPiece() string {
	if p.IsWhite {
		return "White Bishop"
	}
	return "Black Bishop"
}

//IsUnicode prints unicode carachter
func (p *Bishop) IsUnicode() string {
	if p.IsWhite {
		return `♗`
	}
	return `♝`
}
