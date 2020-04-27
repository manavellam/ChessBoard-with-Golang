package board

import "fmt"

//PrintBoard displays board on the screen
func (b Board) PrintBoard() {
	var num = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	var let = []string{"a", "b", "c", "d", "e", "f", "g", "h"}

	for i, n := range num {
		fmt.Printf("%v ", i)
		for j, l := range let {
			if i != 0 {
				if b.Tiles[fmt.Sprint(n, l)].IsWhite {
					fmt.Print("[")
					if b.Tiles[fmt.Sprint(n, l)].Occupied {
						fmt.Printf(`%v`, b.Tiles[fmt.Sprint(n, l)].Piece.IsUnicode())
					} else {
						fmt.Print(" ")
					}
					fmt.Print("]")
				} else {
					fmt.Print(" ")
					if b.Tiles[fmt.Sprint(n, l)].Occupied {
						fmt.Printf(`%v`, b.Tiles[fmt.Sprint(n, l)].Piece.IsUnicode())
					} else {
						fmt.Print(" ")
					}
					fmt.Print(" ")
				}
			} else {
				//Prints the header with the letters in the first cycle
				fmt.Printf(" %v ", let[j])
			}
		}
		fmt.Print("\n")
	}

}
