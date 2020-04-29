package board

import (
	"fmt"
	"strconv"
)

//Queen represents a queen
type Queen struct {
	IsWhite bool
}

//Move moves the piece to its new location if possible
func (p *Queen) Move(pos string, newPos string, b *Board) {
	b.Tiles[newPos].Piece, b.Tiles[pos].Piece = b.Tiles[pos].Piece, b.Tiles[newPos].Piece
	b.Tiles[newPos].Occupied = true
	b.Tiles[pos].Occupied = false
}

//White tells if piece is white
func (p *Queen) White() bool {
	if p.IsWhite {
		return true
	}
	return false
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

//MoveIfValid move the piece if valid
func (p *Queen) MoveIfValid(pos string, newPos string, b *Board) bool {
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

	if x1 == x2 {
		if y1 < y2 {
			for i := y1 + 1; i <= y2; i++ {
				if !b.Tiles[fmt.Sprint(x1, ntol[i])].Occupied {
					continue
				} else if b.Tiles[fmt.Sprint(x1, ntol[i])].Piece.White() != p.IsWhite {
					p.Move(pos, fmt.Sprint(x1, ntol[i]), b)
					fmt.Printf("%v captured in: %v%v\n", b.Tiles[pos].Piece.IsPiece(), x1, ntol[i])
					return true
				} else {
					fmt.Print("Move not valid: There is a piece of the same color in the way")
					return false
				}
			}
			b.Tiles[newPos].Piece, b.Tiles[pos].Piece = b.Tiles[pos].Piece, b.Tiles[newPos].Piece
			b.Tiles[newPos].Occupied = true
			b.Tiles[pos].Occupied = false

			return true
		}
		for i := y1 - 1; i >= y2; i-- {
			if !b.Tiles[fmt.Sprint(x1, ntol[i])].Occupied {
				continue
			} else if b.Tiles[fmt.Sprint(x1, ntol[i])].Piece.White() != p.IsWhite {
				p.Move(pos, fmt.Sprint(x1, ntol[i]), b)
				fmt.Printf("%v captured in: %v%v\n", b.Tiles[pos].Piece.IsPiece(), x1, ntol[i])
				return true
			} else {
				fmt.Print("Move not valid: There is a piece of the same color in the way\n")
				return false
			}
		}
		b.Tiles[newPos].Piece, b.Tiles[pos].Piece = b.Tiles[pos].Piece, b.Tiles[newPos].Piece
		b.Tiles[newPos].Occupied = true
		b.Tiles[pos].Occupied = false

		return true
	} else if y1 == y2 {
		if x1 < x2 {
			for i := x1 + 1; i <= x2; i++ {
				if !b.Tiles[fmt.Sprint(i, ntol[y1])].Occupied {
					continue
				} else if b.Tiles[fmt.Sprint(i, ntol[y1])].Piece.White() != p.IsWhite {
					p.Move(pos, fmt.Sprint(i, ntol[y1]), b)
					fmt.Printf("%v captured in: %v%v\n", b.Tiles[pos].Piece.IsPiece(), i, ntol[y1])
					return true
				} else {
					fmt.Print("Move not valid: There is a piece of the same color in the way\n")
					return false
				}

			}
			p.Move(pos, newPos, b)
			return true
		}
		for i := x1 - 1; i >= x2; i-- {
			if !b.Tiles[fmt.Sprint(i, ntol[y1])].Occupied {
				continue
			} else if b.Tiles[fmt.Sprint(i, ntol[y1])].Piece.White() != p.IsWhite {
				p.Move(pos, fmt.Sprint(i, ntol[y1]), b)
				fmt.Printf("%v captured in: %v%v\n", b.Tiles[pos].Piece.IsPiece(), i, ntol[y1])
				return true
			} else {
				fmt.Print("Move not valid: There is a piece of the same color in the way\n")
				return false
			}
		}
		p.Move(pos, newPos, b)
		return true
	} else if (x1-x2)*(x1-x2) == (y1-y2)*(y1-y2) {
		if x1 < x2 {
			for i := x1 + 1; i <= x2; i++ {
				if y1 < y2 {
					y1++
					if !b.Tiles[fmt.Sprint(i, ntol[y1])].Occupied {
						continue
					} else if b.Tiles[fmt.Sprint(i, ntol[y1])].Piece.White() != p.IsWhite {
						p.Move(pos, fmt.Sprint(i, ntol[y1]), b)
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
						p.Move(pos, fmt.Sprint(i, ntol[y1]), b)
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
						p.Move(pos, fmt.Sprint(i, ntol[y1]), b)
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
						p.Move(pos, fmt.Sprint(i, ntol[y1]), b)
						fmt.Printf("%v captured in: %v%v\n", b.Tiles[pos].Piece.IsPiece(), i, ntol[y1])
						return true
					} else {
						fmt.Print("Move not valid: There is a piece of the same color in the way\n")
						return false
					}
				}
			}
		}
		p.Move(pos, newPos, b)
		return true
	}
	fmt.Println("Move not valid. Queen must move as horse, bishop or tower.")
	return false
}
