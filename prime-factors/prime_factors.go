package prime

// Nth finds the nth prime where n is an int input
func Factors(n int64) []int64 {
	if n <= 1 {
		return []int64{}
	}
	factors := []int64{}
	for n%2 == 0 {
		factors = append(factors, 2)
		n /= 2
	}
	for num := int64(3); n > 1; num += 2 {
		for n%num == 0 {
			factors = append(factors, num)
			n /= num
		}
	}
	return factors
}
