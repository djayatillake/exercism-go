// package cryptosquare contains a function to encode a string
package cryptosquare

import (
	"math"
	"regexp"
	"strings"
)

// Encode a string using cryptosquare method
func Encode(input string) (encoded string) {
	if input == "" {
		return
	}
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	cleaned_input := reg.ReplaceAllString(input, "")
	lower_input := strings.ToLower(cleaned_input)
	cf := math.Sqrt(float64(len(lower_input)))
	c := int(cf)
	if cf > float64(c) {
		c++
	}
	padding := 0
	if len(lower_input)%c > 0 {
		padding = c - (len(lower_input) % c)
	}
	lower_input_pad := lower_input + strings.Repeat(" ", padding)
	rows := []string{}
	for i := 0; i < len(lower_input_pad); i += c {
		rows = append(rows, lower_input_pad[i:i+c])
	}
	rowsc := make([]string, c)
	for i, _ := range rowsc {
		nrow := []rune{}
		for _, row := range rows {
			nrow = append(nrow, rune(row[i]))
		}
		rowsc[i] = string(nrow)
	}
	for i, row := range rowsc {
		if i < len(rowsc)-1 {
			encoded += row + " "
		} else {
			encoded += row
		}
	}
	return
}
