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
	rows := []string{}
	for i := 0; i < len(cleaned_input_pad); i += c {
		rows = append(rows, cleaned_input_pad[i:i+c])
	}
	rowsc := make([]string, c)
	for i, _ := range rowsc {
		nrow := []rune{}
		for _, row := range rows {
			nrow = append(nrow, rune(row[i]))
		}
		rowsc[i] = string(nrow)
	}
	return strings.Join(rowsc, " ")
}
