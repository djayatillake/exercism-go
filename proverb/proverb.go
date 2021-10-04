// Package proverb contains a function which returns a proverb
package proverb

// Proverb returns a prover based on a string slice input
func Proverb(rhyme []string) []string {
	length := len(rhyme)
	if rhyme == nil || length == 0 {
		return []string{}
	}
	proverb := make([]string, length)
	for i := 0; i < length-1; i++ {
		proverb[i] = "For want of a " + rhyme[i] + " the " + rhyme[i+1] + " was lost."
	}
	proverb[length-1] = "And all for the want of a " + rhyme[0] + "."
	return proverb
}
