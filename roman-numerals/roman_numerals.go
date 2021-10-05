// package romannumerals provides a function to convert an int to roman numerals
package romannumerals

import "errors"

// bms is a helper function which takes big, middle and large numerals and outputs the correct pattern for an arabic digit
func bms(to9 int, big, mid, sma string) string {
	switch {
	case to9 == 9:
		return sma + big
	case to9 > 4:
		ret := mid
		for i := 0; i < to9-5; i++ {
			ret += sma
		}
		return ret
	case to9 == 4:
		return sma + mid
	default:
		ret := ""
		for i := 0; i < to9; i++ {
			ret += sma
		}
		return ret
	}
}

// ToRomanNumeral converts an int to roman numerals
func ToRomanNumeral(arabic int) (string, error) {
	if arabic <= 0 || arabic > 3000 {
		return "", errors.New("only ints between 1 and 3000 are allowed")
	}
	return bms(arabic/1000, "", "", "M") + bms((arabic%1000)/100, "M", "D", "C") +
		bms((arabic%100)/10, "C", "L", "X") + bms(arabic%10, "X", "V", "I"), nil
}
