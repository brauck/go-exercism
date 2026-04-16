package main

import (
	b "exercism/blackjack/turns"
	"fmt"
)

var p = fmt.Println

func main() {
	p(b.ParseCard("ace"))

	p(b.FirstTurn("ace", "ace", "jack") == "P")
	p(b.FirstTurn("ace", "king", "ace") == "S")
	p(b.FirstTurn("five", "queen", "ace") == "H")
}
