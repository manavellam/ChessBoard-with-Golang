package board

import "fmt"

//Initialize creates a new board
func (B *Board) Initialize(indexes *[]string) {
	B.Tiles = make(map[string]*Tile)
	var num = []int{1, 2, 3, 4, 5, 6, 7, 8}
	var let = []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	iswhite := false

	for _, num := range num {
		for _, let := range let {
			pos := fmt.Sprint(num, let)
			*indexes = append(*indexes, pos)
			B.Tiles[pos] = new(Tile)
			B.Tiles[pos].IsWhite = iswhite
			iswhite = !iswhite
		}
		iswhite = !iswhite
	}
}
