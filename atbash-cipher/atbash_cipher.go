// package atbash contains a function to encode a string to an atbash cipher
package atbash

import "unicode"

// Atbash returns an atbash cipher from a string
func Atbash(input string) string {
	runes := []rune{}
	for _, r := range input {
		leng, r_dig, r_let := len(runes), unicode.IsDigit(r), unicode.IsLetter(r)
		if leng == 5 || (leng+1)%6 == 0 && leng > 5 && (r_dig || r_let) {
			runes = append(runes, ' ')
		}
		if r_dig {
			runes = append(runes, r)
		} else if r_let {
			ciph_r := 'a' + (25 - (unicode.ToLower(r) - 'a'))
			runes = append(runes, ciph_r)
		}
	}
	return string(runes)
}
