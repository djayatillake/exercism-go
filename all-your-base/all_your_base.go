// package allyourbase contains a function for base conversion
package allyourbase

import (
	"errors"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return []int{}, errors.New("input base must be >= 2")
	} else if outputBase < 2 {
		return []int{}, errors.New("output base must be >= 2")
	} else if len(inputDigits) == 0 {
		return []int{0}, nil
	}
	decimal := 0
	for _, digit := range inputDigits {
		if digit >= inputBase || digit < 0 {
			return []int{}, errors.New("all digits must satisfy 0 <= d < input base")
		}
		decimal = inputBase*decimal + digit
	}
	if decimal == 0 {
		return []int{0}, nil
	}
	output := []int{}
	for ; decimal > 0; decimal /= outputBase {
		output = append([]int{decimal % outputBase}, output...)
	}
	return output, nil
}
