// package protein contains a function to translate an rna string into protein
package protein

import "errors"

var (
	ErrStop        error = errors.New("ErrStop")
	ErrInvalidBase error = errors.New("ErrInvalidBase")
)

// FromCodon translates codon to protein
func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}

// FromRNA translates an RNA string to a slice of proteins
func FromRNA(RNA string) ([]string, error) {
	proteins := []string{}
	for i := 0; i < len(RNA); i += 3 {
		v, err := FromCodon(RNA[i : i+3])
		if err == ErrStop {
			return proteins, nil
		} else if err == ErrInvalidBase {
			return proteins, err
		}
		proteins = append(proteins, v)
	}
	return proteins, nil
}
