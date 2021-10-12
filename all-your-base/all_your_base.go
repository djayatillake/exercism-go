// package allyourbase contains a function for base conversion
package allyourbase

import (
	"errors"
	"math"
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
	for i, digit := range inputDigits {
		if digit >= inputBase || digit < 0 {
			return []int{}, errors.New("all digits must satisfy 0 <= d < input base")
		}
		decimal += digit * int(math.Pow(float64(inputBase), float64(len(inputDigits)-1-i)))
	}
	if decimal == 0 {
		return []int{0}, nil
	}
	dec_exp := math.Log(float64(decimal)) / math.Log(float64(outputBase))
	dec_exp_int := int(dec_exp)
	output := make([]int, dec_exp_int+1)
	for i := 0; i < dec_exp_int+1; i++ {
		base_p := int(math.Pow(float64(outputBase), float64(dec_exp_int-i)))
		digit := decimal / base_p
		output[i] = digit
		decimal -= base_p * digit
	}
	return output, nil
}
