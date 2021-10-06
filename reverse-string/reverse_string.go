// package reverse has a function which returns the reverse of a string
package reverse

// Reverse returns the reverse of a string
func Reverse(input string) (rev string) {
	for _, rune := range input {
		rev = string(rune) + rev
	}
	return
}
