package board

//Horse represents a horse
type Horse struct {
	IsWhite bool
}

//Move moves the piece to its new location if possible
func (p *Horse) Move(pos string, newPos string) {

}

//White tells if piece is white
func (p *Horse) White() bool {
	if p.IsWhite {
		return true
	}
	return false
}

//MoveIfValid move the piece if valid
func (p *Horse) MoveIfValid(pos string, newPos string, B *Board) bool {
	//Piece in the way
	//Piece of the same color

	return true
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
	if !p.IsWhite {
		return `♘`
	}
	return `♞`
}
