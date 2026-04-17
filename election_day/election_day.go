package main

import (
	e "exercism/election_day/electionday"
	"fmt"
)

var p = fmt.Println

func main() {
	var initialVotes int
	initialVotes = 2

	var counter *int
	counter = e.NewVoteCounter(initialVotes)
	p(counter == &initialVotes) // true
	p(*e.NewElectionResult("Peter", 3))
}
