package pieces

//Horse represents a horse
type Horse struct {
	IsWhite bool
}

//Move moves the piece to its new location if possible
func (p *Horse) Move(pos string, newPos string) {

}

//IsMoveValid validates current moove
func (p *Horse) IsMoveValid(pos string, newPos string) (bool, string) {
	//Piece in the way
	//Piece of the same color

	return true, ""
}

//IsPiece returns the kind and colour of the piece
func (p *Horse) IsPiece() string {
	if p.IsWhite {
		return "White Horse"
	}
	return "Black Horse"
}

//IsUnicode prints unicode carachter
func (p *Horse) IsUnicode() string {
	if p.IsWhite {
		return `♘`
	}
	return `♞`
}
