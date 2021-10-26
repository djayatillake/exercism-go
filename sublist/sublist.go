// package sublist containst a function to test if a list a sublist of another
package sublist

type Relation string

// equal determins whether two lists are equal
func equal(list1 []int, list2 []int) bool {
	len1, len2, c := len(list1), len(list2), 0
	if len1 != len2 {
		return false
	} else if len1 == 0 {
		return true
	}
	for i := 0; i < len1; i++ {
		if list1[i] == list2[i] {
			c++
		}
	}
	return c == len1
}

// Sublist tests whether a list1 is a sublist, superlist of list2 or if they are equal or unequal
func Sublist(list1 []int, list2 []int) Relation {
	len1, len2 := len(list1), len(list2)
	if equal(list1, list2) {
		return "equal"
	} else if len1 > len2 && len2 == 0 {
		return "superlist"
	} else if len2 > len1 && len1 == 0 {
		return "sublist"
	}

	if len1 > len2 {
		for i := 0; i <= len1-len2; i++ {
			if equal(list1[i:len2+i], list2) {
				return "superlist"
			}
		}
	} else if len1 < len2 {
		for i := 0; i <= len2-len1; i++ {
			if equal(list2[i:len1+i], list1) {
				return "sublist"
			}
		}
	}
	return "unequal"
}
