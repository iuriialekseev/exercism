// Package space calculates person's age in years related to different planets
package space

type Planet string

var orbitalPeriods = map[Planet]float64{
	"Earth":   1.0,
	"Jupiter": 11.862615,
	"Mars":    1.8808158,
	"Mercury": 0.2408467,
	"Neptune": 164.79132,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Venus":   0.61519726,
}

const secondsInYear = 31557600.0

// Age calculates person's age in years from given seconds and planet
func Age(s float64, p Planet) float64 {
	return s / orbitalPeriods[p] / secondsInYear
}
