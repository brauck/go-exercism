package main

import (
	j "exercism/jedliks_toys/jedlik"
	"fmt"
)

var p = fmt.Println

func main() {
	car := j.NewCar(5, 2)
	p(car.DisplayBattery())

}
