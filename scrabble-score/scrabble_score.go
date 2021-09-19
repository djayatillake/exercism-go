// package scrabble provides functions to calculate the value of a word in the game
package scrabble

import "unicode"

// Score calculates the value of a word in Scrabble depending on letter values
func Score(word string) (score int) {
	for _, v := range word {
		score += letter_val(v)
	}
	return score
}

// Score calculates the value of a word in Scrabble depending on letter values
func ScoreComplex(word string, dblindex, tripindex []bool, dblwd, tripwd int) (score int) {
	for i, v := range word {
		if dblindex[i] {
			score += letter_val(v) * 2
		} else if tripindex[i] {
			score += letter_val(v) * 3
		} else {
			score += letter_val(v)
		}
	}
	switch {
	case dblwd == 1:
		score = score * 2
	case dblwd == 2:
		score = score * 4
	case tripwd == 1:
		score = score * 3
	case tripwd == 2:
		score = score * 9
	}
	return score
}

// letter_val helper function provides the value of any letter in scrabble
func letter_val(letter rune) (val int) {
	switch unicode.ToUpper(letter) {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		val = 1
	case 'D', 'G':
		val = 2
	case 'B', 'C', 'M', 'P':
		val = 3
	case 'F', 'H', 'V', 'W', 'Y':
		val = 4
	case 'K':
		val = 5
	case 'J', 'X':
		val = 8
	case 'Q', 'Z':
		val = 10
	}
	return val
}
