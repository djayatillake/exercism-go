// package ls product contains a function which calculates the largest series product for digits
package lsproduct

import "errors"

// LargestSeriesProduct calculates the largest series product for digits for a specified span
func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span > len(digits) || span < 0 {
		return -1, errors.New("span can't be longer than string or negative")
	}
	highest_product := 0
	for i := 0; i < (len(digits) - span + 1); i++ {
		product := 1
		for _, r := range digits[i : i+span] {
			num := r - '0'
			if num < 0 || num > 9 {
				return -1, errors.New("invalid character in string")
			}
			product *= int(num)
		}
		if product > highest_product {
			highest_product = product
		}
	}
	return int64(highest_product), nil
}
