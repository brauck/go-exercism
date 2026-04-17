package main

import (
	c "exercism/chessboard/chessboard"
	"fmt"
)

var p = fmt.Println

func main() {
	board := c.Chessboard{
		"A": c.File{true, false, true, false, false, false, true, false},
		"B": c.File{false, false, false, false, false, false, false, false},
	}

	p(c.CountInFile(board, "A")) // 3
	p(c.CountInFile(board, "B")) // 0
	p(c.CountInFile(board, "Z")) // 0
}
