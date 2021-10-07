// package wordcount contains functions to count the number of distinct words in a string
package wordcount

import (
	"regexp"
	"strings"
)

// implementing type Frequency
type Frequency map[string]int

// WordCount counts the number of distinct words in a string, case-insensitive
func WordCount(input string) Frequency {
	counts := Frequency{}
	for _, str := range strings.Split(clean(input), " ") {
		if strings.TrimSpace(str) == "" {
			continue
		}
		counts[str]++
	}
	return counts
}

// clean strips the input string of unwanted characters
func clean(input string) string {
	cinput := strings.ReplaceAll(
		strings.ReplaceAll(
			strings.ToLower(" "+input+" "), " '", " "), "' ", " ")
	reg, _ := regexp.Compile("[^a-zA-Z0-9' ]+")
	return reg.ReplaceAllString(cinput, " ")
}
