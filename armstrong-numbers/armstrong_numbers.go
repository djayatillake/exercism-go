// package armstrong contains a function to test if a number is an armstrong no
package armstrong

import (
	"math"
)

// IsNumber tests whether an integer is an armstrong no
func IsNumber(input int) bool {
	if input < 10 {
		return true
	}
	sum := float64(0)
	input2 := input
	power := math.Ceil(math.Log10(float64(input)))
	for ; input2 > 0; input2 /= 10 {
		digit := input2 % 10
		sum += math.Pow(float64(digit), power)
	}
	return int(sum) == input
}
