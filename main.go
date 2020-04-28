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

	//Here the game begins

	for {
		//tm.MoveCursor(1, 1)
		//tm.Clear()
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
			fmt.Print("Choose destiny (destiny=origin to cancel): ")
			fmt.Scanln(&destino)
			if destino == piece {
				break
			}
			if _, ok := MyBoard.Tiles[destino]; ok == true {
				if MyBoard.Tiles[piece].Piece.MoveIfValid(piece, destino, &MyBoard) {
					//here check if move valid or not. Return info about pieces involved.
					break
				}
			} else {
				fmt.Print("Not a valid move \n")
			}
		}
		fmt.Print("Press Enter to continue")
		fmt.Scanln(&piece)
	}
}
