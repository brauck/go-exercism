package main

import (
	s "exercism/airport_robot/airportrobot"
	"fmt"
)

var p = fmt.Println

func main() {
	/* italian := s.Italian{
		Lang:     "Italian",
		Greeting: "Ciao ",
	} */

	/* portuguese := s.Portuguese {
		Lang: "Portuguese",
		Greeting: "Olá ",
	} */

	p(s.SayHello("Flora", s.Italian{}))
}
