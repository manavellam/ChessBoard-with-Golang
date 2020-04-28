package board

import (
	"fmt"
	"strconv"
)

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

//MoveIfValid move the piece if valid
func (p *Horse) MoveIfValid(pos string, newPos string, b *Board) bool {
	var (
		//maps leters to numbers
		lton = map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8}
		//maps numbers to letters
		ntol = make(map[int]string)
	)

	for key, value := range lton {
		ntol[value] = key
	}
	x1, _ := strconv.Atoi(string(pos[0]))
	y1 := lton[string(pos[1])]
	x2, _ := strconv.Atoi(string(newPos[0]))
	y2 := lton[string(newPos[1])]

	if ((y1-y2)*(y1-y2) == 4 && (x1-x2)*(x1-x2) == 1) || ((x1-x2)*(x1-x2) == 4 && (y1-y2)*(y1-y2) == 1) {
		if !b.Tiles[newPos].Occupied {
			b.Tiles[newPos].Piece, b.Tiles[pos].Piece = b.Tiles[pos].Piece, b.Tiles[newPos].Piece
			b.Tiles[newPos].Occupied = true
			b.Tiles[pos].Occupied = false
			return true
		} else if b.Tiles[newPos].Piece.White() != p.IsWhite {
			b.Tiles[newPos].Piece, b.Tiles[pos].Piece = b.Tiles[pos].Piece, b.Tiles[newPos].Piece
			b.Tiles[newPos].Occupied = true
			b.Tiles[pos].Occupied = false
			fmt.Printf("%v captured in: %v\n", b.Tiles[pos].Piece.IsPiece(), newPos)
			return true
		} else {
			fmt.Println("Move not valid. There is a piece of the same color in  destiny")
			return false
		}
	}
	fmt.Println("Move not valid. Horse must move in shape of L (2x1 blocks) in any orientation.")

	return false
}
