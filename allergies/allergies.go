// package allergies contains a function to calculate allergies
package allergies

import (
	"math"
)

func Allergies(score uint) (ret []string) {
	subs := []string{
		"cats",
		"pollen",
		"chocolate",
		"tomatoes",
		"strawberries",
		"shellfish",
		"peanuts",
		"eggs",
	}
	score %= 256
	for i := 0; i < 8; i++ {
		base := uint(math.Pow(float64(2), float64(7-i)))
		if score/base == 1 {
			ret = append(ret, subs[i])
		}
		score %= base
	}
	return
}

func AllergicTo(score uint, response string) bool {
	for _, v := range Allergies(score) {
		if v == response {
			return true
		}
	}
	return false
}
