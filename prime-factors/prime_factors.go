package prime

// Nth finds the nth prime where n is an int input
func Factors(n int64) []int64 {
	if n <= 1 {
		return []int64{}
	}
	num := int64(2)
	primes := []int64{}
	factors := []int64{}
	for n > 1 {
		is_prime := true
		for _, p := range primes {
			if num%p == 0 {
				is_prime = false
				break
			}
		}
		if is_prime {
			primes = append(primes, num)
			for n%num == 0 {
				factors = append(factors, num)
				n /= num
			}
		}
		num++
	}
	return factors
}
