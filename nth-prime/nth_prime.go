// package prime contains a function to find the nth prime
package prime

// Nth finds the nth prime where n is an int input
func Nth(n int) (int, bool) {
	if n == 0 {
		return 0, false
	}
	num := 1
	last_prime := 0
	count := 0
	for count < n {
		if is_prime(num) {
			last_prime = num
			count++
		}
		num++
	}
	return last_prime, true
}

// is prime determines if a number is prime
func is_prime(num int) bool {
	factors := []int{}
	for i := 1; i < num; i++ {
		if num%i == 0 {
			factors = append(factors, i)
		}
	}
	if len(factors) == 1 {
		return true
	}
	return false
}
