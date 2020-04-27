package board

import (
	"Test/pieces"
	"math/rand"
	"time"
)

//AllocatePieces allocates pieces randomly on the board
func (B *Board) AllocatePieces(indexes []string) {

	//Generates Aleatory numbers and shuffles the indexes
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	for i := range indexes {
		pos := r.Intn(len(indexes) - 1)
		indexes[i], indexes[pos] = indexes[pos], indexes[i]
	}

	iswhite := true
	//Now we alocate the pieces to each shuffled tile
	//Allocates pawn
	for i := 0; i < 16; i++ {
		p := new(pieces.Pawn)
		p.IsWhite = iswhite
		B.Tiles[indexes[i]].Piece = p
		B.Tiles[indexes[i]].Occupied = true
		iswhite = !iswhite
	}

	//Allocates horses
	for i := 16; i < 20; i++ {
		h := new(pieces.Horse)
		h.IsWhite = iswhite
		B.Tiles[indexes[i]].Piece = h
		B.Tiles[indexes[i]].Occupied = true
		iswhite = !iswhite
	}

	//Allocates queens
	for i := 20; i < 22; i++ {
		q := new(pieces.Queen)
		q.IsWhite = iswhite
		B.Tiles[indexes[i]].Piece = q
		B.Tiles[indexes[i]].Occupied = true
		iswhite = !iswhite
	}

	//Alocates kings
	for i := 22; i < 24; i++ {
		k := new(pieces.King)
		k.IsWhite = iswhite
		B.Tiles[indexes[i]].Piece = k
		B.Tiles[indexes[i]].Occupied = true
		iswhite = !iswhite
	}

	//Alocates Towers
	for i := 24; i < 28; i++ {
		t := new(pieces.Tower)
		t.IsWhite = iswhite
		B.Tiles[indexes[i]].Piece = t
		B.Tiles[indexes[i]].Occupied = true
		iswhite = !iswhite
	}

	//Allocates Bishops
	i := 28
	blanco := 0
	negro := 0
	whitebishopinwhite := false
	whitebishopinblack := false
	blackbishopinwhite := false
	blackbishopinblack := false

	for blanco+negro < 4 {
		if B.Tiles[indexes[i]].IsWhite && iswhite && !whitebishopinwhite {
			blanco++
			b := new(pieces.Bishop)
			b.IsWhite = iswhite
			B.Tiles[indexes[i]].Piece = b
			B.Tiles[indexes[i]].Occupied = true
			whitebishopinwhite = true
		} else if B.Tiles[indexes[i]].IsWhite && !iswhite && !blackbishopinwhite {
			blanco++
			b := new(pieces.Bishop)
			b.IsWhite = iswhite
			B.Tiles[indexes[i]].Piece = b
			B.Tiles[indexes[i]].Occupied = true
			blackbishopinwhite = true
		} else if !B.Tiles[indexes[i]].IsWhite && iswhite && !whitebishopinblack {
			negro++
			b := new(pieces.Bishop)
			b.IsWhite = iswhite
			B.Tiles[indexes[i]].Piece = b
			B.Tiles[indexes[i]].Occupied = true
			whitebishopinblack = true
		} else if !B.Tiles[indexes[i]].IsWhite && !iswhite && !blackbishopinblack {
			negro++
			b := new(pieces.Bishop)
			b.IsWhite = iswhite
			B.Tiles[indexes[i]].Piece = b
			B.Tiles[indexes[i]].Occupied = true
			blackbishopinblack = true
		}
		i++
		iswhite = !iswhite
	}

}
