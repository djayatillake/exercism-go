// euclids formula: pair of integers m and n with m>n>0
// a=m**2 - n**2, b = 2mn, c=m**2+n**2
// therefore to find a pythagorean triple that adds up to N
// N = 2 m**2 + 2mn
// N/2 = m**2 + mn
// can start with a n = 1 iterate over m = 2 and above until N is hit or exceeded
// then increment n if N is exceed and start again
// storing all triples that work

// package pythagorean contains functions which returns pythagorean triplets
package pythagorean

type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive in lexicographic order
func Range(min, max int) []Triplet {
	trips := []Triplet{}
	for i := min; i < max-1; i++ {
		for j := i + 1; j < max; j++ {
			for k := j + 1; k <= max; k++ {
				if i*i+j*j == k*k {
					trips = append(trips, Triplet{i, j, k})
				}
			}
		}
	}
	return trips
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p in lexicographic order
func Sum(p int) []Triplet {
	max := int(p / 2)
	res := []Triplet{}
	for _, v := range Range(1, max) {
		if v[0]+v[1]+v[2] == p {
			res = append(res, v)
		}
	}
	return res
}
