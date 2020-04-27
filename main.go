package main

import (
	"Test/board"
)

func main() {
	var MyBoard board.Board
	var indexes []string
	MyBoard.Initialize(&indexes)
	MyBoard.AllocatePieces(indexes)
	MyBoard.PrintBoard()
}
