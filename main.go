package main

import (
	"Test/board"
	"fmt"
)

func main() {
	var MyBoard board.Board
	var (
		indexes []string
		piece   string
		destino string
	)
	MyBoard.Initialize(&indexes)
	MyBoard.AllocatePieces(indexes)

	for {
		MyBoard.PrintBoard()
		for {
			fmt.Print("Choose Piece: ")
			fmt.Scanln(&piece)
			if tile, ok := MyBoard.Tiles[piece]; ok == true && (*tile).Occupied {
				break
			} else {
				fmt.Print("There is not a piece there, or tile does not exist \n")
			}
		}

		for {
			fmt.Print("Choose destiny: ")
			fmt.Scanln(&destino)
			if _, ok := MyBoard.Tiles[destino]; ok == true {
				fmt.Printf("Piece in %v to %v \n", piece, destino)
				break
			} else {
				fmt.Print("Not a valid move \n")
			}
		}
	}
}
