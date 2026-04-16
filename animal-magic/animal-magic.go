package main

import (
	a "exercism/animal-magic/chance"
	"fmt"
)

var p = fmt.Println

func main() {
	p(a.RollADie())
	p(a.GenerateWandEnergy())
	p(a.ShuffleAnimals())
}
