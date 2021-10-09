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
