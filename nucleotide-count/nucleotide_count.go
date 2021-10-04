package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
func (d DNA) Counts() (Histogram, error) {
	h := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	if len(d) == 0 {
		return h, nil
	}
	for _, r := range d {
		if r != 'A' && r != 'C' && r != 'G' && r != 'T' {
			return nil, errors.New("invalid dna")
		}
		h[r]++
	}
	return h, nil
}
