// package cryptosquare contains a function to encode a string
package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

// Encode a string using cryptosquare method
func Encode(input string) string {
	if input == "" {
		return ""
	}
	reg, _ := regexp.Compile("[^a-z0-9]+")
	cleaned_input := reg.ReplaceAllString(strings.ToLower(input), "")
	c := int(math.Ceil(math.Sqrt(float64(len(cleaned_input)))))
	padding := 0
	if len(cleaned_input)%c > 0 {
		padding = c - (len(cleaned_input) % c)
	}
	cleaned_input_pad := cleaned_input + strings.Repeat(" ", padding)
	rowsc := make([]string, c)
	for i, r := range cleaned_input_pad {
		rowsc[i%c] += string(r)
	}
	return strings.Join(rowsc, " ")
}
