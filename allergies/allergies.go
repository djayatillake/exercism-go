// package allergies contains a function to calculate allergies
package allergies

func Allergies(score uint) (ret []string) {
	subs := []string{
		"cats",
		"pollen",
		"chocolate",
		"tomatoes",
		"strawberries",
		"shellfish",
		"peanuts",
		"eggs",
	}
	score %= 256
	count := 7
	for ; score > 0; score /= 2 {
		if score%2 == 1 {
			ret = append(ret, subs[count])
		}
		count--
	}
	return
}

func AllergicTo(score uint, response string) bool {
	for _, v := range Allergies(score) {
		if v == response {
			return true
		}
	}
	return false
}
