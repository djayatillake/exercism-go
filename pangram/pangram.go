// package pangram contains a function to check if a sentence is a pangram
package pangram

import (
	"unicode"
)

// IsPangram checks if a string is a pangram
func IsPangram(input string) bool {
	set := 0
	for _, r := range input {
		if !unicode.IsLetter(r) {
			continue
		}
		d := unicode.ToLower(r) - 'a'
		set |= 1 << d
	}
	return set == 67108863
}
