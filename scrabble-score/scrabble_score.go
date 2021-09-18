// package scrabble provides functions to calculate the value of a word in the game
package scrabble

import "strings"

// Score calculates the value of a word in Scrabble depending on letter values
func Score(word string) (score int) {
	word_s := []rune(word)
	word_val := map[string]int{
		"A": 1, "E": 1, "I": 1, "O": 1, "U": 1, "L": 1, "N": 1, "R": 1, "S": 1, "T": 1,
		"D": 2, "G": 2,
		"B": 3, "C": 3, "M": 3, "P": 3,
		"F": 4, "H": 4, "V": 4, "W": 4, "Y": 4,
		"K": 5,
		"J": 8, "X": 8,
		"Q": 10, "Z": 10,
	}

	for i := 0; i < len(word_s); i++ {
		score += word_val[strings.ToUpper(string(word_s[i]))]
	}
	return score
}

// Score calculates the value of a word in Scrabble depending on letter values
func ScoreComplex(word string, dblindex, tripindex []bool, dblwd, tripwd bool) (score int) {
	word_s := []rune(word)
	word_val := map[string]int{
		"A": 1, "E": 1, "I": 1, "O": 1, "U": 1, "L": 1, "N": 1, "R": 1, "S": 1, "T": 1,
		"D": 2, "G": 2,
		"B": 3, "C": 3, "M": 3, "P": 3,
		"F": 4, "H": 4, "V": 4, "W": 4, "Y": 4,
		"K": 5,
		"J": 8, "X": 8,
		"Q": 10, "Z": 10,
	}

	for i := 0; i < len(word_s); i++ {
		if dblindex[i] {
			score += word_val[strings.ToUpper(string(word_s[i]))] * 2
		} else if tripindex[i] {
			score += word_val[strings.ToUpper(string(word_s[i]))] * 3
		} else {
			score += word_val[strings.ToUpper(string(word_s[i]))]
		}
	}

	if dblwd {
		return score * 2
	} else if tripwd {
		return score * 3
	} else {
		return score
	}
}
