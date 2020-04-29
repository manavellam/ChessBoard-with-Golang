package board

import (
	"fmt"
	"strconv"
)

//Pawn represents a single pawn. This struct allow us to create methods for this structure.
type Pawn struct {
	IsWhite bool
}

//Move try to move the pawn to its new position
func (p *Pawn) Move(pos string, newPos string, b *Board) {
	b.Tiles[newPos].Piece, b.Tiles[pos].Piece = b.Tiles[pos].Piece, b.Tiles[newPos].Piece
	b.Tiles[newPos].Occupied = true
	b.Tiles[pos].Occupied = false
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
	if !p.IsWhite {
		return `♙`
	}
	return `♟`
}

//White tells if piece is white
func (p *Pawn) White() bool {
	if p.IsWhite {
		return true
	}
	return false
}

//MoveIfValid move the piece if valid
func (p *Pawn) MoveIfValid(pos string, newPos string, b *Board) bool {
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

	if p.IsWhite {
		if x2-x1 > 0 {
			if y1 == y2 {
				if !b.Tiles[newPos].Occupied {
					p.Move(pos, newPos, b)
					return true
				}
				fmt.Print("Move is not valid: Tile is occupied\n")

			} else {
				if (y2-y1 == 1 || y2-y1 == -1) && b.Tiles[newPos].Occupied && !b.Tiles[newPos].Piece.White() {
					p.Move(pos, newPos, b)
					fmt.Printf("%v captured in: %v\n", b.Tiles[pos].Piece.IsPiece(), newPos)
					return true
				}
				fmt.Print("Move is not valid: Only allowed to capture black pieces\n")

			}
		} else {
			fmt.Print("Move is not valid: White Pawn can only move forward\n")
		}
	} else {
		if x2-x1 < 0 {
			if y1 == y2 {
				if !b.Tiles[newPos].Occupied {
					p.Move(pos, newPos, b)
					return true
				}
				fmt.Print("Move is not valid: Tile is occupied\n")

			} else {
				if (y2-y1 == 1 || y2-y1 == -1) && b.Tiles[newPos].Occupied && b.Tiles[newPos].Piece.White() {
					p.Move(pos, newPos, b)
					fmt.Printf("%v captured in: %v\n", b.Tiles[pos].Piece.IsPiece(), newPos)
					return true
				}
				fmt.Print("Move is not valid: Only allowed to capture white pieces\n")

			}
		} else {
			fmt.Print("Move is not valid: Black Pawn can only move forward\n")
		}
	}
	return false
}
