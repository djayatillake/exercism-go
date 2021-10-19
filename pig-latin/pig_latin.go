package piglatin

import "strings"

// Sentence takes a sentence and returns it in pig latin
func Sentence(input string) (output string) {
	for _, str := range strings.Split(input, " ") {
		output += PigLatin(str) + " "
	}
	output = strings.TrimSpace(output)
	return
}

// PigLatin outputs a word in pig latin
func PigLatin(input string) string {
	first := string(input[0])
	first2 := input[0:2]
	var first3 string
	var endafter3 string
	var endafter2 string
	var secnthird string
	var third string
	if len(input) > 2 {
		first3 = string(input[0:3])
		endafter3 = input[3:]
		endafter2 = input[2:]
		secnthird = input[1:3]
		third = string(input[2])
	}
	endafter1 := input[1:]
	end := "ay"
	var first_v bool
	if first == "a" || first == "e" || first == "i" || first == "o" || first == "u" {
		first_v = true
	}
	if first_v || first2 == "xr" || first2 == "yt" {
		return input + end
	} else if secnthird == "qu" && !first_v || first3 == "thr" || first3 == "sch" {
		return endafter3 + first3 + end
	} else if first2 == "ch" || first2 == "qu" || first2 == "th" || third == "y" {
		return endafter2 + first2 + end
	} else if string(input[1]) == "y" {
		return endafter1 + first + end
	} else {
		return endafter1 + first + end
	}
}
