// package strains contains functions to split a collection based on a predicate
package strain

type (
	Ints    []int
	Lists   [][]int
	Strings []string
)

func (collection Ints) Keep(f func(int) bool) (ret Ints) {
	for _, v := range collection {
		if f(v) {
			ret = append(ret, v)
		}
	}
	return
}

func (collection Ints) Discard(f func(int) bool) (ret Ints) {
	for _, v := range collection {
		if !f(v) {
			ret = append(ret, v)
		}
	}
	return
}

func (collection Lists) Keep(f func([]int) bool) (ret Lists) {
	for _, v := range collection {
		if f(v) {
			ret = append(ret, v)
		}
	}
	return
}

func (collection Strings) Keep(f func(string) bool) (ret Strings) {
	for _, v := range collection {
		if f(v) {
			ret = append(ret, v)
		}
	}
	return
}
