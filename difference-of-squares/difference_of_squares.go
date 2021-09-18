// package diffsquares for finding sum of squares and square of sums plus diff
package diffsquares

// SquareOfSum calculates the Square of the sum of the first n natural numbers input
func SquareOfSum(n int) int {
	Sum := ((n * (n + 1)) / 2)
	return Sum * Sum
}

// SumOfSquares calculates the sum of the squares of the first n natural numbers input
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference calculates the difference between the SquareOfSum and SumOfSquares
func Difference(n int) int {
	return n * (n*n - 1) * (3*n + 2) / 12
}
