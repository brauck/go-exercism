package main

import (
	b "exercism/bird-watcher/birdwatcher"
	"fmt"
)

var p = fmt.Println

func main() {
	birdsPerDay := []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}
	p(b.TotalBirdCount(birdsPerDay))
	p(b.BirdsInWeek(birdsPerDay, 2))

	birdsPerDay2 := []int{3, 0, 5, 1, 0, 4, 1, 0, 3, 4, 3, 0, 8, 0}
	p(b.BirdsInWeek(birdsPerDay2, 1))

	birdsPerDay = []int{2, 5, 0, 7, 4, 1}
	p(b.FixBirdCountLog(birdsPerDay))

}
