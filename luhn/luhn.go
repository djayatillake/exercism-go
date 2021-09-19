// package luhn provides a function to validate an id number based on the Luhn algorithm
package luhn

// Valid tests whether an input string is a valid id number based on the Luhn algorithm
func Valid(id string) bool {
	idx := len(id) - 1
	sum := 0
	digitco := 0

	for i := idx; i >= 0; i-- {
		r := id[i]
		digit := (r - '0')
		if r >= '0' && r <= '9' {
			digitco++
		}
		switch {
		case r == ' ':
			continue
		case digitco%2 == 1 || r == '9' || r == '0':
			sum += int(digit)
		case r >= '1' && r <= '4':
			sum += int(digit << 1)
		case r >= '5' && r <= '8':
			sum += int(digit<<1) - 9
		default:
			return false
		}
	}
	return digitco >= 2 && sum%10 == 0
}
