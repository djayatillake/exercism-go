// package rotational cipher contains a function to encode a string to a cipher
package rotationalcipher

import "unicode"

const letters = "abcdefghijklmnopqrstuvwxyz"

// RotationalCipher returns a cipher from a string and rotation int
func RotationalCipher(input string, shift int) string {
	runes := make([]rune, len(input))
	for i, r := range input {
		if !unicode.IsLetter(r) {
			runes[i] = r
			continue
		}
		lower_r := unicode.ToLower(r)
		shiftr := letters[(int(lower_r-'a')+shift)%26]
		if unicode.IsUpper(r) {
			runes[i] = unicode.ToUpper(rune(shiftr))
		} else {
			runes[i] = rune(shiftr)
		}
	}
	return string(runes)
}
