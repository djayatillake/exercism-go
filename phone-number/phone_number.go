// package phonenumber contains functions to check for valid NANP phone numbers
package phonenumber

import (
	"fmt"
	"unicode"
)

// Number tests if a string contains a valid NANP phone number, and returns it formatted
func Number(input string) (string, error) {
	digits := []rune{}
	for _, r := range input {
		if unicode.IsLetter(r) {
			return "", fmt.Errorf("letters not permitted")
		}
		if unicode.IsDigit(r) {
			digits = append(digits, r)
		}
	}
	NoLen := len(digits)
	if NoLen > 11 {
		return "", fmt.Errorf("more than 11 digits")
	} else if NoLen < 10 {
		return "", fmt.Errorf("incorrect number of digits")
	} else if NoLen == 11 && digits[0] != '1' {
		return "", fmt.Errorf("11 digits must start with 1")
	}
	if NoLen == 10 && digits[0] < '2' || digits[3] < '2' {
		return "", fmt.Errorf("area or exchange code")
	} else if NoLen == 11 && digits[1] < '2' || digits[4] < '2' {
		return "", fmt.Errorf("area or exchange code")
	}
	if NoLen == 11 {
		return string(digits[1:]), nil
	} else {
		return string(digits), nil
	}
}

// AreaCode tests if a an area code is valid
func AreaCode(input string) (ret string, err error) {
	var clean_no string
	clean_no, err = Number(input)
	if err != nil {
		return
	}
	if len(clean_no) == 11 {
		ret = string(clean_no[1:4])
	} else {
		ret = string(clean_no[0:3])
	}
	return
}

// Format outputs a number in the correct format
func Format(input string) (ret string, err error) {
	var clean_no string
	clean_no, err = Number(input)
	if err != nil {
		return
	}
	if len(clean_no) == 11 {
		ret = "1 (" + string(clean_no[1:4]) + ") " + string(clean_no[4:7]) + "-" +
			string(clean_no[7:])
	} else {
		ret = "(" + string(clean_no[0:3]) + ") " + string(clean_no[3:6]) + "-" +
			string(clean_no[6:])
	}
	return
}
