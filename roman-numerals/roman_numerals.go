// package romannumerals provides a function to convert an int to roman numerals
package romannumerals

import "errors"

// ToRomanNumeral converts an int to roman numerals
func ToRomanNumeral(arabic int) (string, error) {
	if arabic <= 0 || arabic > 3000 {
		return "", errors.New("only ints between 1 and 3000 are allowed")
	}
	var M, CMDC, XCLX, IXVX string
	switch arabic / 1000 {
	case 3:
		M = "MMM"
	case 2:
		M = "MM"
	case 1:
		M = "M"
	default:
		M = ""
	}
	switch (arabic % 1000) / 100 {
	case 9:
		CMDC = "CM"
	case 8:
		CMDC = "DCCC"
	case 7:
		CMDC = "DCC"
	case 6:
		CMDC = "DC"
	case 5:
		CMDC = "D"
	case 4:
		CMDC = "CD"
	case 3:
		CMDC = "CCC"
	case 2:
		CMDC = "CC"
	case 1:
		CMDC = "C"
	default:
		CMDC = ""
	}
	switch (arabic % 100) / 10 {
	case 9:
		XCLX = "XC"
	case 8:
		XCLX = "LXXX"
	case 7:
		XCLX = "LXX"
	case 6:
		XCLX = "LX"
	case 5:
		XCLX = "L"
	case 4:
		XCLX = "XL"
	case 3:
		XCLX = "XXX"
	case 2:
		XCLX = "XX"
	case 1:
		XCLX = "X"
	default:
		XCLX = ""
	}
	switch arabic % 10 {
	case 9:
		IXVX = "IX"
	case 8:
		IXVX = "VIII"
	case 7:
		IXVX = "VII"
	case 6:
		IXVX = "VI"
	case 5:
		IXVX = "V"
	case 4:
		IXVX = "IV"
	case 3:
		IXVX = "III"
	case 2:
		IXVX = "II"
	case 1:
		IXVX = "I"
	default:
		IXVX = ""
	}
	return M + CMDC + XCLX + IXVX, nil
}
