// package protein contains a function to translate an rna string into protein
package protein

import "errors"

var (
	ErrStop        error = errors.New("ErrStop")
	ErrInvalidBase error = errors.New("ErrInvalidBase")
)

// FromCodon translates codon to protein
func FromCodon(codon string) (string, error) {
	translate := map[string]string{
		"AUG": "Methionine",
		"UUU": "Phenylalanine",
		"UUC": "Phenylalanine",
		"UUA": "Leucine",
		"UUG": "Leucine",
		"UCU": "Serine",
		"UCC": "Serine",
		"UCA": "Serine",
		"UCG": "Serine",
		"UAU": "Tyrosine",
		"UAC": "Tyrosine",
		"UGU": "Cysteine",
		"UGC": "Cysteine",
		"UGG": "Tryptophan",
		"UAA": "STOP",
		"UAG": "STOP",
		"UGA": "STOP",
	}
	if v, ok := translate[codon]; !ok {
		return "", ErrInvalidBase
	} else if v == "STOP" {
		return "", ErrStop
	}
	return translate[codon], nil
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
