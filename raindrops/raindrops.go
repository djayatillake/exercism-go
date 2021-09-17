// package raindrops has functions to convert raindrops numbers to string resp
package raindrops

import "fmt"

// Convert checks raindrops input integer for factors and outputs string noises
func Convert(drops int) (result string) {
	if drops%3 == 0 {
		result += "Pling"
	}
	if drops%5 == 0 {
		result += "Plang"
	}
	if drops%7 == 0 {
		result += "Plong"
	}
	if result == "" {
		return fmt.Sprint(drops)
	}
	return result
}
