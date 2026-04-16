package speed

import "fmt"

// TODO: define the 'Car' type struct
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:      100,
		batteryDrain: batteryDrain,
		distance:     0,
		speed:        speed,
	}
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if car.batteryDrain > car.battery {
		return car
	}
	return Car{
		speed:        car.speed,
		batteryDrain: car.batteryDrain,
		battery:      car.battery - car.batteryDrain,
		distance:     car.speed + car.distance,
	}
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	distance := track.distance
	car.distance = car.battery / car.batteryDrain * car.speed
	fmt.Println(car.distance)
	if car.distance >= distance {
		return true
	}
	/* if car.battery > car.batteryDrain && track.distance >= car.distance {
		return true
	} */
	return false
}
