// package strains contains functions to split a collection based on a predicate
package strain

type (
	Ints    []int
	Lists   [][]int
	Strings []string
)

// method Keep on ints collection returns the collection filtered on a predicate
func (collection Ints) Keep(f func(int) bool) (ret Ints) {
	for _, v := range collection {
		if f(v) {
			ret = append(ret, v)
		}
	}
	return
}

// method Keep on ints collection returns the collection not filtered on a predicate
func (collection Ints) Discard(f func(int) bool) (ret Ints) {
	for _, v := range collection {
		if !f(v) {
			ret = append(ret, v)
		}
	}
	return
}

// method Keep on Lists collection returns the collection filtered on a predicate
func (collection Lists) Keep(f func([]int) bool) (ret Lists) {
	for _, v := range collection {
		if f(v) {
			ret = append(ret, v)
		}
	}
	return
}

// method Keep on strings collection returns the collection filtered on a predicate
func (collection Strings) Keep(f func(string) bool) (ret Strings) {
	for _, v := range collection {
		if f(v) {
			ret = append(ret, v)
		}
	}
	return
}
