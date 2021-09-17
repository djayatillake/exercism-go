// package hamming provides a function to calculate hamming distance
package hamming

import "errors"

// Distance calculates the hamming distance between two strings of equal length
func Distance(a, b string) (int, error) {
	if len(a) == len(b) && a != "" {
		count := 0
		for i := 0; i < len(a); i++ {
			if string(a[i]) != string(b[i]) {
				count++
			}
		}
		return count, nil
	} else if a == "" && b == "" {
		return 0, nil
	}
	return 0, errors.New("true")
}
