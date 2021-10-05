// package accumulate contains a function to apply a function to each element of a collection
package accumulate

func Accumulate(col []string, f func(string) string) []string {
	output := make([]string, len(col))
	for i, v := range col {
		output[i] = f(v)
	}
	return output
}
