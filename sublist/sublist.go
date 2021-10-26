// package sublist containst a function to test if a list a sublist of another
package sublist

type Relation string

// is_sublist determines whether list1 is a sublist of list2
func is_sublist(list1, list2 []int) bool {
	len1, len2 := len(list1), len(list2)
	for i := 0; i <= len2-len1; i++ {
		c := 0
		for j := 0; j < len1; j++ {
			if list1[j] == list2[i : len1+i][j] {
				c++
			}
		}
		if c == len1 {
			return true
		}
	}
	return false
}

// Sublist tests whether a list1 is a sublist, superlist of list2 or if they are equal or unequal
func Sublist(list1, list2 []int) Relation {
	len1, len2 := len(list1), len(list2)
	is_sub := is_sublist(list1, list2)
	switch {
	case len1 == 0 && len2 == 0 || is_sub && len1 == len2:
		return "equal"
	case len2 > len1 && len1 == 0 || is_sub:
		return "sublist"
	case len1 > len2 && len2 == 0 || is_sublist(list2, list1):
		return "superlist"
	default:
		return "unequal"
	}
}
