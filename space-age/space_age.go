package space

type Planet string

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

const EarthSecs = 31557600

func Age(seconds float64, planet Planet) float64 {
	if _, ok := Planets[planet]; !ok {
		return 0
	}
	return seconds / EarthSecs / Planets[planet]
}
