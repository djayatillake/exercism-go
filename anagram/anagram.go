// package anagram provides functions to test for anagrams
package anagram

import (
	"strings"
	"unicode"
)

// uni_or returns the byte-or of all the letters in a string
func dig_plus(input string) int {
	set := 0
	for _, r := range input {
		if !unicode.IsLetter(r) {
			continue
		}
		d := unicode.ToLower(r) - 'a'
		set += 1 << d
	}
	return set
}

// Detect tests whether the string has anagrams in the slice of strings and returns the slice of them
func Detect(word string, list []string) (anagrams []string) {
	lword := strings.ToLower(word)
	lword_uor := dig_plus(lword)
	for _, item := range list {
		if len(item) != len(word) {
			continue
		}
		litem := strings.ToLower(item)
		if litem == lword {
			continue
		}
		litem_uor := dig_plus(litem)
		if lword_uor == litem_uor {
			anagrams = append(anagrams, item)
		}
	}
	return anagrams
}
