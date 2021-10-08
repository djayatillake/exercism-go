// package isbn contains a function to check if a string is a valid isbn
package isbn

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// IsValidISBN checks if a string is a valid ISBN
func IsValidISBN(input string) bool {
	str := strings.ReplaceAll(input, "-", "")
	if len(str) > 10 || len(str) < 10 {
		return false
	}
	sum := 0
	for i := 0; i < 9; i++ {
		num := str[i] - '0'
		if num < 0 || num > 9 {
			return false
		}
		sum += int(num) * (10 - i)
	}
	num := str[9] - '0'
	is_num := num >= 0 && num <= 9
	if str[9] != 'X' && !is_num {
		return false
	} else if str[9] == 'X' {
		sum += 10
	} else {
		sum += int(num)
	}
	return sum%11 == 0
}

// ToISBN13 generates an ISBN 13 from an ISBN 10
func ToISBN13(input string) string {
	if IsValidISBN(input) {
		str := strings.ReplaceAll(input, "-", "")
		str2 := "978" + str[:9]
		return fmt.Sprintf("%s%d", str2, Mod10(str2))
	}
	return ""
}

// Mod10 returns the output of the mod 10 algo for a string
func Mod10(id string) int {
	idx := len(id) - 1
	sum := 0
	digitco := 0

	for i := idx; i >= 0; i-- {
		r := id[i]
		digit := (r - '0')
		if r >= '0' && r <= '9' {
			digitco++
		}
		switch {
		case r == ' ':
			continue
		case digitco%2 == 1 || r == '9' || r == '0':
			sum += int(digit)
		case r >= '1' && r <= '4':
			sum += int(digit << 1)
		case r >= '5' && r <= '8':
			sum += int(digit<<1) - 9
		}
	}
	return sum % 10
}

const letterBytes = "0123456789"

// GenISBN generates a valid ISBN10
func GenISBN() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 9)
	sum := 0
	for i := 0; i < 9; i++ {
		b[i] = rune(letterBytes[rand.Intn(len(letterBytes))])
		num := b[i] - '0'
		sum += int(num) * (10 - i)
	}
	return fmt.Sprintf("%s%d", string(b), sum%11)
}
