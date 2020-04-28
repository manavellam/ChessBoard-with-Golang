package board

import (
	"fmt"
	"strconv"
)

//Bishop represents a bishop
type Bishop struct {
	IsWhite bool
}

//Move moves the piece to its new location if possible
func (p *Bishop) Move(pos string, newPos string) {

}

//White tells if piece is white
func (p *Bishop) White() bool {
	if p.IsWhite {
		return true
	}
	return false
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
	if !p.IsWhite {
		return `♗`
	}
	return `♝`
}

//MoveIfValid move the piece if valid
func (p *Bishop) MoveIfValid(pos string, newPos string, b *Board) bool {
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
	if (x1-x2)*(x1-x2) == (y1-y2)*(y1-y2) {
		if x1 < x2 {
			for i := x1 + 1; i <= x2; i++ {
				if y1 < y2 {
					y1++
					if !b.Tiles[fmt.Sprint(i, ntol[y1])].Occupied {
						continue
					} else if b.Tiles[fmt.Sprint(i, ntol[y1])].Piece.White() != p.IsWhite {
						b.Tiles[fmt.Sprint(i, ntol[y1])].Piece, b.Tiles[pos].Piece = b.Tiles[pos].Piece, b.Tiles[fmt.Sprint(i, ntol[y1])].Piece
						b.Tiles[fmt.Sprint(i, ntol[y1])].Occupied = true
						b.Tiles[pos].Occupied = false
						fmt.Printf("%v captured in: %v%v\n", b.Tiles[pos].Piece.IsPiece(), i, ntol[y1])
						return true
					} else {
						fmt.Print("Move not valid: There is a piece of the same color in the way\n")
						return false
					}
				} else {
					y1--
					if !b.Tiles[fmt.Sprint(i, ntol[y1])].Occupied {
						continue
					} else if b.Tiles[fmt.Sprint(i, ntol[y1])].Piece.White() != p.IsWhite {
						b.Tiles[fmt.Sprint(i, ntol[y1])].Piece, b.Tiles[pos].Piece = b.Tiles[pos].Piece, b.Tiles[fmt.Sprint(i, ntol[y1])].Piece
						b.Tiles[fmt.Sprint(i, ntol[y1])].Occupied = true
						b.Tiles[pos].Occupied = false
						fmt.Printf("%v captured in: %v%v\n", b.Tiles[pos].Piece.IsPiece(), i, ntol[y1])
						return true
					} else {
						fmt.Print("Move not valid: There is a piece of the same color in the way\n")
						return false
					}
				}
			}

		} else {
			for i := x1 - 1; i >= x2; i-- {
				if y1 < y2 {
					y1++
					if !b.Tiles[fmt.Sprint(i, ntol[y1])].Occupied {
						continue
					} else if b.Tiles[fmt.Sprint(i, ntol[y1])].Piece.White() != p.IsWhite {
						b.Tiles[fmt.Sprint(i, ntol[y1])].Piece, b.Tiles[pos].Piece = b.Tiles[pos].Piece, b.Tiles[fmt.Sprint(i, ntol[y1])].Piece
						b.Tiles[fmt.Sprint(i, ntol[y1])].Occupied = true
						b.Tiles[pos].Occupied = false
						fmt.Printf("%v captured in: %v%v\n", b.Tiles[pos].Piece.IsPiece(), i, ntol[y1])
						return true
					} else {
						fmt.Print("Move not valid: There is a piece of the same color in the way\n")
						return false
					}
				} else {
					y1--
					if !b.Tiles[fmt.Sprint(i, ntol[y1])].Occupied {
						continue
					} else if b.Tiles[fmt.Sprint(i, ntol[y1])].Piece.White() != p.IsWhite {
						b.Tiles[fmt.Sprint(i, ntol[y1])].Piece, b.Tiles[pos].Piece = b.Tiles[pos].Piece, b.Tiles[fmt.Sprint(i, ntol[y1])].Piece
						b.Tiles[fmt.Sprint(i, ntol[y1])].Occupied = true
						b.Tiles[pos].Occupied = false
						fmt.Printf("%v captured in: %v%v\n", b.Tiles[pos].Piece.IsPiece(), i, ntol[y1])
						return true
					} else {
						fmt.Print("Move not valid: There is a piece of the same color in the way\n")
						return false
					}
				}
			}
		}
		b.Tiles[newPos].Piece, b.Tiles[pos].Piece = b.Tiles[pos].Piece, b.Tiles[newPos].Piece
		b.Tiles[newPos].Occupied = true
		b.Tiles[pos].Occupied = false
		return true
	}
	fmt.Print("Move is not valid. Bishop can only move in diagonal lines.")

	return false
}
