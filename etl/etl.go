// package etl contains a function to transform from old to new format
package etl

import "strings"

// Transform converts old format map of letter values to new
func Transform(old map[int][]string) map[string]int {
	new := map[string]int{}
	for k, v := range old {
		for _, s := range v {
			new[strings.ToLower(s)] = k
		}
	}
	return new
}
