// package phonenumber contains functions to check for valid NANP phone numbers
package phonenumber

import "fmt"

// Number tests if a string contains a valid NANP phone number, and returns it formatted
func Number(input string) (string, error) {
	digits := []rune{}
	for _, r := range input {
		if r >= '0' && r <= '9' {
			digits = append(digits, r)
		}
	}
	if len(digits) == 11 && digits[0] == '1' {
		digits = digits[1:]
	}
	if len(digits) != 10 {
		return "", fmt.Errorf("not enough or incorrect digits")
	}
	if digits[0] < '2' || digits[3] < '2' {
		return "", fmt.Errorf("incorrect area or exchange code")
	}
	return string(digits), nil
}

// AreaCode tests if a an area code is valid
func AreaCode(input string) (ret string, err error) {
	clean_no, err := Number(input)
	if err != nil {
		return "", err
	}
	return string(clean_no[0:3]), nil
}

// Format outputs a number in the correct format
func Format(input string) (string, error) {
	clean_no, err := Number(input)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", clean_no[0:3], clean_no[3:6], clean_no[6:]), nil
}
