// Package proverb contains a function which returns a proverb
package proverb

import "fmt"

// Proverb returns a prover based on a string slice input
func Proverb(rhyme []string) []string {
	if rhyme == nil {
		return []string{}
	}
	length := len(rhyme)
	proverb := make([]string, length)
	for i, v := range rhyme {
		if i == length-1 {
			proverb[i] = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
		} else {
			proverb[i] = fmt.Sprintf("For want of a %s the %s was lost.", v, rhyme[i+1])
		}
	}
	return proverb
}
