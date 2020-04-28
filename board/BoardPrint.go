package board

import "fmt"

//PrintBoard displays board on the screen
func (b Board) PrintBoard() {
	var num = []int{0, 8, 7, 6, 5, 4, 3, 2, 1}
	var let = []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	var layout string
	for i, n := range num {
		//fmt.Printf("%v ", n)
		layout = layout + fmt.Sprintf("%v ", n)
		for j, l := range let {
			if i != 0 {
				if b.Tiles[fmt.Sprint(n, l)].IsWhite {
					layout = layout + fmt.Sprint("[")
					if b.Tiles[fmt.Sprint(n, l)].Occupied {
						layout = layout + fmt.Sprintf(`%v`, b.Tiles[fmt.Sprint(n, l)].Piece.IsUnicode())
					} else {
						layout = layout + fmt.Sprint(" ")
					}
					layout = layout + fmt.Sprint("]")
				} else {
					layout = layout + fmt.Sprint(" ")
					if b.Tiles[fmt.Sprint(n, l)].Occupied {
						layout = layout + fmt.Sprintf(`%v`, b.Tiles[fmt.Sprint(n, l)].Piece.IsUnicode())
					} else {
						layout = layout + fmt.Sprint(" ")
					}
					layout = layout + fmt.Sprint(" ")
				}
			} else {
				//Prints the header with the letters in the first cycle
				layout = layout + fmt.Sprintf(" %v ", let[j])
			}
		}
		layout = layout + fmt.Sprint("\n")
	}
	fmt.Print(layout)
}
