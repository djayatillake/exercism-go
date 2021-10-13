// package perfect checks if a numbers are perfect, deficient or abundant
package perfect

import "errors"

type Classification string

var ErrOnlyPositive error = errors.New("ErrOnlyPositive")

var (
	ClassificationPerfect   Classification = "ClassificationPerfect"
	ClassificationDeficient Classification = "ClassificationDeficient"
	ClassificationAbundant  Classification = "ClassificationAbundant"
)

// Classify tests if a number is perfect
func Classify(input int64) (Classification, error) {
	if input <= 0 {
		return Classification(""), ErrOnlyPositive
	}
	return ClassifyDetail(input), nil
}

// ClassifyDetail checks if a number is perfect deficient or abundant
func ClassifyDetail(input int64) Classification {
	sum := int64(0)
	for i := int64(1); i < input; i++ {
		if input%i == 0 {
			sum += i
		}
	}
	switch {
	case sum == input:
		return ClassificationPerfect
	case sum > input:
		return ClassificationAbundant
	default:
		return ClassificationDeficient
	}
}
