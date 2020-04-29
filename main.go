package main

import (
	"fmt"

	tm "github.com/buger/goterm"
	"github.com/manavellam/ChessBoard-with-Golang/board"
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
	tm.MoveCursor(0, 0)
	tm.Clear()
	tm.Flush()
	for {
		tm.MoveCursor(0, 0)
		tm.Clear()
		tm.Flush()
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
					break
				}
			} else {
				fmt.Print("Not a valid move \n")
			}
		}
		fmt.Println("Press Enter to continue")
		fmt.Scanln()
		tm.Flush()
		tm.Clear()
		tm.Flush()
	}
}
