package main

import (
	s "exercism/the_farm/thefarm"
	"fmt"
)

var p = fmt.Println

func main() {
	err := s.ValidateNumberOfCows(-5)
	p(err.Error())

	// p(s.DivideFood(s.FodderCalculator, 10))
}
