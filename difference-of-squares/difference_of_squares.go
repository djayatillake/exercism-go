// package diffsquares for finding sum of squares and square of sums plus diff
package diffsquares

// SquareOfSum
func SquareOfSum(n int) int {
	Sum := ((n * (n + 1)) / 2)
	return Sum * Sum
}

// SumOfSquares
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference
func Difference(n int) int {
	return n * (n*n - 1) * (3*n + 2) / 12
}
