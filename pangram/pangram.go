// package pangram contains a function to check if a sentence is a pangram
package pangram

import (
	"strings"
)

// IsPangram checks if a string is a pangram
func IsPangram(input string) bool {
	runes := map[rune]bool{
		'a': false, 'b': false, 'c': false, 'd': false, 'e': false, 'f': false,
		'g': false, 'h': false, 'i': false, 'j': false, 'k': false, 'l': false,
		'm': false, 'n': false, 'o': false, 'p': false, 'q': false, 'r': false,
		's': false, 't': false, 'u': false, 'v': false, 'w': false, 'x': false,
		'y': false, 'z': false,
	}
	for _, rune := range strings.ToLower(input) {
		runes[rune] = true
	}
	for _, bools := range runes {
		if !bools {
			return false
		}
	}
	return true
}
