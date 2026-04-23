package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type

type Temperature struct {
	Degree int
	Unit   TemperatureUnit
}

func (tm TemperatureUnit) String() string {
	units := []string{"°C", "°F"}
	return units[tm]
}

// Add a String method to the Temperature type
func (t Temperature) String() string {
	return fmt.Sprintf("%v %v", t.Degree, t.Unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (sp SpeedUnit) String() string {
	units := []string{"km/h", "mph"}
	return units[sp]
}

type Speed struct {
	Magnitude int
	Unit      SpeedUnit
	// magnitude int
	// unit      SpeedUnit
}

// Add a String method to Speed
func (s Speed) String() string {
	//return s.Magnitude.String() + " " + s.Unit.String()
	return fmt.Sprintf("%v %v", s.Magnitude, s.Unit)
	// return fmt.Sprintf("%v %v", s.magnitude, s.unit)
}

type MeteorologyData struct {
	Location      string
	Temperature   Temperature
	WindDirection string
	WindSpeed     Speed
	Humidity      int

	/* location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int */
}

// Add a String method to MeteorologyData
func (md MeteorologyData) String() string {
	return fmt.Sprintf("%v: %v, Wind %v at %v, %v%% Humidity",
		md.Location,
		md.Temperature,
		md.WindDirection,
		md.WindSpeed,
		md.Humidity)

	/* md.location,
	md.temperature,
	md.windDirection,
	md.windSpeed,
	md.humidity ) */
}
