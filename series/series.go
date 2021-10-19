// package series contains functions to return series from a string
package series

// All returns a list of all substrings of s with length n.
func All(n int, s string) (output []string) {
	for i := 0; i <= len(s)-n; i++ {
		output = append(output, s[i:i+n])
	}
	return
}

// UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, s string) string {
	return s[:n]
}

// First returns the first substring of s with length n safely
func First(n int, s string) (first string, ok bool) {
	if len(s) < n {
		return "", false
	}
	return s[:n], true
}
