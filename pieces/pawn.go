package pieces

//Pawn represents a single pawn. This struct allow us to create methods for this structure.
type Pawn struct {
	IsWhite bool
}

//Move try to move the pawn to its new position
func (p *Pawn) Move(pos string, newPos string) {

}

//IsMoveValid validates current moove
func (p *Pawn) IsMoveValid(pos string, newPos string) (bool, string) {
	//Piece in the way
	//Piece of the same color

	return true, ""
}

//IsPiece returns the kind and colour of the piece
func (p *Pawn) IsPiece() string {
	if p.IsWhite {
		return "White Pawn"
	}
	return "Black Pawn"
}

//IsUnicode prints unicode carachter
func (p *Pawn) IsUnicode() string {
	if p.IsWhite {
		return `♙`
	}
	return `♟`
}
