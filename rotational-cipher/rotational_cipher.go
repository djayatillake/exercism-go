// package rotational cipher contains a function to encode a string to a cipher
package rotationalcipher

import "unicode"

// RotationalCipher returns a cipher from a string and rotation int
func RotationalCipher(input string, shift int) string {
	runes := make([]rune, len(input))
	for i, r := range input {
		if !unicode.IsLetter(r) {
			runes[i] = r
			continue
		}
		if unicode.IsUpper(r) {
			runes[i] = 'A' + ((r - 'A' + rune(shift)) % 26)
		} else {
			runes[i] = 'a' + ((r - 'a' + rune(shift)) % 26)
		}
	}
	return string(runes)
}
