package cipher

import (
	"regexp"
	"strings"
	"unicode"
)

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type VigenereCipher string

func NewCaesar() *VigenereCipher {
	return NewShift(3)
}

func NewShift(shift int) *VigenereCipher {
	if shift >= 26 || shift <= -26 || shift == 0 {
		return nil
	}
	shift = (shift%26 + 26) % 26
	shiftr := 'a' + rune(shift)
	newscipher := VigenereCipher(strings.Repeat(string(shiftr), 26))
	return &newscipher
}

func NewVigenere(key string) *VigenereCipher {
	if key == "" || len(key) == 1 {
		return nil
	}
	only_a := true
	for _, ru := range key {
		if !unicode.IsLetter(ru) || unicode.IsUpper(ru) {
			return nil
		}
		if ru != 'a' {
			only_a = false
		}
	}
	if only_a {
		return nil
	}
	newvcipher := VigenereCipher(key)
	return &newvcipher
}

func (v VigenereCipher) Encode(input string) string {
	if input == "" {
		return ""
	}
	reg, _ := regexp.Compile("[^a-z]+")
	cleaned_input := reg.ReplaceAllString(strings.ToLower(input), "")
	if cleaned_input == "" {
		return ""
	}
	encoded := make([]rune, len(cleaned_input))
	for i, ru := range cleaned_input {
		ru_from_a, vshift := ru-'a', rune(v[i%len(v)])-'a'
		cod_dist := int(ru_from_a+vshift) % 26
		encoded[i] = 'a' + rune(cod_dist)
	}
	return string(encoded)
}

func (v VigenereCipher) Decode(input string) string {
	if input == "" {
		return ""
	}
	encoded := make([]rune, len(input))
	for i, ru := range input {
		ru_from_a, vshift := ru-'a', rune(v[i%len(v)])-'a'
		cod_dist := (int(ru_from_a-vshift) + 26) % 26
		encoded[i] = 'a' + rune(cod_dist)
	}
	return string(encoded)
}
