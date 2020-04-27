package board

import (
	"fmt"
)

//CreateBoard allocates randomly the pieces on the board
func CreateBoard(Board Board, indexes *[]string) {
	var num = []int{1, 2, 3, 4, 5, 6, 7, 8}
	var let = []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	iswhite := false

	for _, num := range num {
		for _, let := range let {
			pos := fmt.Sprint(num, let)
			*indexes = append(*indexes, pos)
			Board.Tiles[pos] = new(Tile)
			Board.Tiles[pos].IsWhite = iswhite
			iswhite = !iswhite
		}
		iswhite = !iswhite
	}
}
