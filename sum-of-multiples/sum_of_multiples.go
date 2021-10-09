// package summultiples contains a function that returns the sum of multiples
package summultiples

// SumMultiples
func SumMultiples(limit int, divisors ...int) (sum int) {
	if len(divisors) == 0 {
		return
	}
	if limit < divisors[0] {
		return
	}
	factors := []int{}
	for _, n := range divisors {
		if n == 0 {
			continue
		}
		fac := (limit - 1) / n
		for i := 1; i <= fac; i++ {
			already := false
			if len(factors) == 0 {
				factors = append(factors, i*n)
				continue
			}
			for _, v := range factors {
				if i*n == v {
					already = true
				}
			}
			if !already {
				factors = append(factors, i*n)
			}
		}
	}
	for _, v := range factors {
		sum += v
	}
	return
}
