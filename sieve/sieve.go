// package sieve provides a function to implement the sieve of Eratosthenes
package sieve

// Sieve returns all the primes from 2 to an input limit using the Eratosthenes method
func Sieve(limit int) []int {
	if limit < 2 {
		return []int{}
	}
	primes := []int{}
	non_primes := []int{}
	for p := 2; p <= limit; p++ {
		p_comp := false
		for _, np := range non_primes {
			if p == np {
				p_comp = true
				break
			}
		}
		if p_comp {
			continue
		}
		primes = append(primes, p)
		for j := 2; p*j <= limit; j++ {
			non_primes = append(non_primes, p*j)
		}
	}
	return primes
}
