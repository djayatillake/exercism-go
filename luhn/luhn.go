// package luhn provides a function to validate an id number based on the Luhn algorithm
package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

// Valid tests whether an input string is a valid id number based on the Luhn algorithm
func Valid(id string) bool {
	trimmed_id := strings.TrimSpace(id)
	if len(trimmed_id) < 2 {
		return false
	}

	id_slice := []rune(strings.ReplaceAll(trimmed_id, " ", ""))
	sum := 0
	even_len := len(id_slice)%2 == 0

	for i, v := range id_slice {
		string_dig := string(v)

		if !unicode.IsDigit(v) {
			return false
		}

		string_dig_int, _ := strconv.Atoi(string_dig)

		if even_len {
			if i%2 == 0 {
				sum += multi9(string_dig_int)
			} else {
				sum += string_dig_int
			}
		} else {
			if i%2 == 1 {
				sum += multi9(string_dig_int)
			} else {
				sum += string_dig_int
			}
		}
	}
	return sum%10 == 0
}

func multi9(digit int) int {
	if digit*2 > 9 {
		return digit*2 - 9
	} else {
		return digit * 2
	}
}
