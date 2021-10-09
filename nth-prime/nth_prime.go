// package prime contains a function to find the nth prime
package prime

// Nth finds the nth prime where n is an int input
func Nth(n int) (int, bool) {
	if n == 0 {
		return 0, false
	}
	num := 2
	last_prime := 0
	count := 0
	primes := []int{}
	for count < n {
		is_prime := true
		for _, p := range primes {
			if num%p == 0 {
				is_prime = false
				break
			}
		}
		if is_prime {
			last_prime = num
			primes = append(primes, num)
			count++
		}
		num++
	}
	return last_prime, true
}
