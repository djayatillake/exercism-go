// package strand has a function to transcribe dna to rna
package strand

// ToRNA takes a dna string and outputs an rna string
func ToRNA(dna string) string {
	rna := make([]rune, len(dna))
	for i, dr := range dna {
		if dr != 'G' && dr != 'C' && dr != 'T' && dr != 'A' {
			return ""
		}
		switch dr {
		case 'G':
			rna[i] = 'C'
		case 'C':
			rna[i] = 'G'
		case 'T':
			rna[i] = 'A'
		case 'A':
			rna[i] = 'U'
		}
	}
	return string(rna)
}
