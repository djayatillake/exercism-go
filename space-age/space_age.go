// package space provides functions around calculating someones age on different planets
package space

// declaring type planet which is the name of a planet as a string
type Planet string

// declaring the ratio of time for a planet's orbit vs Earth's orbit
var Planets = map[Planet]float64{
	"Mercury":   0.2408467,
	"Venus":     0.61519726,
	"Earth":     1.0,
	"EarthDays": 365.25,
	//"EarthSecs" : 31557600
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// the number of seconds in an Earth Year
const EarthSecs = 31557600

// Age calculates someones age on a given planet with input of seconds and planet
func Age(seconds float64, planet Planet) float64 {
	if _, ok := Planets[planet]; !ok {
		return 0
	}
	return seconds / EarthSecs / Planets[planet]
}
