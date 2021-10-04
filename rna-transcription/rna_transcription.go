// package strand has a function to transcribe dna to rna
package strand

// ToRNA takes a dna string and outputs an rna string
func ToRNA(dna string) string {
	dna2rna := map[rune]rune{'G': 'C', 'C': 'G', 'T': 'A', 'A': 'U'}
	rna := make([]rune, len(dna))
	for i, dr := range dna {
		rr, ok := dna2rna[dr]
		if !ok {
			return ""
		}
		rna[i] = rr
	}
	return string(rna)
}
