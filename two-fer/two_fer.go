// Package twofer is a one function package with an implementation of twofer
package twofer

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}

	return "One for " + name + ", one for me."
}
