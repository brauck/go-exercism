package main

import (
	s "exercism/speed/speed"
	"fmt"
)

var p = fmt.Println

func main() {
	speed := 5
	batteryDrain := 2
	car := s.NewCar(speed, batteryDrain)
	p(car)

	distance := 100
	track := s.NewTrack(distance)
	p(track)

	/* speed = 5
	batteryDrain = 2
	car = s.NewCar(speed, batteryDrain)
	car = s.Drive(car) */
	p(car)

	fmt.Println(s.CanFinish(car, track))
}
