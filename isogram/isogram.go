// package isogram contains a function to check if a word is an isogram
package isogram

import "strings"

// IsIsogram checks whether an input string is an isogram word
func IsIsogram(word string) bool {
	word_c := strings.ToLower(
		strings.TrimSpace(
			strings.ReplaceAll(
				strings.ReplaceAll(word, " ", ""),
				"-", "")))
	letter_count := map[rune]int{}
	word_c_s := []rune(word_c)
	for i := range word_c_s {
		letter_count[word_c_s[i]] += 1
		if letter_count[word_c_s[i]] > 1 {
			return false
		}
	}
	return true
}
