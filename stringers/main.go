package main

import "fmt"

// Stringer is an interface for defining the string format of values.

// The interface consists of a single String method:

// type Stringer interface {
//     String() string
// }

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// implement Stringer interface
func (tu TemperatureUnit) String() string {
    units := []string{"°C", "°F"}
    return units[tu]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// implemenet Stringer interface
func (t Temperature) String() string {
    return fmt.Sprintf("%v %v", t.degree, t.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// implemenet Stringer interface
func (sp SpeedUnit) String() string {
    units := []string{"km/h", "mph"}
    return units[sp]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// implemenet Stringer interface
func (s Speed) String() string {
    return fmt.Sprintf("%v %v", s.magnitude, s.unit)
}


type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// implemenet Stringer interface
func (md MeteorologyData) String() string {
    return fmt.Sprintf("%v: %v, Wind %v at %v, %v%% Humidity", md.location, md.temperature, md.windDirection,  md.windSpeed, md.humidity )
}


func main() {

	md := MeteorologyData{
		location: "Philly",
		temperature: Temperature{
			degree: 32,
			unit: 1,
		},
		windDirection: "NE",
		windSpeed: Speed{
			magnitude: 10,
			unit: SpeedUnit(1),
		},
		humidity: 50,
	}

	fmt.Println(md)
	// should print: Philly: 32 °F, Wind NE at 10 mph, 50% Humidity
}
