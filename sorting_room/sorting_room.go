package main

import (
	s "exercism/sorting_room/sorting"
	"fmt"
)

var p = fmt.Println

func main() {
	p(s.DescribeNumber(-12.345))

	box := s.NumberBoxContaining{12}
	p(s.DescribeNumberBox(box))

	//fn := s.FancyNumber{"10"}
	p(s.ExtractFancyNumber(s.FancyNumber{"10"}))

	p(s.DescribeFancyNumberBox(s.FancyNumber{"10"}))
}
