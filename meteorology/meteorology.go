package main

import (
	s "exercism/meteorology/meteorology"
	"fmt"
)

var p = fmt.Println

func main() {
	//////////// TASK 1 ////////////
	celsiusUnit := s.Celsius
	fahrenheitUnit := s.Fahrenheit

	p(celsiusUnit.String())
	// => °C
	p(fahrenheitUnit.String())
	// => °F
	p(fmt.Sprint(celsiusUnit))
	// Output: °C

	//////////// TASK 2 ////////////
	celsiusTemp := s.Temperature{
		Degree: 21,
		Unit:   s.Celsius,
	}
	p(celsiusTemp.String())
	// => 21 °C
	p(fmt.Sprint(celsiusTemp))
	// Output: 21 °C

	fahrenheitTemp := s.Temperature{
		Degree: 75,
		Unit:   s.Fahrenheit,
	}
	p(fahrenheitTemp.String())
	// => 75 °F
	p(fmt.Sprint(fahrenheitTemp))
	// Output: 75 °F

	//////////// TASK 3 ////////////
	mphUnit := s.MilesPerHour
	p(mphUnit.String())
	// => mph
	p(fmt.Sprint(mphUnit))
	// Output: mph

	kmhUnit := s.KmPerHour
	p(kmhUnit.String())
	// => km/h
	p(fmt.Sprint(kmhUnit))
	// Output: km/h

	//////////// TASK 4 ////////////
	windSpeedNow := s.Speed{
		Magnitude: 18,
		Unit:      s.KmPerHour,
	}
	p(windSpeedNow.String())
	// => 18 km/h
	//p(fmt.Sprint(windSpeedNow))
	//p(fmt.Sprintf(windSpeedNow))
	p(fmt.Sprint(windSpeedNow))
	// Output: 18 km/h

	windSpeedYesterday := s.Speed{
		Magnitude: 22,
		Unit:      s.MilesPerHour,
	}
	p(windSpeedYesterday.String())
	// => 22 mph
	p(fmt.Sprint(windSpeedYesterday))
	// Output: 22 mph

	//////////// TASK 5 ////////////
	sfData := s.MeteorologyData{
		Location: "San Francisco",
		Temperature: s.Temperature{
			Degree: 57,
			Unit:   s.Fahrenheit,
		},
		WindDirection: "NW",
		WindSpeed: s.Speed{
			Magnitude: 19,
			Unit:      s.MilesPerHour,
		},
		Humidity: 60,
	}
	p(sfData.String())
}
